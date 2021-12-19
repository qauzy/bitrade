package entity

func (this *WithdrawRecord) SetId(id int64) (result *WithdrawRecord) {
	this.Id = id
	return this
}
func (this *WithdrawRecord) GetId() (id int64) {
	return this.Id
}
func (this *WithdrawRecord) SetMemberId(memberId int64) (result *WithdrawRecord) {
	this.MemberId = memberId
	return this
}
func (this *WithdrawRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *WithdrawRecord) SetCoin(coin Coin) (result *WithdrawRecord) {
	this.Coin = coin
	return this
}
func (this *WithdrawRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *WithdrawRecord) SetTotalAmount(totalAmount math.BigDecimal) (result *WithdrawRecord) {
	this.TotalAmount = totalAmount
	return this
}
func (this *WithdrawRecord) GetTotalAmount() (totalAmount math.BigDecimal) {
	return this.TotalAmount
}
func (this *WithdrawRecord) SetFee(fee math.BigDecimal) (result *WithdrawRecord) {
	this.Fee = fee
	return this
}
func (this *WithdrawRecord) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *WithdrawRecord) SetArrivedAmount(arrivedAmount math.BigDecimal) (result *WithdrawRecord) {
	this.ArrivedAmount = arrivedAmount
	return this
}
func (this *WithdrawRecord) GetArrivedAmount() (arrivedAmount math.BigDecimal) {
	return this.ArrivedAmount
}
func (this *WithdrawRecord) SetCreateTime(createTime time.Time) (result *WithdrawRecord) {
	this.CreateTime = createTime
	return this
}
func (this *WithdrawRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *WithdrawRecord) SetDealTime(dealTime time.Time) (result *WithdrawRecord) {
	this.DealTime = dealTime
	return this
}
func (this *WithdrawRecord) GetDealTime() (dealTime time.Time) {
	return this.DealTime
}
func (this *WithdrawRecord) SetStatus(status constant.WithdrawStatus) (result *WithdrawRecord) {
	this.Status = status
	return this
}
func (this *WithdrawRecord) GetStatus() (status constant.WithdrawStatus) {
	return this.Status
}
func (this *WithdrawRecord) SetIsAuto(isAuto constant.BooleanEnum) (result *WithdrawRecord) {
	this.IsAuto = isAuto
	return this
}
func (this *WithdrawRecord) GetIsAuto() (isAuto constant.BooleanEnum) {
	return this.IsAuto
}
func (this *WithdrawRecord) SetAdmin(admin Admin) (result *WithdrawRecord) {
	this.Admin = admin
	return this
}
func (this *WithdrawRecord) GetAdmin() (admin Admin) {
	return this.Admin
}
func (this *WithdrawRecord) SetCanAutoWithdraw(canAutoWithdraw constant.BooleanEnum) (result *WithdrawRecord) {
	this.CanAutoWithdraw = canAutoWithdraw
	return this
}
func (this *WithdrawRecord) GetCanAutoWithdraw() (canAutoWithdraw constant.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *WithdrawRecord) SetTransactionNumber(transactionNumber string) (result *WithdrawRecord) {
	this.TransactionNumber = transactionNumber
	return this
}
func (this *WithdrawRecord) GetTransactionNumber() (transactionNumber string) {
	return this.TransactionNumber
}
func (this *WithdrawRecord) SetAddress(address string) (result *WithdrawRecord) {
	this.Address = address
	return this
}
func (this *WithdrawRecord) GetAddress() (address string) {
	return this.Address
}
func (this *WithdrawRecord) SetRemark(remark string) (result *WithdrawRecord) {
	this.Remark = remark
	return this
}
func (this *WithdrawRecord) GetRemark() (remark string) {
	return this.Remark
}

type WithdrawRecord struct {
	Id                int64
	MemberId          int64
	Coin              Coin
	TotalAmount       math.BigDecimal
	Fee               math.BigDecimal
	ArrivedAmount     math.BigDecimal
	CreateTime        time.Time
	DealTime          time.Time
	Status            constant.WithdrawStatus
	IsAuto            constant.BooleanEnum
	Admin             Admin
	CanAutoWithdraw   constant.BooleanEnum
	TransactionNumber string
	Address           string
	Remark            string
}
