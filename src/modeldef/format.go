package modeldef

type Format struct {
	Id     int
	Format string
	Seed   []*Seed                 `orm:"reverse(many)"`
	User   []*UserMainSiteSettings `orm:"reverse(many)"`
}

func (f *Format) GetString() string {
	return f.Format
}

