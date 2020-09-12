package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"

	"personio.com/organization-board/config"
	"personio.com/organization-board/constants"
	"personio.com/organization-board/db"
	"personio.com/organization-board/handlers"
	httpHdlr "personio.com/organization-board/handlers/http"
)

var (
	handlerList = []handlers.IHTTPHandler{}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Connection Interface is used for database portability
	var db db.IDB = new(db.SQLLite)
	// check application able to create connection or not
	conn, err := db.NewConnection()
	if nil != err {
		log.Fatalf("Error while making db connectiion:%s", err.Error())
	}

	handlerList = []handlers.IHTTPHandler{
		httpHdlr.NewLoginHandler(conn),
	}

}

func createRouterGroup(router *chi.Mux, authenticated bool) {
	// Authenticated routes
	router.Group(func(r chi.Router) {
		if authenticated { // set authetication checks
			r.Use( // Seek, verify and validate JWT tokens
				jwtauth.Verifier(constants.AuthToken),
				jwtauth.Authenticator,
			)
		}

		for _, hdlr := range handlerList { // register all handlers
			for _, hlr := range hdlr.GetHTTPHandler() {
				path := fmt.Sprintf("/api/v%d/%s", hlr.Version, hlr.Path)
				if authenticated == hlr.Authenticated {
					log.Println("creating router:", path)
					switch hlr.Method {
					case http.MethodGet:
						r.Get(path, hlr.Func)
					case http.MethodPost:
						r.Post(path, hlr.Func)
					case http.MethodPut:
						r.Put(path, hlr.Func)
					case http.MethodDelete:
						r.Delete(path, hlr.Func)
					default:
						log.Println("Not supported HTTP request method")
					}
				}
			}
		}
	})
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	createRouterGroup(router, true)  // group authenticated routes
	createRouterGroup(router, false) //group unauthenticated routes

	http.ListenAndServe(fmt.Sprintf("%s:%d",
		config.Config().Host, config.Config().Port), router)
}
