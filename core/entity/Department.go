package entity

import "time"

func (this *Department) SetId(id int64) (result *Department) {
	this.Id = id
	return this
}
func (this *Department) GetId() (id int64) {
	return this.Id
}
func (this *Department) SetName(name string) (result *Department) {
	this.Name = name
	return this
}
func (this *Department) GetName() (name string) {
	return this.Name
}
func (this *Department) SetRemark(remark string) (result *Department) {
	this.Remark = remark
	return this
}
func (this *Department) GetRemark() (remark string) {
	return this.Remark
}
func (this *Department) SetLeaderId(leaderId int64) (result *Department) {
	this.LeaderId = leaderId
	return this
}
func (this *Department) GetLeaderId() (leaderId int64) {
	return this.LeaderId
}
func (this *Department) SetCreateTime(createTime time.Time) (result *Department) {
	this.CreateTime = createTime
	return this
}
func (this *Department) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *Department) SetUpdateTime(updateTime time.Time) (result *Department) {
	this.UpdateTime = updateTime
	return this
}
func (this *Department) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}

type Department struct {
	Id         int64
	Name       string
	Remark     string
	LeaderId   int64
	CreateTime time.Time
	UpdateTime time.Time
}
