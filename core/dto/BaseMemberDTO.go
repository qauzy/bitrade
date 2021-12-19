package dto

func (this *BaseMemberDTO) SetMemberId(memberId int64) (result *BaseMemberDTO) {
	this.MemberId = memberId
	return this
}
func (this *BaseMemberDTO) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *BaseMemberDTO) SetRealName(realName string) (result *BaseMemberDTO) {
	this.RealName = realName
	return this
}
func (this *BaseMemberDTO) GetRealName() (realName string) {
	return this.RealName
}
func (this *BaseMemberDTO) SetEmail(email string) (result *BaseMemberDTO) {
	this.Email = email
	return this
}
func (this *BaseMemberDTO) GetEmail() (email string) {
	return this.Email
}
func (this *BaseMemberDTO) SetMobilePhone(mobilePhone string) (result *BaseMemberDTO) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *BaseMemberDTO) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *BaseMemberDTO) SetUsername(username string) (result *BaseMemberDTO) {
	this.Username = username
	return this
}
func (this *BaseMemberDTO) GetUsername() (username string) {
	return this.Username
}

type BaseMemberDTO struct {
	MemberId    int64
	RealName    string
	Email       string
	MobilePhone string
	Username    string
}
