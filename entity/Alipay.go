package entity

var SerialVersionUID int64 = 8317734763036284945

func (this *Alipay) SetAliNo(aliNo string) {
	this.AliNo = aliNo
}
func (this *Alipay) GetAliNo() (aliNo string) {
	return this.AliNo
}
func (this *Alipay) SetQrCodeUrl(qrCodeUrl string) {
	this.QrCodeUrl = qrCodeUrl
}
func (this *Alipay) GetQrCodeUrl() (qrCodeUrl string) {
	return this.QrCodeUrl
}

type Alipay struct {
	AliNo     string
	QrCodeUrl string
}
