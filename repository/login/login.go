package login

import (
	"bytes"
	"context"
	"database/sql"

	"personio.com/organization-board/models"
	"personio.com/organization-board/repository"
)

type LoginRepository struct {
	repository.Repository
	conn *sql.DB
}

func NewLoginRepository(conn *sql.DB) *LoginRepository {
	return &LoginRepository{conn: conn}
}

func (login *LoginRepository) Authenticate(ctx context.Context, obj *models.Login) (user *models.User, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("SELECT id, email")
	buffer.WriteString(" FROM user_detail")
	buffer.WriteString(" WHERE email = ?")
	buffer.WriteString(" AND password = ?")
	buffer.WriteString(" AND deleted = 0")

	rows, err := login.conn.QueryContext(ctx, buffer.String(), obj.Email, obj.Password)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		user = new(models.User)
		err = rows.Scan(
			&user.Id,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
	} else {
		err = models.DBErrNotFound
	}

	return user, err
}
