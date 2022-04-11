package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *TransferRecord) SetId(Id int64) (result *TransferRecord) {
	this.Id = Id
	return this
}
func (this *TransferRecord) GetId() (Id int64) {
	return this.Id
}
func (this *TransferRecord) SetAmount(Amount math.BigDecimal) (result *TransferRecord) {
	this.Amount = Amount
	return this
}
func (this *TransferRecord) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *TransferRecord) SetMemberId(MemberId int64) (result *TransferRecord) {
	this.MemberId = MemberId
	return this
}
func (this *TransferRecord) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *TransferRecord) SetCoin(Coin *Coin) (result *TransferRecord) {
	this.Coin = Coin
	return this
}
func (this *TransferRecord) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *TransferRecord) SetCreateTime(CreateTime xtime.Xtime) (result *TransferRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *TransferRecord) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *TransferRecord) SetFee(Fee math.BigDecimal) (result *TransferRecord) {
	this.Fee = Fee
	return this
}
func (this *TransferRecord) GetFee() (Fee math.BigDecimal) {
	return this.Fee
}
func (this *TransferRecord) SetOrderSn(OrderSn string) (result *TransferRecord) {
	this.OrderSn = OrderSn
	return this
}
func (this *TransferRecord) GetOrderSn() (OrderSn string) {
	return this.OrderSn
}
func (this *TransferRecord) SetAddress(Address string) (result *TransferRecord) {
	this.Address = Address
	return this
}
func (this *TransferRecord) GetAddress() (Address string) {
	return this.Address
}
func (this *TransferRecord) SetRemark(Remark string) (result *TransferRecord) {
	this.Remark = Remark
	return this
}
func (this *TransferRecord) GetRemark() (Remark string) {
	return this.Remark
}
func NewTransferRecord(id int64, amount math.BigDecimal, memberId int64, coin *Coin, createTime xtime.Xtime, fee math.BigDecimal, orderSn string, address string, remark string) (ret *TransferRecord) {
	ret = new(TransferRecord)
	ret.Id = id
	ret.Amount = amount
	ret.MemberId = memberId
	ret.Coin = coin
	ret.CreateTime = createTime
	ret.Fee = fee
	ret.OrderSn = orderSn
	ret.Address = address
	ret.Remark = remark
	return
}

type TransferRecord struct {
	Id         int64           `gorm:"column:id" json:"id"`
	Amount     math.BigDecimal `gorm:"column:amount" json:"amount"`
	MemberId   int64           `gorm:"column:member_id" json:"memberId"`
	Coin       *Coin           `gorm:"column:coin" json:"coin"`
	CreateTime xtime.Xtime     `gorm:"column:create_time" json:"createTime"`
	Fee        math.BigDecimal `gorm:"column:fee" json:"fee"`
	OrderSn    string          `gorm:"column:order_sn" json:"orderSn"`
	Address    string          `gorm:"column:address" json:"address"`
	Remark     string          `gorm:"column:remark" json:"remark"`
}
