package main

import (
	"github.com/gorilla/mux"
	"github.com/sepernol/sim-pos-api/handlers"
)

func CreateRoute() *mux.Router {
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

	/**
	* Suppliers
	* Create, Read, Update, Delete
	**/
	sup := r.PathPrefix("/suppliers/").Subrouter()
	sup.HandleFunc("/{id}", handlers.GetSupplier).Methods("GET")
	sup.HandleFunc("/{id}", handlers.PutSupplier).Methods("PUT")
	sup.HandleFunc("/{id}", handlers.DeleteSupplier).Methods("DELETE")

	sup.HandleFunc("/{size}/page/{page}", handlers.GetSuppliers).Methods("GET")
	sup.HandleFunc("/", handlers.PostSupplier).Methods("POST")
	/**
	* Suppliers end
	**/

	/**
	* Uoms
	* Create, Read, Update, Delete
	**/
	u := r.PathPrefix("/uoms/").Subrouter()
	u.HandleFunc("/{id}", handlers.GetUom).Methods("GET")
	u.HandleFunc("/{id}", handlers.PutUom).Methods("PUT")
	u.HandleFunc("/{id}", handlers.DeleteUom).Methods("DELETE")

	u.HandleFunc("/{size}/page/{page}", handlers.GetUoms).Methods("GET")
	u.HandleFunc("/", handlers.PostUom).Methods("POST")
	/**
	* Uoms end
	**/

	/**
	* Products
	**/
	pr := r.PathPrefix("/products/").Subrouter()
	ProductRouter(pr)

	return r
}

/**
* Products
* Create, Read, Update, Delete
**/
func ProductRouter(pr *mux.Router) {
	pr.HandleFunc("/{size}/page/{page}", handlers.GetProducts).Methods("GET")
	pr.HandleFunc("/", handlers.PostProduct).Methods("POST")

	pr.HandleFunc("/{id}", handlers.GetProduct).Methods("GET")
	pr.HandleFunc("/{id}", handlers.PutProduct).Methods("PUT")
	pr.HandleFunc("/{id}", handlers.DeleteProduct).Methods("DELETE")

	ProductUomRouter(pr.PathPrefix("/{id}/uoms").Subrouter())
}

/**
* Products UOM
* Create, Read, Delete
**/
func ProductUomRouter(pru *mux.Router) {
	pru.HandleFunc("/", handlers.GetProductUoms).Methods("GET")
	pru.HandleFunc("/", handlers.AddProductUom).Methods("POST")
	pru.HandleFunc("/{uom_id}", handlers.DeleteProductUom).Methods("DELETE")
}
