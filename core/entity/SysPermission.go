package entity

import "github.com/qauzy/chocolate/lists/arraylist"

func (this *SysPermission) SetId(Id int64) (result *SysPermission) {
	this.Id = Id
	return this
}
func (this *SysPermission) GetId() (Id int64) {
	return this.Id
}
func (this *SysPermission) SetTitle(Title string) (result *SysPermission) {
	this.Title = Title
	return this
}
func (this *SysPermission) GetTitle() (Title string) {
	return this.Title
}
func (this *SysPermission) SetDescription(Description string) (result *SysPermission) {
	this.Description = Description
	return this
}
func (this *SysPermission) GetDescription() (Description string) {
	return this.Description
}
func (this *SysPermission) SetParentId(ParentId int64) (result *SysPermission) {
	this.ParentId = ParentId
	return this
}
func (this *SysPermission) GetParentId() (ParentId int64) {
	return this.ParentId
}
func (this *SysPermission) SetSort(Sort int) (result *SysPermission) {
	this.Sort = Sort
	return this
}
func (this *SysPermission) GetSort() (Sort int) {
	return this.Sort
}
func (this *SysPermission) SetName(Name string) (result *SysPermission) {
	this.Name = Name
	return this
}
func (this *SysPermission) GetName() (Name string) {
	return this.Name
}
func (this *SysPermission) SetRoles(Roles arraylist.List[SysRole]) (result *SysPermission) {
	this.Roles = Roles
	return this
}
func (this *SysPermission) GetRoles() (Roles arraylist.List[SysRole]) {
	return this.Roles
}
func NewSysPermission(id int64, title string, description string, parentId int64, sort int, name string, roles arraylist.List[SysRole]) (ret *SysPermission) {
	ret = new(SysPermission)
	ret.Id = id
	ret.Title = title
	ret.Description = description
	ret.ParentId = parentId
	ret.Sort = sort
	ret.Name = name
	ret.Roles = roles
	return
}

type SysPermission struct {
	Id          int64                   `gorm:"column:id" json:"id"`
	Title       string                  `gorm:"column:title" json:"title"`
	Description string                  `gorm:"column:description" json:"description"`
	ParentId    int64                   `gorm:"column:parent_id" json:"parentId"`
	Sort        int                     `gorm:"column:sort" json:"sort"`
	Name        string                  `gorm:"column:name" json:"name"`
	Roles       arraylist.List[SysRole] `gorm:"column:roles" json:"roles"`
}
