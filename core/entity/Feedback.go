package entity

import "time"

func (this *Feedback) SetId(id int64) (result *Feedback) {
	this.Id = id
	return this
}
func (this *Feedback) GetId() (id int64) {
	return this.Id
}
func (this *Feedback) SetCreateTime(createTime time.Time) (result *Feedback) {
	this.CreateTime = createTime
	return this
}
func (this *Feedback) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *Feedback) SetRemark(remark string) (result *Feedback) {
	this.Remark = remark
	return this
}
func (this *Feedback) GetRemark() (remark string) {
	return this.Remark
}
func (this *Feedback) SetMember(member Member) (result *Feedback) {
	this.Member = member
	return this
}
func (this *Feedback) GetMember() (member Member) {
	return this.Member
}

type Feedback struct {
	Id         int64
	CreateTime time.Time
	Remark     string
	Member     Member
}
