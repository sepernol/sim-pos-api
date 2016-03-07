package handlers

import (
	h "github.com/sepernol/sim-pos-api/helpers"
	m "github.com/sepernol/sim-pos-api/models"
	"net/http"
	"strconv"
)

func GetProductCategory(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := m.GetProductCategory(id)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func GetProductCategories(w http.ResponseWriter, r *http.Request) {
	paging, err := h.GetPageParam(r)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := m.GetProductCategories(paging)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func PostProductCategory(w http.ResponseWriter, r *http.Request) {
	postData := &m.ProductCategory{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	err = m.InsertProductCategory(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Category Added")
}

func PutProductCategory(w http.ResponseWriter, r *http.Request) {
	postData := &m.ProductCategory{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	id := h.GetQueryParam(r, "id")
	postData.Id, err = strconv.ParseInt(id, 10, 64)
	err = m.UpdateProductCategory(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Category updated")
}

func DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	id := h.GetQueryParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	data, err := m.GetProductCategory(idInt)
	err = m.DeleteProductCategory(idInt)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Product Category deleted")
}
