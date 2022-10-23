package main

import (
	"fmt"
	"server/server"
	
)

func main() {
	fmt.Println("Starting aplication...")
	config, err := server.ReadConfig()
	if err != nil {
		fmt.Println("Error: %v",err)
	}
	fmt.Println("Database address: ",config.DATABASEADDR)
	server.StartServer(config)
}