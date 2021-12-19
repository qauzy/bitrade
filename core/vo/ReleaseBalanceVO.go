package vo

import "time"

func (this *ReleaseBalanceVO) SetId(id []int64) (result *ReleaseBalanceVO) {
	this.Id = id
	return this
}
func (this *ReleaseBalanceVO) GetId() (id []int64) {
	return this.Id
}
func (this *ReleaseBalanceVO) SetMemberName(memberName string) (result *ReleaseBalanceVO) {
	this.MemberName = memberName
	return this
}
func (this *ReleaseBalanceVO) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *ReleaseBalanceVO) SetPhone(phone string) (result *ReleaseBalanceVO) {
	this.Phone = phone
	return this
}
func (this *ReleaseBalanceVO) GetPhone() (phone string) {
	return this.Phone
}
func (this *ReleaseBalanceVO) SetStartTime(startTime time.Time) (result *ReleaseBalanceVO) {
	this.StartTime = startTime
	return this
}
func (this *ReleaseBalanceVO) GetStartTime() (startTime time.Time) {
	return this.StartTime
}
func (this *ReleaseBalanceVO) SetEndTime(endTime time.Time) (result *ReleaseBalanceVO) {
	this.EndTime = endTime
	return this
}
func (this *ReleaseBalanceVO) GetEndTime() (endTime time.Time) {
	return this.EndTime
}
func (this *ReleaseBalanceVO) SetReleaseState(releaseState string) (result *ReleaseBalanceVO) {
	this.ReleaseState = releaseState
	return this
}
func (this *ReleaseBalanceVO) GetReleaseState() (releaseState string) {
	return this.ReleaseState
}
func (this *ReleaseBalanceVO) SetPageNo(pageNo int64) (result *ReleaseBalanceVO) {
	this.PageNo = pageNo
	return this
}
func (this *ReleaseBalanceVO) GetPageNo() (pageNo int64) {
	return this.PageNo
}
func (this *ReleaseBalanceVO) SetPageSize(pageSize int64) (result *ReleaseBalanceVO) {
	this.PageSize = pageSize
	return this
}
func (this *ReleaseBalanceVO) GetPageSize() (pageSize int64) {
	return this.PageSize
}

type ReleaseBalanceVO struct {
	Id           []int64
	MemberName   string
	Phone        string
	StartTime    time.Time
	EndTime      time.Time
	ReleaseState string
	PageNo       int64
	PageSize     int64
}
