package PayMode

var (
	ALIPAY = PayMode{"支付宝", 0}
	WECHAT = PayMode{"微信", 1}
	BANK   = PayMode{"银联", 2}
	PAYPAL = PayMode{"Paypal", 3}
)

func (this *PayMode) Ordinal() (result int) {
	return this.ordinal
}
func (this *PayMode) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PayMode struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
