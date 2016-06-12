package handlers

import (
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
	repo "github.com/sepernol/sim-pos-api/repositories"
	"net/http"
	"strconv"
)

//AddProductUnitPrice handles POST method of /products/:id/unit-prices/:uom_id
func AddProductUnitPrice(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	productID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	postData := &e.ProductUnitPrice{}
	err = h.DecodeJSON(r, postData)
	postData.ProductID = productID
	if err != nil {
		handleError(err, w)
		return
	}
	err = repo.InsertProductUnitPrice(postData)
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
	h.ResponseJSONAndMessage(w, postData, "Product Unit Price Added")
}

//GetProductUnitPrices handles GET method of /products/:id/unit-prices
func GetProductUnitPrices(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	productID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := repo.GetProductUnitPrices(productID)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

//DeleteProductUnitPrice handles DELETE method
//for /products/:id/unit-prices/:uom_id
func DeleteProductUnitPrice(w http.ResponseWriter, r *http.Request) {
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
	data, err := repo.GetProductUnitPrice(productID, uomID)
	if err != nil {
		handleError(err, w)
		return
	}
	err = repo.DeleteProductUnitPrice(productID, uomID)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Product Unit Price deleted")
}

//PutProductUnitPrice handles PUT method of /products/:id/unit-prices/:uom_id
func PutProductUnitPrice(w http.ResponseWriter, r *http.Request) {
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
	postData := &e.ProductUnitPrice{}
	err = h.DecodeJSON(r, postData)
	postData.ProductID = productID
	postData.UomID = uomID
	err = repo.UpdateProductUnitPrice(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Unit Price updated")
}
