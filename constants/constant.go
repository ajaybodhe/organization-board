package constants

import (
	"github.com/go-chi/jwtauth"
)

const (
	authToken           = "#1fsyHGFY8842sfg%$"
	AuthorizationHeader = "Authorization"
	MapClaimUser        = "user"
)

var (
	AuthToken = jwtauth.New("HS256", []byte(authToken), nil)
)
