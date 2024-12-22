package mongo_infra

import (
	config_dto "AnA-Roaming/repo-dto/config-dto"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"strings"
)

type MongoInfraImpl struct {
	MasterMongoDB *mongo.Database
	SlaveMongoDB  *mongo.Database
	Logger        *zap.SugaredLogger
}

func (m MongoInfraImpl) GetSlaveDB() *mongo.Database {
	return m.SlaveMongoDB
}

func (m MongoInfraImpl) GetMasterDB() *mongo.Database {
	return m.MasterMongoDB
}

func NewMongoInfra(lc fx.Lifecycle, config *config_dto.Config, logger *zap.SugaredLogger) (MongoInfra, error) {
	mongodbClient := MongoInfraImpl{
		Logger: logger,
	}

	// Connect to the master MongoDB
	mongoMasterURI := strings.Replace(config.MongoDB.Master.URI, "{mongo-user-name}", config.MongoDB.Master.Username, 1)
	mongoMasterURI = strings.Replace(mongoMasterURI, "{mongo-pswd}", config.MongoDB.Master.Password, 1)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoMasterURI))
	if err != nil {
		fmt.Println(fmt.Sprintf("Error connecting to MongoDB: %s", err.Error()))
		return mongodbClient, err
	}
	mongodbClient.MasterMongoDB = client.Database(config.MongoDB.Database)

	// Connect to the slave MongoDB
	mongoSlaveURI := strings.Replace(config.MongoDB.Slave.URI, "{mongo-user-name}", config.MongoDB.Slave.Username, 1)
	mongoSlaveURI = strings.Replace(mongoSlaveURI, "{mongo-pswd}", config.MongoDB.Slave.Password, 1)

	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoSlaveURI))
	if err != nil {
		fmt.Println(fmt.Sprintf("Error connecting to MongoDB: %s", err.Error()))
		return mongodbClient, err
	}
	mongodbClient.SlaveMongoDB = client.Database(config.MongoDB.Database)

	// Check if connected
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error pinging MongoDB: %s", err.Error()))
		return mongodbClient, err
	} else {
		mongodbClient.Logger.Info("NewMongoInfra: Connected to MongoDB")
	}

	// Close Connection on shut-down
	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			if err := client.Disconnect(context.Background()); err != nil {
				fmt.Println(fmt.Sprintf("Error disconnecting from MongoDB: %s", err.Error()))
				return err
			}
			return nil
		},
	})

	return mongodbClient, nil
}
