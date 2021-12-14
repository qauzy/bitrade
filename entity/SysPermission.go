package entity

func (this *SysPermission) SetId(id int64) {
	this.Id = id
}
func (this *SysPermission) GetId() (id int64) {
	return this.Id
}
func (this *SysPermission) SetTitle(title string) {
	this.Title = title
}
func (this *SysPermission) GetTitle() (title string) {
	return this.Title
}
func (this *SysPermission) SetDescription(description string) {
	this.Description = description
}
func (this *SysPermission) GetDescription() (description string) {
	return this.Description
}
func (this *SysPermission) SetParentId(parentId int64) {
	this.ParentId = parentId
}
func (this *SysPermission) GetParentId() (parentId int64) {
	return this.ParentId
}
func (this *SysPermission) SetSort(sort int64) {
	this.Sort = sort
}
func (this *SysPermission) GetSort() (sort int64) {
	return this.Sort
}
func (this *SysPermission) SetName(name string) {
	this.Name = name
}
func (this *SysPermission) GetName() (name string) {
	return this.Name
}
func (this *SysPermission) SetRoles(roles []SysRole) {
	this.Roles = roles
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
