package entity

func (this *SysPermission) SetId(id int64) (result *SysPermission) {
	this.Id = id
	return this
}
func (this *SysPermission) GetId() (id int64) {
	return this.Id
}
func (this *SysPermission) SetTitle(title string) (result *SysPermission) {
	this.Title = title
	return this
}
func (this *SysPermission) GetTitle() (title string) {
	return this.Title
}
func (this *SysPermission) SetDescription(description string) (result *SysPermission) {
	this.Description = description
	return this
}
func (this *SysPermission) GetDescription() (description string) {
	return this.Description
}
func (this *SysPermission) SetParentId(parentId int64) (result *SysPermission) {
	this.ParentId = parentId
	return this
}
func (this *SysPermission) GetParentId() (parentId int64) {
	return this.ParentId
}
func (this *SysPermission) SetSort(sort int64) (result *SysPermission) {
	this.Sort = sort
	return this
}
func (this *SysPermission) GetSort() (sort int64) {
	return this.Sort
}
func (this *SysPermission) SetName(name string) (result *SysPermission) {
	this.Name = name
	return this
}
func (this *SysPermission) GetName() (name string) {
	return this.Name
}
func (this *SysPermission) SetRoles(roles []SysRole) (result *SysPermission) {
	this.Roles = roles
	return this
}
func (this *SysPermission) GetRoles() (roles []SysRole) {
	return this.Roles
}

type SysPermission struct {
	Id          int64
	Title       string
	Description string
	ParentId    int64
	Sort        int64
	Name        string
	Roles       []SysRole
}
