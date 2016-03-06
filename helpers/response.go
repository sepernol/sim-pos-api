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
