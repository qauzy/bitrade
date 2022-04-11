package entity

func (this *WechatPay) SetWechat(Wechat string) (result *WechatPay) {
	this.Wechat = Wechat
	return this
}
func (this *WechatPay) GetWechat() (Wechat string) {
	return this.Wechat
}
func (this *WechatPay) SetQrWeCodeUrl(QrWeCodeUrl string) (result *WechatPay) {
	this.QrWeCodeUrl = QrWeCodeUrl
	return this
}
func (this *WechatPay) GetQrWeCodeUrl() (QrWeCodeUrl string) {
	return this.QrWeCodeUrl
}
func NewWechatPay(wechat string, qrWeCodeUrl string) (ret *WechatPay) {
	ret = new(WechatPay)
	ret.Wechat = wechat
	ret.QrWeCodeUrl = qrWeCodeUrl
	return
}

type WechatPay struct {
	Wechat      string `gorm:"column:wechat" json:"wechat"`
	QrWeCodeUrl string `gorm:"column:qr_we_code_url" json:"qrWeCodeUrl"`
}
