package dto

func (this *SmsDTO) SetId(Id int64) (result *SmsDTO) {
	this.Id = Id
	return this
}
func (this *SmsDTO) GetId() (Id int64) {
	return this.Id
}
func (this *SmsDTO) SetKeyId(KeyId string) (result *SmsDTO) {
	this.KeyId = KeyId
	return this
}
func (this *SmsDTO) GetKeyId() (KeyId string) {
	return this.KeyId
}
func (this *SmsDTO) SetKeySecret(KeySecret string) (result *SmsDTO) {
	this.KeySecret = KeySecret
	return this
}
func (this *SmsDTO) GetKeySecret() (KeySecret string) {
	return this.KeySecret
}
func (this *SmsDTO) SetSignId(SignId string) (result *SmsDTO) {
	this.SignId = SignId
	return this
}
func (this *SmsDTO) GetSignId() (SignId string) {
	return this.SignId
}
func (this *SmsDTO) SetTemplateId(TemplateId string) (result *SmsDTO) {
	this.TemplateId = TemplateId
	return this
}
func (this *SmsDTO) GetTemplateId() (TemplateId string) {
	return this.TemplateId
}
func (this *SmsDTO) SetSmsStatus(SmsStatus string) (result *SmsDTO) {
	this.SmsStatus = SmsStatus
	return this
}
func (this *SmsDTO) GetSmsStatus() (SmsStatus string) {
	return this.SmsStatus
}
func (this *SmsDTO) SetSmsName(SmsName string) (result *SmsDTO) {
	this.SmsName = SmsName
	return this
}
func (this *SmsDTO) GetSmsName() (SmsName string) {
	return this.SmsName
}

type SmsDTO struct {
	Id         int64  `gorm:"column:id" json:"id"`
	KeyId      string `gorm:"column:key_id" json:"keyId"`
	KeySecret  string `gorm:"column:key_secret" json:"keySecret"`
	SignId     string `gorm:"column:sign_id" json:"signId"`
	TemplateId string `gorm:"column:template_id" json:"templateId"`
	SmsStatus  string `gorm:"column:sms_status" json:"smsStatus"`
	SmsName    string `gorm:"column:sms_name" json:"smsName"`
}
