package sqldb

import (
	"database/sql"
	"fmt"
	"love-date/entity"
)

func (d *MySQLDB) CreatePartner(partner entity.Partner) (entity.Partner, error) {
	res, err := d.db.Exec(`insert into partners(user_id,name,birthday,first_date,is_deleted) values(?,?,?,?,?)`,
		partner.UserID, partner.Name, partner.Birthday, partner.FirstDate, false)
	if err != nil {
		return entity.Partner{}, fmt.Errorf("can't executed command %w", err)
	}

	id, _ := res.LastInsertId()
	partner.ID = int(id)

	return partner, nil
}

func (d *MySQLDB) UpdatePartner(partnerID int, partner entity.Partner) (entity.Partner, error) {
	_, err := d.db.Exec(`update  partners set name=? ,birthday=? ,first_date=?,is_deleted=? where id=?`,
		partner.Name, partner.Birthday, partner.FirstDate, partner.GetDeletedStatus(), partnerID)
	if err != nil {
		return entity.Partner{}, fmt.Errorf("can't execute command: %w", err)
	}
	//TODO: what can i do to return partner updated -- in all repo with mysql
	row := d.db.QueryRow(`select * from partners where id =? and is_deleted=?`, partnerID, false)
	var isDeleted bool
	rErr := row.Scan(&partner.ID, &partner.UserID, &partner.Name, &partner.Birthday, &partner.FirstDate, &isDeleted)
	if rErr != nil {
		if rErr == sql.ErrNoRows {

			return entity.Partner{}, nil
		}

		return entity.Partner{}, fmt.Errorf("can't scan query result: %w", rErr)
	}

	return partner, nil
}

func (d *MySQLDB) DoesUserHaveActivePartner(userID int) (bool, entity.Partner, error) {
	partner := entity.Partner{}
	// TODO : ask: is better "isDeleted" check in repository or service?
	row := d.db.QueryRow(`select * from partners where user_id =? and is_deleted=?`, userID, false)
	var isDeleted bool
	err := row.Scan(&partner.ID, &partner.UserID, &partner.Name, &partner.Birthday, &partner.FirstDate, &isDeleted)
	if err != nil {
		if err == sql.ErrNoRows {

			return false, entity.Partner{}, nil
		}

		return false, entity.Partner{}, fmt.Errorf("can't scan query result: %w", err)
	}
	return true, partner, nil
}
