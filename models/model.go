package models

import (
	"encoding/json"
)

// IModel : prototype of db model
type IModel interface {
	String() string
}

// Stringify : string version of an obj
func Stringify(obj interface{}) string {
	byts, _ := json.Marshal(obj)
	return string(byts)
}
