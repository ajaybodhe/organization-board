package login

import (
	"bytes"
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"personio.com/organization-board/constants"
	"personio.com/organization-board/models"
)

// TestCache_Authenticate_Success: test for login success
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

	rows := sqlmock.NewRows([]string{"id", "email"}).AddRow(user.ID, user.Email)
	mock.ExpectQuery(regexp.QuoteMeta(constants.LoginDetailsSelectQuery)).WithArgs(user.Email, user.Password).WillReturnRows(rows)

	loginRepository := NewLoginRepository(db)

	loginModel := &models.Login{
		Email:    "personia@personio.com",
		Password: "personia",
	}

	cntx := context.Background()
	dbuser, err := loginRepository.Authenticate(cntx, loginModel)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, dbuser.ID)
	assert.Equal(t, user.Email, dbuser.Email)
}

// TestCache_Authenticate_Fail: test for login fail
func TestAuthenticate_Fail(t *testing.T) {
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

	rows := sqlmock.NewRows([]string{"id", "email"})
	mock.ExpectQuery(regexp.QuoteMeta(buffer.String())).WithArgs(user.Email, user.Password).WillReturnRows(rows)

	loginRepository := NewLoginRepository(db)

	loginModel := &models.Login{
		Email:    "personia@personio.com",
		Password: "personia",
	}

	cntx := context.Background()
	_, err = loginRepository.Authenticate(cntx, loginModel)
	assert.NotNil(t, err)
}
