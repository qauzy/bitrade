package entity

import (
	"bitrade/core/constant/AuditStatus"
	"bitrade/core/enums/CredentialsType"
	"github.com/qauzy/chocolate/xtime"
)

func (this *MemberApplication) SetId(Id int64) (result *MemberApplication) {
	this.Id = Id
	return this
}
func (this *MemberApplication) GetId() (Id int64) {
	return this.Id
}
func (this *MemberApplication) SetRealName(RealName string) (result *MemberApplication) {
	this.RealName = RealName
	return this
}
func (this *MemberApplication) GetRealName() (RealName string) {
	return this.RealName
}
func (this *MemberApplication) SetIdCard(IdCard string) (result *MemberApplication) {
	this.IdCard = IdCard
	return this
}
func (this *MemberApplication) GetIdCard() (IdCard string) {
	return this.IdCard
}
func (this *MemberApplication) SetType(Type CredentialsType.CredentialsType) (result *MemberApplication) {
	this.Type = Type
	return this
}
func (this *MemberApplication) GetType() (Type CredentialsType.CredentialsType) {
	return this.Type
}
func (this *MemberApplication) SetIdentityCardImgFront(IdentityCardImgFront string) (result *MemberApplication) {
	this.IdentityCardImgFront = IdentityCardImgFront
	return this
}
func (this *MemberApplication) GetIdentityCardImgFront() (IdentityCardImgFront string) {
	return this.IdentityCardImgFront
}
func (this *MemberApplication) SetIdentityCardImgReverse(IdentityCardImgReverse string) (result *MemberApplication) {
	this.IdentityCardImgReverse = IdentityCardImgReverse
	return this
}
func (this *MemberApplication) GetIdentityCardImgReverse() (IdentityCardImgReverse string) {
	return this.IdentityCardImgReverse
}
func (this *MemberApplication) SetIdentityCardImgInHand(IdentityCardImgInHand string) (result *MemberApplication) {
	this.IdentityCardImgInHand = IdentityCardImgInHand
	return this
}
func (this *MemberApplication) GetIdentityCardImgInHand() (IdentityCardImgInHand string) {
	return this.IdentityCardImgInHand
}
func (this *MemberApplication) SetAuditStatus(AuditStatus AuditStatus.AuditStatus) (result *MemberApplication) {
	this.AuditStatus = AuditStatus
	return this
}
func (this *MemberApplication) GetAuditStatus() (AuditStatus AuditStatus.AuditStatus) {
	return this.AuditStatus
}
func (this *MemberApplication) SetMember(Member *Member) (result *MemberApplication) {
	this.Member = Member
	return this
}
func (this *MemberApplication) GetMember() (Member *Member) {
	return this.Member
}
func (this *MemberApplication) SetRejectReason(RejectReason string) (result *MemberApplication) {
	this.RejectReason = RejectReason
	return this
}
func (this *MemberApplication) GetRejectReason() (RejectReason string) {
	return this.RejectReason
}
func (this *MemberApplication) SetCreateTime(CreateTime xtime.Xtime) (result *MemberApplication) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberApplication) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *MemberApplication) SetUpdateTime(UpdateTime xtime.Xtime) (result *MemberApplication) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *MemberApplication) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func (this *MemberApplication) SetVideoUrl(VideoUrl string) (result *MemberApplication) {
	this.VideoUrl = VideoUrl
	return this
}
func (this *MemberApplication) GetVideoUrl() (VideoUrl string) {
	return this.VideoUrl
}
func (this *MemberApplication) SetKycStatus(KycStatus int) (result *MemberApplication) {
	this.KycStatus = KycStatus
	return this
}
func (this *MemberApplication) GetKycStatus() (KycStatus int) {
	return this.KycStatus
}
func (this *MemberApplication) SetVideoRandom(VideoRandom string) (result *MemberApplication) {
	this.VideoRandom = VideoRandom
	return this
}
func (this *MemberApplication) GetVideoRandom() (VideoRandom string) {
	return this.VideoRandom
}
func NewMemberApplication(id int64, realName string, idCard string, oType CredentialsType.CredentialsType, identityCardImgFront string, identityCardImgReverse string, identityCardImgInHand string, auditStatus AuditStatus.AuditStatus, member *Member, rejectReason string, createTime xtime.Xtime, updateTime xtime.Xtime, videoUrl string, kycStatus int, videoRandom string) (ret *MemberApplication) {
	ret = new(MemberApplication)
	ret.Id = id
	ret.RealName = realName
	ret.IdCard = idCard
	ret.Type = oType
	ret.IdentityCardImgFront = identityCardImgFront
	ret.IdentityCardImgReverse = identityCardImgReverse
	ret.IdentityCardImgInHand = identityCardImgInHand
	ret.AuditStatus = auditStatus
	ret.Member = member
	ret.RejectReason = rejectReason
	ret.CreateTime = createTime
	ret.UpdateTime = updateTime
	ret.VideoUrl = videoUrl
	ret.KycStatus = kycStatus
	ret.VideoRandom = videoRandom
	return
}

type MemberApplication struct {
	Id                     int64                           `gorm:"column:id" json:"id"`
	RealName               string                          `gorm:"column:real_name" json:"realName"`
	IdCard                 string                          `gorm:"column:id_card" json:"idCard"`
	Type                   CredentialsType.CredentialsType `gorm:"column:type" json:"type"`
	IdentityCardImgFront   string                          `gorm:"column:identity_card_img_front" json:"identityCardImgFront"`
	IdentityCardImgReverse string                          `gorm:"column:identity_card_img_reverse" json:"identityCardImgReverse"`
	IdentityCardImgInHand  string                          `gorm:"column:identity_card_img_in_hand" json:"identityCardImgInHand"`
	AuditStatus            AuditStatus.AuditStatus         `gorm:"column:audit_status" json:"auditStatus"`
	Member                 *Member                         `gorm:"column:member" json:"member"`
	RejectReason           string                          `gorm:"column:reject_reason" json:"rejectReason"`
	CreateTime             xtime.Xtime                     `gorm:"column:create_time" json:"createTime"`
	UpdateTime             xtime.Xtime                     `gorm:"column:update_time" json:"updateTime"`
	VideoUrl               string                          `gorm:"column:video_url" json:"videoUrl"`
	KycStatus              int                             `gorm:"column:kyc_status" json:"kycStatus"`
	VideoRandom            string                          `gorm:"column:video_random" json:"videoRandom"`
}
