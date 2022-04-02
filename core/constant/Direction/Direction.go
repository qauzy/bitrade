package Direction

var (
	ASC  = Direction{"升序", 0}
	DESC = Direction{"降序", 1}
)

func (this *Direction) Ordinal() (result int) {
	return this.ordinal
}
func (this *Direction) GetOrdinal() (result int) {
	return this.Ordinal()
}

type Direction struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
