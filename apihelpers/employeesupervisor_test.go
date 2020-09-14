package apihelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"personio.com/organization-board/models"
)

// TestCreateSuperVisorResponse : tests supervisor response for an emplouyee
func TestCreateSupervisorResponse(t *testing.T) {
	supervisors := []string{"Nick", "Sophie"}
	expectedResponse := &models.EmployeeSupervisorResponse{
		Supervisor:             "Nick",
		SupervisorOfsupervisor: "Sophie",
	}
	assert.Equal(t, expectedResponse, CreateSupervisorResponse(supervisors))
}
