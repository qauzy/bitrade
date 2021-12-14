package entity

func (this *Feedback) SetId(id int64) {
	this.Id = id
}
func (this *Feedback) GetId() (id int64) {
	return this.Id
}
func (this *Feedback) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *Feedback) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *Feedback) SetRemark(remark string) {
	this.Remark = remark
}
func (this *Feedback) GetRemark() (remark string) {
	return this.Remark
}
func (this *Feedback) SetMember(member Member) {
	this.Member = member
}
func (this *Feedback) GetMember() (member Member) {
	return this.Member
}

type Feedback struct {
	Id         int64
	CreateTime Date
	Remark     string
	Member     Member
}
