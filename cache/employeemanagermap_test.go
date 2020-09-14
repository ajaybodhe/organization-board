package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"personio.com/organization-board/constants"
	"personio.com/organization-board/models"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestCache_GetEmployeeMgrMap : tests initialization of cache and asserts cached value of employee-manager map
func TestCache_GetEmployeeMgrMap(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"employee_name", "manager_name"}).AddRow("Nick", "Sophie").AddRow("Sophie", "Jonas")
	mock.ExpectQuery(constants.EmployeeManagerMappingSelectQuery).WillReturnRows(rows)

	Init(db)
	assert.Equal(t, models.EmployeeManagerMap{"Nick": "Sophie", "Sophie": "Jonas"}, GetEmployeeMgrMap())
}
