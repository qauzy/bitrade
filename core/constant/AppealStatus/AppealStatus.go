package AppealStatus

var (
	NOT_PROCESSED = AppealStatus{"未处理", 0}
	PROCESSED     = AppealStatus{"已处理", 1}
)

func (this *AppealStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AppealStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AppealStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
