package apihelpers

import (
	"personio.com/organization-board/models"
)

func convertEmployeeHierarchyToTree(rootEmployee string, mgrEmplyList map[string][]string, response map[string]interface{}) {
	for _, employee := range mgrEmplyList[rootEmployee] {
		if response[employee] == nil {
			nextResponse := make(map[string]interface{})
			response[employee] = nextResponse
			convertEmployeeHierarchyToTree(employee, mgrEmplyList, nextResponse)
		}
	}
}

// CreateEmployeeRelationshipResponseTree : creates employee relationship hierarchy tree
// TODO : we can expose this functionality as an interface.
// So that response format can be changed as needed.
// In that case, return value should just be interface{}.
func CreateEmployeeRelationshipResponseTree(emplyMgrMap *models.EmployeeManagerMap) map[string]interface{} {
	mgrEmplyList := emplyMgrMap.CreateManagerToEmployeeList()
	rootEmployee := emplyMgrMap.GetRootEmployee(mgrEmplyList)
	response := make(map[string]interface{})
	nextResponse := make(map[string]interface{})
	response[rootEmployee] = nextResponse
	convertEmployeeHierarchyToTree(rootEmployee, mgrEmplyList, nextResponse)
	return response
}
