
package entity

func (this *IntegrationRecord) SetId(id int64) {
	this.Id = id
}
func (this *IntegrationRecord) GetId() (id int64) {
	return this.Id
}
func (this *IntegrationRecord) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *IntegrationRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *IntegrationRecord) SetType(type IntegrationRecordType) {
	this.Type = type
}
func (this *IntegrationRecord) GetType() (type IntegrationRecordType) {
	return this.Type
}
func (this *IntegrationRecord) SetAmount(amount int64) {
	this.Amount = amount
}
func (this *IntegrationRecord) GetAmount() (amount int64) {
	return this.Amount
}
func (this *IntegrationRecord) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *IntegrationRecord) GetCreateTime() (createTime Date) {
	return this.CreateTime
}

type IntegrationRecord struct {
	Id         int64
	MemberId   int64
	Type       IntegrationRecordType
	Amount     int64
	CreateTime Date
}

