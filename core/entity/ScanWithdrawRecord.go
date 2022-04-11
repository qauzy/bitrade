package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/WithdrawStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *ScanWithdrawRecord) SetTotalAmount(totalAmount math.BigDecimal) (result *ScanWithdrawRecord) {
	this.TotalAmount = totalAmount
	return this
}
func (this *ScanWithdrawRecord) GetTotalAmount() (totalAmount math.BigDecimal) {
	return this.TotalAmount
}
func (this *ScanWithdrawRecord) SetFee(fee math.BigDecimal) (result *ScanWithdrawRecord) {
	this.Fee = fee
	return this
}
func (this *ScanWithdrawRecord) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *ScanWithdrawRecord) SetArrivedAmount(arrivedAmount math.BigDecimal) (result *ScanWithdrawRecord) {
	this.ArrivedAmount = arrivedAmount
	return this
}
func (this *ScanWithdrawRecord) GetArrivedAmount() (arrivedAmount math.BigDecimal) {
	return this.ArrivedAmount
}
func (this *ScanWithdrawRecord) SetCreateTime(createTime xtime.Xtime) (result *ScanWithdrawRecord) {
	this.CreateTime = createTime
	return this
}
func (this *ScanWithdrawRecord) GetCreateTime() (createTime xtime.Xtime) {
	return this.CreateTime
}
func (this *ScanWithdrawRecord) SetDealTime(dealTime xtime.Xtime) (result *ScanWithdrawRecord) {
	this.DealTime = dealTime
	return this
}
func (this *ScanWithdrawRecord) GetDealTime() (dealTime xtime.Xtime) {
	return this.DealTime
}
func (this *ScanWithdrawRecord) SetStatus(status WithdrawStatus.WithdrawStatus) (result *ScanWithdrawRecord) {
	this.Status = status
	return this
}
func (this *ScanWithdrawRecord) GetStatus() (status WithdrawStatus.WithdrawStatus) {
	return this.Status
}
func (this *ScanWithdrawRecord) SetIsAuto(isAuto BooleanEnum.BooleanEnum) (result *ScanWithdrawRecord) {
	this.IsAuto = isAuto
	return this
}
func (this *ScanWithdrawRecord) GetIsAuto() (isAuto BooleanEnum.BooleanEnum) {
	return this.IsAuto
}
func (this *ScanWithdrawRecord) SetUnit(unit string) (result *ScanWithdrawRecord) {
	this.Unit = unit
	return this
}
func (this *ScanWithdrawRecord) GetUnit() (unit string) {
	return this.Unit
}
func (this *ScanWithdrawRecord) SetCanAutoWithdraw(canAutoWithdraw BooleanEnum.BooleanEnum) (result *ScanWithdrawRecord) {
	this.CanAutoWithdraw = canAutoWithdraw
	return this
}
func (this *ScanWithdrawRecord) GetCanAutoWithdraw() (canAutoWithdraw BooleanEnum.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *ScanWithdrawRecord) SetTransactionNumber(transactionNumber string) (result *ScanWithdrawRecord) {
	this.TransactionNumber = transactionNumber
	return this
}
func (this *ScanWithdrawRecord) GetTransactionNumber() (transactionNumber string) {
	return this.TransactionNumber
}
func (this *ScanWithdrawRecord) SetAddress(address string) (result *ScanWithdrawRecord) {
	this.Address = address
	return this
}
func (this *ScanWithdrawRecord) GetAddress() (address string) {
	return this.Address
}
func (this *ScanWithdrawRecord) SetRemark(remark string) (result *ScanWithdrawRecord) {
	this.Remark = remark
	return this
}
func (this *ScanWithdrawRecord) GetRemark() (remark string) {
	return this.Remark
}
func ToScanWithdrawRecord(withdrawRecord *WithdrawRecord) (result *ScanWithdrawRecord) {
	return new(ScanWithdrawRecord).SetTotalAmount(withdrawRecord.GetTotalAmount()).SetCreateTime(withdrawRecord.GetCreateTime()).SetUnit(withdrawRecord.GetCoin().GetUnit()).SetDealTime(withdrawRecord.GetDealTime()).SetFee(withdrawRecord.GetFee()).SetArrivedAmount(withdrawRecord.GetArrivedAmount()).SetStatus(withdrawRecord.GetStatus()).SetIsAuto(withdrawRecord.GetIsAuto()).SetAddress(withdrawRecord.GetAddress()).SetRemark(withdrawRecord.GetRemark()).SetCanAutoWithdraw(withdrawRecord.GetCanAutoWithdraw()).SetTransactionNumber(withdrawRecord.GetTransactionNumber())
}

type ScanWithdrawRecord struct {
	TotalAmount       math.BigDecimal
	Fee               math.BigDecimal
	ArrivedAmount     math.BigDecimal
	CreateTime        xtime.Xtime
	DealTime          xtime.Xtime
	Status            WithdrawStatus.WithdrawStatus
	IsAuto            BooleanEnum.BooleanEnum
	Unit              string
	CanAutoWithdraw   BooleanEnum.BooleanEnum
	TransactionNumber string
	Address           string
	Remark            string
}
