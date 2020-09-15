package apihelpers

import (
	"encoding/json"
	"net/http"

	"personio.com/organization-board/models"
)

// WriteJSONResponse : send http response to client
func WriteJSONResponse(w http.ResponseWriter,
	r *http.Request,
	payload interface{},
	code int,
	err error) {

	var response []byte

	if nil != err {
		errResponse := &models.ErrorResponse{
			ErrorMessage: err.Error(),
		}
		response, _ = json.Marshal(errResponse)
	} else {
		response, _ = json.Marshal(payload)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}
