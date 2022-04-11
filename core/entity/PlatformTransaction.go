package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *PlatformTransaction) SetId(Id int64) (result *PlatformTransaction) {
	this.Id = Id
	return this
}
func (this *PlatformTransaction) GetId() (Id int64) {
	return this.Id
}
func (this *PlatformTransaction) SetAmount(Amount math.BigDecimal) (result *PlatformTransaction) {
	this.Amount = Amount
	return this
}
func (this *PlatformTransaction) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *PlatformTransaction) SetBizOrderId(BizOrderId string) (result *PlatformTransaction) {
	this.BizOrderId = BizOrderId
	return this
}
func (this *PlatformTransaction) GetBizOrderId() (BizOrderId string) {
	return this.BizOrderId
}
func (this *PlatformTransaction) SetType(Type int) (result *PlatformTransaction) {
	this.Type = Type
	return this
}
func (this *PlatformTransaction) GetType() (Type int) {
	return this.Type
}
func (this *PlatformTransaction) SetDirection(Direction int) (result *PlatformTransaction) {
	this.Direction = Direction
	return this
}
func (this *PlatformTransaction) GetDirection() (Direction int) {
	return this.Direction
}
func (this *PlatformTransaction) SetDay(Day string) (result *PlatformTransaction) {
	this.Day = Day
	return this
}
func (this *PlatformTransaction) GetDay() (Day string) {
	return this.Day
}
func (this *PlatformTransaction) SetTime(Time xtime.Xtime) (result *PlatformTransaction) {
	this.Time = Time
	return this
}
func (this *PlatformTransaction) GetTime() (Time xtime.Xtime) {
	return this.Time
}
func NewPlatformTransaction(id int64, amount math.BigDecimal, bizOrderId string, oType int, direction int, day string, time xtime.Xtime) (ret *PlatformTransaction) {
	ret = new(PlatformTransaction)
	ret.Id = id
	ret.Amount = amount
	ret.BizOrderId = bizOrderId
	ret.Type = oType
	ret.Direction = direction
	ret.Day = day
	ret.Time = time
	return
}

type PlatformTransaction struct {
	Id         int64           `gorm:"column:id" json:"id"`
	Amount     math.BigDecimal `gorm:"column:amount" json:"amount"`
	BizOrderId string          `gorm:"column:biz_order_id" json:"bizOrderId"`
	Type       int             `gorm:"column:type" json:"type"`
	Direction  int             `gorm:"column:direction" json:"direction"`
	Day        string          `gorm:"column:day" json:"day"`
	Time       xtime.Xtime     `gorm:"column:time" json:"time"`
}
