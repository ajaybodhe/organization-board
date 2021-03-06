package emplymgrmap

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"personio.com/organization-board/constants"
	"personio.com/organization-board/models"
)

func getTestEmployeeManagerMap() *models.EmployeeManagerMap {
	emplyMgrMap := make(models.EmployeeManagerMap)
	emplyMgrMap["Peter"] = "Nick"
	emplyMgrMap["Barbara"] = "Nick"
	emplyMgrMap["Nick"] = "Sophie"
	emplyMgrMap["Sophie"] = "Jonas"
	return &emplyMgrMap
}

// TestCreate_Success: test for insert mapping success
func TestCreate_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mocking for delete statement
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))

	// mocking for insert statement
	mock.ExpectPrepare(regexp.QuoteMeta(constants.EmployeeManagerMappingInsertQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()
	employeeManagerMap := getTestEmployeeManagerMap()

	_, err = employeeManagerMapRepository.Create(cntx, *employeeManagerMap)

	assert.Nil(t, err)
}

// TestCreate_FailOnDelete: test for insert mapping failure (delete fail)
func TestCreate_FailOnDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mocking for delete statement
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnError(errors.New("MOCK ERROR: In delete"))

	// mocking for insert statement
	mock.ExpectPrepare(regexp.QuoteMeta(constants.EmployeeManagerMappingInsertQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()
	employeeManagerMap := getTestEmployeeManagerMap()

	_, err = employeeManagerMapRepository.Create(cntx, *employeeManagerMap)

	assert.NotNil(t, err)
}

// TestCreate_FailOnInsert: test for insert mapping failure (insert fail)
func TestCreate_FailOnInsert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mocking for delete statement
	mock.ExpectExec(constants.EmployeeManagerMappingDeleteQuery).WillReturnResult(sqlmock.NewResult(0, 0))

	// mocking for insert statement
	mock.ExpectPrepare(regexp.QuoteMeta(constants.EmployeeManagerMappingInsertQuery)).ExpectExec().WillReturnError(errors.New("MOCK ERROR: In delete"))

	employeeManagerMapRepository := NewEmployeeManagerMapRepository(db)
	cntx := context.Background()
	employeeManagerMap := getTestEmployeeManagerMap()

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
		{EmployeeName: "Peter", ManagerName: "Nick"},
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

// TestGetAll_Fail: test for error while fetching data from db
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
