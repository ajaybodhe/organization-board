package apihelpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"personio.com/organization-board/models"
)

// TestWriteJSONResponse_Valid : tests if valid json response reaches the client or not
func TestWriteJSONResponse_Valid(t *testing.T) {
	w := httptest.NewRecorder()
	expectedResponse := &models.EmployeeSupervisorResponse{
		Supervisor:             "Nick",
		SupervisorOfsupervisor: "Sophie",
	}
	expectedResponseJSON, err := json.Marshal(expectedResponse)
	assert.Nil(t, err)
	WriteJSONResponse(w, nil, expectedResponse, http.StatusOK, nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectedResponseJSON), w.Body.String())
}

// TestWriteJSONResponse_ErrorResponse : tests if error json response reaches the client or not
func TestWriteJSONResponse_ErrorResponse(t *testing.T) {
	w := httptest.NewRecorder()
	expectedError := errors.New("Your input is incorrect!")
	expectedErrorResponse := &models.ErrorResponse{
		ErrorMessage: expectedError.Error(),
	}
	expectedErrorResponseJSON, err := json.Marshal(expectedErrorResponse)
	assert.Nil(t, err)
	WriteJSONResponse(w, nil, nil, http.StatusBadRequest, expectedError)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expectedErrorResponseJSON), w.Body.String())
}
