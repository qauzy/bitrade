package entity

import "github.com/qauzy/chocolate/xtime"

func (this *Feedback) SetId(Id int64) (result *Feedback) {
	this.Id = Id
	return this
}
func (this *Feedback) GetId() (Id int64) {
	return this.Id
}
func (this *Feedback) SetCreateTime(CreateTime xtime.Xtime) (result *Feedback) {
	this.CreateTime = CreateTime
	return this
}
func (this *Feedback) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *Feedback) SetRemark(Remark string) (result *Feedback) {
	this.Remark = Remark
	return this
}
func (this *Feedback) GetRemark() (Remark string) {
	return this.Remark
}
func (this *Feedback) SetMember(Member *Member) (result *Feedback) {
	this.Member = Member
	return this
}
func (this *Feedback) GetMember() (Member *Member) {
	return this.Member
}
func NewFeedback(id int64, createTime xtime.Xtime, remark string, member *Member) (ret *Feedback) {
	ret = new(Feedback)
	ret.Id = id
	ret.CreateTime = createTime
	ret.Remark = remark
	ret.Member = member
	return
}

type Feedback struct {
	Id         int64       `gorm:"column:id" json:"id"`
	CreateTime xtime.Xtime `gorm:"column:create_time" json:"createTime"`
	Remark     string      `gorm:"column:remark" json:"remark"`
	Member     *Member     `gorm:"column:member" json:"member"`
}
