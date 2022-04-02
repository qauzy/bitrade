package dto

func (this *BaseMemberDTO) SetMemberId(MemberId int64) (result *BaseMemberDTO) {
	this.MemberId = MemberId
	return this
}
func (this *BaseMemberDTO) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *BaseMemberDTO) SetRealName(RealName string) (result *BaseMemberDTO) {
	this.RealName = RealName
	return this
}
func (this *BaseMemberDTO) GetRealName() (RealName string) {
	return this.RealName
}
func (this *BaseMemberDTO) SetEmail(Email string) (result *BaseMemberDTO) {
	this.Email = Email
	return this
}
func (this *BaseMemberDTO) GetEmail() (Email string) {
	return this.Email
}
func (this *BaseMemberDTO) SetMobilePhone(MobilePhone string) (result *BaseMemberDTO) {
	this.MobilePhone = MobilePhone
	return this
}
func (this *BaseMemberDTO) GetMobilePhone() (MobilePhone string) {
	return this.MobilePhone
}
func (this *BaseMemberDTO) SetUsername(Username string) (result *BaseMemberDTO) {
	this.Username = Username
	return this
}
func (this *BaseMemberDTO) GetUsername() (Username string) {
	return this.Username
}

type BaseMemberDTO struct {
	MemberId    int64  `gorm:"column:member_id" json:"memberId"`
	RealName    string `gorm:"column:real_name" json:"realName"`
	Email       string `gorm:"column:email" json:"email"`
	MobilePhone string `gorm:"column:mobile_phone" json:"mobilePhone"`
	Username    string `gorm:"column:username" json:"username"`
}
