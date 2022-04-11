package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *HotTransferRecord) SetId(Id int64) (result *HotTransferRecord) {
	this.Id = Id
	return this
}
func (this *HotTransferRecord) GetId() (Id int64) {
	return this.Id
}
func (this *HotTransferRecord) SetAdminId(AdminId int64) (result *HotTransferRecord) {
	this.AdminId = AdminId
	return this
}
func (this *HotTransferRecord) GetAdminId() (AdminId int64) {
	return this.AdminId
}
func (this *HotTransferRecord) SetTransferTime(TransferTime xtime.Xtime) (result *HotTransferRecord) {
	this.TransferTime = TransferTime
	return this
}
func (this *HotTransferRecord) GetTransferTime() (TransferTime xtime.Xtime) {
	return this.TransferTime
}
func (this *HotTransferRecord) SetAmount(Amount math.BigDecimal) (result *HotTransferRecord) {
	this.Amount = Amount
	return this
}
func (this *HotTransferRecord) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *HotTransferRecord) SetBalance(Balance math.BigDecimal) (result *HotTransferRecord) {
	this.Balance = Balance
	return this
}
func (this *HotTransferRecord) GetBalance() (Balance math.BigDecimal) {
	return this.Balance
}
func (this *HotTransferRecord) SetMinerFee(MinerFee math.BigDecimal) (result *HotTransferRecord) {
	this.MinerFee = MinerFee
	return this
}
func (this *HotTransferRecord) GetMinerFee() (MinerFee math.BigDecimal) {
	return this.MinerFee
}
func (this *HotTransferRecord) SetUnit(Unit string) (result *HotTransferRecord) {
	this.Unit = Unit
	return this
}
func (this *HotTransferRecord) GetUnit() (Unit string) {
	return this.Unit
}
func (this *HotTransferRecord) SetAdminName(AdminName string) (result *HotTransferRecord) {
	this.AdminName = AdminName
	return this
}
func (this *HotTransferRecord) GetAdminName() (AdminName string) {
	return this.AdminName
}
func (this *HotTransferRecord) SetColdAddress(ColdAddress string) (result *HotTransferRecord) {
	this.ColdAddress = ColdAddress
	return this
}
func (this *HotTransferRecord) GetColdAddress() (ColdAddress string) {
	return this.ColdAddress
}
func (this *HotTransferRecord) SetTransactionNumber(TransactionNumber string) (result *HotTransferRecord) {
	this.TransactionNumber = TransactionNumber
	return this
}
func (this *HotTransferRecord) GetTransactionNumber() (TransactionNumber string) {
	return this.TransactionNumber
}
func NewHotTransferRecord() (this *HotTransferRecord) {
	this = new(HotTransferRecord)
	return
}
func NewHotTransferRecordEx(id int64, adminId int64, transferTime xtime.Xtime, amount math.BigDecimal, balance math.BigDecimal, minerFee math.BigDecimal, unit string, adminName string, coldAddress string, transactionNumber string) (ret *HotTransferRecord) {
	ret = new(HotTransferRecord)
	ret.Id = id
	ret.AdminId = adminId
	ret.TransferTime = transferTime
	ret.Amount = amount
	ret.Balance = balance
	ret.MinerFee = minerFee
	ret.Unit = unit
	ret.AdminName = adminName
	ret.ColdAddress = coldAddress
	ret.TransactionNumber = transactionNumber
	return
}

type HotTransferRecord struct {
	Id                int64           `gorm:"column:id" json:"id"`
	AdminId           int64           `gorm:"column:admin_id" json:"adminId"`
	TransferTime      xtime.Xtime     `gorm:"column:transfer_time" json:"transferTime"`
	Amount            math.BigDecimal `gorm:"column:amount" json:"amount"`
	Balance           math.BigDecimal `gorm:"column:balance" json:"balance"`
	MinerFee          math.BigDecimal `gorm:"column:miner_fee" json:"minerFee"`
	Unit              string          `gorm:"column:unit" json:"unit"`
	AdminName         string          `gorm:"column:admin_name" json:"adminName"`
	ColdAddress       string          `gorm:"column:cold_address" json:"coldAddress"`
	TransactionNumber string          `gorm:"column:transaction_number" json:"transactionNumber"`
}
