package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLogin_Valid : tests valid login
func TestLogin_Valid(t *testing.T) {
	login := &Login{
		Email:    "testemail",
		Password: "testpassword",
	}
	assert.True(t, login.Valid())
}

// TestLogin_InValid : tests invalid login
func TestLogin_InValid(t *testing.T) {
	login := &Login{
		Email: "testemail",
	}
	assert.False(t, login.Valid())
}
