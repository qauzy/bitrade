package entity

func (this *BindAli) SetRealName(realName string) (result *BindAli) {
	this.RealName = realName
	return this
}
func (this *BindAli) GetRealName() (realName string) {
	return this.RealName
}
func (this *BindAli) SetAli(ali string) (result *BindAli) {
	this.Ali = ali
	return this
}
func (this *BindAli) GetAli() (ali string) {
	return this.Ali
}
func (this *BindAli) SetJyPassword(jyPassword string) (result *BindAli) {
	this.JyPassword = jyPassword
	return this
}
func (this *BindAli) GetJyPassword() (jyPassword string) {
	return this.JyPassword
}
func (this *BindAli) SetQrCodeUrl(qrCodeUrl string) (result *BindAli) {
	this.QrCodeUrl = qrCodeUrl
	return this
}
func (this *BindAli) GetQrCodeUrl() (qrCodeUrl string) {
	return this.QrCodeUrl
}

type BindAli struct {
	RealName   string
	Ali        string
	JyPassword string
	QrCodeUrl  string
}
