package OrderStatus

var (
	CANCELLED  = OrderStatus{"已取消", 0}
	NONPAYMENT = OrderStatus{"未付款", 1}
	PAID       = OrderStatus{"已付款", 2}
	COMPLETED  = OrderStatus{"已完成", 3}
	APPEAL     = OrderStatus{"申诉中", 4}
)

func (this *OrderStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *OrderStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type OrderStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
