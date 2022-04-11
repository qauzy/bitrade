package entity

import (
	"bitrade/core/constant/DepositStatusEnum"
	"github.com/qauzy/math"
)

func (this *DepositRecord) SetId(Id string) (result *DepositRecord) {
	this.Id = Id
	return this
}
func (this *DepositRecord) GetId() (Id string) {
	return this.Id
}
func (this *DepositRecord) SetMember(Member *Member) (result *DepositRecord) {
	this.Member = Member
	return this
}
func (this *DepositRecord) GetMember() (Member *Member) {
	return this.Member
}
func (this *DepositRecord) SetCoin(Coin *Coin) (result *DepositRecord) {
	this.Coin = Coin
	return this
}
func (this *DepositRecord) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *DepositRecord) SetAmount(Amount math.BigDecimal) (result *DepositRecord) {
	this.Amount = Amount
	return this
}
func (this *DepositRecord) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *DepositRecord) SetStatus(Status DepositStatusEnum.DepositStatusEnum) (result *DepositRecord) {
	this.Status = Status
	return this
}
func (this *DepositRecord) GetStatus() (Status DepositStatusEnum.DepositStatusEnum) {
	return this.Status
}
func NewDepositRecord(id string, member *Member, coin *Coin, amount math.BigDecimal, status DepositStatusEnum.DepositStatusEnum) (ret *DepositRecord) {
	ret = new(DepositRecord)
	ret.Id = id
	ret.Member = member
	ret.Coin = coin
	ret.Amount = amount
	ret.Status = status
	return
}

type DepositRecord struct {
	Id     string                              `gorm:"column:id" json:"id"`
	Member *Member                             `gorm:"column:member" json:"member"`
	Coin   *Coin                               `gorm:"column:coin" json:"coin"`
	Amount math.BigDecimal                     `gorm:"column:amount" json:"amount"`
	Status DepositStatusEnum.DepositStatusEnum `gorm:"column:status" json:"status"`
}
