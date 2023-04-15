package sqldb

import (
	"database/sql"
	"love-date/entity"
	"love-date/pkg/errhandling/richerror"
)

func (d *MySQLDB) CreatePartner(partner entity.Partner) (entity.Partner, error) {
	const op = "sqldb.CreatePartner"

	res, err := d.db.Exec(`insert into partners(user_id,name,birthday,first_date,is_deleted) values(?,?,?,?,?)`,
		partner.UserID, partner.Name, partner.Birthday, partner.FirstDate, false)
	if err != nil {

		return entity.Partner{}, richerror.New(op).WithWrapError(err).
			WithMessage(err.Error()).WithKind(richerror.KindUnexpected)
	}

	id, _ := res.LastInsertId()
	partner.ID = int(id)

	return partner, nil
}

func (d *MySQLDB) UpdatePartner(partnerID int, partner entity.Partner) (entity.Partner, error) {
	const op = "sqldb.UpdatePartner"

	_, err := d.db.Exec(`update  partners set name=? ,birthday=? ,first_date=?,is_deleted=? where id=?`,
		partner.Name, partner.Birthday, partner.FirstDate, partner.GetDeletedStatus(), partnerID)
	if err != nil {

		return entity.Partner{}, richerror.New(op).WithWrapError(err).
			WithMessage(err.Error()).WithKind(richerror.KindUnexpected)
	}

	row := d.db.QueryRow(`select * from partners where id =? and is_deleted=?`, partnerID, false)

	var isDeleted bool

	rErr := row.Scan(&partner.ID, &partner.UserID, &partner.Name, &partner.Birthday, &partner.FirstDate, &isDeleted)
	if rErr != nil {
		if rErr == sql.ErrNoRows {
			return entity.Partner{}, nil
		}

		return entity.Partner{}, richerror.New(op).WithWrapError(err).
			WithMessage(err.Error()).WithKind(richerror.KindUnexpected)
	}

	return partner, nil
}

func (d *MySQLDB) DoesUserHaveActivePartner(userID int) (bool, entity.Partner, error) {
	const op = "sqldb.DoesUserHaveActivePartner"

	partner := entity.Partner{}

	row := d.db.QueryRow(`select * from partners where user_id =? and is_deleted=?`, userID, false)

	var isDeleted bool

	err := row.Scan(&partner.ID, &partner.UserID, &partner.Name, &partner.Birthday, &partner.FirstDate, &isDeleted)
	if err != nil {
		if err == sql.ErrNoRows {

			return false, entity.Partner{}, nil
		}

		return false, entity.Partner{}, richerror.New(op).WithWrapError(err).
			WithMessage(err.Error()).WithKind(richerror.KindUnexpected)
	}

	return true, partner, nil
}
