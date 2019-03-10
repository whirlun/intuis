package modeldef

type User struct {
	Id                  int
	Name                string `orm:"unique;index"`
	Email               string `orm:"unique"`
	PasswordDigest      string
	Avatar              string
	PasswordQuestion    string
	PasswordAnswer      string
	PersonalInstruction string
	Privacy             bool
	Seed                []*Seed               `orm:"reverse(many)"`
	Threads             []*Thread             `orm:"reverse(many)"`
	Reply               []*Reply              `orm:"reverse(many)"`
	MainSiteSettings    *UserMainSiteSettings `orm:"rel(one)"`
	PmSettings          *UserPmSettings       `orm:"rel(one)"`
	ForumSettings       *UserForumSettings    `orm:"rel(one)"`
	Codes               []*InvitationCode     `orm:"reverse(many)"`
	UseCode             *InvitationCode       `orm:"reverse(one)"`
	LastForumThread     *Thread               `orm:"reverse(one)"`
}


