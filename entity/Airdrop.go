package entity

func (this *Airdrop) SetId(id int64) {
	this.Id = id
}
func (this *Airdrop) GetId() (id int64) {
	return this.Id
}
func (this *Airdrop) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *Airdrop) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *Airdrop) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *Airdrop) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *Airdrop) SetErrorIndex(errorIndex int64) {
	this.ErrorIndex = errorIndex
}
func (this *Airdrop) GetErrorIndex() (errorIndex int64) {
	return this.ErrorIndex
}
func (this *Airdrop) SetSuccessCount(successCount int64) {
	this.SuccessCount = successCount
}
func (this *Airdrop) GetSuccessCount() (successCount int64) {
	return this.SuccessCount
}
func (this *Airdrop) SetFileName(fileName string) {
	this.FileName = fileName
}
func (this *Airdrop) GetFileName() (fileName string) {
	return this.FileName
}
func (this *Airdrop) SetStatus(status int64) {
	this.Status = status
}
func (this *Airdrop) GetStatus() (status int64) {
	return this.Status
}
func (this *Airdrop) SetErrorMsg(errorMsg string) {
	this.ErrorMsg = errorMsg
}
func (this *Airdrop) GetErrorMsg() (errorMsg string) {
	return this.ErrorMsg
}

type Airdrop struct {
	Id           int64
	CreateTime   Date
	Admin        Admin
	ErrorIndex   int64
	SuccessCount int64
	FileName     string
	Status       int64
	ErrorMsg     string
}
