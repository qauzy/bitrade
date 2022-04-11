package entity

import "github.com/qauzy/chocolate/xtime"

func (this *Airdrop) SetId(Id int64) (result *Airdrop) {
	this.Id = Id
	return this
}
func (this *Airdrop) GetId() (Id int64) {
	return this.Id
}
func (this *Airdrop) SetCreateTime(CreateTime xtime.Xtime) (result *Airdrop) {
	this.CreateTime = CreateTime
	return this
}
func (this *Airdrop) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *Airdrop) SetAdmin(Admin *Admin) (result *Airdrop) {
	this.Admin = Admin
	return this
}
func (this *Airdrop) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func (this *Airdrop) SetErrorIndex(ErrorIndex int) (result *Airdrop) {
	this.ErrorIndex = ErrorIndex
	return this
}
func (this *Airdrop) GetErrorIndex() (ErrorIndex int) {
	return this.ErrorIndex
}
func (this *Airdrop) SetSuccessCount(SuccessCount int) (result *Airdrop) {
	this.SuccessCount = SuccessCount
	return this
}
func (this *Airdrop) GetSuccessCount() (SuccessCount int) {
	return this.SuccessCount
}
func (this *Airdrop) SetFileName(FileName string) (result *Airdrop) {
	this.FileName = FileName
	return this
}
func (this *Airdrop) GetFileName() (FileName string) {
	return this.FileName
}
func (this *Airdrop) SetStatus(Status int) (result *Airdrop) {
	this.Status = Status
	return this
}
func (this *Airdrop) GetStatus() (Status int) {
	return this.Status
}
func (this *Airdrop) SetErrorMsg(ErrorMsg string) (result *Airdrop) {
	this.ErrorMsg = ErrorMsg
	return this
}
func (this *Airdrop) GetErrorMsg() (ErrorMsg string) {
	return this.ErrorMsg
}
func NewAirdrop(id int64, createTime xtime.Xtime, admin *Admin, errorIndex int, successCount int, fileName string, status int, errorMsg string) (ret *Airdrop) {
	ret = new(Airdrop)
	ret.Id = id
	ret.CreateTime = createTime
	ret.Admin = admin
	ret.ErrorIndex = errorIndex
	ret.SuccessCount = successCount
	ret.FileName = fileName
	ret.Status = status
	ret.ErrorMsg = errorMsg
	return
}

type Airdrop struct {
	Id           int64       `gorm:"column:id" json:"id"`
	CreateTime   xtime.Xtime `gorm:"column:create_time" json:"createTime"`
	Admin        *Admin      `gorm:"column:admin" json:"admin"`
	ErrorIndex   int         `gorm:"column:error_index" json:"errorIndex"`
	SuccessCount int         `gorm:"column:success_count" json:"successCount"`
	FileName     string      `gorm:"column:file_name" json:"fileName"`
	Status       int         `gorm:"column:status" json:"status"`
	ErrorMsg     string      `gorm:"column:error_msg" json:"errorMsg"`
}
