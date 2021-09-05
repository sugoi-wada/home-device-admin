package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bamzi/jobrunner"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sugoi-wada/home-device-admin/env"
	"github.com/sugoi-wada/home-device-admin/graph"
	"github.com/sugoi-wada/home-device-admin/graph/generated"
	"github.com/sugoi-wada/home-device-admin/worker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	time.Local = time.FixedZone("UTC", 0)
	db, err := gorm.Open(postgres.Open(env.Get("DB_DSN")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("CP", "Cp"),
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	jobrunner.Start()
	jobrunner.Now(worker.RefreshCPToken{DB: db})
	jobrunner.In(5*time.Minute, worker.FetchCPDeviceList{DB: db})
	jobrunner.Every(10*time.Minute, worker.FetchCPDeviceInfo{DB: db})
	jobrunner.Every(1*time.Hour, worker.RefreshCPToken{DB: db})

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

	err = e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}

func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}
