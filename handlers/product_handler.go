package handlers

import (
	h "github.com/sepernol/sim-pos-api/helpers"
	m "github.com/sepernol/sim-pos-api/models"
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
	result, err := m.GetProduct(id)
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
	result, err := m.GetProducts(paging)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	postData := &m.Product{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	err = m.InsertProduct(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Added")
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	postData := &m.Product{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	id := h.GetQueryParam(r, "id")
	postData.Id, err = strconv.ParseInt(id, 10, 64)
	err = m.UpdateProduct(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product updated")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := h.GetQueryParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	data, err := m.GetProduct(idInt)
	err = m.DeleteProduct(idInt)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Product deleted")
}
