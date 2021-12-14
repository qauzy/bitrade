package entity

func (this *BusinessAuthApplyDetailVO) SetId(id int64) {
	this.Id = id
}
func (this *BusinessAuthApplyDetailVO) GetId() (id int64) {
	return this.Id
}
func (this *BusinessAuthApplyDetailVO) SetInfo(info JSONObject) {
	this.Info = info
}
func (this *BusinessAuthApplyDetailVO) GetInfo() (info JSONObject) {
	return this.Info
}
func (this *BusinessAuthApplyDetailVO) SetStatus(status CertifiedBusinessStatus) {
	this.Status = status
}
func (this *BusinessAuthApplyDetailVO) GetStatus() (status CertifiedBusinessStatus) {
	return this.Status
}
func (this *BusinessAuthApplyDetailVO) SetCheckTime(checkTime Date) {
	this.CheckTime = checkTime
}
func (this *BusinessAuthApplyDetailVO) GetCheckTime() (checkTime Date) {
	return this.CheckTime
}
func (this *BusinessAuthApplyDetailVO) SetRealName(realName string) {
	this.RealName = realName
}
func (this *BusinessAuthApplyDetailVO) GetRealName() (realName string) {
	return this.RealName
}
func (this *BusinessAuthApplyDetailVO) SetDetail(detail string) {
	this.Detail = detail
}
func (this *BusinessAuthApplyDetailVO) GetDetail() (detail string) {
	return this.Detail
}
func (this *BusinessAuthApplyDetailVO) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *BusinessAuthApplyDetailVO) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthApplyDetailVO) SetAuthInfo(authInfo string) {
	this.AuthInfo = authInfo
}
func (this *BusinessAuthApplyDetailVO) GetAuthInfo() (authInfo string) {
	return this.AuthInfo
}

type BusinessAuthApplyDetailVO struct {
	Id        int64
	Info      JSONObject
	Status    CertifiedBusinessStatus
	CheckTime Date
	RealName  string
	Detail    string
	Amount    BigDecimal
	AuthInfo  string
}
