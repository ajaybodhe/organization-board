package emplymgrmap

import (
	"bytes"
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"personio.com/organization-board/constants"
	"personio.com/organization-board/models"
)

// TestCache_Authenticate_Success: test for insert mapping success
func TestCreate_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mocking for delete statement
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))

	// mocking for insert statement
	var buffer bytes.Buffer
	buffer.WriteString(constants.EmployeeManagerMappingInsertQuery)
	mock.ExpectPrepare(regexp.QuoteMeta(buffer.String())).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()
	employeeManagerMap := &models.EmployeeManagerMap{}
	(*employeeManagerMap)["Pete"] = "Nick"
	(*employeeManagerMap)["Barbara"] = "Nick"
	(*employeeManagerMap)["Nick"] = "Sophie"
	(*employeeManagerMap)["Sophie"] = "Jonas"

	_, err = employeeManagerMapRepository.Create(cntx, *employeeManagerMap)

	assert.Nil(t, err)
}

// TestCreate_Fail: test for insert mapping failure (delete fail)
func TestCreate_Fail_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mocking for delete statement
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnError(errors.New("MOCK ERROR: In delete"))

	// mocking for insert statement
	var buffer bytes.Buffer
	buffer.WriteString(constants.EmployeeManagerMappingInsertQuery)
	mock.ExpectPrepare(regexp.QuoteMeta(buffer.String())).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()
	employeeManagerMap := &models.EmployeeManagerMap{}
	(*employeeManagerMap)["Pete"] = "Nick"
	(*employeeManagerMap)["Barbara"] = "Nick"
	(*employeeManagerMap)["Nick"] = "Sophie"
	(*employeeManagerMap)["Sophie"] = "Jonas"

	_, err = employeeManagerMapRepository.Create(cntx, *employeeManagerMap)

	assert.NotNil(t, err)
}

// TestCreate_Fail_Insert: test for insert mapping failure (insert fail)
func TestCreate_Fail_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mocking for delete statement
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))

	// mocking for insert statement
	var buffer bytes.Buffer
	buffer.WriteString(constants.EmployeeManagerMappingInsertQuery)
	mock.ExpectPrepare(regexp.QuoteMeta(buffer.String())).ExpectExec().WillReturnError(errors.New("MOCK ERROR: In delete"))

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()
	employeeManagerMap := &models.EmployeeManagerMap{}
	(*employeeManagerMap)["Pete"] = "Nick"
	(*employeeManagerMap)["Barbara"] = "Nick"
	(*employeeManagerMap)["Nick"] = "Sophie"
	(*employeeManagerMap)["Sophie"] = "Jonas"

	_, err = employeeManagerMapRepository.Create(cntx, *employeeManagerMap)

	assert.NotNil(t, err)
}

// TestGetAll_Sucess: test for get all rows from employee manager map
func TestGetAll_Sucess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mocking for get all statement
	type emplyMgrMap struct {
		EmployeeName string
		ManagerName  string
	}
	emplyMgrMapList := []emplyMgrMap{
		{EmployeeName: "Pete", ManagerName: "Nick"},
		{EmployeeName: "Barbara", ManagerName: "Nick"},
	}

	rows := sqlmock.NewRows([]string{"employee_name", "manager_name"}).AddRow(
		emplyMgrMapList[0].EmployeeName,
		emplyMgrMapList[0].ManagerName).AddRow(
		emplyMgrMapList[1].EmployeeName,
		emplyMgrMapList[1].ManagerName)
	mock.ExpectQuery(constants.EmployeeManagerMappingSelectQuery).WillReturnRows(rows)

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()

	obj, err := employeeManagerMapRepository.GetAll(cntx)
	assert.Nil(t, err)
	dbEmplyMgrMapList := obj.(models.EmployeeManagerMap)
	// check number of element in map, it must be 2
	assert.Equal(t, len(dbEmplyMgrMapList), 2)
}

// TestGetAll_Sucess: test for error while fetching data from db
func TestGetAll_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(constants.EmployeeManagerMappingSelectQuery).WillReturnError(errors.New("MOCK ERROR: In selct"))

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()

	_, err = employeeManagerMapRepository.GetAll(cntx)
	assert.NotNil(t, err)

}
