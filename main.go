package main

import (
	"github.com/gorilla/mux"
	"github.com/sepernol/sim-pos-api/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	/**
	* Product categories
	* Create, Read, Update, Delete
	**/
	pc := r.PathPrefix("/product_categories/").Subrouter()
	pc.HandleFunc("/{id}", handlers.GetProductCategory).Methods("GET")
	pc.HandleFunc("/{id}", handlers.PutProductCategory).Methods("PUT")
	pc.HandleFunc("/{id}", handlers.DeleteProductCategory).Methods("DELETE")

	pc.HandleFunc("/{size}/page/{page}", handlers.GetProductCategories).Methods("GET")
	pc.HandleFunc("/", handlers.PostProductCategory).Methods("POST")
	/**
	* Product categories end
	**/

	log.Println("Server started")

	log.Fatal(http.ListenAndServe(":8080", r))
}
