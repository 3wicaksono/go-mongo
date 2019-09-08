package infrastructures

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig config mongo struct
type MongoConfig struct {
	Host         string
	Port         int
	DatabaseName string
	DBTimeout    int
	User         string
	Password     string
}

// MongoConnect mongo connect object
type MongoConnect struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// MongoOpen start open connection to mongo db
func MongoOpen(config MongoConfig) (connect MongoConnect) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", config.User, config.Password, config.Host, config.Port, config.DatabaseName)))
	if err != nil {
		log.Fatalf("Couldn't connect to Mongo, err: %v", err.Error())
		return
	}
	database := client.Database(config.DatabaseName)

	connect.Client = client
	connect.Database = database
	return
}
