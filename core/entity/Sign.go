package entity

import (
	"bitrade/core/constant/SignStatus"
	"github.com/qauzy/math"
	"time"
)

func (this *Sign) SetId(id int64) (result *Sign) {
	this.Id = id
	return this
}
func (this *Sign) GetId() (id int64) {
	return this.Id
}
func (this *Sign) SetCoin(coin *Coin) (result *Sign) {
	this.Coin = coin
	return this
}
func (this *Sign) GetCoin() (coin *Coin) {
	return this.Coin
}
func (this *Sign) SetAmount(amount math.BigDecimal) (result *Sign) {
	this.Amount = amount
	return this
}
func (this *Sign) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *Sign) SetEndDate(endDate time.Time) (result *Sign) {
	this.EndDate = endDate
	return this
}
func (this *Sign) GetEndDate() (endDate time.Time) {
	return this.EndDate
}
func (this *Sign) SetCreationTime(creationTime time.Time) (result *Sign) {
	this.CreationTime = creationTime
	return this
}
func (this *Sign) GetCreationTime() (creationTime time.Time) {
	return this.CreationTime
}
func (this *Sign) SetStatus(status SignStatus.SignStatus) (result *Sign) {
	this.Status = status
	return this
}
func (this *Sign) GetStatus() (status SignStatus.SignStatus) {
	return this.Status
}

type Sign struct {
	Id           int64
	Coin         *Coin
	Amount       math.BigDecimal
	EndDate      time.Time
	CreationTime time.Time
	Status       SignStatus.SignStatus
}
