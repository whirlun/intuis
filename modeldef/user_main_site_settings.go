package modeldef

type UserMainSiteSettings struct {
	Id               int
	Freeze           bool
	UseHttps         bool
	Gender           int
	Country          string
	Downstream       string
	Upstream         string
	ISP              string
	Categories       []*Category        `orm:"rel(m2m)"`
	Media            []*Medium          `orm:"rel(m2m)"`
	Format           []*Format          `orm:"rel(m2m)"`
	VideoEncodes     []*VideoEncode     `orm:"rel(m2m)"`
	AudioEncodes     []*AudioEncode     `orm:"rel(m2m)"`
	ProductionGroup  []*ProductionGroup `orm:"rel(m2m)"`
	ActiveSetting    string             //seed status:active,inactive or both
	FreeSetting      string             //free status:all,normal,free,2xfree,2x,50%,50%free,30%
	BookmarkSetting  string             //mark status:marked,unmarked,all
	Language         string
	FontSize         int
	RecommendClassic bool  //whether show classic recommandation on main page
	RecommendHot     bool  //whether show hot recommandation on main page
	ShowAdvertisment bool  // whether show advertisment
	ShowContent      bool
	ShowNFO          bool
	ShowIMDB         bool
	ShowDouban       bool
	ShowComment      bool
	SeedPerPage      int
	User             *User `orm:"reverse(one)"`
}