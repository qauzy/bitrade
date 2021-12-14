
package entity

func (this *PlatformTransaction) SetId(id int64) {
	this.Id = id
}
func (this *PlatformTransaction) GetId() (id int64) {
	return this.Id
}
func (this *PlatformTransaction) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *PlatformTransaction) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *PlatformTransaction) SetBizOrderId(bizOrderId string) {
	this.BizOrderId = bizOrderId
}
func (this *PlatformTransaction) GetBizOrderId() (bizOrderId string) {
	return this.BizOrderId
}
func (this *PlatformTransaction) SetType(type int) {
	this.Type = type
}
func (this *PlatformTransaction) GetType() (type int) {
	return this.Type
}
func (this *PlatformTransaction) SetDirection(direction int) {
	this.Direction = direction
}
func (this *PlatformTransaction) GetDirection() (direction int) {
	return this.Direction
}
func (this *PlatformTransaction) SetDay(day string) {
	this.Day = day
}
func (this *PlatformTransaction) GetDay() (day string) {
	return this.Day
}
func (this *PlatformTransaction) SetTime(time Date) {
	this.Time = time
}
func (this *PlatformTransaction) GetTime() (time Date) {
	return this.Time
}

type PlatformTransaction struct {
	Id         int64
	Amount     BigDecimal
	BizOrderId string
	Type       int
	Direction  int
	Day        string
	Time       Date
}

