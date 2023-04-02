package sqldb

import (
	"database/sql"
	"fmt"
	"love-date/entity"
)

func (d *MySQLDB) DoesThisUserProfileExist(userID int) (bool, entity.Profile, error) {
	profile := entity.Profile{}
	var vipActive bool
	row := d.db.QueryRow(`select * from profiles where user_id = ?`, userID)
	err := row.Scan(&profile.ID, &profile.UserID, &profile.Name, &profile.SpecialDaysNotifyActive, &profile.BirthdayNotifyActive, &vipActive)
	profile.UserAsVIP(vipActive)
	if err != nil {
		if err == sql.ErrNoRows {

			return false, entity.Profile{}, nil
		}

		return false, entity.Profile{}, fmt.Errorf("can't scan query result: %w", err)
	}
	return true, profile, nil
}

func (d *MySQLDB) CreateProfile(profile entity.Profile) (entity.Profile, error) {
	res, err := d.db.Exec(`insert into profiles(user_id,name,special_days_notify_active,birthday_notify_active,vip_active)
values(?,?,?,?,?)`, profile.UserID, profile.Name, profile.SpecialDaysNotifyActive, profile.BirthdayNotifyActive, false)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("can't execute command: %w", err)
	}

	id, _ := res.LastInsertId()
	profile.ID = int(id)

	return profile, nil
}

func (d *MySQLDB) Update(profileID int, profile entity.Profile) (entity.Profile, error) {
	_, err := d.db.Exec(`update  profiles set name=? ,special_days_notify_active=? ,birthday_notify_active=?  where id=?`,
		profile.Name, profile.SpecialDaysNotifyActive, profile.BirthdayNotifyActive, profileID)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("can't execute command: %w", err)
	}

	var vipActive bool
	row := d.db.QueryRow(`select * from profiles where id = ?`, profileID)
	rErr := row.Scan(&profile.ID, &profile.UserID, &profile.Name, &profile.SpecialDaysNotifyActive, &profile.BirthdayNotifyActive, &vipActive)
	if rErr != nil {
		if rErr == sql.ErrNoRows {

			return entity.Profile{}, nil
		}

		return entity.Profile{}, fmt.Errorf("can't scan query result: %w", rErr)
	}

	profile.UserAsVIP(vipActive)

	return profile, nil
}
