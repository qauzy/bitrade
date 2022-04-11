package entity

func (this *Alipay) SetAliNo(AliNo string) (result *Alipay) {
	this.AliNo = AliNo
	return this
}
func (this *Alipay) GetAliNo() (AliNo string) {
	return this.AliNo
}
func (this *Alipay) SetQrCodeUrl(QrCodeUrl string) (result *Alipay) {
	this.QrCodeUrl = QrCodeUrl
	return this
}
func (this *Alipay) GetQrCodeUrl() (QrCodeUrl string) {
	return this.QrCodeUrl
}
func NewAlipay(aliNo string, qrCodeUrl string) (ret *Alipay) {
	ret = new(Alipay)
	ret.AliNo = aliNo
	ret.QrCodeUrl = qrCodeUrl
	return
}

type Alipay struct {
	AliNo     string `gorm:"column:ali_no" json:"aliNo"`
	QrCodeUrl string `gorm:"column:qr_code_url" json:"qrCodeUrl"`
}
