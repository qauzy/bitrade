package entity

import "github.com/qauzy/chocolate/lists/arraylist"

func (this *SysRole) SetId(Id int64) (result *SysRole) {
	this.Id = Id
	return this
}
func (this *SysRole) GetId() (Id int64) {
	return this.Id
}
func (this *SysRole) SetRole(Role string) (result *SysRole) {
	this.Role = Role
	return this
}
func (this *SysRole) GetRole() (Role string) {
	return this.Role
}
func (this *SysRole) SetDescription(Description string) (result *SysRole) {
	this.Description = Description
	return this
}
func (this *SysRole) GetDescription() (Description string) {
	return this.Description
}
func (this *SysRole) SetPermissions(Permissions arraylist.List[SysPermission]) (result *SysRole) {
	this.Permissions = Permissions
	return this
}
func (this *SysRole) GetPermissions() (Permissions arraylist.List[SysPermission]) {
	return this.Permissions
}
func NewSysRoleV2(id int64, role string) (this *SysRole) {
	this = new(SysRole)
	this.Id = this.Id
	this.Role = this.Role
	return
}
func NewSysRoleV3(id int64, role string, description string) (this *SysRole) {
	this = new(SysRole)
	this.Id = this.Id
	this.Role = this.Role
	this.Description = this.Description
	return
}
func NewSysRole() (this *SysRole) {
	this = new(SysRole)
	return
}
func NewSysRoleEx(id int64, role string, description string, permissions arraylist.List[SysPermission]) (ret *SysRole) {
	ret = new(SysRole)
	ret.Id = id
	ret.Role = role
	ret.Description = description
	ret.Permissions = permissions
	return
}

type SysRole struct {
	Id          int64                         `gorm:"column:id" json:"id"`
	Role        string                        `gorm:"column:role" json:"role"`
	Description string                        `gorm:"column:description" json:"description"`
	Permissions arraylist.List[SysPermission] `gorm:"column:permissions" json:"permissions"`
}
