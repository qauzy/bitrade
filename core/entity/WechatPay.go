package entity

func (this *WechatPay) SetWechat(wechat string) (result *WechatPay) {
	this.Wechat = wechat
	return this
}
func (this *WechatPay) GetWechat() (wechat string) {
	return this.Wechat
}
func (this *WechatPay) SetQrWeCodeUrl(qrWeCodeUrl string) (result *WechatPay) {
	this.QrWeCodeUrl = qrWeCodeUrl
	return this
}
func (this *WechatPay) GetQrWeCodeUrl() (qrWeCodeUrl string) {
	return this.QrWeCodeUrl
}

type WechatPay struct {
	Wechat      string
	QrWeCodeUrl string
}
