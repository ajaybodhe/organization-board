package handlers

import (
	"encoding/json"
	"net/http"
)

// IHTTPHandler : interface  http handler
type IHTTPHandler interface {
	GetHTTPHandler() []*HTTPHandler
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

// HTTPHandler : implements IHTTPHandler
type HTTPHandler struct {
	Authenticated bool
	Method        string
	Version       int
	Path          string
	Func          func(http.ResponseWriter, *http.Request)
}

type response struct {
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func (hdlr *HTTPHandler) GetHTTPHandler() []HTTPHandler {
	return []HTTPHandler{}
}

func (hdlr *HTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	return
}

func (hdlr *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	return
}

func WriteJSONResponse(w http.ResponseWriter,
	r *http.Request,
	payload interface{},
	code int,
	err error) {

	resp := &response{
		Data: payload,
	}

	if nil != err {
		resp.ErrorMessage = err.Error()
	}

	response, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}
