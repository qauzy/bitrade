package entity

import (
	"bitrade/core/constant/CommonStatus"
	"github.com/qauzy/math"
)

func (this *TransferAddress) SetId(Id int64) (result *TransferAddress) {
	this.Id = Id
	return this
}
func (this *TransferAddress) GetId() (Id int64) {
	return this.Id
}
func (this *TransferAddress) SetCoin(Coin *Coin) (result *TransferAddress) {
	this.Coin = Coin
	return this
}
func (this *TransferAddress) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *TransferAddress) SetAddress(Address string) (result *TransferAddress) {
	this.Address = Address
	return this
}
func (this *TransferAddress) GetAddress() (Address string) {
	return this.Address
}
func (this *TransferAddress) SetTransferFee(TransferFee math.BigDecimal) (result *TransferAddress) {
	this.TransferFee = TransferFee
	return this
}
func (this *TransferAddress) GetTransferFee() (TransferFee math.BigDecimal) {
	return this.TransferFee
}
func (this *TransferAddress) SetMinAmount(MinAmount math.BigDecimal) (result *TransferAddress) {
	this.MinAmount = MinAmount
	return this
}
func (this *TransferAddress) GetMinAmount() (MinAmount math.BigDecimal) {
	return this.MinAmount
}
func (this *TransferAddress) SetStatus(Status CommonStatus.CommonStatus) (result *TransferAddress) {
	this.Status = Status
	return this
}
func (this *TransferAddress) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func NewTransferAddress(id int64, coin *Coin, address string, transferFee math.BigDecimal, minAmount math.BigDecimal, status CommonStatus.CommonStatus) (ret *TransferAddress) {
	ret = new(TransferAddress)
	ret.Id = id
	ret.Coin = coin
	ret.Address = address
	ret.TransferFee = transferFee
	ret.MinAmount = minAmount
	ret.Status = status
	return
}

type TransferAddress struct {
	Id          int64                     `gorm:"column:id" json:"id"`
	Coin        *Coin                     `gorm:"column:coin" json:"coin"`
	Address     string                    `gorm:"column:address" json:"address"`
	TransferFee math.BigDecimal           `gorm:"column:transfer_fee" json:"transferFee"`
	MinAmount   math.BigDecimal           `gorm:"column:min_amount" json:"minAmount"`
	Status      CommonStatus.CommonStatus `gorm:"column:status" json:"status"`
}
