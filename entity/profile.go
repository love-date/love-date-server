package entity

type Profile struct {
	ID                      int    `json:"id"`
	UserID                  int    `json:"user_id"`
	Name                    string `json:"name"`
	BirthdayNotifyActive    bool   `json:"birthday_notify_active"`
	SpecialDaysNotifyActive bool   `json:"special_days_notify_active"`
	vipActive               bool
}

func (p *Profile) UserAsVIP(status bool) {
	p.vipActive = status
}

func (p *Profile) IsUserVIP() bool {
	return p.vipActive
}
