package entity

func (this *AdminAccessLog) SetId(id int64) {
	this.Id = id
}
func (this *AdminAccessLog) GetId() (id int64) {
	return this.Id
}
func (this *AdminAccessLog) SetAdminId(adminId int64) {
	this.AdminId = adminId
}
func (this *AdminAccessLog) GetAdminId() (adminId int64) {
	return this.AdminId
}
func (this *AdminAccessLog) SetAdminName(adminName string) {
	this.AdminName = adminName
}
func (this *AdminAccessLog) GetAdminName() (adminName string) {
	return this.AdminName
}
func (this *AdminAccessLog) SetUri(uri string) {
	this.Uri = uri
}
func (this *AdminAccessLog) GetUri() (uri string) {
	return this.Uri
}
func (this *AdminAccessLog) SetModule(module AdminModule) {
	this.Module = module
}
func (this *AdminAccessLog) GetModule() (module AdminModule) {
	return this.Module
}
func (this *AdminAccessLog) SetOperation(operation string) {
	this.Operation = operation
}
func (this *AdminAccessLog) GetOperation() (operation string) {
	return this.Operation
}
func (this *AdminAccessLog) SetAccessIp(accessIp string) {
	this.AccessIp = accessIp
}
func (this *AdminAccessLog) GetAccessIp() (accessIp string) {
	return this.AccessIp
}
func (this *AdminAccessLog) SetAccessMethod(accessMethod string) {
	this.AccessMethod = accessMethod
}
func (this *AdminAccessLog) GetAccessMethod() (accessMethod string) {
	return this.AccessMethod
}
func (this *AdminAccessLog) SetAccessTime(accessTime Date) {
	this.AccessTime = accessTime
}
func (this *AdminAccessLog) GetAccessTime() (accessTime Date) {
	return this.AccessTime
}

type AdminAccessLog struct {
	Id           int64
	AdminId      int64
	AdminName    string
	Uri          string
	Module       AdminModule
	Operation    string
	AccessIp     string
	AccessMethod string
	AccessTime   Date
}
