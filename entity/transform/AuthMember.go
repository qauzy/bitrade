package transform

var SerialVersionUID int64 = -4199550203850153635

func (this *AuthMember) SetId(id int64) {
	this.Id = id
}
func (this *AuthMember) GetId() (id int64) {
	return this.Id
}
func (this *AuthMember) SetName(name string) {
	this.Name = name
}
func (this *AuthMember) GetName() (name string) {
	return this.Name
}
func (this *AuthMember) SetRealName(realName string) {
	this.RealName = realName
}
func (this *AuthMember) GetRealName() (realName string) {
	return this.RealName
}
func (this *AuthMember) SetLocation(location Location) {
	this.Location = location
}
func (this *AuthMember) GetLocation() (location Location) {
	return this.Location
}
func (this *AuthMember) SetMobilePhone(mobilePhone string) {
	this.MobilePhone = mobilePhone
}
func (this *AuthMember) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *AuthMember) SetEmail(email string) {
	this.Email = email
}
func (this *AuthMember) GetEmail() (email string) {
	return this.Email
}
func (this *AuthMember) SetMemberLevel(memberLevel MemberLevelEnum) {
	this.MemberLevel = memberLevel
}
func (this *AuthMember) GetMemberLevel() (memberLevel MemberLevelEnum) {
	return this.MemberLevel
}
func (this *AuthMember) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *AuthMember) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *AuthMember) SetMemberGradeId(memberGradeId int64) {
	this.MemberGradeId = memberGradeId
}
func (this *AuthMember) GetMemberGradeId() (memberGradeId int64) {
	return this.MemberGradeId
}

/**
 * 如需添加信息在{@link #toAuthMember(Member)}方法中添加
 *
 * @param member
 * @return
 */
func ToAuthMember(member Member) (result AuthMember) {
	return AuthMember.Builder().Id(member.GetId()).Name(member.GetUsername()).RealName(member.GetRealName()).Location(member.GetLocation()).MobilePhone(member.GetMobilePhone()).Email(member.GetEmail()).MemberLevel(member.GetMemberLevel()).Status(member.GetStatus()).MemberGradeId(member.GetMemberGradeId()).Build()
}

type AuthMember struct {
	Id            int64
	Name          string
	RealName      string
	Location      Location
	MobilePhone   string
	Email         string
	MemberLevel   MemberLevelEnum
	Status        CommonStatus
	MemberGradeId int64
}
