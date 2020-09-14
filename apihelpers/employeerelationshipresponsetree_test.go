package apihelpers

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"personio.com/organization-board/models"
)

func getTestEmployeeManagerMap() models.EmployeeManagerMap {
	emplyMgrMap := make(models.EmployeeManagerMap)
	emplyMgrMap["Pete"] = "Nick"
	emplyMgrMap["Barbara"] = "Nick"
	emplyMgrMap["Nick"] = "Sophie"
	emplyMgrMap["Sophie"] = "Jonas"
	return emplyMgrMap
}

// TestCreateRemployeeRelationshipResponseTree_EmptyList : tests response tree for emty map
func TestCreateRemployeeRelationshipResponseTree_EmptyList(t *testing.T) {
	emplyMgrMap := make(models.EmployeeManagerMap)
	responseTree := CreateEmployeeRelationshipResponseTree(&emplyMgrMap)
	response, err := json.Marshal(responseTree)
	assert.Nil(t, err)
	assert.Equal(t, string(response), "{\"\":{}}")
}

// TestCreateRemployeeRelationshipResponseTree_Valid : tests response tree for valid map
func TestCreateRemployeeRelationshipResponseTree_Valid(t *testing.T) {
	emplyMgrMap := getTestEmployeeManagerMap()
	responseTree := CreateEmployeeRelationshipResponseTree(&emplyMgrMap)
	response, err := json.Marshal(responseTree)
	assert.Nil(t, err)
	assert.Equal(t, string(response), "{\"Jonas\":{\"Sophie\":{\"Nick\":{\"Barbara\":{},\"Pete\":{}}}}}")
}
