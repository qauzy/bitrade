package entity

func (this *BindAli) SetRealName(RealName string) (result *BindAli) {
	this.RealName = RealName
	return this
}
func (this *BindAli) GetRealName() (RealName string) {
	return this.RealName
}
func (this *BindAli) SetAli(Ali string) (result *BindAli) {
	this.Ali = Ali
	return this
}
func (this *BindAli) GetAli() (Ali string) {
	return this.Ali
}
func (this *BindAli) SetJyPassword(JyPassword string) (result *BindAli) {
	this.JyPassword = JyPassword
	return this
}
func (this *BindAli) GetJyPassword() (JyPassword string) {
	return this.JyPassword
}
func (this *BindAli) SetQrCodeUrl(QrCodeUrl string) (result *BindAli) {
	this.QrCodeUrl = QrCodeUrl
	return this
}
func (this *BindAli) GetQrCodeUrl() (QrCodeUrl string) {
	return this.QrCodeUrl
}

type BindAli struct {
	RealName   string `gorm:"column:real_name" json:"realName"`
	Ali        string `gorm:"column:ali" json:"ali"`
	JyPassword string `gorm:"column:jy_password" json:"jyPassword"`
	QrCodeUrl  string `gorm:"column:qr_code_url" json:"qrCodeUrl"`
}
