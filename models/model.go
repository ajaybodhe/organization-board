package models

import (
	"encoding/json"
)

// prototype of db model
type IModel interface {
	String() string
}

func Stringify(obj interface{}) string {
	byts, _ := json.Marshal(obj)
	return string(byts)
}
