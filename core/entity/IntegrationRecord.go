package entity

func (this *IntegrationRecord) SetId(id int64) (result *IntegrationRecord) {
	this.Id = id
	return this
}
func (this *IntegrationRecord) GetId() (id int64) {
	return this.Id
}
func (this *IntegrationRecord) SetMemberId(memberId int64) (result *IntegrationRecord) {
	this.MemberId = memberId
	return this
}
func (this *IntegrationRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *IntegrationRecord) SetType(oType IntegrationRecordType.IntegrationRecordType) (result *IntegrationRecord) {
	this.Type = oType
	return this
}
func (this *IntegrationRecord) GetType() (oType IntegrationRecordType.IntegrationRecordType) {
	return this.Type
}
func (this *IntegrationRecord) SetAmount(amount int64) (result *IntegrationRecord) {
	this.Amount = amount
	return this
}
func (this *IntegrationRecord) GetAmount() (amount int64) {
	return this.Amount
}
func (this *IntegrationRecord) SetCreateTime(createTime time.Time) (result *IntegrationRecord) {
	this.CreateTime = createTime
	return this
}
func (this *IntegrationRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}

type IntegrationRecord struct {
	Id         int64
	MemberId   int64
	Type       IntegrationRecordType.IntegrationRecordType
	Amount     int64
	CreateTime time.Time
}
