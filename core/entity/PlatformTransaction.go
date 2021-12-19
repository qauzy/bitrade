package entity

func (this *PlatformTransaction) SetId(id int64) (result *PlatformTransaction) {
	this.Id = id
	return this
}
func (this *PlatformTransaction) GetId() (id int64) {
	return this.Id
}
func (this *PlatformTransaction) SetAmount(amount math.BigDecimal) (result *PlatformTransaction) {
	this.Amount = amount
	return this
}
func (this *PlatformTransaction) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *PlatformTransaction) SetBizOrderId(bizOrderId string) (result *PlatformTransaction) {
	this.BizOrderId = bizOrderId
	return this
}
func (this *PlatformTransaction) GetBizOrderId() (bizOrderId string) {
	return this.BizOrderId
}
func (this *PlatformTransaction) SetType(oType int) (result *PlatformTransaction) {
	this.Type = oType
	return this
}
func (this *PlatformTransaction) GetType() (oType int) {
	return this.Type
}
func (this *PlatformTransaction) SetDirection(direction int) (result *PlatformTransaction) {
	this.Direction = direction
	return this
}
func (this *PlatformTransaction) GetDirection() (direction int) {
	return this.Direction
}
func (this *PlatformTransaction) SetDay(day string) (result *PlatformTransaction) {
	this.Day = day
	return this
}
func (this *PlatformTransaction) GetDay() (day string) {
	return this.Day
}
func (this *PlatformTransaction) SetTime(time time.Time) (result *PlatformTransaction) {
	this.Time = time
	return this
}
func (this *PlatformTransaction) GetTime() (time time.Time) {
	return this.Time
}

type PlatformTransaction struct {
	Id         int64
	Amount     math.BigDecimal
	BizOrderId string
	Type       int
	Direction  int
	Day        string
	Time       time.Time
}
