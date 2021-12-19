package vo

func (this *EmptionRecrodVO) SetUserId(userId int64) (result *EmptionRecrodVO) {
	this.UserId = userId
	return this
}
func (this *EmptionRecrodVO) GetUserId() (userId int64) {
	return this.UserId
}
func (this *EmptionRecrodVO) SetStartTime(startTime string) (result *EmptionRecrodVO) {
	this.StartTime = startTime
	return this
}
func (this *EmptionRecrodVO) GetStartTime() (startTime string) {
	return this.StartTime
}
func (this *EmptionRecrodVO) SetEndTime(endTime string) (result *EmptionRecrodVO) {
	this.EndTime = endTime
	return this
}
func (this *EmptionRecrodVO) GetEndTime() (endTime string) {
	return this.EndTime
}
func (this *EmptionRecrodVO) SetUserName(userName string) (result *EmptionRecrodVO) {
	this.UserName = userName
	return this
}
func (this *EmptionRecrodVO) GetUserName() (userName string) {
	return this.UserName
}
func (this *EmptionRecrodVO) SetUserMobile(userMobile string) (result *EmptionRecrodVO) {
	this.UserMobile = userMobile
	return this
}
func (this *EmptionRecrodVO) GetUserMobile() (userMobile string) {
	return this.UserMobile
}
func (this *EmptionRecrodVO) SetIeoName(ieoName string) (result *EmptionRecrodVO) {
	this.IeoName = ieoName
	return this
}
func (this *EmptionRecrodVO) GetIeoName() (ieoName string) {
	return this.IeoName
}
func (this *EmptionRecrodVO) SetStatus(status string) (result *EmptionRecrodVO) {
	this.Status = status
	return this
}
func (this *EmptionRecrodVO) GetStatus() (status string) {
	return this.Status
}

type EmptionRecrodVO struct {
	UserId     int64
	StartTime  string
	EndTime    string
	UserName   string
	UserMobile string
	IeoName    string
	Status     string
}
