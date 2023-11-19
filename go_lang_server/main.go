package main

import (
	"ai_dev_private_api/ownapi"
	"ai_dev_private_api/ownapipro"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}
	//TODO support Graceful Shutdown
	serverRouter := mux.NewRouter()
	serverRouter.HandleFunc("/ownapi/{uuid}", ownapi.ProcessRequest).Methods("POST")
	serverRouter.HandleFunc("/ownapipro/{uuid}", ownapipro.ProcessRequest).Methods("POST")
	portNum := os.Getenv("PORT")

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C")

	// Spinning up the server.
	err = http.ListenAndServe(":"+portNum, serverRouter)
	if err != nil {
		log.Fatal(err)
	}
}
