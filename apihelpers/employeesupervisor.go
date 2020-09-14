package apihelpers

import (
	"personio.com/organization-board/models"
)

//GetSupervisor: return the supervisor and suprvisor of supervisor for employeeName
func GetSupervisor(employeeName string, employeeMap models.EmployeeManagerMap) []string {
	var supervisors []string
	if supervisor, found := employeeMap[employeeName]; found {
		supervisors = append(supervisors, supervisor)
		if supervisor, found := employeeMap[supervisor]; found {
			supervisors = append(supervisors, supervisor)
		}
	}

	return supervisors
}

func CreateSuperVisorResponse(supervisors []string) interface{} {
	type response struct {
		Supervisor             string `json:"supervisor"`
		SupervisorOfsupervisor string `json:"supervisor_of_supervisor"`
	}

	resp := &response{}
	for idx, name := range supervisors {
		switch idx {
		case 0:
			resp.Supervisor = name
		case 1:
			resp.SupervisorOfsupervisor = name
		}
	}

	return resp
}
