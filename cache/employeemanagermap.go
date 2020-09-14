package cache

import (
	"database/sql"
	"log"

	"personio.com/organization-board/models"
	"personio.com/organization-board/repository/emplymgrmap"
)

// in process caching for employee manager map
var (
	employeeMgrMap models.EmployeeManagerMap
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
	return employeeMgrMap
}

// SetEmployeeMgrMap : write EmployeeManagerMap to cache
func SetEmployeeMgrMap(updatedEmployeeMgrMap models.EmployeeManagerMap) {
	employeeMgrMap = updatedEmployeeMgrMap
}
