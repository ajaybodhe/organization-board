package models

import (
	"errors"
)

var (
	ErrInvalidRequest          = errors.New("Error:: Invalid Request")
	ErrUnauthorizedAccess      = errors.New("Error:: User is unauthorized to perform the operation.")
	ErrDBRecordNotFound        = errors.New("Error:: Record does not exist")
	ErrDBRecordCreationFailure = errors.New("Error:: Record can not be added to DB")
)
