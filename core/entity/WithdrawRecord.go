package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/WithdrawStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *WithdrawRecord) SetId(Id int64) (result *WithdrawRecord) {
	this.Id = Id
	return this
}
func (this *WithdrawRecord) GetId() (Id int64) {
	return this.Id
}
func (this *WithdrawRecord) SetMemberId(MemberId int64) (result *WithdrawRecord) {
	this.MemberId = MemberId
	return this
}
func (this *WithdrawRecord) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *WithdrawRecord) SetCoin(Coin *Coin) (result *WithdrawRecord) {
	this.Coin = Coin
	return this
}
func (this *WithdrawRecord) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *WithdrawRecord) SetTotalAmount(TotalAmount math.BigDecimal) (result *WithdrawRecord) {
	this.TotalAmount = TotalAmount
	return this
}
func (this *WithdrawRecord) GetTotalAmount() (TotalAmount math.BigDecimal) {
	return this.TotalAmount
}
func (this *WithdrawRecord) SetFee(Fee math.BigDecimal) (result *WithdrawRecord) {
	this.Fee = Fee
	return this
}
func (this *WithdrawRecord) GetFee() (Fee math.BigDecimal) {
	return this.Fee
}
func (this *WithdrawRecord) SetArrivedAmount(ArrivedAmount math.BigDecimal) (result *WithdrawRecord) {
	this.ArrivedAmount = ArrivedAmount
	return this
}
func (this *WithdrawRecord) GetArrivedAmount() (ArrivedAmount math.BigDecimal) {
	return this.ArrivedAmount
}
func (this *WithdrawRecord) SetCreateTime(CreateTime xtime.Xtime) (result *WithdrawRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *WithdrawRecord) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *WithdrawRecord) SetDealTime(DealTime xtime.Xtime) (result *WithdrawRecord) {
	this.DealTime = DealTime
	return this
}
func (this *WithdrawRecord) GetDealTime() (DealTime xtime.Xtime) {
	return this.DealTime
}
func (this *WithdrawRecord) SetStatus(Status WithdrawStatus.WithdrawStatus) (result *WithdrawRecord) {
	this.Status = Status
	return this
}
func (this *WithdrawRecord) GetStatus() (Status WithdrawStatus.WithdrawStatus) {
	return this.Status
}
func (this *WithdrawRecord) SetIsAuto(IsAuto BooleanEnum.BooleanEnum) (result *WithdrawRecord) {
	this.IsAuto = IsAuto
	return this
}
func (this *WithdrawRecord) GetIsAuto() (IsAuto BooleanEnum.BooleanEnum) {
	return this.IsAuto
}
func (this *WithdrawRecord) SetAdmin(Admin *Admin) (result *WithdrawRecord) {
	this.Admin = Admin
	return this
}
func (this *WithdrawRecord) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func (this *WithdrawRecord) SetCanAutoWithdraw(CanAutoWithdraw BooleanEnum.BooleanEnum) (result *WithdrawRecord) {
	this.CanAutoWithdraw = CanAutoWithdraw
	return this
}
func (this *WithdrawRecord) GetCanAutoWithdraw() (CanAutoWithdraw BooleanEnum.BooleanEnum) {
	return this.CanAutoWithdraw
}
func (this *WithdrawRecord) SetTransactionNumber(TransactionNumber string) (result *WithdrawRecord) {
	this.TransactionNumber = TransactionNumber
	return this
}
func (this *WithdrawRecord) GetTransactionNumber() (TransactionNumber string) {
	return this.TransactionNumber
}
func (this *WithdrawRecord) SetAddress(Address string) (result *WithdrawRecord) {
	this.Address = Address
	return this
}
func (this *WithdrawRecord) GetAddress() (Address string) {
	return this.Address
}
func (this *WithdrawRecord) SetRemark(Remark string) (result *WithdrawRecord) {
	this.Remark = Remark
	return this
}
func (this *WithdrawRecord) GetRemark() (Remark string) {
	return this.Remark
}
func NewWithdrawRecord(id int64, memberId int64, coin *Coin, totalAmount math.BigDecimal, fee math.BigDecimal, arrivedAmount math.BigDecimal, createTime xtime.Xtime, dealTime xtime.Xtime, status WithdrawStatus.WithdrawStatus, isAuto BooleanEnum.BooleanEnum, admin *Admin, canAutoWithdraw BooleanEnum.BooleanEnum, transactionNumber string, address string, remark string) (ret *WithdrawRecord) {
	ret = new(WithdrawRecord)
	ret.Id = id
	ret.MemberId = memberId
	ret.Coin = coin
	ret.TotalAmount = totalAmount
	ret.Fee = fee
	ret.ArrivedAmount = arrivedAmount
	ret.CreateTime = createTime
	ret.DealTime = dealTime
	ret.Status = status
	ret.IsAuto = isAuto
	ret.Admin = admin
	ret.CanAutoWithdraw = canAutoWithdraw
	ret.TransactionNumber = transactionNumber
	ret.Address = address
	ret.Remark = remark
	return
}

type WithdrawRecord struct {
	Id                int64                         `gorm:"column:id" json:"id"`
	MemberId          int64                         `gorm:"column:member_id" json:"memberId"`
	Coin              *Coin                         `gorm:"column:coin" json:"coin"`
	TotalAmount       math.BigDecimal               `gorm:"column:total_amount" json:"totalAmount"`
	Fee               math.BigDecimal               `gorm:"column:fee" json:"fee"`
	ArrivedAmount     math.BigDecimal               `gorm:"column:arrived_amount" json:"arrivedAmount"`
	CreateTime        xtime.Xtime                   `gorm:"column:create_time" json:"createTime"`
	DealTime          xtime.Xtime                   `gorm:"column:deal_time" json:"dealTime"`
	Status            WithdrawStatus.WithdrawStatus `gorm:"column:status" json:"status"`
	IsAuto            BooleanEnum.BooleanEnum       `gorm:"column:is_auto" json:"isAuto"`
	Admin             *Admin                        `gorm:"column:admin" json:"admin"`
	CanAutoWithdraw   BooleanEnum.BooleanEnum       `gorm:"column:can_auto_withdraw" json:"canAutoWithdraw"`
	TransactionNumber string                        `gorm:"column:transaction_number" json:"transactionNumber"`
	Address           string                        `gorm:"column:address" json:"address"`
	Remark            string                        `gorm:"column:remark" json:"remark"`
}
