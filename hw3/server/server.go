package server

import (
	"encoding/json"
	"fmt"
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
	}
}

func StartServer() {

	server := &http.Server{
		Addr: ":10004",
	}
	http.HandleFunc("/", mainHandler)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
