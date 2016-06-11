package handlers

import (
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
	repo "github.com/sepernol/sim-pos-api/repositories"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := repo.GetProduct(id)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	paging, err := h.GetPageParam(r)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := repo.GetProducts(paging)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	postData := &e.Product{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	err = repo.InsertProduct(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Added")
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	postData := &e.Product{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	id := h.GetQueryParam(r, "id")
	postData.ID, err = strconv.ParseInt(id, 10, 64)
	err = repo.UpdateProduct(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product updated")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := h.GetQueryParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	data, err := repo.GetProduct(idInt)
	err = repo.DeleteProduct(idInt)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Product deleted")
}
