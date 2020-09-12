package http

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"personio.com/organization-board/handlers"
	"personio.com/organization-board/models"
	"personio.com/organization-board/repository/emplymgrmap"
)

type EmployeeManagerMap struct {
	handlers.HTTPHandler
	repo *emplymgrmap.EmployeeManagerMapRepository
}

func NewEmployeeManagerMapHandler(conn *sql.DB) *EmployeeManagerMap {
	return &EmployeeManagerMap{
		repo: emplymgrmap.NewEmployeeManagerMapRepository(conn),
	}
}

func (emplyMgrMap *EmployeeManagerMap) GetHTTPHandler() []*handlers.HTTPHandler {
	return []*handlers.HTTPHandler{
		&handlers.HTTPHandler{Authenticated: true,
			Method:  http.MethodPost,
			Version: 1,
			Path:    "emplymgrmap",
			Func:    emplyMgrMap.Create,
		},
	}
}

func (emplyMgrMap *EmployeeManagerMap) Create(w http.ResponseWriter, r *http.Request) {
	var reqEmplyMgrMap models.EmployeeManagerMap
	err := json.NewDecoder(r.Body).Decode(&reqEmplyMgrMap)
	for {
		if nil != err {
			break
		}

		_, err = emplyMgrMap.repo.Create(r.Context(), reqEmplyMgrMap)
		break
	}

	handlers.WriteJSONResponse(w, r, reqEmplyMgrMap, http.StatusOK, err)
}
