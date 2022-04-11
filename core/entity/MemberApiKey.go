package entity

import "github.com/qauzy/chocolate/xtime"

func (this *MemberApiKey) SetId(Id int64) (result *MemberApiKey) {
	this.Id = Id
	return this
}
func (this *MemberApiKey) GetId() (Id int64) {
	return this.Id
}
func (this *MemberApiKey) SetMemberId(MemberId int64) (result *MemberApiKey) {
	this.MemberId = MemberId
	return this
}
func (this *MemberApiKey) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *MemberApiKey) SetApiKey(ApiKey string) (result *MemberApiKey) {
	this.ApiKey = ApiKey
	return this
}
func (this *MemberApiKey) GetApiKey() (ApiKey string) {
	return this.ApiKey
}
func (this *MemberApiKey) SetSecretKey(SecretKey string) (result *MemberApiKey) {
	this.SecretKey = SecretKey
	return this
}
func (this *MemberApiKey) GetSecretKey() (SecretKey string) {
	return this.SecretKey
}
func (this *MemberApiKey) SetBindIp(BindIp string) (result *MemberApiKey) {
	this.BindIp = BindIp
	return this
}
func (this *MemberApiKey) GetBindIp() (BindIp string) {
	return this.BindIp
}
func (this *MemberApiKey) SetApiName(ApiName string) (result *MemberApiKey) {
	this.ApiName = ApiName
	return this
}
func (this *MemberApiKey) GetApiName() (ApiName string) {
	return this.ApiName
}
func (this *MemberApiKey) SetRemark(Remark string) (result *MemberApiKey) {
	this.Remark = Remark
	return this
}
func (this *MemberApiKey) GetRemark() (Remark string) {
	return this.Remark
}
func (this *MemberApiKey) SetExpireTime(ExpireTime xtime.Xtime) (result *MemberApiKey) {
	this.ExpireTime = ExpireTime
	return this
}
func (this *MemberApiKey) GetExpireTime() (ExpireTime xtime.Xtime) {
	return this.ExpireTime
}
func (this *MemberApiKey) SetCreateTime(CreateTime xtime.Xtime) (result *MemberApiKey) {
	this.CreateTime = CreateTime
	return this
}
func (this *MemberApiKey) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *MemberApiKey) SetCode(Code string) (result *MemberApiKey) {
	this.Code = Code
	return this
}
func (this *MemberApiKey) GetCode() (Code string) {
	return this.Code
}
func NewMemberApiKey() (this *MemberApiKey) {
	this = new(MemberApiKey)
	return
}
func NewMemberApiKeyV9(memberId int64, apiKey string, secretKey string, bindIp string, apiName string, remark string, expireTime xtime.Xtime, id int64, createTime xtime.Xtime) (this *MemberApiKey) {
	this = new(MemberApiKey)
	this.Id = this.Id
	this.MemberId = this.MemberId
	this.ApiKey = this.ApiKey
	this.SecretKey = this.SecretKey
	this.BindIp = this.BindIp
	this.ApiName = this.ApiName
	this.Remark = this.Remark
	this.ExpireTime = this.ExpireTime
	this.CreateTime = this.CreateTime
	return
}
func NewMemberApiKeyV8(memberId int64, apiKey string, bindIp string, apiName string, remark string, expireTime xtime.Xtime, id int64, createTime xtime.Xtime) (this *MemberApiKey) {
	this = new(MemberApiKey)
	this.Id = this.Id
	this.MemberId = this.MemberId
	this.ApiKey = this.ApiKey
	this.BindIp = this.BindIp
	this.ApiName = this.ApiName
	this.Remark = this.Remark
	this.ExpireTime = this.ExpireTime
	this.CreateTime = this.CreateTime
	return
}
func NewMemberApiKeyV5(memberId int64, bindIp string, apiName string, remark string, createTime xtime.Xtime) (this *MemberApiKey) {
	this = new(MemberApiKey)
	this.MemberId = memberId
	this.BindIp = bindIp
	this.ApiName = apiName
	this.Remark = remark
	this.CreateTime = createTime
	return
}
func NewMemberApiKeyEx(id int64, memberId int64, apiKey string, secretKey string, bindIp string, apiName string, remark string, expireTime xtime.Xtime, createTime xtime.Xtime, code string) (ret *MemberApiKey) {
	ret = new(MemberApiKey)
	ret.Id = id
	ret.MemberId = memberId
	ret.ApiKey = apiKey
	ret.SecretKey = secretKey
	ret.BindIp = bindIp
	ret.ApiName = apiName
	ret.Remark = remark
	ret.ExpireTime = expireTime
	ret.CreateTime = createTime
	ret.Code = code
	return
}

type MemberApiKey struct {
	Id         int64       `gorm:"column:id" json:"id"`
	MemberId   int64       `gorm:"column:member_id" json:"memberId"`
	ApiKey     string      `gorm:"column:api_key" json:"apiKey"`
	SecretKey  string      `gorm:"column:secret_key" json:"secretKey"`
	BindIp     string      `gorm:"column:bind_ip" json:"bindIp"`
	ApiName    string      `gorm:"column:api_name" json:"apiName"`
	Remark     string      `gorm:"column:remark" json:"remark"`
	ExpireTime xtime.Xtime `gorm:"column:expire_time" json:"expireTime"`
	CreateTime xtime.Xtime `gorm:"column:create_time" json:"createTime"`
	Code       string      `gorm:"column:code" json:"code"`
}
