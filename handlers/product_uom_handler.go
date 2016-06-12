package handlers

import (
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
	repo "github.com/sepernol/sim-pos-api/repositories"
	"net/http"
	"strconv"
)

//AddProductUom handles POST method of /products/:id/uoms
func AddProductUom(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	productID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	postData := &e.ProductUom{}
	err = h.DecodeJSON(r, postData)
	postData.ProductID = productID
	if err != nil {
		handleError(err, w)
		return
	}
	err = repo.InsertProductUom(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	uom, err := repo.GetUom(postData.UomID)
	if err != nil {
		handleError(err, w)
		return
	}
	postData.UomDesc = uom.Description
	h.ResponseJSONAndMessage(w, postData, "Product UOM Added")
}

//GetProductUoms handles GET method of /products/:id/uoms
func GetProductUoms(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := repo.GetProductUoms(id)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

//DeleteProductUom handles DELETE method of /products/:id/uoms/:uom_id
func DeleteProductUom(w http.ResponseWriter, r *http.Request) {
	idStrProd := h.GetQueryParam(r, "id")
	productID, err := strconv.ParseInt(idStrProd, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}

	idStrUom := h.GetQueryParam(r, "uom_id")
	uomID, err := strconv.ParseInt(idStrUom, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	data, err := repo.GetProductUom(productID, uomID)
	if err != nil {
		handleError(err, w)
		return
	}
	err = repo.DeleteProductUom(productID, uomID)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Uom deleted")
}
