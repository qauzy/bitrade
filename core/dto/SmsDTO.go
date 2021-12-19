package dto

func (this *SmsDTO) SetId(id int64) (result *SmsDTO) {
	this.Id = id
	return this
}
func (this *SmsDTO) GetId() (id int64) {
	return this.Id
}
func (this *SmsDTO) SetKeyId(keyId string) (result *SmsDTO) {
	this.KeyId = keyId
	return this
}
func (this *SmsDTO) GetKeyId() (keyId string) {
	return this.KeyId
}
func (this *SmsDTO) SetKeySecret(keySecret string) (result *SmsDTO) {
	this.KeySecret = keySecret
	return this
}
func (this *SmsDTO) GetKeySecret() (keySecret string) {
	return this.KeySecret
}
func (this *SmsDTO) SetSignId(signId string) (result *SmsDTO) {
	this.SignId = signId
	return this
}
func (this *SmsDTO) GetSignId() (signId string) {
	return this.SignId
}
func (this *SmsDTO) SetTemplateId(templateId string) (result *SmsDTO) {
	this.TemplateId = templateId
	return this
}
func (this *SmsDTO) GetTemplateId() (templateId string) {
	return this.TemplateId
}
func (this *SmsDTO) SetSmsStatus(smsStatus string) (result *SmsDTO) {
	this.SmsStatus = smsStatus
	return this
}
func (this *SmsDTO) GetSmsStatus() (smsStatus string) {
	return this.SmsStatus
}
func (this *SmsDTO) SetSmsName(smsName string) (result *SmsDTO) {
	this.SmsName = smsName
	return this
}
func (this *SmsDTO) GetSmsName() (smsName string) {
	return this.SmsName
}

type SmsDTO struct {
	Id         int64
	KeyId      string
	KeySecret  string
	SignId     string
	TemplateId string
	SmsStatus  string
	SmsName    string
}
