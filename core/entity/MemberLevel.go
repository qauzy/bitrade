package entity

func (this *MemberLevel) SetId(id int64) (result *MemberLevel) {
	this.Id = id
	return this
}
func (this *MemberLevel) GetId() (id int64) {
	return this.Id
}
func (this *MemberLevel) SetName(name string) (result *MemberLevel) {
	this.Name = name
	return this
}
func (this *MemberLevel) GetName() (name string) {
	return this.Name
}
func (this *MemberLevel) SetIsDefault(isDefault bool) (result *MemberLevel) {
	this.IsDefault = isDefault
	return this
}
func (this *MemberLevel) GetIsDefault() (isDefault bool) {
	return this.IsDefault
}
func (this *MemberLevel) SetRemark(remark string) (result *MemberLevel) {
	this.Remark = remark
	return this
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
