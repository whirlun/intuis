package modeldef

type Category struct {
	Id       int
	Category string
	Seed     []*Seed                 `orm:"reverse(many)"`
	User     []*UserMainSiteSettings `orm:"reverse(many)"`
}

func (c *Category) GetString() string {
	return c.Category
}
