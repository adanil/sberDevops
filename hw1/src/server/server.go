package server

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"time"
	"encoding/json"
)

type SuccessfulResponse struct {
	Code		int32 	`json:"code" bson:"code"`
	SystemTime 	string 	`json:"time" bson:"time"`
	Message		string 	`json:"message" bson:"message"`
	Service 	string 	`json:"service" bson:"service"`
}

func mainHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("mainHandler")
	switch r.Method {
	case "GET":
		content, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error %v",err)
		}
		sResp := SuccessfulResponse{200,time.Now().String(),string(content),"GO server"}
		Insert(&sResp)

		json, err := json.Marshal(sResp)
		if err != nil {
			log.Printf("Error %v", err)
		}

		w.Header().Set("Content-Type","application/json")
		w.Write(json)
		w.Write([]byte("\n"))
	}
}


func StartServer(config *Config){

	server := &http.Server{
		Addr:           ":10003",
	}

	err := InitConnection(config.DATABASEADDR)
	if err != nil {
		log.Printf("Error: %v",err)
	}

	http.HandleFunc("/",mainHandler)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}