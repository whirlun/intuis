package modeldef

import "time"

type Loves struct {
	Id          int
	OperateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Operator    *User     `orm:"rel(fk)"`
	Receiver    *User     `orm:"rel(fk)"`
	Thread      *Thread   `orm:"rel(fk)"`
}
