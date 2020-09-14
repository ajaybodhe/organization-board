package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEmployeeManagerMap_Valid : tests valid EmployeeManagerMap
func TestEmployeeManagerMap_Valid(t *testing.T) {
	emm := make(EmployeeManagerMap)
	emm["person1"] = "person2"
	assert.True(t, emm.Valid())
}

// TestEmployeeManagerMap_InValid : tests invalid EmployeeManagerMap
func TestEmployeeManagerMap_InValid(t *testing.T) {
	emm := make(EmployeeManagerMap)
	assert.False(t, emm.Valid())

	emm["person1"] = ""
	assert.False(t, emm.Valid())
}
