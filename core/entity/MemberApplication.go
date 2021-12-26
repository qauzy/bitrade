package entity

func (this *MemberApplication) SetId(id int64) (result *MemberApplication) {
	this.Id = id
	return this
}
func (this *MemberApplication) GetId() (id int64) {
	return this.Id
}
func (this *MemberApplication) SetRealName(realName string) (result *MemberApplication) {
	this.RealName = realName
	return this
}
func (this *MemberApplication) GetRealName() (realName string) {
	return this.RealName
}
func (this *MemberApplication) SetIdCard(idCard string) (result *MemberApplication) {
	this.IdCard = idCard
	return this
}
func (this *MemberApplication) GetIdCard() (idCard string) {
	return this.IdCard
}
func (this *MemberApplication) SetType(oType enums.CredentialsType) (result *MemberApplication) {
	this.Type = oType
	return this
}
func (this *MemberApplication) GetType() (oType enums.CredentialsType) {
	return this.Type
}
func (this *MemberApplication) SetIdentityCardImgFront(identityCardImgFront string) (result *MemberApplication) {
	this.IdentityCardImgFront = identityCardImgFront
	return this
}
func (this *MemberApplication) GetIdentityCardImgFront() (identityCardImgFront string) {
	return this.IdentityCardImgFront
}
func (this *MemberApplication) SetIdentityCardImgReverse(identityCardImgReverse string) (result *MemberApplication) {
	this.IdentityCardImgReverse = identityCardImgReverse
	return this
}
func (this *MemberApplication) GetIdentityCardImgReverse() (identityCardImgReverse string) {
	return this.IdentityCardImgReverse
}
func (this *MemberApplication) SetIdentityCardImgInHand(identityCardImgInHand string) (result *MemberApplication) {
	this.IdentityCardImgInHand = identityCardImgInHand
	return this
}
func (this *MemberApplication) GetIdentityCardImgInHand() (identityCardImgInHand string) {
	return this.IdentityCardImgInHand
}
func (this *MemberApplication) SetAuditStatus(auditStatus AuditStatus.AuditStatus) (result *MemberApplication) {
	this.AuditStatus = auditStatus
	return this
}
func (this *MemberApplication) GetAuditStatus() (auditStatus AuditStatus.AuditStatus) {
	return this.AuditStatus
}
func (this *MemberApplication) SetMember(member Member) (result *MemberApplication) {
	this.Member = member
	return this
}
func (this *MemberApplication) GetMember() (member Member) {
	return this.Member
}
func (this *MemberApplication) SetRejectReason(rejectReason string) (result *MemberApplication) {
	this.RejectReason = rejectReason
	return this
}
func (this *MemberApplication) GetRejectReason() (rejectReason string) {
	return this.RejectReason
}
func (this *MemberApplication) SetCreateTime(createTime time.Time) (result *MemberApplication) {
	this.CreateTime = createTime
	return this
}
func (this *MemberApplication) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *MemberApplication) SetUpdateTime(updateTime time.Time) (result *MemberApplication) {
	this.UpdateTime = updateTime
	return this
}
func (this *MemberApplication) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}
func (this *MemberApplication) SetVideoUrl(videoUrl string) (result *MemberApplication) {
	this.VideoUrl = videoUrl
	return this
}
func (this *MemberApplication) GetVideoUrl() (videoUrl string) {
	return this.VideoUrl
}
func (this *MemberApplication) SetKycStatus(kycStatus int64) (result *MemberApplication) {
	this.KycStatus = kycStatus
	return this
}
func (this *MemberApplication) GetKycStatus() (kycStatus int64) {
	return this.KycStatus
}
func (this *MemberApplication) SetVideoRandom(videoRandom string) (result *MemberApplication) {
	this.VideoRandom = videoRandom
	return this
}
func (this *MemberApplication) GetVideoRandom() (videoRandom string) {
	return this.VideoRandom
}

type MemberApplication struct {
	Id                     int64
	RealName               string
	IdCard                 string
	Type                   enums.CredentialsType
	IdentityCardImgFront   string
	IdentityCardImgReverse string
	IdentityCardImgInHand  string
	AuditStatus            AuditStatus.AuditStatus
	Member                 Member
	RejectReason           string
	CreateTime             time.Time
	UpdateTime             time.Time
	VideoUrl               string
	KycStatus              int64
	VideoRandom            string
}
