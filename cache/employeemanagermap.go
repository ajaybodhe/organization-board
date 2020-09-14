package cache

import (
	"database/sql"
	"log"
	"sync"

	"personio.com/organization-board/models"
	"personio.com/organization-board/repository/emplymgrmap"
)

// in process caching for employee manager map
var (
	employeeMgrMap models.EmployeeManagerMap
	mux            sync.Mutex
)

// Init : init cache values from DB on startup
func Init(conn *sql.DB) {
	dbEmployeeMgrMap, err := emplymgrmap.NewEmployeeManagerMapRepository(conn).GetAll(nil)
	if nil != err {
		log.Fatalf("Error while fetching database employee manager releationship:%s", err.Error())
		return
	}
	employeeMgrMap = dbEmployeeMgrMap.(models.EmployeeManagerMap)
}

// GetEmployeeMgrMap : read EmployeeManagerMap from cache
func GetEmployeeMgrMap() models.EmployeeManagerMap {
	mux.Lock()
	newEmployeeMgrMap := make(models.EmployeeManagerMap)
	for employee, manager := range employeeMgrMap {
		newEmployeeMgrMap[employee] = manager
	}
	mux.Unlock()
	return newEmployeeMgrMap
}

// SetEmployeeMgrMap : write EmployeeManagerMap to cache
func SetEmployeeMgrMap(updatedEmployeeMgrMap models.EmployeeManagerMap) {
	mux.Lock()
	employeeMgrMap = updatedEmployeeMgrMap
	mux.Unlock()
}
