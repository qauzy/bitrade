package transform

import (
	"bitrade/core/constant"
	"bitrade/core/entity"
)

func (this *AuthMember) SetId(id int64) (result *AuthMember) {
	this.Id = id
	return this
}
func (this *AuthMember) GetId() (id int64) {
	return this.Id
}
func (this *AuthMember) SetName(name string) (result *AuthMember) {
	this.Name = name
	return this
}
func (this *AuthMember) GetName() (name string) {
	return this.Name
}
func (this *AuthMember) SetRealName(realName string) (result *AuthMember) {
	this.RealName = realName
	return this
}
func (this *AuthMember) GetRealName() (realName string) {
	return this.RealName
}
func (this *AuthMember) SetLocation(location entity.Location) (result *AuthMember) {
	this.Location = location
	return this
}
func (this *AuthMember) GetLocation() (location entity.Location) {
	return this.Location
}
func (this *AuthMember) SetMobilePhone(mobilePhone string) (result *AuthMember) {
	this.MobilePhone = mobilePhone
	return this
}
func (this *AuthMember) GetMobilePhone() (mobilePhone string) {
	return this.MobilePhone
}
func (this *AuthMember) SetEmail(email string) (result *AuthMember) {
	this.Email = email
	return this
}
func (this *AuthMember) GetEmail() (email string) {
	return this.Email
}
func (this *AuthMember) SetMemberLevel(memberLevel constant.MemberLevelEnum) (result *AuthMember) {
	this.MemberLevel = memberLevel
	return this
}
func (this *AuthMember) GetMemberLevel() (memberLevel constant.MemberLevelEnum) {
	return this.MemberLevel
}
func (this *AuthMember) SetStatus(status constant.CommonStatus) (result *AuthMember) {
	this.Status = status
	return this
}
func (this *AuthMember) GetStatus() (status constant.CommonStatus) {
	return this.Status
}
func (this *AuthMember) SetMemberGradeId(memberGradeId int64) (result *AuthMember) {
	this.MemberGradeId = memberGradeId
	return this
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
func ToAuthMember(member *entity.Member) (result *AuthMember) {
	return new(AuthMember).SetId(member.GetId()).SetName(member.GetUsername()).SetRealName(member.GetRealName()).SetLocation(member.GetLocation()).SetMobilePhone(member.GetMobilePhone()).SetEmail(member.GetEmail()).SetMemberLevel(member.GetMemberLevel()).SetStatus(member.GetStatus()).SetMemberGradeId(member.GetMemberGradeId())
}

type AuthMember struct {
	Id            int64
	Name          string
	RealName      string
	Location      entity.Location
	MobilePhone   string
	Email         string
	MemberLevel   constant.MemberLevelEnum
	Status        constant.CommonStatus
	MemberGradeId int64
}
