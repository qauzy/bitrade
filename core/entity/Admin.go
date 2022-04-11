package entity

import (
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/chocolate/xtime"
)

func (this *Admin) SetId(Id int64) (result *Admin) {
	this.Id = Id
	return this
}
func (this *Admin) GetId() (Id int64) {
	return this.Id
}
func (this *Admin) SetUsername(Username string) (result *Admin) {
	this.Username = Username
	return this
}
func (this *Admin) GetUsername() (Username string) {
	return this.Username
}
func (this *Admin) SetPassword(Password string) (result *Admin) {
	this.Password = Password
	return this
}
func (this *Admin) GetPassword() (Password string) {
	return this.Password
}
func (this *Admin) SetEnable(Enable CommonStatus.CommonStatus) (result *Admin) {
	this.Enable = Enable
	return this
}
func (this *Admin) GetEnable() (Enable CommonStatus.CommonStatus) {
	return this.Enable
}
func (this *Admin) SetLastLoginTime(LastLoginTime xtime.Xtime) (result *Admin) {
	this.LastLoginTime = LastLoginTime
	return this
}
func (this *Admin) GetLastLoginTime() (LastLoginTime xtime.Xtime) {
	return this.LastLoginTime
}
func (this *Admin) SetLastLoginIp(LastLoginIp string) (result *Admin) {
	this.LastLoginIp = LastLoginIp
	return this
}
func (this *Admin) GetLastLoginIp() (LastLoginIp string) {
	return this.LastLoginIp
}
func (this *Admin) SetRoleId(RoleId int64) (result *Admin) {
	this.RoleId = RoleId
	return this
}
func (this *Admin) GetRoleId() (RoleId int64) {
	return this.RoleId
}
func (this *Admin) SetDepartment(Department *Department) (result *Admin) {
	this.Department = Department
	return this
}
func (this *Admin) GetDepartment() (Department *Department) {
	return this.Department
}
func (this *Admin) SetRealName(RealName string) (result *Admin) {
	this.RealName = RealName
	return this
}
func (this *Admin) GetRealName() (RealName string) {
	return this.RealName
}
func (this *Admin) SetMobilePhone(MobilePhone string) (result *Admin) {
	this.MobilePhone = MobilePhone
	return this
}
func (this *Admin) GetMobilePhone() (MobilePhone string) {
	return this.MobilePhone
}
func (this *Admin) SetQq(Qq string) (result *Admin) {
	this.Qq = Qq
	return this
}
func (this *Admin) GetQq() (Qq string) {
	return this.Qq
}
func (this *Admin) SetEmail(Email string) (result *Admin) {
	this.Email = Email
	return this
}
func (this *Admin) GetEmail() (Email string) {
	return this.Email
}
func (this *Admin) SetAvatar(Avatar string) (result *Admin) {
	this.Avatar = Avatar
	return this
}
func (this *Admin) GetAvatar() (Avatar string) {
	return this.Avatar
}
func (this *Admin) SetStatus(Status CommonStatus.CommonStatus) (result *Admin) {
	this.Status = Status
	return this
}
func (this *Admin) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *Admin) SetRoleName(RoleName string) (result *Admin) {
	this.RoleName = RoleName
	return this
}
func (this *Admin) GetRoleName() (RoleName string) {
	return this.RoleName
}
func (this *Admin) SetGoogleKey(GoogleKey string) (result *Admin) {
	this.GoogleKey = GoogleKey
	return this
}
func (this *Admin) GetGoogleKey() (GoogleKey string) {
	return this.GoogleKey
}
func (this *Admin) SetGoogleState(GoogleState int) (result *Admin) {
	this.GoogleState = GoogleState
	return this
}
func (this *Admin) GetGoogleState() (GoogleState int) {
	return this.GoogleState
}
func (this *Admin) SetGoogleDate(GoogleDate xtime.Xtime) (result *Admin) {
	this.GoogleDate = GoogleDate
	return this
}
func (this *Admin) GetGoogleDate() (GoogleDate xtime.Xtime) {
	return this.GoogleDate
}
func (this *Admin) SetAreaCode(AreaCode string) (result *Admin) {
	this.AreaCode = AreaCode
	return this
}
func (this *Admin) GetAreaCode() (AreaCode string) {
	return this.AreaCode
}
func NewAdmin(id int64, username string, password string, enable CommonStatus.CommonStatus, lastLoginTime xtime.Xtime, lastLoginIp string, roleId int64, department *Department, realName string, mobilePhone string, qq string, email string, avatar string, status CommonStatus.CommonStatus, roleName string, googleKey string, googleState int, googleDate xtime.Xtime, areaCode string) (ret *Admin) {
	ret = new(Admin)
	ret.Id = id
	ret.Username = username
	ret.Password = password
	ret.Enable = enable
	ret.LastLoginTime = lastLoginTime
	ret.LastLoginIp = lastLoginIp
	ret.RoleId = roleId
	ret.Department = department
	ret.RealName = realName
	ret.MobilePhone = mobilePhone
	ret.Qq = qq
	ret.Email = email
	ret.Avatar = avatar
	ret.Status = status
	ret.RoleName = roleName
	ret.GoogleKey = googleKey
	ret.GoogleState = googleState
	ret.GoogleDate = googleDate
	ret.AreaCode = areaCode
	return
}

type Admin struct {
	Id            int64                     `gorm:"column:id" json:"id"`
	Username      string                    `gorm:"column:username" json:"username"`
	Password      string                    `gorm:"column:password" json:"password"`
	Enable        CommonStatus.CommonStatus `gorm:"column:enable" json:"enable"`
	LastLoginTime xtime.Xtime               `gorm:"column:last_login_time" json:"lastLoginTime"`
	LastLoginIp   string                    `gorm:"column:last_login_ip" json:"lastLoginIp"`
	RoleId        int64                     `gorm:"column:role_id" json:"roleId"`
	Department    *Department               `gorm:"column:department" json:"department"`
	RealName      string                    `gorm:"column:real_name" json:"realName"`
	MobilePhone   string                    `gorm:"column:mobile_phone" json:"mobilePhone"`
	Qq            string                    `gorm:"column:qq" json:"qq"`
	Email         string                    `gorm:"column:email" json:"email"`
	Avatar        string                    `gorm:"column:avatar" json:"avatar"`
	Status        CommonStatus.CommonStatus `gorm:"column:status" json:"status"`
	RoleName      string                    `gorm:"column:role_name" json:"roleName"`
	GoogleKey     string                    `gorm:"column:google_key" json:"googleKey"`
	GoogleState   int                       `gorm:"column:google_state" json:"googleState"`
	GoogleDate    xtime.Xtime               `gorm:"column:google_date" json:"googleDate"`
	AreaCode      string                    `gorm:"column:area_code" json:"areaCode"`
}
