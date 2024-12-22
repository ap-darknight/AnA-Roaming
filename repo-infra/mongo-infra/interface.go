package mongo_infra

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInfra interface {
	GetSlaveDB() *mongo.Database
	GetMasterDB() *mongo.Database
}
