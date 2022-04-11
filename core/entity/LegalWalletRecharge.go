package entity

import (
	"bitrade/core/constant/LegalWalletState"
	"bitrade/core/constant/PayMode"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *LegalWalletRecharge) SetId(Id int64) (result *LegalWalletRecharge) {
	this.Id = Id
	return this
}
func (this *LegalWalletRecharge) GetId() (Id int64) {
	return this.Id
}
func (this *LegalWalletRecharge) SetMember(Member *Member) (result *LegalWalletRecharge) {
	this.Member = Member
	return this
}
func (this *LegalWalletRecharge) GetMember() (Member *Member) {
	return this.Member
}
func (this *LegalWalletRecharge) SetAmount(Amount math.BigDecimal) (result *LegalWalletRecharge) {
	this.Amount = Amount
	return this
}
func (this *LegalWalletRecharge) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *LegalWalletRecharge) SetPaymentInstrument(PaymentInstrument string) (result *LegalWalletRecharge) {
	this.PaymentInstrument = PaymentInstrument
	return this
}
func (this *LegalWalletRecharge) GetPaymentInstrument() (PaymentInstrument string) {
	return this.PaymentInstrument
}
func (this *LegalWalletRecharge) SetCoin(Coin *Coin) (result *LegalWalletRecharge) {
	this.Coin = Coin
	return this
}
func (this *LegalWalletRecharge) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *LegalWalletRecharge) SetState(State LegalWalletState.LegalWalletState) (result *LegalWalletRecharge) {
	this.State = State
	return this
}
func (this *LegalWalletRecharge) GetState() (State LegalWalletState.LegalWalletState) {
	return this.State
}
func (this *LegalWalletRecharge) SetPayMode(PayMode PayMode.PayMode) (result *LegalWalletRecharge) {
	this.PayMode = PayMode
	return this
}
func (this *LegalWalletRecharge) GetPayMode() (PayMode PayMode.PayMode) {
	return this.PayMode
}
func (this *LegalWalletRecharge) SetRemark(Remark string) (result *LegalWalletRecharge) {
	this.Remark = Remark
	return this
}
func (this *LegalWalletRecharge) GetRemark() (Remark string) {
	return this.Remark
}
func (this *LegalWalletRecharge) SetCreationTime(CreationTime xtime.Xtime) (result *LegalWalletRecharge) {
	this.CreationTime = CreationTime
	return this
}
func (this *LegalWalletRecharge) GetCreationTime() (CreationTime xtime.Xtime) {
	return this.CreationTime
}
func (this *LegalWalletRecharge) SetDealTime(DealTime xtime.Xtime) (result *LegalWalletRecharge) {
	this.DealTime = DealTime
	return this
}
func (this *LegalWalletRecharge) GetDealTime() (DealTime xtime.Xtime) {
	return this.DealTime
}
func (this *LegalWalletRecharge) SetAdmin(Admin *Admin) (result *LegalWalletRecharge) {
	this.Admin = Admin
	return this
}
func (this *LegalWalletRecharge) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func (this *LegalWalletRecharge) SetUpdateTime(UpdateTime xtime.Xtime) (result *LegalWalletRecharge) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *LegalWalletRecharge) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func NewLegalWalletRecharge(id int64, member *Member, amount math.BigDecimal, paymentInstrument string, coin *Coin, state LegalWalletState.LegalWalletState, payMode PayMode.PayMode, remark string, creationTime xtime.Xtime, dealTime xtime.Xtime, admin *Admin, updateTime xtime.Xtime) (ret *LegalWalletRecharge) {
	ret = new(LegalWalletRecharge)
	ret.Id = id
	ret.Member = member
	ret.Amount = amount
	ret.PaymentInstrument = paymentInstrument
	ret.Coin = coin
	ret.State = state
	ret.PayMode = payMode
	ret.Remark = remark
	ret.CreationTime = creationTime
	ret.DealTime = dealTime
	ret.Admin = admin
	ret.UpdateTime = updateTime
	return
}

type LegalWalletRecharge struct {
	Id                int64                             `gorm:"column:id" json:"id"`
	Member            *Member                           `gorm:"column:member" json:"member"`
	Amount            math.BigDecimal                   `gorm:"column:amount" json:"amount"`
	PaymentInstrument string                            `gorm:"column:payment_instrument" json:"paymentInstrument"`
	Coin              *Coin                             `gorm:"column:coin" json:"coin"`
	State             LegalWalletState.LegalWalletState `gorm:"column:state" json:"state"`
	PayMode           PayMode.PayMode                   `gorm:"column:pay_mode" json:"payMode"`
	Remark            string                            `gorm:"column:remark" json:"remark"`
	CreationTime      xtime.Xtime                       `gorm:"column:creation_time" json:"creationTime"`
	DealTime          xtime.Xtime                       `gorm:"column:deal_time" json:"dealTime"`
	Admin             *Admin                            `gorm:"column:admin" json:"admin"`
	UpdateTime        xtime.Xtime                       `gorm:"column:update_time" json:"updateTime"`
}
