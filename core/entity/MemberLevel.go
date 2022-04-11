package entity

func (this *MemberLevel) SetId(Id int64) (result *MemberLevel) {
	this.Id = Id
	return this
}
func (this *MemberLevel) GetId() (Id int64) {
	return this.Id
}
func (this *MemberLevel) SetName(Name string) (result *MemberLevel) {
	this.Name = Name
	return this
}
func (this *MemberLevel) GetName() (Name string) {
	return this.Name
}
func (this *MemberLevel) SetIsDefault(IsDefault bool) (result *MemberLevel) {
	this.IsDefault = IsDefault
	return this
}
func (this *MemberLevel) GetIsDefault() (IsDefault bool) {
	return this.IsDefault
}
func (this *MemberLevel) SetRemark(Remark string) (result *MemberLevel) {
	this.Remark = Remark
	return this
}
func (this *MemberLevel) GetRemark() (Remark string) {
	return this.Remark
}
func NewMemberLevel(id int64, name string, isDefault bool, remark string) (ret *MemberLevel) {
	ret = new(MemberLevel)
	ret.Id = id
	ret.Name = name
	ret.IsDefault = isDefault
	ret.Remark = remark
	return
}

type MemberLevel struct {
	Id        int64  `gorm:"column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	IsDefault bool   `gorm:"column:is_default" json:"isDefault"`
	Remark    string `gorm:"column:remark" json:"remark"`
}
