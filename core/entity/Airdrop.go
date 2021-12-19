package entity

import "time"

func (this *Airdrop) SetId(id int64) (result *Airdrop) {
	this.Id = id
	return this
}
func (this *Airdrop) GetId() (id int64) {
	return this.Id
}
func (this *Airdrop) SetCreateTime(createTime time.Time) (result *Airdrop) {
	this.CreateTime = createTime
	return this
}
func (this *Airdrop) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *Airdrop) SetAdmin(admin Admin) (result *Airdrop) {
	this.Admin = admin
	return this
}
func (this *Airdrop) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *Airdrop) SetErrorIndex(errorIndex int64) (result *Airdrop) {
	this.ErrorIndex = errorIndex
	return this
}
func (this *Airdrop) GetErrorIndex() (errorIndex int64) {
	return this.ErrorIndex
}
func (this *Airdrop) SetSuccessCount(successCount int64) (result *Airdrop) {
	this.SuccessCount = successCount
	return this
}
func (this *Airdrop) GetSuccessCount() (successCount int64) {
	return this.SuccessCount
}
func (this *Airdrop) SetFileName(fileName string) (result *Airdrop) {
	this.FileName = fileName
	return this
}
func (this *Airdrop) GetFileName() (fileName string) {
	return this.FileName
}
func (this *Airdrop) SetStatus(status int64) (result *Airdrop) {
	this.Status = status
	return this
}
func (this *Airdrop) GetStatus() (status int64) {
	return this.Status
}
func (this *Airdrop) SetErrorMsg(errorMsg string) (result *Airdrop) {
	this.ErrorMsg = errorMsg
	return this
}
func (this *Airdrop) GetErrorMsg() (errorMsg string) {
	return this.ErrorMsg
}

type Airdrop struct {
	Id           int64
	CreateTime   time.Time
	Admin        Admin
	ErrorIndex   int64
	SuccessCount int64
	FileName     string
	Status       int64
	ErrorMsg     string
}
