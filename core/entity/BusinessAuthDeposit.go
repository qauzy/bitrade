package entity

import (
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/math"
	"time"
)

func (this *BusinessAuthDeposit) SetId(id int64) (result *BusinessAuthDeposit) {
	this.Id = id
	return this
}
func (this *BusinessAuthDeposit) GetId() (id int64) {
	return this.Id
}
func (this *BusinessAuthDeposit) SetCoin(coin Coin) (result *BusinessAuthDeposit) {
	this.Coin = coin
	return this
}
func (this *BusinessAuthDeposit) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *BusinessAuthDeposit) SetAmount(amount math.BigDecimal) (result *BusinessAuthDeposit) {
	this.Amount = amount
	return this
}
func (this *BusinessAuthDeposit) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthDeposit) SetCreateTime(createTime time.Time) (result *BusinessAuthDeposit) {
	this.CreateTime = createTime
	return this
}
func (this *BusinessAuthDeposit) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *BusinessAuthDeposit) SetAdmin(admin Admin) (result *BusinessAuthDeposit) {
	this.Admin = admin
	return this
}
func (this *BusinessAuthDeposit) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *BusinessAuthDeposit) SetStatus(status CommonStatus.CommonStatus) (result *BusinessAuthDeposit) {
	this.Status = status
	return this
}
func (this *BusinessAuthDeposit) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}

type BusinessAuthDeposit struct {
	Id         int64
	Coin       Coin
	Amount     math.BigDecimal
	CreateTime time.Time
	Admin      Admin
	Status     CommonStatus.CommonStatus
}
