package modeldef

type ReferRule struct {
	Id        int
	ReferRule string
	Seed      []*Seed `orm:"reverse(many)"`
}

func (r *ReferRule) GetString() string {
	return r.ReferRule
}

