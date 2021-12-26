package entity

func (this *MemberPromotion) SetId(id int64) (result *MemberPromotion) {
	this.Id = id
	return this
}
func (this *MemberPromotion) GetId() (id int64) {
	return this.Id
}
func (this *MemberPromotion) SetInviterId(inviterId int64) (result *MemberPromotion) {
	this.InviterId = inviterId
	return this
}
func (this *MemberPromotion) GetInviterId() (inviterId int64) {
	return this.InviterId
}
func (this *MemberPromotion) SetInviteesId(inviteesId int64) (result *MemberPromotion) {
	this.InviteesId = inviteesId
	return this
}
func (this *MemberPromotion) GetInviteesId() (inviteesId int64) {
	return this.InviteesId
}
func (this *MemberPromotion) SetLevel(level PromotionLevel.PromotionLevel) (result *MemberPromotion) {
	this.Level = level
	return this
}
func (this *MemberPromotion) GetLevel() (level PromotionLevel.PromotionLevel) {
	return this.Level
}

type MemberPromotion struct {
	Id         int64
	InviterId  int64
	InviteesId int64
	Level      PromotionLevel.PromotionLevel
}
