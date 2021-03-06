package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEmployeeManagerMap_Valid : tests valid EmployeeManagerMap
func TestEmployeeManagerMap_Valid(t *testing.T) {
	emplyMgrMap := getTestEmployeeManagerMap()
	assert.Nil(t, emplyMgrMap.Valid())
}

// TestEmployeeManagerMap_MultipleRoots : tests EmployeeManagerMap with multiple roots
func TestEmployeeManagerMap_MultipleRoots(t *testing.T) {
	emplyMgrMap := getTestEmployeeManagerMap()
	(*emplyMgrMap)["John"] = "Johnie"
	assert.True(t, strings.HasPrefix(emplyMgrMap.Valid().Error(), "There are at least two root employees"))
}

// TestEmployeeManagerMap_Loop : tests EmployeeManagerMap with loops
func TestEmployeeManagerMap_Loop(t *testing.T) {
	emplyMgrMap := getTestEmployeeManagerMap()
	(*emplyMgrMap)["Jonas"] = "Barbara"
	assert.True(t, strings.HasPrefix(emplyMgrMap.Valid().Error(), "Adding this relationship results in loop"))
}

// TestEmployeeManagerMap_InValid : tests invalid EmployeeManagerMap
func TestEmployeeManagerMap_InValid(t *testing.T) {
	emplyMgrMap := make(EmployeeManagerMap)
	assert.NotNil(t, emplyMgrMap.Valid())

	emplyMgrMap["person1"] = ""
	assert.NotNil(t, emplyMgrMap.Valid())
}

func getTestEmployeeManagerMap() *EmployeeManagerMap {
	emplyMgrMap := make(EmployeeManagerMap)
	emplyMgrMap["Peter"] = "Nick"
	emplyMgrMap["Barbara"] = "Nick"
	emplyMgrMap["Nick"] = "Sophie"
	emplyMgrMap["Sophie"] = "Jonas"
	return &emplyMgrMap
}

// TestCreateManagerToEmployeeList : tests conevrsion of employe->manager map to manager->employees map
func TestCreateManagerToEmployeeList(t *testing.T) {
	emplyMgrMap := getTestEmployeeManagerMap()

	expectedMgrEmployeeList := map[string][]string{
		"Nick":   []string{"Peter", "Barbara"},
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
