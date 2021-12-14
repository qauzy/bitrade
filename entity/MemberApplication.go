
package entity

func (this *MemberApplication) SetId(id int64) {
	this.Id = id
}
func (this *MemberApplication) GetId() (id int64) {
	return this.Id
}
func (this *MemberApplication) SetRealName(realName string) {
	this.RealName = realName
}
func (this *MemberApplication) GetRealName() (realName string) {
	return this.RealName
}
func (this *MemberApplication) SetIdCard(idCard string) {
	this.IdCard = idCard
}
func (this *MemberApplication) GetIdCard() (idCard string) {
	return this.IdCard
}
func (this *MemberApplication) SetType(type CredentialsType) {
	this.Type = type
}
func (this *MemberApplication) GetType() (type CredentialsType) {
	return this.Type
}
func (this *MemberApplication) SetIdentityCardImgFront(identityCardImgFront string) {
	this.IdentityCardImgFront = identityCardImgFront
}
func (this *MemberApplication) GetIdentityCardImgFront() (identityCardImgFront string) {
	return this.IdentityCardImgFront
}
func (this *MemberApplication) SetIdentityCardImgReverse(identityCardImgReverse string) {
	this.IdentityCardImgReverse = identityCardImgReverse
}
func (this *MemberApplication) GetIdentityCardImgReverse() (identityCardImgReverse string) {
	return this.IdentityCardImgReverse
}
func (this *MemberApplication) SetIdentityCardImgInHand(identityCardImgInHand string) {
	this.IdentityCardImgInHand = identityCardImgInHand
}
func (this *MemberApplication) GetIdentityCardImgInHand() (identityCardImgInHand string) {
	return this.IdentityCardImgInHand
}
func (this *MemberApplication) SetAuditStatus(auditStatus AuditStatus) {
	this.AuditStatus = auditStatus
}
func (this *MemberApplication) GetAuditStatus() (auditStatus AuditStatus) {
	return this.AuditStatus
}
func (this *MemberApplication) SetMember(member Member) {
	this.Member = member
}
func (this *MemberApplication) GetMember() (member Member) {
	return this.Member
}
func (this *MemberApplication) SetRejectReason(rejectReason string) {
	this.RejectReason = rejectReason
}
func (this *MemberApplication) GetRejectReason() (rejectReason string) {
	return this.RejectReason
}
func (this *MemberApplication) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *MemberApplication) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *MemberApplication) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *MemberApplication) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}
func (this *MemberApplication) SetVideoUrl(videoUrl string) {
	this.VideoUrl = videoUrl
}
func (this *MemberApplication) GetVideoUrl() (videoUrl string) {
	return this.VideoUrl
}
func (this *MemberApplication) SetKycStatus(kycStatus int64) {
	this.KycStatus = kycStatus
}
func (this *MemberApplication) GetKycStatus() (kycStatus int64) {
	return this.KycStatus
}
func (this *MemberApplication) SetVideoRandom(videoRandom string) {
	this.VideoRandom = videoRandom
}
func (this *MemberApplication) GetVideoRandom() (videoRandom string) {
	return this.VideoRandom
}

type MemberApplication struct {
	Id                     int64
	RealName               string
	IdCard                 string
	Type                   CredentialsType
	IdentityCardImgFront   string
	IdentityCardImgReverse string
	IdentityCardImgInHand  string
	AuditStatus            AuditStatus
	Member                 Member
	RejectReason           string
	CreateTime             Date
	UpdateTime             Date
	VideoUrl               string
	KycStatus              int64
	VideoRandom            string
}

