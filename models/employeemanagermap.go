package models

import (
	"bytes"
	"errors"
	"fmt"
)

// EmployeeManagerMap : typedef for employee->manager relationship
type EmployeeManagerMap map[string]string

// CreateManagerToEmployeeList : converts EmployeeManagerMap to ManagerToEmployeeList
func (emplyMgrMap *EmployeeManagerMap) CreateManagerToEmployeeList() map[string][]string {
	mgrEmplyList := make(map[string][]string)
	for empl, mgr := range *emplyMgrMap {
		mgrEmplyList[mgr] = append(mgrEmplyList[mgr], empl)
	}
	return mgrEmplyList
}

func (emplyMgrMap *EmployeeManagerMap) GetRootEmployee(mgrEmplyList map[string][]string) string {
	for manager := range mgrEmplyList {
		if (*emplyMgrMap)[manager] == "" {
			return manager
		}
	}
	return ""
}

// dfs over Manager to Employee list. Returns error if loop exists.
func dfsToDetectLoop(mgrEmplyList map[string][]string, visited map[string]bool, manager string) error {
	visited[manager] = true
	for _, employee := range mgrEmplyList[manager] {
		if visited[employee] == true {
			return errors.New("Adding this relationship results in loop : " + employee + "->" + manager)
		}
		if err := dfsToDetectLoop(mgrEmplyList, visited, employee); nil != err {
			return err
		}
	}
	return nil
}

// Returns error if there is loop in employee hierarchy
func (emplyMgrMap *EmployeeManagerMap) detectLoopInHierarchy() error {
	visited := make(map[string]bool)
	mgrEmplyList := emplyMgrMap.CreateManagerToEmployeeList()
	rootEmployee := emplyMgrMap.GetRootEmployee(mgrEmplyList)
	return dfsToDetectLoop(mgrEmplyList, visited, rootEmployee)
}

// Returns error if there are multiple roots in employee hierarchy
func (emplyMgrMap *EmployeeManagerMap) detectMulipleRootsInHierarchy() error {
	visited := make(map[string]bool)
	rootEmployee := ""
	for employee, manager := range *emplyMgrMap {
		if visited[employee] == false {
			visited[employee] = true
			newManager := manager
			for visited[newManager] != true {
				visited[newManager] = true
				if (*emplyMgrMap)[newManager] == "" {
					if rootEmployee == "" {
						rootEmployee = newManager
						break
					}
					return errors.New("There are at least two root employees : " + rootEmployee + " & " + newManager)
				}
				newManager = (*emplyMgrMap)[newManager]
			}
		}
	}
	return nil
}

// Valid : checks EmployeeManagerMap if struct is valid
func (emplyMgrMap *EmployeeManagerMap) Valid() error {
	var buffer bytes.Buffer

	if 0 == len(*emplyMgrMap) {
		buffer.WriteString("Empty Mapping Object")
	}

	for key, val := range *emplyMgrMap {
		if "" == key || "" == val {
			buffer.WriteString(fmt.Sprintf("Invalid key/value pair %s:%s", key, val))
		}
	}

	if buffer.Len() > 0 {
		return fmt.Errorf("%s", buffer.String())
	}

	if err := emplyMgrMap.detectLoopInHierarchy(); nil != err {
		return err
	}

	if err := emplyMgrMap.detectMulipleRootsInHierarchy(); nil != err {
		return err
	}

	return nil
}
