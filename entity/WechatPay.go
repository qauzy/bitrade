package entity

var SerialVersionUID int64 = 1511509989072675896

func (this *WechatPay) SetWechat(wechat string) {
	this.Wechat = wechat
}
func (this *WechatPay) GetWechat() (wechat string) {
	return this.Wechat
}
func (this *WechatPay) SetQrWeCodeUrl(qrWeCodeUrl string) {
	this.QrWeCodeUrl = qrWeCodeUrl
}
func (this *WechatPay) GetQrWeCodeUrl() (qrWeCodeUrl string) {
	return this.QrWeCodeUrl
}

type WechatPay struct {
	Wechat      string
	QrWeCodeUrl string
}
