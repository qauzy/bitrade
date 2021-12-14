package entity

func (this *MemberApiKey) SetId(id int64) {
	this.Id = id
}
func (this *MemberApiKey) GetId() (id int64) {
	return this.Id
}
func (this *MemberApiKey) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *MemberApiKey) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberApiKey) SetApiKey(apiKey string) {
	this.ApiKey = apiKey
}
func (this *MemberApiKey) GetApiKey() (apiKey string) {
	return this.ApiKey
}
func (this *MemberApiKey) SetSecretKey(secretKey string) {
	this.SecretKey = secretKey
}
func (this *MemberApiKey) GetSecretKey() (secretKey string) {
	return this.SecretKey
}
func (this *MemberApiKey) SetBindIp(bindIp string) {
	this.BindIp = bindIp
}
func (this *MemberApiKey) GetBindIp() (bindIp string) {
	return this.BindIp
}
func (this *MemberApiKey) SetApiName(apiName string) {
	this.ApiName = apiName
}
func (this *MemberApiKey) GetApiName() (apiName string) {
	return this.ApiName
}
func (this *MemberApiKey) SetRemark(remark string) {
	this.Remark = remark
}
func (this *MemberApiKey) GetRemark() (remark string) {
	return this.Remark
}
func (this *MemberApiKey) SetExpireTime(expireTime Date) {
	this.ExpireTime = expireTime
}
func (this *MemberApiKey) GetExpireTime() (expireTime Date) {
	return this.ExpireTime
}
func (this *MemberApiKey) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *MemberApiKey) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *MemberApiKey) SetCode(code string) {
	this.Code = code
}
func (this *MemberApiKey) GetCode() (code string) {
	return this.Code
}
func NewMemberApiKey() (this *MemberApiKey) {
	this = new(MemberApiKey)
	return
}
func NewMemberApiKey(memberId int64, apiKey string, secretKey string, bindIp string, apiName string, remark string, expireTime Date, id int64, createTime Date) (this *MemberApiKey) {
	this = new(MemberApiKey)
	this.Id = id
	this.MemberId = memberId
	this.ApiKey = apiKey
	this.SecretKey = secretKey
	this.BindIp = bindIp
	this.ApiName = apiName
	this.Remark = remark
	this.ExpireTime = expireTime
	this.CreateTime = createTime
	return
}
func NewMemberApiKey(memberId int64, apiKey string, bindIp string, apiName string, remark string, expireTime Date, id int64, createTime Date) (this *MemberApiKey) {
	this = new(MemberApiKey)
	this.Id = id
	this.MemberId = memberId
	this.ApiKey = apiKey
	this.BindIp = bindIp
	this.ApiName = apiName
	this.Remark = remark
	this.ExpireTime = expireTime
	this.CreateTime = createTime
	return
}
func NewMemberApiKey(memberId int64, bindIp string, apiName string, remark string, createTime Date) (this *MemberApiKey) {
	this = new(MemberApiKey)
	this.MemberId = memberId
	this.BindIp = bindIp
	this.ApiName = apiName
	this.Remark = remark
	this.CreateTime = createTime
	return
}

type MemberApiKey struct {
	Id         int64
	MemberId   int64
	ApiKey     string
	SecretKey  string
	BindIp     string
	ApiName    string
	Remark     string
	ExpireTime Date
	CreateTime Date
	Code       string
}
