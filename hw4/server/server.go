package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
	"time"
)

type SuccessfulResponse struct {
	Code       int32  `json:"code" bson:"code"`
	SystemTime string `json:"time" bson:"time"`
	Message    string `json:"message" bson:"message"`
	Service    string `json:"service" bson:"service"`
}

var databaseAddr string
var client *mongo.Client

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mainHandler")
	switch r.Method {
	case "GET":
		content, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error %v", err)
		}
		sResp := SuccessfulResponse{200, time.Now().String(), string(content), "GO server"}

		json, err := json.Marshal(sResp)
		if err != nil {
			log.Printf("Error %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
		w.Write([]byte("\n"))

		collection := client.Database("HW4").Collection("logs")
		_, err = collection.InsertOne(context.TODO(), sResp)
		if err != nil {
			fmt.Printf("Couldn't insert log into mongodb: %v\n", err)
		}
	}
}

func initConntectionDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(databaseAddr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err
}

func StartServer(cfg *config) {
	databaseAddr = cfg.DatabaseAddr
	var err error
	client, err = initConntectionDB()
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr: ":9091",
	}
	http.HandleFunc("/", mainHandler)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
