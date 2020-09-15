package http

import (
	"bytes"
	gohttp "net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"personio.com/organization-board/constants"

	"context"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestingEmployeeManagerMapCreate : tests handler for creating Employee Manager Map
func TestEmployeeManagerMapCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employee_name", "manager_name"}).AddRow("Nick", "Sophie").AddRow("Sophie", "Jonas")
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectPrepare(regexp.QuoteMeta(constants.EmployeeManagerMappingInsertQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery(constants.EmployeeManagerMappingSelectQuery).WillReturnRows(rows)

	emplyManagerMap := NewEmployeeManagerMapHandler(db)

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"Peter":"Nick", "Barbara":"Nick", "Nick":"Sophie", "Sophie":"Jonas"}`)
	r := httptest.NewRequest("POST", "http://localhost:9090/api/v1/emplymgrmap", bytes.NewBuffer(jsonStr))
	r = r.WithContext(context.Background())
	emplyManagerMap.Create(w, r)

	expectedResponse := `{"Jonas":{"Sophie":{"Nick":{"Barbara":{},"Peter":{}}}}}`
	assert.Equal(t, gohttp.StatusOK, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

// TestEmployeeManagerMapCreate_BadRequest : tests for bad request from client
func TestEmployeeManagerMapCreate_BadRequest(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	emplyManagerMap := NewEmployeeManagerMapHandler(db)

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"invalidjson":}`)
	r := httptest.NewRequest("POST", "http://localhost:9090/api/v1/emplymgrmap", bytes.NewBuffer(jsonStr))
	r = r.WithContext(context.Background())
	emplyManagerMap.Create(w, r)

	expectedResponse := `{"error_message":"Error:: Invalid Request"}`
	assert.Equal(t, gohttp.StatusBadRequest, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

// TestEmployeeManagerMapCreate_DBError : tests for internal server error
func TestEmployeeManagerMapCreate_DBError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	emplyManagerMap := NewEmployeeManagerMapHandler(db)

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"Peter":"Nick", "Barbar":"Nick", "Nick":"Sophie", "Sophie":"Jonas"}`)
	r := httptest.NewRequest("POST", "http://localhost:9090/api/v1/emplymgrmap", bytes.NewBuffer(jsonStr))
	r = r.WithContext(context.Background())
	emplyManagerMap.Create(w, r)

	expectedResponse := `{"error_message":"Error:: Record can not be added to DB"}`
	assert.Equal(t, gohttp.StatusInternalServerError, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

// TestEmployeeManagerMapGetByID : test the handler response for get supervisor info for an employee
func TestEmployeeManagerMapGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employee_name", "manager_name"}).AddRow("Nick", "Sophie").AddRow("Sophie", "Jonas")
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectPrepare(regexp.QuoteMeta(constants.EmployeeManagerMappingInsertQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery(constants.EmployeeManagerMappingSelectQuery).WillReturnRows(rows)

	emplyManagerMap := NewEmployeeManagerMapHandler(db)

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"Peter":"Nick", "Barbara":"Nick", "Nick":"Sophie", "Sophie":"Jonas"}`)
	r := httptest.NewRequest("POST", "http://localhost:9090/api/v1/emplymgrmap", bytes.NewBuffer(jsonStr))
	r = r.WithContext(context.Background())
	emplyManagerMap.Create(w, r)

	r = httptest.NewRequest("GET", "http://localhost:9090/api/v1/emplymgrmap/Nick?supervisor=true&name=Nick", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "Nick")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	w = httptest.NewRecorder()
	emplyManagerMap.GetByID(w, r)
	expectedResponse := `{"supervisor":"Sophie","supervisor_of_supervisor":"Jonas"}`
	assert.Equal(t, gohttp.StatusOK, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

// TestEmployeeManagerMapGetAll : test the handler response for get the entire employee manager relationship map
func TestEmployeeManagerMapGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employee_name", "manager_name"}).AddRow("Nick", "Sophie").AddRow("Sophie", "Jonas")
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectPrepare(regexp.QuoteMeta(constants.EmployeeManagerMappingInsertQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery(constants.EmployeeManagerMappingSelectQuery).WillReturnRows(rows)

	emplyManagerMap := NewEmployeeManagerMapHandler(db)

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"Peter":"Nick", "Barbara":"Nick", "Nick":"Sophie", "Sophie":"Jonas"}`)
	r := httptest.NewRequest("POST", "http://localhost:9090/api/v1/emplymgrmap", bytes.NewBuffer(jsonStr))
	r = r.WithContext(context.Background())
	emplyManagerMap.Create(w, r)

	w = httptest.NewRecorder()
	emplyManagerMap.GetAll(w, r)
	expectedResponse := `{"Jonas":{"Sophie":{"Nick":{"Barbara":{},"Peter":{}}}}}`
	assert.Equal(t, gohttp.StatusOK, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}
