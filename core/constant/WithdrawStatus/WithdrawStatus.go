package WithdrawStatus

var (
	PROCESSING = WithdrawStatus{"审核中", 0}
	WAITING    = WithdrawStatus{"等待放币", 1}
	FAIL       = WithdrawStatus{"失败", 2}
	SUCCESS    = WithdrawStatus{"成功", 3}
	TRANSFER   = WithdrawStatus{"转账中", 4}
)

func (this *WithdrawStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *WithdrawStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type WithdrawStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
