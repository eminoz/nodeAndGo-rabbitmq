package database

import (
	"context"
	"github.com/eminoz/rabbitmq/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Database *mongo.Database

func SetDatabase() error {
	getConfig := config.GetConfig()
	var database *mongo.Database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getConfig.MongoDb))

	if err != nil {
		panic(err)
	}

	database = client.Database("rabbitUser")

	Database = database
	return nil
}
func GetDatabase() *mongo.Database {
	return Database
}
