package models

// EmployeeSupervisorResponse : answer to who is my supervisor and super-supervisor
type EmployeeSupervisorResponse struct {
	Supervisor             string `json:"supervisor"`
	SupervisorOfsupervisor string `json:"supervisor_of_supervisor"`
}
