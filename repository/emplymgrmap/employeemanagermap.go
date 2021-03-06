package emplymgrmap

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"

	"personio.com/organization-board/constants"
	"personio.com/organization-board/models"
)

// EmployeeManagerMapRepository : deals with DB(CRUD) operations for EmployeeManagerMap
type EmployeeManagerMapRepository struct {
	conn *sql.DB
}

// NewEmployeeManagerMapRepository : constructor for EmployeeManagerMapRepository
func NewEmployeeManagerMapRepository(conn *sql.DB) *EmployeeManagerMapRepository {
	return &EmployeeManagerMapRepository{conn: conn}
}

// Create : add records of EmployeeManagerMap into DB
func (emplymgr *EmployeeManagerMapRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	reqEmplyMgrMap, ok := obj.(models.EmployeeManagerMap)
	if !ok {
		return nil, errors.New("Object is not of type EmployeeManagerMap")
	}

	if err := emplymgr.deleteAllEmployeeManager(); nil != err {
		return nil, err
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(constants.EmployeeManagerMappingInsertQuery)

	var params []interface{}
	var placeholders []string
	for employee, manager := range reqEmplyMgrMap {
		params = append(params, employee)
		params = append(params, manager)
		placeholders = append(placeholders, "(?, ?)")
	}

	queryBuffer.WriteString(strings.Join(placeholders, ","))

	query := queryBuffer.String()
	stmt, err := emplymgr.conn.Prepare(query)
	if nil != err {
		log.Printf("Insert Syntax Error: %s\n\tError Query: %s\n",
			err.Error(), query)
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(params...)
	if nil != err {
		log.Printf("Insert Execute Error: %s\nError Query: %s\n",
			err.Error(), query)
		return nil, err
	}

	return nil, nil
}

func (emplymgr *EmployeeManagerMapRepository) deleteAllEmployeeManager() error {
	if _, err := emplymgr.conn.Exec(constants.EmployeeManagerMappingDeleteQuery); nil != err {
		log.Printf("Error while deleting employee_manager_mapping:%s", err)
		return err
	}

	return nil
}

// GetAll : real all records for EmployeeManagerMap from DB
func (emplymgr *EmployeeManagerMapRepository) GetAll(cntx context.Context) (interface{}, error) {
	employeeMgrMap := make(models.EmployeeManagerMap)

	row, err := emplymgr.conn.Query(constants.EmployeeManagerMappingSelectQuery)
	if nil != err {
		log.Printf("Error while fetching employee_manager_mapping:%s\n", err.Error())
		return employeeMgrMap, err
	}

	defer row.Close()
	for row.Next() {
		var employeeName string
		var mgrName string
		err = row.Scan(&employeeName, &mgrName)
		if nil != err {
			log.Printf("Error in employee_manager_mapping row scan: %s\n", err.Error())
			continue
		}

		employeeMgrMap[employeeName] = mgrName
	}

	return employeeMgrMap, nil
}
