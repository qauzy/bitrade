package entity

func (this *MemberPromotion) SetId(id int64) {
	this.Id = id
}
func (this *MemberPromotion) GetId() (id int64) {
	return this.Id
}
func (this *MemberPromotion) SetInviterId(inviterId int64) {
	this.InviterId = inviterId
}
func (this *MemberPromotion) GetInviterId() (inviterId int64) {
	return this.InviterId
}
func (this *MemberPromotion) SetInviteesId(inviteesId int64) {
	this.InviteesId = inviteesId
}
func (this *MemberPromotion) GetInviteesId() (inviteesId int64) {
	return this.InviteesId
}
func (this *MemberPromotion) SetLevel(level PromotionLevel) {
	this.Level = level
}
func (this *MemberPromotion) GetLevel() (level PromotionLevel) {
	return this.Level
}

type MemberPromotion struct {
	Id         int64
	InviterId  int64
	InviteesId int64
	Level      PromotionLevel
}
