package modeldef

type LovesCount struct {
	Id       int
	Operator string `orm:"index"`
	Receiver string `orm:"index"'`
	Count    int64  `orm:"index"'`
}
