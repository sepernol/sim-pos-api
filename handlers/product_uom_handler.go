package handlers

import (
	h "github.com/sepernol/sim-pos-api/helpers"
	m "github.com/sepernol/sim-pos-api/models"
	"net/http"
	"strconv"
)

func AddProductUom(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	productId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	postData := &m.ProductUom{}
	err = h.DecodeJSON(r, postData)
	postData.ProductId = productId
	if err != nil {
		handleError(err, w)
		return
	}
	err = m.InsertProductUom(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	uom, err := m.GetUom(postData.UomId)
	if err != nil {
		handleError(err, w)
		return
	}
	postData.UomDesc = uom.Description
	h.ResponseJSONAndMessage(w, postData, "Product UOM Added")
}

func GetProductUoms(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := m.GetProductUoms(id)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func DeleteProductUom(w http.ResponseWriter, r *http.Request) {
	idStrProd := h.GetQueryParam(r, "id")
	productId, err := strconv.ParseInt(idStrProd, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}

	idStrUom := h.GetQueryParam(r, "uom_id")
	uomId, err := strconv.ParseInt(idStrUom, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	data, err := m.GetProductUom(productId, uomId)
	if err != nil {
		handleError(err, w)
		return
	}
	err = m.DeleteProductUom(productId, uomId)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Uom deleted")
}
