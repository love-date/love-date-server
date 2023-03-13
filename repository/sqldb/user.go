package sqldb

import (
	"database/sql"
	"fmt"
	"love-date/entity"
)

func (d *MySQLDB) CreateUser(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(email) values(?)`, user.Email)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	id, _ := res.LastInsertId()
	user.ID = int(id)

	return user, nil
}

func (d *MySQLDB) DoesThisUserEmailExist(email string) (bool, entity.User, error) {
	user := entity.User{}

	row := d.db.QueryRow(`select * from users where email = ?`, email)
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {

			return false, entity.User{}, nil
		}

		return false, entity.User{}, fmt.Errorf("can't scan query result: %w", err)
	}

	return true, user, nil
}

//func (p Partner) AppendNames(req AppendPartnerNameRequest) (AppendPartnerNameResponse, error) {
//	partnerExist, partner, err := p.repo.DoesUserHaveActivePartner(req.AuthenticatedUserID)
//	if err != nil {
//
//		return AppendPartnerNameResponse{}, fmt.Errorf("unexpected error : %w", err)
//	}
//	if !partnerExist {
//
//		return AppendPartnerNameResponse{}, fmt.Errorf("this user hasn't any active partner")
//	}
//
//	if appendNames, aErr := p.repo.AppendNameWithPartner(partner.ID); aErr != nil {
//
//		return AppendPartnerNameResponse{}, fmt.Errorf("unexpected error : %w", err)
//	} else {
//
//		return AppendPartnerNameResponse{appendNames}, nil
//	}
//}
