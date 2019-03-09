package modeldef

type IMDBData struct {
	Id int
	Name string
	Alias string
	Pubyear int
	Length int
	Vote int
	Score int
	Language string
	Area string
	Genre string
	Director string
	ScriptWriter string
	Introduction string `orm:"type(text)"`
	Actors string
	Seed *Seed `orm:"rel(fk)"`
}
