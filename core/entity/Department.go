package entity

import "github.com/qauzy/chocolate/xtime"

func (this *Department) SetId(Id int64) (result *Department) {
	this.Id = Id
	return this
}
func (this *Department) GetId() (Id int64) {
	return this.Id
}
func (this *Department) SetName(Name string) (result *Department) {
	this.Name = Name
	return this
}
func (this *Department) GetName() (Name string) {
	return this.Name
}
func (this *Department) SetRemark(Remark string) (result *Department) {
	this.Remark = Remark
	return this
}
func (this *Department) GetRemark() (Remark string) {
	return this.Remark
}
func (this *Department) SetLeaderId(LeaderId int64) (result *Department) {
	this.LeaderId = LeaderId
	return this
}
func (this *Department) GetLeaderId() (LeaderId int64) {
	return this.LeaderId
}
func (this *Department) SetCreateTime(CreateTime xtime.Xtime) (result *Department) {
	this.CreateTime = CreateTime
	return this
}
func (this *Department) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *Department) SetUpdateTime(UpdateTime xtime.Xtime) (result *Department) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *Department) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func NewDepartment(id int64, name string, remark string, leaderId int64, createTime xtime.Xtime, updateTime xtime.Xtime) (ret *Department) {
	ret = new(Department)
	ret.Id = id
	ret.Name = name
	ret.Remark = remark
	ret.LeaderId = leaderId
	ret.CreateTime = createTime
	ret.UpdateTime = updateTime
	return
}

type Department struct {
	Id         int64       `gorm:"column:id" json:"id"`
	Name       string      `gorm:"column:name" json:"name"`
	Remark     string      `gorm:"column:remark" json:"remark"`
	LeaderId   int64       `gorm:"column:leader_id" json:"leaderId"`
	CreateTime xtime.Xtime `gorm:"column:create_time" json:"createTime"`
	UpdateTime xtime.Xtime `gorm:"column:update_time" json:"updateTime"`
}
