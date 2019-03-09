package modeldef

type TrackerData struct {
	Id int `orm:"index"`
	Comment int `orm:"index"`
	CreateDate int64 `orm:"index"`
	Download int `orm:"index"`
	Upload int `orm:"index"`
	Completed int `orm:"index"`
}
