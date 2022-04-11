package transform

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/entity"
)

func (this *AuthMember) SetId(Id int64) (result *AuthMember) {
	this.Id = Id
	return this
}
func (this *AuthMember) GetId() (Id int64) {
	return this.Id
}
func (this *AuthMember) SetName(Name string) (result *AuthMember) {
	this.Name = Name
	return this
}
func (this *AuthMember) GetName() (Name string) {
	return this.Name
}
func (this *AuthMember) SetRealName(RealName string) (result *AuthMember) {
	this.RealName = RealName
	return this
}
func (this *AuthMember) GetRealName() (RealName string) {
	return this.RealName
}
func (this *AuthMember) SetLocation(Location *entity.Location) (result *AuthMember) {
	this.Location = Location
	return this
}
func (this *AuthMember) GetLocation() (Location *entity.Location) {
	return this.Location
}
func (this *AuthMember) SetMobilePhone(MobilePhone string) (result *AuthMember) {
	this.MobilePhone = MobilePhone
	return this
}
func (this *AuthMember) GetMobilePhone() (MobilePhone string) {
	return this.MobilePhone
}
func (this *AuthMember) SetEmail(Email string) (result *AuthMember) {
	this.Email = Email
	return this
}
func (this *AuthMember) GetEmail() (Email string) {
	return this.Email
}
func (this *AuthMember) SetMemberLevel(MemberLevel MemberLevelEnum.MemberLevelEnum) (result *AuthMember) {
	this.MemberLevel = MemberLevel
	return this
}
func (this *AuthMember) GetMemberLevel() (MemberLevel MemberLevelEnum.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *AuthMember) SetStatus(Status CommonStatus.CommonStatus) (result *AuthMember) {
	this.Status = Status
	return this
}
func (this *AuthMember) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *AuthMember) SetMemberGradeId(MemberGradeId int64) (result *AuthMember) {
	this.MemberGradeId = MemberGradeId
	return this
}
func (this *AuthMember) GetMemberGradeId() (MemberGradeId int64) {
	return this.MemberGradeId
}
func ToAuthMember(member *entity.Member) (result *AuthMember) {
	return new(AuthMember).SetId(member.GetId()).SetName(member.GetUsername()).SetRealName(member.GetRealName()).SetLocation(member.GetLocation()).SetMobilePhone(member.GetMobilePhone()).SetEmail(member.GetEmail()).SetMemberLevel(member.GetMemberLevel()).SetStatus(member.GetStatus()).SetMemberGradeId(member.GetMemberGradeId())
}
func NewAuthMember(id int64, name string, realName string, location *entity.Location, mobilePhone string, email string, memberLevel MemberLevelEnum.MemberLevelEnum, status CommonStatus.CommonStatus, memberGradeId int64) (ret *AuthMember) {
	ret = new(AuthMember)
	ret.Id = id
	ret.Name = name
	ret.RealName = realName
	ret.Location = location
	ret.MobilePhone = mobilePhone
	ret.Email = email
	ret.MemberLevel = memberLevel
	ret.Status = status
	ret.MemberGradeId = memberGradeId
	return
}

type AuthMember struct {
	Id            int64                           `gorm:"column:id" json:"id"`
	Name          string                          `gorm:"column:name" json:"name"`
	RealName      string                          `gorm:"column:real_name" json:"realName"`
	Location      *entity.Location                `gorm:"column:location" json:"location"`
	MobilePhone   string                          `gorm:"column:mobile_phone" json:"mobilePhone"`
	Email         string                          `gorm:"column:email" json:"email"`
	MemberLevel   MemberLevelEnum.MemberLevelEnum `gorm:"column:member_level" json:"memberLevel"`
	Status        CommonStatus.CommonStatus       `gorm:"column:status" json:"status"`
	MemberGradeId int64                           `gorm:"column:member_grade_id" json:"memberGradeId"`
}
