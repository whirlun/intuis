package modeldef

type VideoEncode struct {
	Id          int
	VideoEncode string
	Seed        []*Seed                 `orm:"reverse(many)"`
	User        []*UserMainSiteSettings `orm:"reverse(many)"`
}

func (v *VideoEncode) GetString() string {
	return v.VideoEncode
}
