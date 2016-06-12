package main

import (
	"log"
	"net/http"
	"os"
)

const (
	//DefaultPort is used when env SIM_POS_LISTEN_PORT is not defined
	DefaultPort = "8080"
)

func main() {
	r := CreateRoute()

	port := os.Getenv("SIM_POS_LISTEN_PORT")
	if port == "" {
		port = DefaultPort
	}
	log.Println("Server started on Port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
