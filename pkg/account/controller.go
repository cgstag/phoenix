package account

import (
	"phoenix/api"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type resource struct {
	db  repository
	log *zap.SugaredLogger
}

type repository struct {
	mongo    *mongo.Client
	database string
	log      *zap.SugaredLogger
}

func ServeResources(env *api.Env, router *echo.Group) {
	r := &resource{repository{mongo: env.DBClient, database: env.DBName}, env.Log}
	rg := router.Group("/account")

	// CRUD ACCOUNT - account.go
	rg.POST("/random", r.random)
	rg.POST("/", r.new)
	rg.GET("/:uuid", r.get)
	rg.GET("/", r.getAll)
	rg.PUT("/:uuid", r.set)
	rg.DELETE("/:uuid", r.delete)
}
