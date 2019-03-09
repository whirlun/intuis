package modeldef

type AudioEncode struct {
	Id          int
	AudioEncode string
	Seed        []*Seed                 `orm:"reverse(many)"`
	User        []*UserMainSiteSettings `orm:"reverse(many)"`
}

func (a *AudioEncode) GetString() string {
	return a.AudioEncode
}

