package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/WithdrawStatus"
	"bitrade/core/entity"
	"github.com/qauzy/math"
	"time"
)

func (this *ScanWithdrawRecord) SetTotalAmount(TotalAmount math.BigDecimal) (result *ScanWithdrawRecord) {
	this.TotalAmount = TotalAmount
	return this
}
func (this *ScanWithdrawRecord) GetTotalAmount() (TotalAmount math.BigDecimal) {
	return this.TotalAmount
}
func (this *ScanWithdrawRecord) SetFee(Fee math.BigDecimal) (result *ScanWithdrawRecord) {
	this.Fee = Fee
	return this
}
func (this *ScanWithdrawRecord) GetFee() (Fee math.BigDecimal) {
	return this.Fee
}
func (this *ScanWithdrawRecord) SetArrivedAmount(ArrivedAmount math.BigDecimal) (result *ScanWithdrawRecord) {
	this.ArrivedAmount = ArrivedAmount
	return this
}
func (this *ScanWithdrawRecord) GetArrivedAmount() (ArrivedAmount math.BigDecimal) {
	return this.ArrivedAmount
}
func (this *ScanWithdrawRecord) SetCreateTime(CreateTime time.Time) (result *ScanWithdrawRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *ScanWithdrawRecord) GetCreateTime() (CreateTime time.Time) {
	return this.CreateTime
}
func (this *ScanWithdrawRecord) SetDealTime(DealTime time.Time) (result *ScanWithdrawRecord) {
	this.DealTime = DealTime
	return this
}
func (this *ScanWithdrawRecord) GetDealTime() (DealTime time.Time) {
	return this.DealTime
}
func (this *ScanWithdrawRecord) SetStatus(Status WithdrawStatus.WithdrawStatus) (result *ScanWithdrawRecord) {
	this.Status = Status
	return this
}
func (this *ScanWithdrawRecord) GetStatus() (Status WithdrawStatus.WithdrawStatus) {
	return this.Status
}
func (this *ScanWithdrawRecord) SetIsAuto(IsAuto BooleanEnum.BooleanEnum) (result *ScanWithdrawRecord) {
	this.IsAuto = IsAuto
	return this
}
func (this *ScanWithdrawRecord) GetIsAuto() (IsAuto BooleanEnum.BooleanEnum) {
	return this.IsAuto
}
func (this *ScanWithdrawRecord) SetUnit(Unit string) (result *ScanWithdrawRecord) {
	this.Unit = Unit
	return this
}
func (this *ScanWithdrawRecord) GetUnit() (Unit string) {
	return this.Unit
}
func (this *ScanWithdrawRecord) SetCanAutoWithdraw(CanAutoWithdraw BooleanEnum.BooleanEnum) (result *ScanWithdrawRecord) {
	this.CanAutoWithdraw = CanAutoWithdraw
	return this
}
func (this *ScanWithdrawRecord) GetCanAutoWithdraw() (CanAutoWithdraw BooleanEnum.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *ScanWithdrawRecord) SetTransactionNumber(TransactionNumber string) (result *ScanWithdrawRecord) {
	this.TransactionNumber = TransactionNumber
	return this
}
func (this *ScanWithdrawRecord) GetTransactionNumber() (TransactionNumber string) {
	return this.TransactionNumber
}
func (this *ScanWithdrawRecord) SetAddress(Address string) (result *ScanWithdrawRecord) {
	this.Address = Address
	return this
}
func (this *ScanWithdrawRecord) GetAddress() (Address string) {
	return this.Address
}
func (this *ScanWithdrawRecord) SetRemark(Remark string) (result *ScanWithdrawRecord) {
	this.Remark = Remark
	return this
}
func (this *ScanWithdrawRecord) GetRemark() (Remark string) {
	return this.Remark
}
func ToScanWithdrawRecord(withdrawRecord entity.WithdrawRecord) (result *ScanWithdrawRecord) {
	return new(ScanWithdrawRecord).SetTotalAmount(withdrawRecord.GetTotalAmount()).SetCreateTime(withdrawRecord.GetCreateTime()).SetUnit(withdrawRecord.GetCoin().GetUnit()).SetDealTime(withdrawRecord.GetDealTime()).SetFee(withdrawRecord.GetFee()).SetArrivedAmount(withdrawRecord.GetArrivedAmount()).SetStatus(withdrawRecord.GetStatus()).SetIsAuto(withdrawRecord.GetIsAuto()).SetAddress(withdrawRecord.GetAddress()).SetRemark(withdrawRecord.GetRemark()).SetCanAutoWithdraw(withdrawRecord.GetCanAutoWithdraw()).SetTransactionNumber(withdrawRecord.GetTransactionNumber())
}

type ScanWithdrawRecord struct {
	TotalAmount       math.BigDecimal               `gorm:"column:total_amount" json:"totalAmount"`
	Fee               math.BigDecimal               `gorm:"column:fee" json:"fee"`
	ArrivedAmount     math.BigDecimal               `gorm:"column:arrived_amount" json:"arrivedAmount"`
	CreateTime        time.Time                     `gorm:"column:create_time" json:"createTime"`
	DealTime          time.Time                     `gorm:"column:deal_time" json:"dealTime"`
	Status            WithdrawStatus.WithdrawStatus `gorm:"column:status" json:"status"`
	IsAuto            BooleanEnum.BooleanEnum       `gorm:"column:is_auto" json:"isAuto"`
	Unit              string                        `gorm:"column:unit" json:"unit"`
	CanAutoWithdraw   BooleanEnum.BooleanEnum       `gorm:"column:can_auto_withdraw" json:"canAutoWithdraw"`
	TransactionNumber string                        `gorm:"column:transaction_number" json:"transactionNumber"`
	Address           string                        `gorm:"column:address" json:"address"`
	Remark            string                        `gorm:"column:remark" json:"remark"`
}
