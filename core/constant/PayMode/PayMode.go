package PayMode

type PayMode int

const (
	ALIPAY PayMode = iota
	WECHAT
	BANK
	PAYPAL
)

func (this PayMode) String() string {
	switch this {
	case ALIPAY:
		return "支付宝"
	case WECHAT:
		return "微信"
	case BANK:
		return "银联"
	case PAYPAL:
		return "Paypal"
	}
	return ""
}
