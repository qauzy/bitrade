package entity

import "bitrade/core/constant/PromotionLevel"

func (this *MemberPromotion) SetId(Id int64) (result *MemberPromotion) {
	this.Id = Id
	return this
}
func (this *MemberPromotion) GetId() (Id int64) {
	return this.Id
}
func (this *MemberPromotion) SetInviterId(InviterId int64) (result *MemberPromotion) {
	this.InviterId = InviterId
	return this
}
func (this *MemberPromotion) GetInviterId() (InviterId int64) {
	return this.InviterId
}
func (this *MemberPromotion) SetInviteesId(InviteesId int64) (result *MemberPromotion) {
	this.InviteesId = InviteesId
	return this
}
func (this *MemberPromotion) GetInviteesId() (InviteesId int64) {
	return this.InviteesId
}
func (this *MemberPromotion) SetLevel(Level PromotionLevel.PromotionLevel) (result *MemberPromotion) {
	this.Level = Level
	return this
}
func (this *MemberPromotion) GetLevel() (Level PromotionLevel.PromotionLevel) {
	return this.Level
}
func NewMemberPromotion(id int64, inviterId int64, inviteesId int64, level PromotionLevel.PromotionLevel) (ret *MemberPromotion) {
	ret = new(MemberPromotion)
	ret.Id = id
	ret.InviterId = inviterId
	ret.InviteesId = inviteesId
	ret.Level = level
	return
}

type MemberPromotion struct {
	Id         int64                         `gorm:"column:id" json:"id"`
	InviterId  int64                         `gorm:"column:inviter_id" json:"inviterId"`
	InviteesId int64                         `gorm:"column:invitees_id" json:"inviteesId"`
	Level      PromotionLevel.PromotionLevel `gorm:"column:level" json:"level"`
}
