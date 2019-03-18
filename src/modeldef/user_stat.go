package modeldef

type UserStat struct {
	Id           int
	Visited      int64
	ReadThread   int64
	CreateThread int64
	Commented    int64
	Loved        int64
	Beloved      int64
	User         *User `orm:"reverse(one)"`
}
