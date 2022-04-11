package entity

import (
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *BusinessAuthDeposit) SetId(Id int64) (result *BusinessAuthDeposit) {
	this.Id = Id
	return this
}
func (this *BusinessAuthDeposit) GetId() (Id int64) {
	return this.Id
}
func (this *BusinessAuthDeposit) SetCoin(Coin *Coin) (result *BusinessAuthDeposit) {
	this.Coin = Coin
	return this
}
func (this *BusinessAuthDeposit) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *BusinessAuthDeposit) SetAmount(Amount math.BigDecimal) (result *BusinessAuthDeposit) {
	this.Amount = Amount
	return this
}
func (this *BusinessAuthDeposit) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthDeposit) SetCreateTime(CreateTime xtime.Xtime) (result *BusinessAuthDeposit) {
	this.CreateTime = CreateTime
	return this
}
func (this *BusinessAuthDeposit) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *BusinessAuthDeposit) SetAdmin(Admin *Admin) (result *BusinessAuthDeposit) {
	this.Admin = Admin
	return this
}
func (this *BusinessAuthDeposit) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func (this *BusinessAuthDeposit) SetStatus(Status CommonStatus.CommonStatus) (result *BusinessAuthDeposit) {
	this.Status = Status
	return this
}
func (this *BusinessAuthDeposit) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func NewBusinessAuthDeposit(id int64, coin *Coin, amount math.BigDecimal, createTime xtime.Xtime, admin *Admin, status CommonStatus.CommonStatus) (ret *BusinessAuthDeposit) {
	ret = new(BusinessAuthDeposit)
	ret.Id = id
	ret.Coin = coin
	ret.Amount = amount
	ret.CreateTime = createTime
	ret.Admin = admin
	ret.Status = status
	return
}

type BusinessAuthDeposit struct {
	Id         int64                     `gorm:"column:id" json:"id"`
	Coin       *Coin                     `gorm:"column:coin" json:"coin"`
	Amount     math.BigDecimal           `gorm:"column:amount" json:"amount"`
	CreateTime xtime.Xtime               `gorm:"column:create_time" json:"createTime"`
	Admin      *Admin                    `gorm:"column:admin" json:"admin"`
	Status     CommonStatus.CommonStatus `gorm:"column:status" json:"status"`
}
