package modeldef

type ProductionGroup struct {
	Id              int
	ProductionGroup string
	Seed            []*Seed                 `orm:"reverse(many)"`
	User            []*UserMainSiteSettings `orm:"reverse(many)"`
}

func (p *ProductionGroup) GetString() string {
	return p.ProductionGroup
}

