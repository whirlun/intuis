package modeldef

type Medium struct {
	Id     int
	Medium string
	Seed   []*Seed                 `orm:"reverse(many)"`
	User   []*UserMainSiteSettings `orm:"reverse(many)"`
}

func (m *Medium) GetString() string {
	return m.Medium
}

