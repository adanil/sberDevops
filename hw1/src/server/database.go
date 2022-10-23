package server
import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

func InitConnection(addr string) error {
	clientOptions := options.Client().ApplyURI(addr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("Logs").Collection("get_logs")
	return nil
}

func Insert(resp *SuccessfulResponse) error {
	_, err := collection.InsertOne(ctx, resp)
	if err != nil {
		log.Printf("Error: %v",err)
	}
  	return err
}