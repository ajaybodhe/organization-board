package apihelpers

import (
	"encoding/json"
	"net/http"
)

// WriteJSONResponse : send http response to client
func WriteJSONResponse(w http.ResponseWriter,
	r *http.Request,
	payload interface{},
	code int,
	err error) {

	var response []byte

	if nil != err {
		response = []byte(err.Error())
	} else {
		response, _ = json.Marshal(payload)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}
