package entity

func (this *SysRole) SetId(id int64) {
	this.Id = id
}
func (this *SysRole) GetId() (id int64) {
	return this.Id
}
func (this *SysRole) SetRole(role string) {
	this.Role = role
}
func (this *SysRole) GetRole() (role string) {
	return this.Role
}
func (this *SysRole) SetDescription(description string) {
	this.Description = description
}
func (this *SysRole) GetDescription() (description string) {
	return this.Description
}
func (this *SysRole) SetPermissions(permissions []SysPermission) {
	this.Permissions = permissions
}
func (this *SysRole) GetPermissions() (permissions []SysPermission) {
	return this.Permissions
}
func NewSysRole(id int64, role string) (this *SysRole) {
	this = new(SysRole)
	this.Id = id
	this.Role = role
	return
}
func NewSysRole(id int64, role string, description string) (this *SysRole) {
	this = new(SysRole)
	this.Id = id
	this.Role = role
	this.Description = description
	return
}
func NewSysRole() (this *SysRole) {
	this = new(SysRole)
	return
}

type SysRole struct {
	Id          int64
	Role        string
	Description string
	Permissions []SysPermission
}
