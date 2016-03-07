package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, value interface{}) (err error) {
	js, err := json.Marshal(value)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

type ResponseDataMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseJSONAndMessage(w http.ResponseWriter, value interface{}, message string) (err error) {
	var response ResponseDataMessage
	response.Message = message
	response.Data = value
	return ResponseJSON(w, response)
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func ResponseMessageOnly(w http.ResponseWriter, message string) error {
	var response ResponseMessage
	response.Message = message
	return ResponseJSON(w, response)
}
