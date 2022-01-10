package entity

func (this *BindWechat) SetRealName(realName string) (result *BindWechat) {
	this.RealName = realName
	return this
}
func (this *BindWechat) GetRealName() (realName string) {
	return this.RealName
}
func (this *BindWechat) SetWechat(wechat string) (result *BindWechat) {
	this.Wechat = wechat
	return this
}
func (this *BindWechat) GetWechat() (wechat string) {
	return this.Wechat
}
func (this *BindWechat) SetJyPassword(jyPassword string) (result *BindWechat) {
	this.JyPassword = jyPassword
	return this
}
func (this *BindWechat) GetJyPassword() (jyPassword string) {
	return this.JyPassword
}
func (this *BindWechat) SetQrCodeUrl(qrCodeUrl string) (result *BindWechat) {
	this.QrCodeUrl = qrCodeUrl
	return this
}
func (this *BindWechat) GetQrCodeUrl() (qrCodeUrl string) {
	return this.QrCodeUrl
}

type BindWechat struct {
	RealName   string
	Wechat     string
	JyPassword string
	QrCodeUrl  string
}
