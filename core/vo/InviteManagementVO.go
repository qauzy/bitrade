package vo

func (this *InviteManagementVO) SetId(id int64) (result *InviteManagementVO) {
	this.Id = id
	return this
}
func (this *InviteManagementVO) GetId() (id int64) {
	return this.Id
}
func (this *InviteManagementVO) SetRealName(realName string) (result *InviteManagementVO) {
	this.RealName = realName
	return this
}
func (this *InviteManagementVO) GetRealName() (realName string) {
	return this.RealName
}
func (this *InviteManagementVO) SetMobilePhone(mobilePhone string) (result *InviteManagementVO) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *InviteManagementVO) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *InviteManagementVO) SetEmail(email string) (result *InviteManagementVO) {
	this.Email = email
	return this
}
func (this *InviteManagementVO) GetEmail() (email string) {
	return this.Email
}
func (this *InviteManagementVO) SetPageNo(pageNo int64) (result *InviteManagementVO) {
	this.PageNo = pageNo
	return this
}
func (this *InviteManagementVO) GetPageNo() (pageNo int64) {
	return this.PageNo
}
func (this *InviteManagementVO) SetPageNumber(pageNumber int64) (result *InviteManagementVO) {
	this.PageNumber = pageNumber
	return this
}
func (this *InviteManagementVO) GetPageNumber() (pageNumber int64) {
	return this.PageNumber
}
func (this *InviteManagementVO) SetPageSize(pageSize int64) (result *InviteManagementVO) {
	this.PageSize = pageSize
	return this
}
func (this *InviteManagementVO) GetPageSize() (pageSize int64) {
	return this.PageSize
}

type InviteManagementVO struct {
	Id          int64
	RealName    string
	MobilePhone string
	Email       string
	PageNo      int64
	PageNumber  int64
	PageSize    int64
}
