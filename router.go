package main

import (
	"github.com/gorilla/mux"
	"github.com/sepernol/sim-pos-api/handlers"
)

//CreateRoute creates application route
func CreateRoute() *mux.Router {
	r := mux.NewRouter()

	pc := r.PathPrefix("/product-categories/").Subrouter()
	productCategoriesRouter(pc)

	sup := r.PathPrefix("/suppliers/").Subrouter()
	supplierRouter(sup)

	u := r.PathPrefix("/uoms/").Subrouter()
	uomRouter(u)

	pr := r.PathPrefix("/products/").Subrouter()
	productRouter(pr)

	return r
}

func productCategoriesRouter(pc *mux.Router) {
	pc.HandleFunc("/{id}", handlers.GetProductCategory).Methods("GET")
	pc.HandleFunc("/{id}", handlers.PutProductCategory).Methods("PUT")
	pc.HandleFunc("/{id}", handlers.DeleteProductCategory).Methods("DELETE")

	pc.HandleFunc("/{size}/page/{page}", handlers.GetProductCategories).Methods("GET")
	pc.HandleFunc("/", handlers.PostProductCategory).Methods("POST")
}

func supplierRouter(sup *mux.Router) {
	sup.HandleFunc("/{id}", handlers.GetSupplier).Methods("GET")
	sup.HandleFunc("/{id}", handlers.PutSupplier).Methods("PUT")
	sup.HandleFunc("/{id}", handlers.DeleteSupplier).Methods("DELETE")

	sup.HandleFunc("/{size}/page/{page}", handlers.GetSuppliers).Methods("GET")
	sup.HandleFunc("/", handlers.PostSupplier).Methods("POST")
}

func uomRouter(u *mux.Router) {
	u.HandleFunc("/{id}", handlers.GetUom).Methods("GET")
	u.HandleFunc("/{id}", handlers.PutUom).Methods("PUT")
	u.HandleFunc("/{id}", handlers.DeleteUom).Methods("DELETE")

	u.HandleFunc("/{size}/page/{page}", handlers.GetUoms).Methods("GET")
	u.HandleFunc("/", handlers.PostUom).Methods("POST")
}

func productRouter(pr *mux.Router) {
	pr.HandleFunc("/{size}/page/{page}", handlers.GetProducts).Methods("GET")
	pr.HandleFunc("/", handlers.PostProduct).Methods("POST")

	pr.HandleFunc("/{id}", handlers.GetProduct).Methods("GET")
	pr.HandleFunc("/{id}", handlers.PutProduct).Methods("PUT")
	pr.HandleFunc("/{id}", handlers.DeleteProduct).Methods("DELETE")

	productUomRouter(pr.PathPrefix("/{id}/uoms").Subrouter())
	productUnitPriceRouter(pr.PathPrefix("/{id}/unit-prices").Subrouter())
}

func productUomRouter(pru *mux.Router) {
	pru.HandleFunc("/", handlers.GetProductUoms).Methods("GET")
	pru.HandleFunc("/", handlers.AddProductUom).Methods("POST")
	pru.HandleFunc("/{uom_id}", handlers.DeleteProductUom).Methods("DELETE")
}

func productUnitPriceRouter(up *mux.Router) {
	up.HandleFunc("/", handlers.GetProductUnitPrices).Methods("GET")
	up.HandleFunc("/", handlers.AddProductUnitPrice).Methods("POST")
	up.HandleFunc("/{uom_id}", handlers.PutProductUnitPrice).Methods("PUT")
	up.HandleFunc("/{uom_id}", handlers.DeleteProductUnitPrice).Methods("DELETE")
}
