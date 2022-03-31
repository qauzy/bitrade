package CommonStatus

var (
	NORMAL  = CommonStatus{"正常", 0}
	ILLEGAL = CommonStatus{"非法", 1}
)

func (this *CommonStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *CommonStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type CommonStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
