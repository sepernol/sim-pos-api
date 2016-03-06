package helpers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PageParams struct {
	Size   int
	Page   int
	Offset int
}

func GetPageParam(r *http.Request) (pageParam PageParams, err error) {
	size := GetQueryParam(r, "size")
	page := GetQueryParam(r, "page")

	pageParam.Size, err = strconv.Atoi(size)
	if err != nil {
		return
	}
	pageParam.Page, err = strconv.Atoi(page)
	if err != nil {
		return
	}

	if pageParam.Page < 1 {
		err = errors.New("Page cannot less than 1")
		return
	}
	if pageParam.Size < 1 {
		err = errors.New("Size cannot less than 1")
		return
	}

	pageParam.Offset = (pageParam.Page - 1) * pageParam.Size
	return
}

func GetQueryParam(r *http.Request, key string) string {
	params := mux.Vars(r)
	value := params[key]
	return value
}

func DecodeJSON(r *http.Request, postData interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(postData)
	if err != nil {
		return
	}
	return
}
