package entity

import (
	"bitrade/core/constant/SignStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *Sign) SetId(Id int64) (result *Sign) {
	this.Id = Id
	return this
}
func (this *Sign) GetId() (Id int64) {
	return this.Id
}
func (this *Sign) SetCoin(Coin *Coin) (result *Sign) {
	this.Coin = Coin
	return this
}
func (this *Sign) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *Sign) SetAmount(Amount math.BigDecimal) (result *Sign) {
	this.Amount = Amount
	return this
}
func (this *Sign) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *Sign) SetEndDate(EndDate xtime.Xtime) (result *Sign) {
	this.EndDate = EndDate
	return this
}
func (this *Sign) GetEndDate() (EndDate xtime.Xtime) {
	return this.EndDate
}
func (this *Sign) SetCreationTime(CreationTime xtime.Xtime) (result *Sign) {
	this.CreationTime = CreationTime
	return this
}
func (this *Sign) GetCreationTime() (CreationTime xtime.Xtime) {
	return this.CreationTime
}
func (this *Sign) SetStatus(Status SignStatus.SignStatus) (result *Sign) {
	this.Status = Status
	return this
}
func (this *Sign) GetStatus() (Status SignStatus.SignStatus) {
	return this.Status
}
func NewSign(id int64, coin *Coin, amount math.BigDecimal, endDate xtime.Xtime, creationTime xtime.Xtime, status SignStatus.SignStatus) (ret *Sign) {
	ret = new(Sign)
	ret.Id = id
	ret.Coin = coin
	ret.Amount = amount
	ret.EndDate = endDate
	ret.CreationTime = creationTime
	ret.Status = status
	return
}

type Sign struct {
	Id           int64                 `gorm:"column:id" json:"id"`
	Coin         *Coin                 `gorm:"column:coin" json:"coin"`
	Amount       math.BigDecimal       `gorm:"column:amount" json:"amount"`
	EndDate      xtime.Xtime           `gorm:"column:end_date" json:"endDate"`
	CreationTime xtime.Xtime           `gorm:"column:creation_time" json:"creationTime"`
	Status       SignStatus.SignStatus `gorm:"column:status" json:"status"`
}
