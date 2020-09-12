package models

type EmployeeManagerMap map[string]string

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
