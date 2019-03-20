package modeldef

type Thread struct {
	Id            int
	Title         string
	Content       string
	ContentImage  string
	PostTime      int64
	LastReplyTime int64
	Reply         int64
	Read          int64
	Love          int64
	Star          int64
	Pinned        bool
	Locked        bool
	Category      *ForumCategory `orm:"rel(fk)"`
	Replies       []*Reply       `orm:"reverse(many)"`
	LastReplyUser *User          `orm:"rel(one)"`
	User          *User          `orm:"rel(fk)"`
}
