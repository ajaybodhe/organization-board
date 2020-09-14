package constants

import (
	"github.com/go-chi/jwtauth"
)

// JWT constants
const (
	authToken           = "#1fsyHGFY8842sfg%$"
	AuthorizationHeader = "Authorization"
	MapClaimUser        = "user"
)

var (
	AuthToken = jwtauth.New("HS256", []byte(authToken), nil)
)

// SQL Queries
const (
	EmployeeManagerMappingSelectQuery = "SELECT employee_name, manager_name FROM employee_manager_mapping"
	EmployeeManagerMappingInsertQuery = `INSERT INTO employee_manager_mapping
	(employee_name, manager_name)
		VALUES`
	EmployeeManagerMappingDeleteQuery = "DELETE FROM employee_manager_mapping"
	LoginDetailsSelectQuery           = `SELECT id, email
	FROM user_detail
	WHERE email = ?
	AND password = ?
	AND deleted = 0`
)
