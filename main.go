package main

import (
	"context"
	"fmt"
	"net/http"
	"phoenix/api"
	"phoenix/pkg/account"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {

	// Load Config @TODO Docker Config
	//configuration := config.MustLoadConfig()

	// Load mongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// Initialize Logger
	log = zap.NewExample().Sugar()
	defer log.Sync()

	// Initialize Echo
	e := echo.New()

	// Initialize Middleware
	e.Use(middleware.Recover())
	router := e.Group("/v1")

	env := &api.Env{DBClient: client, DBName: "bank", Log: log}

	// Serve Routes
	account.ServeResources(env, router)

	// Healthcheck
	e.GET("/", func(c echo.Context) error {
		log.Infow("Calling Hello World...")
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.Match([]string{"GET", "HEAD"}, "/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	// Start server
	//address := fmt.Sprintf("%v:%v", configuration.Host, configuration.Port)
	address := fmt.Sprintf("%v:%v", "localhost", "8081")
	e.Logger.Fatal(e.Start(address))
}
