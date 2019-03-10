package modeldef

type Reply struct {
	Id int
	Content string
	PostTime int64
	Thread *Thread `orm:"rel(fk)"`
	User *User `orm:"rel(fk)"`
}