package entity

func (this *Alipay) SetAliNo(aliNo string) (result *Alipay) {
	this.AliNo = aliNo
	return this
}
func (this *Alipay) GetAliNo() (aliNo string) {
	return this.AliNo
}
func (this *Alipay) SetQrCodeUrl(qrCodeUrl string) (result *Alipay) {
	this.QrCodeUrl = qrCodeUrl
	return this
}
func (this *Alipay) GetQrCodeUrl() (qrCodeUrl string) {
	return this.QrCodeUrl
}

type Alipay struct {
	AliNo     string
	QrCodeUrl string
}
