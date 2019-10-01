package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Use to connect to mongodb instance
func ConnectDb() *mongo.Client {
	dbconn := viper.GetString("mongo_uri")
	client, err := mongo.NewClient(options.Client().ApplyURI(dbconn))
	if err != nil {
		log.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Error(err)
	}
	return client
}
