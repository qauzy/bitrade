package constant

type PayMode int

const (
	PAYMODE_ALIPAY PayMode = iota
	PAYMODE_WECHAT
	PAYMODE_BANK
	PAYMODE_PAYPAL
)

func (this PayMode) String() string {
	switch this {
	case PAYMODE_ALIPAY:
		return "支付宝"
	case PAYMODE_WECHAT:
		return "微信"
	case PAYMODE_BANK:
		return "银联"
	case PAYMODE_PAYPAL:
		return "Paypal"
	}
	return ""
}
