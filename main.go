package main

import (
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	r := CreateRoute()

	port := os.Getenv("SIM_POS_LISTEN_PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	log.Println("Server started on Port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
