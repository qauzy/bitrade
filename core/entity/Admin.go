package entity

import (
	"bitrade/core/constant/CommonStatus"
	"time"
)

func (this *Admin) SetId(id int64) (result *Admin) {
	this.Id = id
	return this
}
func (this *Admin) GetId() (id int64) {
	return this.Id
}
func (this *Admin) SetUsername(username string) (result *Admin) {
	this.Username = username
	return this
}
func (this *Admin) GetUsername() (username string) {
	return this.Username
}
func (this *Admin) SetPassword(password string) (result *Admin) {
	this.Password = password
	return this
}
func (this *Admin) GetPassword() (password string) {
	return this.Password
}
func (this *Admin) SetEnable(enable CommonStatus.CommonStatus) (result *Admin) {
	this.Enable = enable
	return this
}
func (this *Admin) GetEnable() (enable CommonStatus.CommonStatus) {
	return this.Enable
}
func (this *Admin) SetLastLoginTime(lastLoginTime time.Time) (result *Admin) {
	this.LastLoginTime = lastLoginTime
	return this
}
func (this *Admin) GetLastLoginTime() (lastLoginTime time.Time) {
	return this.LastLoginTime
}
func (this *Admin) SetLastLoginIp(lastLoginIp string) (result *Admin) {
	this.LastLoginIp = lastLoginIp
	return this
}
func (this *Admin) GetLastLoginIp() (lastLoginIp string) {
	return this.LastLoginIp
}
func (this *Admin) SetRoleId(roleId int64) (result *Admin) {
	this.RoleId = roleId
	return this
}
func (this *Admin) GetRoleId() (roleId int64) {
	return this.RoleId
}
func (this *Admin) SetDepartment(department *Department) (result *Admin) {
	this.Department = department
	return this
}
func (this *Admin) GetDepartment() (department *Department) {
	return this.Department
}
func (this *Admin) SetRealName(realName string) (result *Admin) {
	this.RealName = realName
	return this
}
func (this *Admin) GetRealName() (realName string) {
	return this.RealName
}
func (this *Admin) SetMobilePhone(mobilePhone string) (result *Admin) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *Admin) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *Admin) SetQq(qq string) (result *Admin) {
	this.Qq = qq
	return this
}
func (this *Admin) GetQq() (qq string) {
	return this.Qq
}
func (this *Admin) SetEmail(email string) (result *Admin) {
	this.Email = email
	return this
}
func (this *Admin) GetEmail() (email string) {
	return this.Email
}
func (this *Admin) SetAvatar(avatar string) (result *Admin) {
	this.Avatar = avatar
	return this
}
func (this *Admin) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *Admin) SetStatus(status CommonStatus.CommonStatus) (result *Admin) {
	this.Status = status
	return this
}
func (this *Admin) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *Admin) SetRoleName(roleName string) (result *Admin) {
	this.RoleName = roleName
	return this
}
func (this *Admin) GetRoleName() (roleName string) {
	return this.RoleName
}
func (this *Admin) SetGoogleKey(googleKey string) (result *Admin) {
	this.GoogleKey = googleKey
	return this
}
func (this *Admin) GetGoogleKey() (googleKey string) {
	return this.GoogleKey
}
func (this *Admin) SetGoogleState(googleState int64) (result *Admin) {
	this.GoogleState = googleState
	return this
}
func (this *Admin) GetGoogleState() (googleState int64) {
	return this.GoogleState
}
func (this *Admin) SetGoogleDate(googleDate time.Time) (result *Admin) {
	this.GoogleDate = googleDate
	return this
}
func (this *Admin) GetGoogleDate() (googleDate time.Time) {
	return this.GoogleDate
}
func (this *Admin) SetAreaCode(areaCode string) (result *Admin) {
	this.AreaCode = areaCode
	return this
}
func (this *Admin) GetAreaCode() (areaCode string) {
	return this.AreaCode
}

type Admin struct {
	Id            int64
	Username      string
	Password      string
	Enable        CommonStatus.CommonStatus
	LastLoginTime time.Time
	LastLoginIp   string
	RoleId        int64
	Department    *Department
	RealName      string
	MobilePhone   string
	Qq            string
	Email         string
	Avatar        string
	Status        CommonStatus.CommonStatus
	RoleName      string
	GoogleKey     string
	GoogleState   int64
	GoogleDate    time.Time
	AreaCode      string
}
