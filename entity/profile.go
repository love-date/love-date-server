package entity

type Profile struct {
	ID                      int
	UserID                  int
	Name                    string
	BirthdayNotifyActive    bool
	SpecialDaysNotifyActive bool
	vipActive               bool
}

func (p *Profile) UserAsVIP(status bool) {
	p.vipActive = status
}

func (p *Profile) IsUserVIP() bool {
	return p.vipActive
}
