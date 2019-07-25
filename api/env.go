package api

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Env structure to inject dependencies of global scopes
type Env struct {
	DBClient *mongo.Client
	DBName   string
	Log      *zap.SugaredLogger
}
