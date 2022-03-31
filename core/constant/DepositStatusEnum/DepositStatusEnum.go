package DepositStatusEnum

var (
	PAY      = DepositStatusEnum{"缴纳", 0}
	GET_BACK = DepositStatusEnum{"索回", 1}
)

func (this *DepositStatusEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *DepositStatusEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type DepositStatusEnum struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
