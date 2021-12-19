package vo

import "bitrade/core/constant"

func (this *IntegrationRecordVO) SetUserId(userId int64) (result *IntegrationRecordVO) {
	this.UserId = userId
	return this
}
func (this *IntegrationRecordVO) GetUserId() (userId int64) {
	return this.UserId
}
func (this *IntegrationRecordVO) SetType(oType constant.IntegrationRecordType) (result *IntegrationRecordVO) {
	this.Type = oType
	return this
}
func (this *IntegrationRecordVO) GetType() (oType constant.IntegrationRecordType) {
	return this.Type
}
func (this *IntegrationRecordVO) SetCreateStartTime(createStartTime string) (result *IntegrationRecordVO) {
	this.CreateStartTime = createStartTime
	return this
}
func (this *IntegrationRecordVO) GetCreateStartTime() (createStartTime string) {
	return this.CreateStartTime
}
func (this *IntegrationRecordVO) SetCreateEndTime(createEndTime string) (result *IntegrationRecordVO) {
	this.CreateEndTime = createEndTime
	return this
}
func (this *IntegrationRecordVO) GetCreateEndTime() (createEndTime string) {
	return this.CreateEndTime
}

type IntegrationRecordVO struct {
	UserId          int64
	Type            constant.IntegrationRecordType
	CreateStartTime string
	CreateEndTime   string
}
