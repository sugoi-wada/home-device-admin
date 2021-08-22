package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bamzi/jobrunner"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sugoi-wada/home-device-admin/db/model"
	"github.com/sugoi-wada/home-device-admin/graph"
	"github.com/sugoi-wada/home-device-admin/graph/generated"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	time.Local = time.FixedZone("UTC", 0)
	dsn := "host=localhost user=hikaru.wada dbname=home-device-admin-dev port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	jobrunner.Start()
	jobrunner.Now(UpdateCPDeviceStatus{DB: db})
	jobrunner.Every(10*time.Minute, UpdateCPDeviceStatus{DB: db})

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

type UpdateCPDeviceStatus struct {
	DB *gorm.DB
}

func (data UpdateCPDeviceStatus) Run() {
	fmt.Println("[Run] Update cp devices status...")
	timestamp := time.Now()
	devices := []model.CPDevice{{
		GatewayID: "test-gateway_id",
		Auth:      "test_auth",
		DeviceID:  "test_device_id",
		Nickname:  "test_nickname",
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}}
	for _, device := range devices {
		data.DB.Create(&device)
	}
}
