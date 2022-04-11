package entity

import (
	"bitrade/core/constant/PayMode"
	"bitrade/core/constant/WithdrawStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *LegalWalletWithdraw) SetId(Id int64) (result *LegalWalletWithdraw) {
	this.Id = Id
	return this
}
func (this *LegalWalletWithdraw) GetId() (Id int64) {
	return this.Id
}
func (this *LegalWalletWithdraw) SetMember(Member *Member) (result *LegalWalletWithdraw) {
	this.Member = Member
	return this
}
func (this *LegalWalletWithdraw) GetMember() (Member *Member) {
	return this.Member
}
func (this *LegalWalletWithdraw) SetAmount(Amount math.BigDecimal) (result *LegalWalletWithdraw) {
	this.Amount = Amount
	return this
}
func (this *LegalWalletWithdraw) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletWithdraw) SetCreateTime(CreateTime xtime.Xtime) (result *LegalWalletWithdraw) {
	this.CreateTime = CreateTime
	return this
}
func (this *LegalWalletWithdraw) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *LegalWalletWithdraw) SetDealTime(DealTime xtime.Xtime) (result *LegalWalletWithdraw) {
	this.DealTime = DealTime
	return this
}
func (this *LegalWalletWithdraw) GetDealTime() (DealTime xtime.Xtime) {
	return this.DealTime
}
func (this *LegalWalletWithdraw) SetStatus(Status WithdrawStatus.WithdrawStatus) (result *LegalWalletWithdraw) {
	this.Status = Status
	return this
}
func (this *LegalWalletWithdraw) GetStatus() (Status WithdrawStatus.WithdrawStatus) {
	return this.Status
}
func (this *LegalWalletWithdraw) SetAdmin(Admin *Admin) (result *LegalWalletWithdraw) {
	this.Admin = Admin
	return this
}
func (this *LegalWalletWithdraw) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func (this *LegalWalletWithdraw) SetRemark(Remark string) (result *LegalWalletWithdraw) {
	this.Remark = Remark
	return this
}
func (this *LegalWalletWithdraw) GetRemark() (Remark string) {
	return this.Remark
}
func (this *LegalWalletWithdraw) SetPayMode(PayMode PayMode.PayMode) (result *LegalWalletWithdraw) {
	this.PayMode = PayMode
	return this
}
func (this *LegalWalletWithdraw) GetPayMode() (PayMode PayMode.PayMode) {
	return this.PayMode
}
func (this *LegalWalletWithdraw) SetCoin(Coin *Coin) (result *LegalWalletWithdraw) {
	this.Coin = Coin
	return this
}
func (this *LegalWalletWithdraw) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *LegalWalletWithdraw) SetPaymentInstrument(PaymentInstrument string) (result *LegalWalletWithdraw) {
	this.PaymentInstrument = PaymentInstrument
	return this
}
func (this *LegalWalletWithdraw) GetPaymentInstrument() (PaymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletWithdraw) SetRemitTime(RemitTime xtime.Xtime) (result *LegalWalletWithdraw) {
	this.RemitTime = RemitTime
	return this
}
func (this *LegalWalletWithdraw) GetRemitTime() (RemitTime xtime.Xtime) {
	return this.RemitTime
}
func (this *LegalWalletWithdraw) SetAccount(Account string) (result *LegalWalletWithdraw) {
	this.Account = Account
	return this
}
func (this *LegalWalletWithdraw) GetAccount() (Account string) {
	return this.Account
}
func NewLegalWalletWithdraw(id int64, member *Member, amount math.BigDecimal, createTime xtime.Xtime, dealTime xtime.Xtime, status WithdrawStatus.WithdrawStatus, admin *Admin, remark string, payMode PayMode.PayMode, coin *Coin, paymentInstrument string, remitTime xtime.Xtime, account string) (ret *LegalWalletWithdraw) {
	ret = new(LegalWalletWithdraw)
	ret.Id = id
	ret.Member = member
	ret.Amount = amount
	ret.CreateTime = createTime
	ret.DealTime = dealTime
	ret.Status = status
	ret.Admin = admin
	ret.Remark = remark
	ret.PayMode = payMode
	ret.Coin = coin
	ret.PaymentInstrument = paymentInstrument
	ret.RemitTime = remitTime
	ret.Account = account
	return
}

type LegalWalletWithdraw struct {
	Id                int64                         `gorm:"column:id" json:"id"`
	Member            *Member                       `gorm:"column:member" json:"member"`
	Amount            math.BigDecimal               `gorm:"column:amount" json:"amount"`
	CreateTime        xtime.Xtime                   `gorm:"column:create_time" json:"createTime"`
	DealTime          xtime.Xtime                   `gorm:"column:deal_time" json:"dealTime"`
	Status            WithdrawStatus.WithdrawStatus `gorm:"column:status" json:"status"`
	Admin             *Admin                        `gorm:"column:admin" json:"admin"`
	Remark            string                        `gorm:"column:remark" json:"remark"`
	PayMode           PayMode.PayMode               `gorm:"column:pay_mode" json:"payMode"`
	Coin              *Coin                         `gorm:"column:coin" json:"coin"`
	PaymentInstrument string                        `gorm:"column:payment_instrument" json:"paymentInstrument"`
	RemitTime         xtime.Xtime                   `gorm:"column:remit_time" json:"remitTime"`
	Account           string                        `gorm:"column:account" json:"account"`
}
