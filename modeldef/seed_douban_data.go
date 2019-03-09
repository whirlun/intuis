package modeldef

type DoubanData struct {
	Id int
	Name string
	Actors string
	Genre string
	Area string
	Director string
	Pubyear string
	Vote int
	Score float64
	Introduction string `orm:"type(text)"`
	Seed *Seed `orm:"rel(fk)"`
}
