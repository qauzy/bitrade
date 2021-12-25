package entity

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
func (this *ScanWithdrawRecord) SetCreateTime(createTime time.Time) (result *ScanWithdrawRecord) {
	this.CreateTime = createTime
	return this
}
func (this *ScanWithdrawRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *ScanWithdrawRecord) SetDealTime(dealTime time.Time) (result *ScanWithdrawRecord) {
	this.DealTime = dealTime
	return this
}
func (this *ScanWithdrawRecord) GetDealTime() (dealTime time.Time) {
	return this.DealTime
}
func (this *ScanWithdrawRecord) SetStatus(status constant.WithdrawStatus) (result *ScanWithdrawRecord) {
	this.Status = status
	return this
}
func (this *ScanWithdrawRecord) GetStatus() (status constant.WithdrawStatus) {
	return this.Status
}
func (this *ScanWithdrawRecord) SetIsAuto(isAuto constant.BooleanEnum) (result *ScanWithdrawRecord) {
	this.IsAuto = isAuto
	return this
}
func (this *ScanWithdrawRecord) GetIsAuto() (isAuto constant.BooleanEnum) {
	return this.IsAuto
}
func (this *ScanWithdrawRecord) SetUnit(unit string) (result *ScanWithdrawRecord) {
	this.Unit = unit
	return this
}
func (this *ScanWithdrawRecord) GetUnit() (unit string) {
	return this.Unit
}
func (this *ScanWithdrawRecord) SetCanAutoWithdraw(canAutoWithdraw constant.BooleanEnum) (result *ScanWithdrawRecord) {
	this.CanAutoWithdraw = canAutoWithdraw
	return this
}
func (this *ScanWithdrawRecord) GetCanAutoWithdraw() (canAutoWithdraw constant.BooleanEnum) {
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
func ToScanWithdrawRecord(withdrawRecord WithdrawRecord) (result ScanWithdrawRecord) {
	return new(ScanWithdrawRecord).SetTotalAmount(withdrawRecord.GetTotalAmount()).SetCreateTime(withdrawRecord.GetCreateTime()).SetUnit(withdrawRecord.GetCoin().GetUnit()).SetDealTime(withdrawRecord.GetDealTime()).SetFee(withdrawRecord.GetFee()).SetArrivedAmount(withdrawRecord.GetArrivedAmount()).SetStatus(withdrawRecord.GetStatus()).SetIsAuto(withdrawRecord.GetIsAuto()).SetAddress(withdrawRecord.GetAddress()).SetRemark(withdrawRecord.GetRemark()).SetCanAutoWithdraw(withdrawRecord.GetCanAutoWithdraw()).SetTransactionNumber(withdrawRecord.GetTransactionNumber())
}

type ScanWithdrawRecord struct {
	TotalAmount       math.BigDecimal
	Fee               math.BigDecimal
	ArrivedAmount     math.BigDecimal
	CreateTime        time.Time
	DealTime          time.Time
	Status            constant.WithdrawStatus
	IsAuto            constant.BooleanEnum
	Unit              string
	CanAutoWithdraw   constant.BooleanEnum
	TransactionNumber string
	Address           string
	Remark            string
}
