package modeldef

type ForumCategory struct {
	Id           int
	Name         string
	Thread       []*Thread `orm:"reverse(many)"`
	Introduction string
	Count        int64
	Color        int32
}
