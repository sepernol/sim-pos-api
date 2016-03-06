package main

import (
	"github.com/gorilla/mux"
	"github.com/sepernol/sim-pos/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/product_categories/{id}", handlers.GetProductCategory).Methods("GET")
	r.HandleFunc("/product_categories/{size}/page/{page}", handlers.GetProductCategories).Methods("GET")
	r.HandleFunc("/product_categories", handlers.PostProductCategory).Methods("POST")

	log.Println("Server started")

	http.ListenAndServe(":8080", r)
}
