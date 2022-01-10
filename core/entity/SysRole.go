package entity

func (this *SysRole) SetId(id int64) (result *SysRole) {
	this.Id = id
	return this
}
func (this *SysRole) GetId() (id int64) {
	return this.Id
}
func (this *SysRole) SetRole(role string) (result *SysRole) {
	this.Role = role
	return this
}
func (this *SysRole) GetRole() (role string) {
	return this.Role
}
func (this *SysRole) SetDescription(description string) (result *SysRole) {
	this.Description = description
	return this
}
func (this *SysRole) GetDescription() (description string) {
	return this.Description
}
func (this *SysRole) SetPermissions(permissions []SysPermission) (result *SysRole) {
	this.Permissions = permissions
	return this
}
func (this *SysRole) GetPermissions() (permissions []SysPermission) {
	return this.Permissions
}
func NewSysRoleV2(id int64, role string) (this *SysRole) {
	this = new(SysRole)
	this.Id = id
	this.Role = role
	return
}
func NewSysRoleV3(id int64, role string, description string) (this *SysRole) {
	this = new(SysRole)
	this.Id = id
	this.Role = role
	this.Description = description
	return
}
func NewSysRoleV0() (this *SysRole) {
	this = new(SysRole)
	return
}

type SysRole struct {
	Id          int64
	Role        string
	Description string
	Permissions []SysPermission
}
