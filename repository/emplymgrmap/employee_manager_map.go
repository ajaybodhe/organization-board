package emplymgrmap

import (
	"context"
	"database/sql"
)

type EmployeeManagerMapRepository struct {
	conn *sql.DB
}

func NewEmployeeManagerMapRepository(conn *sql.DB) *EmployeeManagerMapRepository {
	return &EmployeeManagerMapRepository{conn: conn}
}

func (emplymgr *EmployeeManagerMapRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	return nil, nil
}

func (emplymgr *EmployeeManagerMapRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	return nil, nil
}

func (emplymgr *EmployeeManagerMapRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	return nil, nil
}

func (emplymgr *EmployeeManagerMapRepository) Delete(cntx context.Context, id int64) error {

	return nil
}

func (emplymgr *EmployeeManagerMapRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	return nil, nil
}
