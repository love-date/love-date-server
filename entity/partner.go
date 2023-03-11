package entity

import "time"

type Partnert struct {
	ID        int
	UserID    int
	Name      string
	Birthday  time.Time
	FirstDate time.Time
}
