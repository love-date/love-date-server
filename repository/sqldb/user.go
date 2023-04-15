package sqldb

import (
	"database/sql"
	"love-date/entity"
	"love-date/pkg/errhandling/richerror"
)

func (d *MySQLDB) CreateUser(user entity.User) (entity.User, error) {
	const op = "sqldb.CreateUser"

	res, err := d.db.Exec(`insert into users(email) values(?)`, user.Email)
	if err != nil {
		return entity.User{}, richerror.New(op).WithWrapError(err).
			WithMessage(err.Error()).WithKind(richerror.KindUnexpected)
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)

	return user, nil
}

func (d *MySQLDB) DoesThisUserEmailExist(email string) (bool, entity.User, error) {
	const op = "sqldb.DoesThisUserEmailExist"

	user := entity.User{}

	row := d.db.QueryRow(`select * from users where email = ?`, email)
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {

			return false, entity.User{}, nil
		}

		return false, entity.User{}, richerror.New(op).WithWrapError(err).
			WithMessage(err.Error()).WithKind(richerror.KindUnexpected)
	}

	return true, user, nil
}
