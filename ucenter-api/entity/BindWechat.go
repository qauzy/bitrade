package entity

func (this *BindWechat) SetRealName(RealName string) (result *BindWechat) {
	this.RealName = RealName
	return this
}
func (this *BindWechat) GetRealName() (RealName string) {
	return this.RealName
}
func (this *BindWechat) SetWechat(Wechat string) (result *BindWechat) {
	this.Wechat = Wechat
	return this
}
func (this *BindWechat) GetWechat() (Wechat string) {
	return this.Wechat
}
func (this *BindWechat) SetJyPassword(JyPassword string) (result *BindWechat) {
	this.JyPassword = JyPassword
	return this
}
func (this *BindWechat) GetJyPassword() (JyPassword string) {
	return this.JyPassword
}
func (this *BindWechat) SetQrCodeUrl(QrCodeUrl string) (result *BindWechat) {
	this.QrCodeUrl = QrCodeUrl
	return this
}
func (this *BindWechat) GetQrCodeUrl() (QrCodeUrl string) {
	return this.QrCodeUrl
}

type BindWechat struct {
	RealName   string `gorm:"column:real_name" json:"realName"`
	Wechat     string `gorm:"column:wechat" json:"wechat"`
	JyPassword string `gorm:"column:jy_password" json:"jyPassword"`
	QrCodeUrl  string `gorm:"column:qr_code_url" json:"qrCodeUrl"`
}
