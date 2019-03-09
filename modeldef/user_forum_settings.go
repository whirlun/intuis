package modeldef

type UserForumSettings struct {
	Id            int
	ThreadPerPage int //default 20 maximum 100
	TopicPerPage  int //default 20 maximum 100
	ShowAvatar    bool
	ShowSignature bool
	Signature     string
	User          *User `orm:"reverse(one)"`
}
