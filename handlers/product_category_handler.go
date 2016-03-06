package handlers

import (
	h "github.com/sepernol/sim-pos/helpers"
	m "github.com/sepernol/sim-pos/models"
	"net/http"
	"strconv"
)

func GetProductCategory(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	id, err := strconv.Atoi(idStr)
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
	h.ResponseJSON(w, postData)
}
