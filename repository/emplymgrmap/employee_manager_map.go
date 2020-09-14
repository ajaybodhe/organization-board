package emplymgrmap

import (
	"bytes"
	"context"
	"database/sql"
	"log"
	"strings"

	"personio.com/organization-board/models"
)

type EmployeeManagerMapRepository struct {
	conn *sql.DB
}

func NewEmployeeManagerMapRepository(conn *sql.DB) *EmployeeManagerMapRepository {
	return &EmployeeManagerMapRepository{conn: conn}
}

func (emplymgr *EmployeeManagerMapRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	return nil, nil
}

func (emplymgr *EmployeeManagerMapRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	reqEmplyMgrMap := obj.(models.EmployeeManagerMap)
	if err := emplymgr.deleteAllEmployeeManager(); nil != err {
		return nil, err
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("INSERT INTO employee_manager_mapping")
	queryBuffer.WriteString("(employee_name, manager_name)")
	queryBuffer.WriteString("VALUES")

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
	query := "DELETE FROM employee_manager_mapping;"

	if _, err := emplymgr.conn.Exec(query); nil != err {
		log.Printf("Error while deleting employee_manager_mapping:%s", err)
		return err
	}

	return nil
}

func (emplymgr *EmployeeManagerMapRepository) Delete(cntx context.Context, id int64) error {
	return nil
}

func (emplymgr *EmployeeManagerMapRepository) GetAll(cntx context.Context) (interface{}, error) {
	query := "SELECT employee_name, manager_name FROM employee_manager_mapping"

	employeeMgrMap := make(models.EmployeeManagerMap)

	row, err := emplymgr.conn.Query(query)
	if nil != err {
		log.Printf("Error while fetching employee_manager_mapping:%s\n", err.Error())
		return employeeMgrMap, nil
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
