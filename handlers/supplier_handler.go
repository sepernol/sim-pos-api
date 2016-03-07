package handlers

import (
	h "github.com/sepernol/sim-pos-api/helpers"
	m "github.com/sepernol/sim-pos-api/models"
	"net/http"
	"strconv"
)

func GetSupplier(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := m.GetSupplier(id)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	paging, err := h.GetPageParam(r)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := m.GetSuppliers(paging)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func PostSupplier(w http.ResponseWriter, r *http.Request) {
	postData := &m.Supplier{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	err = m.InsertSupplier(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Supplier Added")
}

func PutSupplier(w http.ResponseWriter, r *http.Request) {
	postData := &m.Supplier{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	id := h.GetQueryParam(r, "id")
	postData.Id, err = strconv.ParseInt(id, 10, 64)
	err = m.UpdateSupplier(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Supplier updated")
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	id := h.GetQueryParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	data, err := m.GetSupplier(idInt)
	err = m.DeleteSupplier(idInt)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Supplier deleted")
}
