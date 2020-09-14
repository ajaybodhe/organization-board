package models

import (
	"bytes"
	"errors"
	"fmt"
)

// EmployeeManagerMap : typedef for employee->manager relationship
type EmployeeManagerMap map[string]string

// Returns error if there is loop
func (emplyMgrMap *EmployeeManagerMap) detectLoopInHierarchy() error {
	visited := make(map[string]bool)
	for employee, manager := range *emplyMgrMap {
		if visited[employee] && visited[manager] {
			fmt.Println("Adding this relationship results in loop : ", employee, manager)
			return errors.New("Adding this relationship results in loop : " + employee + "->" + manager)
		}
		visited[employee] = true
		visited[manager] = true
	}
	return nil
}

// Returns error if there is multiple roots are there
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
					fmt.Println("There are at least two root employees : ", rootEmployee, newManager)
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

	if err := emplyMgrMap.detectMulipleRootsInHierarchy(); nil != err {
		return err
	}

	if err := emplyMgrMap.detectLoopInHierarchy(); nil != err {
		return err
	}

	return nil
}
