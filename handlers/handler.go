package handlers

import (
	"net/http"
)

// IHTTPHandler : interface for http handlers
type IHTTPHandler interface {
	GetHTTPHandler() []*HTTPHandler
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

// HTTPHandler : Base Class for all HTTP Handler classes, implements IHTTPHandler
type HTTPHandler struct {
	Authenticated bool
	Method        string
	Version       int
	Path          string
	Func          func(http.ResponseWriter, *http.Request)
}

// GetHTTPHandler : Returns set of http handlers for path/version/method
func (hdlr *HTTPHandler) GetHTTPHandler() []HTTPHandler {
	return []HTTPHandler{}
}

// GetByID : GET method for a resource
func (hdlr *HTTPHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	return
}

// Create : POST method for a resource
func (hdlr *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	return
}

// Update : PUT method for a resource
func (hdlr *HTTPHandler) Update(w http.ResponseWriter, r *http.Request) {
	return
}

// Delete : DELETE method for a resource
func (hdlr *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	return
}

// GetAll : GET method for all the resources in a collection
func (hdlr *HTTPHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	return
}
