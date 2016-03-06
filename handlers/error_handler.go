package handlers

import (
	h "github.com/sepernol/sim-pos-api/helpers"
	"log"
	"net/http"
)

type ErrorMessage struct {
	UserMessage      string `json:"user_message"`
	DeveloperMessage string `json:"developer_message"`
	ErrorCode        int    `json:"error_code"`
	MoreInfo         string `json:"more_info"`
}

func handleError(err error, w http.ResponseWriter) {
	msg := &ErrorMessage{UserMessage: err.Error(), DeveloperMessage: err.Error(), ErrorCode: 500, MoreInfo: "www.google.com"}

	log.Println(err)
	h.ResponseJSON(w, msg)
}
