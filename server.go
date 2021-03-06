package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bamzi/jobrunner"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sugoi-wada/home-device-admin/client/cp_client"
	db_config "github.com/sugoi-wada/home-device-admin/db/config"
	"github.com/sugoi-wada/home-device-admin/graph"
	"github.com/sugoi-wada/home-device-admin/graph/generated"
	"github.com/sugoi-wada/home-device-admin/job"
	"github.com/sugoi-wada/home-device-admin/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	time.Local = time.FixedZone("UTC", 0)
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), db_config.GetConf())

	if err != nil {
		log.Fatalln(err)
	}

	jobrunner.Start()
	if os.Getenv("DEBUG") != "true" {
		client := cp_client.NewClient()
		repo := &repo.CPDeviceRepo{DB: db, Client: client}
		jobrunner.Now(job.RefreshCPToken{Repo: repo})
		jobrunner.In(5*time.Minute, job.FetchCPDeviceList{Repo: repo})
		jobrunner.Every(10*time.Minute, job.FetchCPDeviceInfo{Repo: repo})
		jobrunner.Every(1*time.Hour, job.RefreshCPToken{Repo: repo})
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", welcome())
	e.GET("/jobrunner/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, jobrunner.StatusJson())
	})

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err = e.Start(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}
}

func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}
