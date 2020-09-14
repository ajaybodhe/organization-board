package cache

import (
	"database/sql"
	"log"

	"personio.com/organization-board/models"
	"personio.com/organization-board/repository/emplymgrmap"
)

var (
	employeeMgrMap models.EmployeeManagerMap
)

func Init(conn *sql.DB) {
	dbEmployeeMgrMap, err := emplymgrmap.NewEmployeeManagerMapRepository(conn).GetAll(nil)
	if nil != err {
		log.Fatalf("Error while fetching database employee manager releationship:%s", err.Error())
	}
	employeeMgrMap = dbEmployeeMgrMap.(models.EmployeeManagerMap)
}

func GetEmployeeMgrMap() models.EmployeeManagerMap {
	return employeeMgrMap
}

func SetEmployeeMgrMap(updatedEmployeeMgrMap models.EmployeeManagerMap) {
	employeeMgrMap = updatedEmployeeMgrMap
}
