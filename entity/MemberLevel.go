package entity

func (this *MemberLevel) SetId(id int64) {
	this.Id = id
}
func (this *MemberLevel) GetId() (id int64) {
	return this.Id
}
func (this *MemberLevel) SetName(name string) {
	this.Name = name
}
func (this *MemberLevel) GetName() (name string) {
	return this.Name
}
func (this *MemberLevel) SetIsDefault(isDefault bool) {
	this.IsDefault = isDefault
}
func (this *MemberLevel) GetIsDefault() (isDefault bool) {
	return this.IsDefault
}
func (this *MemberLevel) SetRemark(remark string) {
	this.Remark = remark
}
func (this *MemberLevel) GetRemark() (remark string) {
	return this.Remark
}

type MemberLevel struct {
	Id        int64
	Name      string
	IsDefault bool
	Remark    string
}
