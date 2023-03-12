package entity

import "time"

type Partner struct {
	ID        int
	UserID    int
	Name      string
	Birthday  time.Time
	FirstDate time.Time
	isDeleted bool
}

//func (p *Partner) GetDeletedStatus() bool {
//	return p.isDeleted
//}

func (p *Partner) Delete() {
	p.isDeleted = true
}
