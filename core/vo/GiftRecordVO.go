package vo

func (this *GiftRecordVO) SetStartTime(startTime string) (result *GiftRecordVO) {
	this.StartTime = startTime
	return this
}
func (this *GiftRecordVO) GetStartTime() (startTime string) {
	return this.StartTime
}
func (this *GiftRecordVO) SetEndTime(endTime string) (result *GiftRecordVO) {
	this.EndTime = endTime
	return this
}
func (this *GiftRecordVO) GetEndTime() (endTime string) {
	return this.EndTime
}
func (this *GiftRecordVO) SetUserId(userId int64) (result *GiftRecordVO) {
	this.UserId = userId
	return this
}
func (this *GiftRecordVO) GetUserId() (userId int64) {
	return this.UserId
}
func (this *GiftRecordVO) SetUserName(userName string) (result *GiftRecordVO) {
	this.UserName = userName
	return this
}
func (this *GiftRecordVO) GetUserName() (userName string) {
	return this.UserName
}
func (this *GiftRecordVO) SetMobile(mobile string) (result *GiftRecordVO) {
	this.Mobile = mobile
	return this
}
func (this *GiftRecordVO) GetMobile() (mobile string) {
	return this.Mobile
}

type GiftRecordVO struct {
	StartTime string
	EndTime   string
	UserId    int64
	UserName  string
	Mobile    string
}
