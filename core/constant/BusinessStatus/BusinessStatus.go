package BusinessStatus

var (
	ZERO    = BusinessStatus{"为成交", 0}
	PORTION = BusinessStatus{"部分成交", 1}
	ALL     = BusinessStatus{"全部成交", 2}
)

func (this *BusinessStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *BusinessStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type BusinessStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
