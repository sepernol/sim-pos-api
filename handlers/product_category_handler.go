package handlers

import (
	e "github.com/sepernol/sim-pos-api/entities"
	h "github.com/sepernol/sim-pos-api/helpers"
	repo "github.com/sepernol/sim-pos-api/repositories"
	"net/http"
	"strconv"
)

//GetProductCategory handles GET method of /product-categories/:id
func GetProductCategory(w http.ResponseWriter, r *http.Request) {
	idStr := h.GetQueryParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := repo.GetProductCategory(id)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

//GetProductCategories handles GET method of
// /product-categories/:size/page/:page_no
func GetProductCategories(w http.ResponseWriter, r *http.Request) {
	paging, err := h.GetPageParam(r)
	if err != nil {
		handleError(err, w)
		return
	}
	result, err := repo.GetProductCategories(paging)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSON(w, result)
}

//PostProductCategory handles POST method of /product-categories
func PostProductCategory(w http.ResponseWriter, r *http.Request) {
	postData := &e.ProductCategory{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	err = repo.InsertProductCategory(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Category Added")
}

//PutProductCategory handles PUT method of /product-categories/:id
func PutProductCategory(w http.ResponseWriter, r *http.Request) {
	postData := &e.ProductCategory{}
	err := h.DecodeJSON(r, postData)
	if err != nil {
		handleError(err, w)
		return
	}
	id := h.GetQueryParam(r, "id")
	postData.ID, err = strconv.ParseInt(id, 10, 64)
	err = repo.UpdateProductCategory(postData)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, postData, "Product Category updated")
}

//DeleteProductCategory handles DELETE method of /product-categories/:id
func DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	id := h.GetQueryParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	data, err := repo.GetProductCategory(idInt)
	err = repo.DeleteProductCategory(idInt)
	if err != nil {
		handleError(err, w)
		return
	}
	h.ResponseJSONAndMessage(w, data, "Product Category deleted")
}
