package handlers

import (
	h "github.com/sepernol/sim-pos-api/helpers"
	m "github.com/sepernol/sim-pos-api/models"
	repo "github.com/sepernol/sim-pos-api/repositories"
	"net/http"
	"strconv"
)

func AddProductUnitPrice(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	productId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	postData := &m.ProductUnitPrice{}
	err = h.DecodeJSON(r, postData)
	postData.ProductId = productId
	if err != nil {
		handleError(err, w)
		return
	}
	err = m.InsertProductUnitPrice(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	uom, err := repo.GetUom(postData.UomId)
	if err != nil {
		handleError(err, w)
		return
	}
	postData.UomDesc = uom.Description
	h.ResponseJSONAndMessage(w, postData, "Product Unit Price Added")
}

func GetProductUnitPrices(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	productId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := m.GetProductUnitPrices(productId)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

func DeleteProductUnitPrice(w http.ResponseWriter, r *http.Request) {
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
	data, err := m.GetProductUnitPrice(productId, uomId)
	if err != nil {
		handleError(err, w)
		return
	}
	err = m.DeleteProductUnitPrice(productId, uomId)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Product Unit Price deleted")
}

func PutProductUnitPrice(w http.ResponseWriter, r *http.Request) {
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
	postData := &m.ProductUnitPrice{}
	err = h.DecodeJSON(r, postData)
	postData.ProductId = productId
	postData.UomId = uomId
	err = m.UpdateProductUnitPrice(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Unit Price updated")
}
