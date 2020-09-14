package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"personio.com/organization-board/apihelpers"
	"personio.com/organization-board/constants"
	"personio.com/organization-board/handlers"
	"personio.com/organization-board/models"
	"personio.com/organization-board/repository/login"
)

// Login : structure to facilitate login into the system
type Login struct {
	handlers.HTTPHandler
	repo *login.LoginRepository
}

// NewLoginHandler : Constructor for Login
func NewLoginHandler(conn *sql.DB) *Login {
	return &Login{
		repo: login.NewLoginRepository(conn),
	}
}

// GetHTTPHandler : return POST http handler for Login resource
func (lgn *Login) GetHTTPHandler() []*handlers.HTTPHandler {
	return []*handlers.HTTPHandler{
		&handlers.HTTPHandler{Authenticated: false,
			Method:  http.MethodPost,
			Version: 1,
			Path:    "login",
			Func:    lgn.Authenticate,
		},
	}
}

// Authenticate : Authenticate login request
func (lgn *Login) Authenticate(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	obj := new(models.Login)
	err := json.NewDecoder(r.Body).Decode(&obj)
	if nil != err {
		log.Printf("Error while reading the login request:%s", err.Error())
		apihelpers.WriteJSONResponse(w, r, user, http.StatusBadRequest, models.ErrInvalidRequest)
		return
	}

	if !obj.Valid() {
		log.Printf("Invalid login request object:%s", models.Stringify(obj))
		apihelpers.WriteJSONResponse(w, r, user, http.StatusBadRequest, models.ErrInvalidRequest)
		return
	}

	user, err = lgn.repo.Authenticate(r.Context(), obj)
	if nil != err {
		log.Printf("Error in user authentication: %s", err.Error())
		apihelpers.WriteJSONResponse(w, r, user, http.StatusUnauthorized, models.ErrUnauthorizedAccess)
		return
	}

	mapClaims := jwt.MapClaims{constants.MapClaimUser: user}
	_, tokenString, _ := constants.AuthToken.Encode(mapClaims)
	w.Header().Set(constants.AuthorizationHeader, tokenString)
	apihelpers.WriteJSONResponse(w, r, user, http.StatusOK, nil)
}
