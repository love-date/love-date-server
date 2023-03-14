package entity

import "time"

type Partner struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Birthday  time.Time `json:"birthday"`
	FirstDate time.Time `json:"first_date"`
	isDeleted bool
}

//func (p *Partner) GetDeletedStatus() bool {
//	return p.isDeleted
//}

func (p *Partner) Delete() {
	p.isDeleted = true
}
