package models

// EmployeeManagerMap : typedef for employee->manager relationship
type EmployeeManagerMap map[string]string

// Valid : checks EmployeeManagerMap if struct is valid
func (empMgrMap *EmployeeManagerMap) Valid() bool {
	if 0 == len(*empMgrMap) {
		return false
	}

	for key, val := range *empMgrMap {
		if "" == key || "" == val {
			return false
		}
	}

	return true
}
