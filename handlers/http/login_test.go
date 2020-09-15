package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"personio.com/organization-board/constants"
	"personio.com/organization-board/models"
)

// TestCache_Authenticate_Success: test for authentication success
func TestAuthenticate_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	user := models.User{
		ID:       1,
		Email:    "personia@personio.com",
		Password: "personia",
	}

	var buffer bytes.Buffer
	buffer.WriteString(constants.LoginDetailsSelectQuery)

	rows := sqlmock.NewRows([]string{"id", "email"}).AddRow(user.ID, user.Email)
	mock.ExpectQuery(regexp.QuoteMeta(buffer.String())).WithArgs(user.Email, user.Password).WillReturnRows(rows)

	lgnHandler := NewLoginHandler(db)

	reqBody, _ := json.Marshal(models.Login{Email: "personia@personio.com", Password: "personia"})
	r, err := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	lgnHandler.Authenticate(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusOK)
	assert.NotEmpty(t, w.Header().Get("Authorization"))
}

// TestAuthenticate_Failed: test for authentication failed
func TestAuthenticate_Failed(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	user := models.User{
		ID:       1,
		Email:    "personia@personio.com",
		Password: "personia",
	}

	var buffer bytes.Buffer
	buffer.WriteString(constants.LoginDetailsSelectQuery)

	mock.ExpectQuery(regexp.QuoteMeta(buffer.String())).WithArgs(user.Email, user.Password).WillReturnError(errors.New("MOCK ERROR: Authentication Failed"))

	lgnHandler := NewLoginHandler(db)

	reqBody, _ := json.Marshal(models.Login{Email: "personia@personio.com", Password: "personia"})
	r, err := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	lgnHandler.Authenticate(w, r)
	assert.Equal(t, w.Result().StatusCode, http.StatusUnauthorized)
	assert.Empty(t, w.Header().Get("Authorization"))
}
