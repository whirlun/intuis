package modeldef


type SeedFilesProfile struct {
	Id          int
	Path        string
	Length      int64
	SeedProfile *SeedProfile `orm:"rel(fk)"`
}

