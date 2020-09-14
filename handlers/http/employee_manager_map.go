package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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
			Method:  http.MethodPost,
			Version: 1, // TODO version your API , it should be date rather than version no
			Path:    "emplymgrmap",
			Func:    emplyMgrMap.Create,
		},
	}
}

func updateEmployeeManagerMap(reqEmplyMgrMap models.EmployeeManagerMap) (models.EmployeeManagerMap, error) {
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
		return nil, err
	}

	return newEmployeeMgrMap, nil
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

	updatedEmployeeManagerMap, err := updateEmployeeManagerMap(reqEmplyMgrMap)

	if nil != err {
		handlers.WriteJSONResponse(w, r, nil, http.StatusInternalServerError, err)
		return
	}

	// update into cache & db
	cache.SetEmployeeMgrMap(updatedEmployeeManagerMap)
	_, err = emplyMgrMap.repo.Create(r.Context(), updatedEmployeeManagerMap)

	if nil != err {
		handlers.WriteJSONResponse(w, r, nil, http.StatusInternalServerError, models.ErrDBRecordCreationFailure)
		return
	}

	handlers.WriteJSONResponse(w, r, reqEmplyMgrMap, http.StatusOK, nil)
}

// UPDATE : supports PUT/update semantics on EmployeeManagerMap resource
func (emplyMgrMap *EmployeeManagerMap) Update(w http.ResponseWriter, r *http.Request) {
	emplyMgrMap.Create(w, r)
}
