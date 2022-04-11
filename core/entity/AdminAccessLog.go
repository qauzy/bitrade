package entity

import (
	"bitrade/core/constant/AdminModule"
	"github.com/qauzy/chocolate/xtime"
)

func (this *AdminAccessLog) SetId(Id int64) (result *AdminAccessLog) {
	this.Id = Id
	return this
}
func (this *AdminAccessLog) GetId() (Id int64) {
	return this.Id
}
func (this *AdminAccessLog) SetAdminId(AdminId int64) (result *AdminAccessLog) {
	this.AdminId = AdminId
	return this
}
func (this *AdminAccessLog) GetAdminId() (AdminId int64) {
	return this.AdminId
}
func (this *AdminAccessLog) SetAdminName(AdminName string) (result *AdminAccessLog) {
	this.AdminName = AdminName
	return this
}
func (this *AdminAccessLog) GetAdminName() (AdminName string) {
	return this.AdminName
}
func (this *AdminAccessLog) SetUri(Uri string) (result *AdminAccessLog) {
	this.Uri = Uri
	return this
}
func (this *AdminAccessLog) GetUri() (Uri string) {
	return this.Uri
}
func (this *AdminAccessLog) SetModule(Module AdminModule.AdminModule) (result *AdminAccessLog) {
	this.Module = Module
	return this
}
func (this *AdminAccessLog) GetModule() (Module AdminModule.AdminModule) {
	return this.Module
}
func (this *AdminAccessLog) SetOperation(Operation string) (result *AdminAccessLog) {
	this.Operation = Operation
	return this
}
func (this *AdminAccessLog) GetOperation() (Operation string) {
	return this.Operation
}
func (this *AdminAccessLog) SetAccessIp(AccessIp string) (result *AdminAccessLog) {
	this.AccessIp = AccessIp
	return this
}
func (this *AdminAccessLog) GetAccessIp() (AccessIp string) {
	return this.AccessIp
}
func (this *AdminAccessLog) SetAccessMethod(AccessMethod string) (result *AdminAccessLog) {
	this.AccessMethod = AccessMethod
	return this
}
func (this *AdminAccessLog) GetAccessMethod() (AccessMethod string) {
	return this.AccessMethod
}
func (this *AdminAccessLog) SetAccessTime(AccessTime xtime.Xtime) (result *AdminAccessLog) {
	this.AccessTime = AccessTime
	return this
}
func (this *AdminAccessLog) GetAccessTime() (AccessTime xtime.Xtime) {
	return this.AccessTime
}
func NewAdminAccessLog(id int64, adminId int64, adminName string, uri string, module AdminModule.AdminModule, operation string, accessIp string, accessMethod string, accessTime xtime.Xtime) (ret *AdminAccessLog) {
	ret = new(AdminAccessLog)
	ret.Id = id
	ret.AdminId = adminId
	ret.AdminName = adminName
	ret.Uri = uri
	ret.Module = module
	ret.Operation = operation
	ret.AccessIp = accessIp
	ret.AccessMethod = accessMethod
	ret.AccessTime = accessTime
	return
}

type AdminAccessLog struct {
	Id           int64                   `gorm:"column:id" json:"id"`
	AdminId      int64                   `gorm:"column:admin_id" json:"adminId"`
	AdminName    string                  `gorm:"column:admin_name" json:"adminName"`
	Uri          string                  `gorm:"column:uri" json:"uri"`
	Module       AdminModule.AdminModule `gorm:"column:module" json:"module"`
	Operation    string                  `gorm:"column:operation" json:"operation"`
	AccessIp     string                  `gorm:"column:access_ip" json:"accessIp"`
	AccessMethod string                  `gorm:"column:access_method" json:"accessMethod"`
	AccessTime   xtime.Xtime             `gorm:"column:access_time" json:"accessTime"`
}
