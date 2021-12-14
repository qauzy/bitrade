package entity

func (this *Admin) SetId(id int64) {
	this.Id = id
}
func (this *Admin) GetId() (id int64) {
	return this.Id
}
func (this *Admin) SetUsername(username string) {
	this.Username = username
}
func (this *Admin) GetUsername() (username string) {
	return this.Username
}
func (this *Admin) SetPassword(password string) {
	this.Password = password
}
func (this *Admin) GetPassword() (password string) {
	return this.Password
}
func (this *Admin) SetEnable(enable CommonStatus) {
	this.Enable = enable
}
func (this *Admin) GetEnable() (enable CommonStatus) {
	return this.Enable
}
func (this *Admin) SetLastLoginTime(lastLoginTime Date) {
	this.LastLoginTime = lastLoginTime
}
func (this *Admin) GetLastLoginTime() (lastLoginTime Date) {
	return this.LastLoginTime
}
func (this *Admin) SetLastLoginIp(lastLoginIp string) {
	this.LastLoginIp = lastLoginIp
}
func (this *Admin) GetLastLoginIp() (lastLoginIp string) {
	return this.LastLoginIp
}
func (this *Admin) SetRoleId(roleId int64) {
	this.RoleId = roleId
}
func (this *Admin) GetRoleId() (roleId int64) {
	return this.RoleId
}
func (this *Admin) SetDepartment(department Department) {
	this.Department = department
}
func (this *Admin) GetDepartment() (department Department) {
	return this.Department
}
func (this *Admin) SetRealName(realName string) {
	this.RealName = realName
}
func (this *Admin) GetRealName() (realName string) {
	return this.RealName
}
func (this *Admin) SetMobilePhone(mobilePhone string) {
	this.MobilePhone = mobilePhone
}
func (this *Admin) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *Admin) SetQq(qq string) {
	this.Qq = qq
}
func (this *Admin) GetQq() (qq string) {
	return this.Qq
}
func (this *Admin) SetEmail(email string) {
	this.Email = email
}
func (this *Admin) GetEmail() (email string) {
	return this.Email
}
func (this *Admin) SetAvatar(avatar string) {
	this.Avatar = avatar
}
func (this *Admin) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *Admin) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *Admin) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *Admin) SetRoleName(roleName string) {
	this.RoleName = roleName
}
func (this *Admin) GetRoleName() (roleName string) {
	return this.RoleName
}
func (this *Admin) SetGoogleKey(googleKey string) {
	this.GoogleKey = googleKey
}
func (this *Admin) GetGoogleKey() (googleKey string) {
	return this.GoogleKey
}
func (this *Admin) SetGoogleState(googleState int64) {
	this.GoogleState = googleState
}
func (this *Admin) GetGoogleState() (googleState int64) {
	return this.GoogleState
}
func (this *Admin) SetGoogleDate(googleDate Date) {
	this.GoogleDate = googleDate
}
func (this *Admin) GetGoogleDate() (googleDate Date) {
	return this.GoogleDate
}
func (this *Admin) SetAreaCode(areaCode string) {
	this.AreaCode = areaCode
}
func (this *Admin) GetAreaCode() (areaCode string) {
	return this.AreaCode
}

type Admin struct {
	Id            int64
	Username      string
	Password      string
	Enable        CommonStatus
	LastLoginTime Date
	LastLoginIp   string
	RoleId        int64
	Department    Department
	RealName      string
	MobilePhone   string
	Qq            string
	Email         string
	Avatar        string
	Status        CommonStatus
	RoleName      string
	GoogleKey     string
	GoogleState   int64
	GoogleDate    Date
	AreaCode      string
}
