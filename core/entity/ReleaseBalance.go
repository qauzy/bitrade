package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *ReleaseBalance) SetId(Id int64) (result *ReleaseBalance) {
	this.Id = Id
	return this
}
func (this *ReleaseBalance) GetId() (Id int64) {
	return this.Id
}
func (this *ReleaseBalance) SetMemberId(MemberId int64) (result *ReleaseBalance) {
	this.MemberId = MemberId
	return this
}
func (this *ReleaseBalance) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *ReleaseBalance) SetCoinName(CoinName string) (result *ReleaseBalance) {
	this.CoinName = CoinName
	return this
}
func (this *ReleaseBalance) GetCoinName() (CoinName string) {
	return this.CoinName
}
func (this *ReleaseBalance) SetCoinUnit(CoinUnit string) (result *ReleaseBalance) {
	this.CoinUnit = CoinUnit
	return this
}
func (this *ReleaseBalance) GetCoinUnit() (CoinUnit string) {
	return this.CoinUnit
}
func (this *ReleaseBalance) SetRegisterTime(RegisterTime xtime.Xtime) (result *ReleaseBalance) {
	this.RegisterTime = RegisterTime
	return this
}
func (this *ReleaseBalance) GetRegisterTime() (RegisterTime xtime.Xtime) {
	return this.RegisterTime
}
func (this *ReleaseBalance) SetUserName(UserName string) (result *ReleaseBalance) {
	this.UserName = UserName
	return this
}
func (this *ReleaseBalance) GetUserName() (UserName string) {
	return this.UserName
}
func (this *ReleaseBalance) SetPhone(Phone string) (result *ReleaseBalance) {
	this.Phone = Phone
	return this
}
func (this *ReleaseBalance) GetPhone() (Phone string) {
	return this.Phone
}
func (this *ReleaseBalance) SetEmail(Email string) (result *ReleaseBalance) {
	this.Email = Email
	return this
}
func (this *ReleaseBalance) GetEmail() (Email string) {
	return this.Email
}
func (this *ReleaseBalance) SetReleaseBalance(ReleaseBalance math.BigDecimal) (result *ReleaseBalance) {
	this.ReleaseBalance = ReleaseBalance
	return this
}
func (this *ReleaseBalance) GetReleaseBalance() (ReleaseBalance math.BigDecimal) {
	return this.ReleaseBalance
}
func (this *ReleaseBalance) SetReleaseState(ReleaseState string) (result *ReleaseBalance) {
	this.ReleaseState = ReleaseState
	return this
}
func (this *ReleaseBalance) GetReleaseState() (ReleaseState string) {
	return this.ReleaseState
}
func (this *ReleaseBalance) SetReleaseTime(ReleaseTime xtime.Xtime) (result *ReleaseBalance) {
	this.ReleaseTime = ReleaseTime
	return this
}
func (this *ReleaseBalance) GetReleaseTime() (ReleaseTime xtime.Xtime) {
	return this.ReleaseTime
}
func NewReleaseBalance(id int64, memberId int64, coinName string, coinUnit string, registerTime xtime.Xtime, userName string, phone string, email string, releaseBalance math.BigDecimal, releaseState string, releaseTime xtime.Xtime) (ret *ReleaseBalance) {
	ret = new(ReleaseBalance)
	ret.Id = id
	ret.MemberId = memberId
	ret.CoinName = coinName
	ret.CoinUnit = coinUnit
	ret.RegisterTime = registerTime
	ret.UserName = userName
	ret.Phone = phone
	ret.Email = email
	ret.ReleaseBalance = releaseBalance
	ret.ReleaseState = releaseState
	ret.ReleaseTime = releaseTime
	return
}

type ReleaseBalance struct {
	Id             int64           `gorm:"column:id" json:"id"`
	MemberId       int64           `gorm:"column:member_id" json:"memberId"`
	CoinName       string          `gorm:"column:coin_name" json:"coinName"`
	CoinUnit       string          `gorm:"column:coin_unit" json:"coinUnit"`
	RegisterTime   xtime.Xtime     `gorm:"column:register_time" json:"registerTime"`
	UserName       string          `gorm:"column:user_name" json:"userName"`
	Phone          string          `gorm:"column:phone" json:"phone"`
	Email          string          `gorm:"column:email" json:"email"`
	ReleaseBalance math.BigDecimal `gorm:"column:release_balance" json:"releaseBalance"`
	ReleaseState   string          `gorm:"column:release_state" json:"releaseState"`
	ReleaseTime    xtime.Xtime     `gorm:"column:release_time" json:"releaseTime"`
}
