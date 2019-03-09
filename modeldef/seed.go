package modeldef

type Seed struct {
	Id              int
	Title           string
	Subtitle        string
	DoubanLink      string
	IMDBLink        string
	Poster          string
	Torrent         string
	Nfo             string
	Is3D            bool
	Top				bool
	ContentImage    string
	Locked          bool
	FreeSetting     string  //free status:all,normal,free,2xfree,2x,50%,50%free,30%
	Category        *Category        `orm:"rel(fk)"`
	AudioEncode     *AudioEncode     `orm:"rel(fk)"`
	Format          *Format          `orm:"rel(fk)"`
	ProductionGroup *ProductionGroup `orm:"rel(fk)"`
	Medium          *Medium          `orm:"rel(fk)"`
	ReferRule       *ReferRule       `orm:"rel(fk)"`
	VideoEncode     *VideoEncode     `orm:"rel(fk)"`
	User            *User            `orm:"rel(fk)"`
	Content         string           `orm:"type(text)"`
	SeedProfile     *SeedProfile     `orm:"rel(one)"`
	SeedIMDBData    []*IMDBData      `orm:"reverse(many);null"`
	SeedDoubanData  []*DoubanData    `orm:"reverse(many);null"`
	SeedTrackerData *TrackerData     `orm:"rel(one)"`
}

type SeedPage struct {
	Id int
	Title string
	Subtitle string
	IMDBPoint string
	FreeSetting string
	ReferRule string
	Comments string
	Datetime int64
	Size string
	Upload string
	Download string
	Completed string
	Author string
}


