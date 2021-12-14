package entity

func (this *WithdrawRecord) SetId(id int64) {
	this.Id = id
}
func (this *WithdrawRecord) GetId() (id int64) {
	return this.Id
}
func (this *WithdrawRecord) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *WithdrawRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *WithdrawRecord) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *WithdrawRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *WithdrawRecord) SetTotalAmount(totalAmount BigDecimal) {
	this.TotalAmount = totalAmount
}
func (this *WithdrawRecord) GetTotalAmount() (totalAmount BigDecimal) {
	return this.TotalAmount
}
func (this *WithdrawRecord) SetFee(fee BigDecimal) {
	this.Fee = fee
}
func (this *WithdrawRecord) GetFee() (fee BigDecimal) {
	return this.Fee
}
func (this *WithdrawRecord) SetArrivedAmount(arrivedAmount BigDecimal) {
	this.ArrivedAmount = arrivedAmount
}
func (this *WithdrawRecord) GetArrivedAmount() (arrivedAmount BigDecimal) {
	return this.ArrivedAmount
}
func (this *WithdrawRecord) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *WithdrawRecord) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *WithdrawRecord) SetDealTime(dealTime Date) {
	this.DealTime = dealTime
}
func (this *WithdrawRecord) GetDealTime() (dealTime Date) {
	return this.DealTime
}
func (this *WithdrawRecord) SetStatus(status WithdrawStatus) {
	this.Status = status
}
func (this *WithdrawRecord) GetStatus() (status WithdrawStatus) {
	return this.Status
}
func (this *WithdrawRecord) SetIsAuto(isAuto BooleanEnum) {
	this.IsAuto = isAuto
}
func (this *WithdrawRecord) GetIsAuto() (isAuto BooleanEnum) {
	return this.IsAuto
}
func (this *WithdrawRecord) SetAdmin(admin Admin) {
	this.Admin = admin
}
func (this *WithdrawRecord) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *WithdrawRecord) SetCanAutoWithdraw(canAutoWithdraw BooleanEnum) {
	this.CanAutoWithdraw = canAutoWithdraw
}
func (this *WithdrawRecord) GetCanAutoWithdraw() (canAutoWithdraw BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *WithdrawRecord) SetTransactionNumber(transactionNumber string) {
	this.TransactionNumber = transactionNumber
}
func (this *WithdrawRecord) GetTransactionNumber() (transactionNumber string) {
	return this.TransactionNumber
}
func (this *WithdrawRecord) SetAddress(address string) {
	this.Address = address
}
func (this *WithdrawRecord) GetAddress() (address string) {
	return this.Address
}
func (this *WithdrawRecord) SetRemark(remark string) {
	this.Remark = remark
}
func (this *WithdrawRecord) GetRemark() (remark string) {
	return this.Remark
}

type WithdrawRecord struct {
	Id                int64
	MemberId          int64
	Coin              Coin
	TotalAmount       BigDecimal
	Fee               BigDecimal
	ArrivedAmount     BigDecimal
	CreateTime        Date
	DealTime          Date
	Status            WithdrawStatus
	IsAuto            BooleanEnum
	Admin             Admin
	CanAutoWithdraw   BooleanEnum
	TransactionNumber string
	Address           string
	Remark            string
}
