package entity

func (this *Department) SetId(id int64) {
	this.Id = id
}
func (this *Department) GetId() (id int64) {
	return this.Id
}
func (this *Department) SetName(name string) {
	this.Name = name
}
func (this *Department) GetName() (name string) {
	return this.Name
}
func (this *Department) SetRemark(remark string) {
	this.Remark = remark
}
func (this *Department) GetRemark() (remark string) {
	return this.Remark
}
func (this *Department) SetLeaderId(leaderId int64) {
	this.LeaderId = leaderId
}
func (this *Department) GetLeaderId() (leaderId int64) {
	return this.LeaderId
}
func (this *Department) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *Department) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *Department) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *Department) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}

type Department struct {
	Id         int64
	Name       string
	Remark     string
	LeaderId   int64
	CreateTime Date
	UpdateTime Date
}
