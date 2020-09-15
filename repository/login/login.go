package login

import (
	"bytes"
	"context"
	"database/sql"

	"personio.com/organization-board/constants"
	"personio.com/organization-board/models"
	"personio.com/organization-board/repository"
)

// LoginRepository : deals with DB(CRUD) operations for Login Data
type LoginRepository struct {
	repository.Repository
	conn *sql.DB
}

// NewLoginRepository : constructor for LoginRepository
func NewLoginRepository(conn *sql.DB) *LoginRepository {
	return &LoginRepository{conn: conn}
}

// Authenticate : Returns authentication information for give Login Details
func (login *LoginRepository) Authenticate(ctx context.Context, obj *models.Login) (user *models.User, err error) {
	var buffer bytes.Buffer
	buffer.WriteString(constants.LoginDetailsSelectQuery)
	rows, err := login.conn.QueryContext(ctx, buffer.String(), obj.Email, obj.Password)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, models.ErrDBRecordNotFound
	}

	user = new(models.User)
	err = rows.Scan(
		&user.ID,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
