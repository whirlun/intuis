package modeldef

type SeedProfile struct {
	Id           int
	Announce     string
	CreationDate int64
	Comment      string
	CreatedBy    string
	Encoding     string
	Files        []*SeedFilesProfile `orm:"reverse(many)"`
	PieceLength  int64
	Pieces       string
	Private      bool
	Seed         *Seed `orm:"reverse(one)"`
}

