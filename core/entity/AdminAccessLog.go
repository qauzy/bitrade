package entity

import (
	"bitrade/core/constant/AdminModule"
	"time"
)

func (this *AdminAccessLog) SetId(id int64) (result *AdminAccessLog) {
	this.Id = id
	return this
}
func (this *AdminAccessLog) GetId() (id int64) {
	return this.Id
}
func (this *AdminAccessLog) SetAdminId(adminId int64) (result *AdminAccessLog) {
	this.AdminId = adminId
	return this
}
func (this *AdminAccessLog) GetAdminId() (adminId int64) {
	return this.AdminId
}
func (this *AdminAccessLog) SetAdminName(adminName string) (result *AdminAccessLog) {
	this.AdminName = adminName
	return this
}
func (this *AdminAccessLog) GetAdminName() (adminName string) {
	return this.AdminName
}
func (this *AdminAccessLog) SetUri(uri string) (result *AdminAccessLog) {
	this.Uri = uri
	return this
}
func (this *AdminAccessLog) GetUri() (uri string) {
	return this.Uri
}
func (this *AdminAccessLog) SetModule(module AdminModule.AdminModule) (result *AdminAccessLog) {
	this.Module = module
	return this
}
func (this *AdminAccessLog) GetModule() (module AdminModule.AdminModule) {
	return this.Module
}
func (this *AdminAccessLog) SetOperation(operation string) (result *AdminAccessLog) {
	this.Operation = operation
	return this
}
func (this *AdminAccessLog) GetOperation() (operation string) {
	return this.Operation
}
func (this *AdminAccessLog) SetAccessIp(accessIp string) (result *AdminAccessLog) {
	this.AccessIp = accessIp
	return this
}
func (this *AdminAccessLog) GetAccessIp() (accessIp string) {
	return this.AccessIp
}
func (this *AdminAccessLog) SetAccessMethod(accessMethod string) (result *AdminAccessLog) {
	this.AccessMethod = accessMethod
	return this
}
func (this *AdminAccessLog) GetAccessMethod() (accessMethod string) {
	return this.AccessMethod
}
func (this *AdminAccessLog) SetAccessTime(accessTime time.Time) (result *AdminAccessLog) {
	this.AccessTime = accessTime
	return this
}
func (this *AdminAccessLog) GetAccessTime() (accessTime time.Time) {
	return this.AccessTime
}

type AdminAccessLog struct {
	Id           int64
	AdminId      int64
	AdminName    string
	Uri          string
	Module       AdminModule.AdminModule
	Operation    string
	AccessIp     string
	AccessMethod string
	AccessTime   time.Time
}
