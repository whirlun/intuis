package modeldef

type InvitationCode struct {
	Id        int
	Code      string `orm:"unique;index"`
	HostUser  *User  `orm:"rel(fk)"`
	Used      bool
	GuestUser *User `orm:"rel(one)"`
}
