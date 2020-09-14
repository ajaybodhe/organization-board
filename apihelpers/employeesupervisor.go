package apihelpers

// CreateSuperVisorResponse : converts list of supervisors for an employee into a json response
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
