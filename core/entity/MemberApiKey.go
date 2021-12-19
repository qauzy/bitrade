package entity

import "time"

func (this *MemberApiKey) SetId(id int64) (result *MemberApiKey) {
	this.Id = id
	return this
}
func (this *MemberApiKey) GetId() (id int64) {
	return this.Id
}
func (this *MemberApiKey) SetMemberId(memberId int64) (result *MemberApiKey) {
	this.MemberId = memberId
	return this
}
func (this *MemberApiKey) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberApiKey) SetApiKey(apiKey string) (result *MemberApiKey) {
	this.ApiKey = apiKey
	return this
}
func (this *MemberApiKey) GetApiKey() (apiKey string) {
	return this.ApiKey
}
func (this *MemberApiKey) SetSecretKey(secretKey string) (result *MemberApiKey) {
	this.SecretKey = secretKey
	return this
}
func (this *MemberApiKey) GetSecretKey() (secretKey string) {
	return this.SecretKey
}
func (this *MemberApiKey) SetBindIp(bindIp string) (result *MemberApiKey) {
	this.BindIp = bindIp
	return this
}
func (this *MemberApiKey) GetBindIp() (bindIp string) {
	return this.BindIp
}
func (this *MemberApiKey) SetApiName(apiName string) (result *MemberApiKey) {
	this.ApiName = apiName
	return this
}
func (this *MemberApiKey) GetApiName() (apiName string) {
	return this.ApiName
}
func (this *MemberApiKey) SetRemark(remark string) (result *MemberApiKey) {
	this.Remark = remark
	return this
}
func (this *MemberApiKey) GetRemark() (remark string) {
	return this.Remark
}
func (this *MemberApiKey) SetExpireTime(expireTime time.Time) (result *MemberApiKey) {
	this.ExpireTime = expireTime
	return this
}
func (this *MemberApiKey) GetExpireTime() (expireTime time.Time) {
	return this.ExpireTime
}
func (this *MemberApiKey) SetCreateTime(createTime time.Time) (result *MemberApiKey) {
	this.CreateTime = createTime
	return this
}
func (this *MemberApiKey) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *MemberApiKey) SetCode(code string) (result *MemberApiKey) {
	this.Code = code
	return this
}
func (this *MemberApiKey) GetCode() (code string) {
	return this.Code
}

func NewMemberApiKey(memberId int64, apiKey string, secretKey string, bindIp string, apiName string, remark string, expireTime time.Time, id int64, createTime time.Time) (this *MemberApiKey) {
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

//func NewMemberApiKey(memberId int64, apiKey string, bindIp string, apiName string, remark string, expireTime time.Time, id int64, createTime time.Time) (this *MemberApiKey) {
//	this = new(MemberApiKey)
//	this.Id = id
//	this.MemberId = memberId
//	this.ApiKey = apiKey
//	this.BindIp = bindIp
//	this.ApiName = apiName
//	this.Remark = remark
//	this.ExpireTime = expireTime
//	this.CreateTime = createTime
//	return
//}
//func NewMemberApiKey(memberId int64, bindIp string, apiName string, remark string, createTime time.Time) (this *MemberApiKey) {
//	this = new(MemberApiKey)
//	this.MemberId = memberId
//	this.BindIp = bindIp
//	this.ApiName = apiName
//	this.Remark = remark
//	this.CreateTime = createTime
//	return
//}

type MemberApiKey struct {
	Id         int64
	MemberId   int64
	ApiKey     string
	SecretKey  string
	BindIp     string
	ApiName    string
	Remark     string
	ExpireTime time.Time
	CreateTime time.Time
	Code       string
}
