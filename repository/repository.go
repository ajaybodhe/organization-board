package repository

import (
	"context"
)

// IRepository : interface for DB CRUD Operations
type IRepository interface {
	GetByID(context.Context, int64) (interface{}, error)
	Create(context.Context, interface{}) (interface{}, error)
	Update(context.Context, interface{}) (interface{}, error)
	Delete(context.Context, int64) error
	GetAll(context.Context) ([]interface{}, error)
}

// Repository : Base class for all CRUD objects, implements IRepository
type Repository struct {
}

// GetByID : Read operation on a DB resource
func (repo *Repository) GetByID(cntx context.Context, id int64) (obj interface{}, err error) {
	return
}

// Create : Create operation on a DB resource
func (repo *Repository) Create(cntx context.Context, obj interface{}) (cobj interface{}, err error) {
	return
}

// Update : Update operation on a DB resource
func (repo *Repository) Update(cntx context.Context, obj interface{}) (uobj interface{}, err error) {
	return
}

// Delete : Delete operation on a DB resource
func (repo *Repository) Delete(cntx context.Context, id int64) (deleted bool, err error) {
	return
}

// GetAll : Get all resources from a collection
func (repo *Repository) GetAll(cntx context.Context) (obj interface{}, err error) {
	return
}
