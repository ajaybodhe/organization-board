package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"personio.com/organization-board/apihelpers"
	"personio.com/organization-board/cache"
	"personio.com/organization-board/handlers"
	"personio.com/organization-board/models"
	"personio.com/organization-board/repository/emplymgrmap"
)

// EmployeeManagerMap : provides http handlers for EmployeeManagerMap resource
type EmployeeManagerMap struct {
	handlers.HTTPHandler
	repo *emplymgrmap.EmployeeManagerMapRepository
}

// NewEmployeeManagerMapHandler : constructor for EmployeeManagerMap
func NewEmployeeManagerMapHandler(conn *sql.DB) *EmployeeManagerMap {
	return &EmployeeManagerMap{
		repo: emplymgrmap.NewEmployeeManagerMapRepository(conn),
	}
}

// GetHTTPHandler : returns POST http handler for EmployeeManagerMap resource
func (emplyMgrMap *EmployeeManagerMap) GetHTTPHandler() []*handlers.HTTPHandler {
	return []*handlers.HTTPHandler{
		&handlers.HTTPHandler{Authenticated: true,
			Method:  http.MethodGet,
			Version: 1, // TODO version your API , it should be date rather than version no
			Path:    "emplymgrmap/{name}",
			Func:    emplyMgrMap.GetByID,
		},
		&handlers.HTTPHandler{Authenticated: true,
			Method:  http.MethodPost,
			Version: 1, // TODO version your API , it should be date rather than version no
			Path:    "emplymgrmap",
			Func:    emplyMgrMap.Create,
		},
		&handlers.HTTPHandler{Authenticated: true,
			Method:  http.MethodPut,
			Version: 1, // TODO version your API , it should be date rather than version no
			Path:    "emplymgrmap",
			Func:    emplyMgrMap.Update,
		},
		&handlers.HTTPHandler{Authenticated: true,
			Method:  http.MethodGet,
			Version: 1, // TODO version your API , it should be date rather than version no
			Path:    "emplymgrmap",
			Func:    emplyMgrMap.GetAll,
		},
	}
}

// GetByID : get supervisor and supervisor of supervisor for an employee
func (emplyMgrMap *EmployeeManagerMap) GetByID(w http.ResponseWriter, r *http.Request) {
	employeeName := chi.URLParam(r, "name")
	employeeMap := cache.GetEmployeeMgrMap()

	type response struct {
		Supervisor             string `json:"supervisor"`
		SupervisorOfsupervisor string `json:"supervisor_of_supervisor"`
	}

	resp := &response{}

	if supervisor, found := employeeMap[employeeName]; found {
		resp.Supervisor = supervisor
		if supervisor, found := employeeMap[employeeName]; found {
			resp.SupervisorOfsupervisor = supervisor
		}
	}

	handlers.WriteJSONResponse(w, r, resp, http.StatusOK, nil)
}

// GetAll : get the entire employee hierarchy
func (emplyMgrMap *EmployeeManagerMap) GetAll(w http.ResponseWriter, r *http.Request) {
	employeeMap := cache.GetEmployeeMgrMap()
	response := apihelpers.CreateRemployeeRelationshipResponseTree(employeeMap)
	handlers.WriteJSONResponse(w, r, response, http.StatusOK, nil)
}

// Create : supports POST/create semantics on EmployeeManagerMap resource
func (emplyMgrMap *EmployeeManagerMap) Create(w http.ResponseWriter, r *http.Request) {
	var reqEmplyMgrMap models.EmployeeManagerMap

	err := json.NewDecoder(r.Body).Decode(&reqEmplyMgrMap)
	if nil != err {
		log.Printf("Error while reading the EmployeeManagerMap request:%s", err.Error())
		handlers.WriteJSONResponse(w, r, nil, http.StatusBadRequest, models.ErrInvalidRequest)
		return
	}

	if err := reqEmplyMgrMap.Valid(); nil != err {
		handlers.WriteJSONResponse(w, r, nil, http.StatusBadRequest, err)
		return
	}

	// update into cache & db
	cache.SetEmployeeMgrMap(reqEmplyMgrMap)
	_, err = emplyMgrMap.repo.Create(r.Context(), reqEmplyMgrMap)

	if nil != err {
		handlers.WriteJSONResponse(w, r, nil, http.StatusInternalServerError, models.ErrDBRecordCreationFailure)
		return
	}

	response := apihelpers.CreateRemployeeRelationshipResponseTree(reqEmplyMgrMap)
	handlers.WriteJSONResponse(w, r, response, http.StatusOK, nil)
}

// UPDATE : supports PUT/update semantics on EmployeeManagerMap resource
func (emplyMgrMap *EmployeeManagerMap) Update(w http.ResponseWriter, r *http.Request) {
	var reqEmplyMgrMap models.EmployeeManagerMap

	err := json.NewDecoder(r.Body).Decode(&reqEmplyMgrMap)
	if nil != err {
		log.Printf("Error while reading the EmployeeManagerMap request:%s", err.Error())
		handlers.WriteJSONResponse(w, r, nil, http.StatusBadRequest, models.ErrInvalidRequest)
		return
	}

	currentEmployeeMgrMap := cache.GetEmployeeMgrMap()

	// create new map by merging current(i.e. db) and request map
	newEmployeeMgrMap := make(models.EmployeeManagerMap)

	for employee, manager := range currentEmployeeMgrMap {
		newEmployeeMgrMap[employee] = manager
	}

	for employee, manager := range reqEmplyMgrMap {
		newEmployeeMgrMap[employee] = manager
	}

	if err := newEmployeeMgrMap.Valid(); nil != err {
		handlers.WriteJSONResponse(w, r, nil, http.StatusBadRequest, err)
		return
	}

	// update into cache & db
	cache.SetEmployeeMgrMap(newEmployeeMgrMap)
	_, err = emplyMgrMap.repo.Create(r.Context(), newEmployeeMgrMap)

	if nil != err {
		handlers.WriteJSONResponse(w, r, nil, http.StatusInternalServerError, models.ErrDBRecordCreationFailure)
		return
	}

	response := apihelpers.CreateRemployeeRelationshipResponseTree(reqEmplyMgrMap)
	handlers.WriteJSONResponse(w, r, response, http.StatusOK, nil)
}
