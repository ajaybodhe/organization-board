package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEmployeeManagerMap_Valid : tests valid EmployeeManagerMap
func TestEmployeeManagerMap_Valid(t *testing.T) {
	emm := make(EmployeeManagerMap)
	emm["person1"] = "person2"
	assert.Nil(t, emm.Valid())
}

// TestEmployeeManagerMap_InValid : tests invalid EmployeeManagerMap
func TestEmployeeManagerMap_InValid(t *testing.T) {
	emm := make(EmployeeManagerMap)
	assert.NotNil(t, emm.Valid())

	emm["person1"] = ""
	assert.NotNil(t, emm.Valid())
}

func getTestEmployeeManagerMap() *EmployeeManagerMap {
	emplyMgrMap := make(EmployeeManagerMap)
	emplyMgrMap["Pete"] = "Nick"
	emplyMgrMap["Barbara"] = "Nick"
	emplyMgrMap["Nick"] = "Sophie"
	emplyMgrMap["Sophie"] = "Jonas"
	return &emplyMgrMap
}

// TestCreateManagerToEmployeeList : tests conevrsion of employe->manager map to manager->employees map
func TestCreateManagerToEmployeeList(t *testing.T) {
	emplyMgrMap := getTestEmployeeManagerMap()

	expectedMgrEmployeeList := map[string][]string{
		"Nick":   []string{"Pete", "Barbara"},
		"Jonas":  []string{"Sophie"},
		"Sophie": []string{"Nick"},
	}
	assert.Equal(t, expectedMgrEmployeeList, emplyMgrMap.CreateManagerToEmployeeList())
}

// TestGetRootEmployee : tests for retrieval of employee at root
func TestGetRootEmployee(t *testing.T) {
	emplyMgrMap := getTestEmployeeManagerMap()
	mgrEmplyList := emplyMgrMap.CreateManagerToEmployeeList()
	assert.Equal(t, "Jonas", emplyMgrMap.GetRootEmployee(mgrEmplyList))
}
