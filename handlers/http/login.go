package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"personio.com/organization-board/constants"
	"personio.com/organization-board/handlers"
	"personio.com/organization-board/models"
	"personio.com/organization-board/repository/login"
)

type Login struct {
	handlers.HTTPHandler
	repo *login.LoginRepository
}

func NewLoginHandler(conn *sql.DB) *Login {
	return &Login{
		repo: login.NewLoginRepository(conn),
	}
}

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

func (lgn *Login) Authenticate(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	obj := new(models.Login)
	err := json.NewDecoder(r.Body).Decode(&obj)
	for {
		if nil != err {
			log.Printf("Error while reading the login request:%s", err.Error())
			break
		}

		if !obj.Valid() {
			err = models.InvalidRequest
			log.Printf("Invalid login request object:%s", models.Stringify(obj))
			break
		}

		user, err = lgn.repo.Authenticate(r.Context(), obj)
		if nil != err {
			log.Printf("Error in user authentication: %s", err.Error())
			break
		}

		mapClaims := jwt.MapClaims{constants.MapClaimUser: user}
		_, tokenString, _ := constants.AuthToken.Encode(mapClaims)
		w.Header().Set(constants.AuthorizationHeader, tokenString)
		break
	}

	handlers.WriteJSONResponse(w, r, user, http.StatusOK, err)
}
