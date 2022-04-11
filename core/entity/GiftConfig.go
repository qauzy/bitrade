package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *GiftConfig) SetId(Id int64) (result *GiftConfig) {
	this.Id = Id
	return this
}
func (this *GiftConfig) GetId() (Id int64) {
	return this.Id
}
func (this *GiftConfig) SetGiftName(GiftName string) (result *GiftConfig) {
	this.GiftName = GiftName
	return this
}
func (this *GiftConfig) GetGiftName() (GiftName string) {
	return this.GiftName
}
func (this *GiftConfig) SetGiftCoin(GiftCoin string) (result *GiftConfig) {
	this.GiftCoin = GiftCoin
	return this
}
func (this *GiftConfig) GetGiftCoin() (GiftCoin string) {
	return this.GiftCoin
}
func (this *GiftConfig) SetAmount(Amount math.BigDecimal) (result *GiftConfig) {
	this.Amount = Amount
	return this
}
func (this *GiftConfig) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *GiftConfig) SetHaveCoin(HaveCoin string) (result *GiftConfig) {
	this.HaveCoin = HaveCoin
	return this
}
func (this *GiftConfig) GetHaveCoin() (HaveCoin string) {
	return this.HaveCoin
}
func (this *GiftConfig) SetHaveAmount(HaveAmount math.BigDecimal) (result *GiftConfig) {
	this.HaveAmount = HaveAmount
	return this
}
func (this *GiftConfig) GetHaveAmount() (HaveAmount math.BigDecimal) {
	return this.HaveAmount
}
func (this *GiftConfig) SetCreateTime(CreateTime xtime.Xtime) (result *GiftConfig) {
	this.CreateTime = CreateTime
	return this
}
func (this *GiftConfig) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func NewGiftConfig(id int64, giftName string, giftCoin string, amount math.BigDecimal, haveCoin string, haveAmount math.BigDecimal, createTime xtime.Xtime) (ret *GiftConfig) {
	ret = new(GiftConfig)
	ret.Id = id
	ret.GiftName = giftName
	ret.GiftCoin = giftCoin
	ret.Amount = amount
	ret.HaveCoin = haveCoin
	ret.HaveAmount = haveAmount
	ret.CreateTime = createTime
	return
}

type GiftConfig struct {
	Id         int64           `gorm:"column:id" json:"id"`
	GiftName   string          `gorm:"column:gift_name" json:"giftName"`
	GiftCoin   string          `gorm:"column:gift_coin" json:"giftCoin"`
	Amount     math.BigDecimal `gorm:"column:amount" json:"amount"`
	HaveCoin   string          `gorm:"column:have_coin" json:"haveCoin"`
	HaveAmount math.BigDecimal `gorm:"column:have_amount" json:"haveAmount"`
	CreateTime xtime.Xtime     `gorm:"column:create_time" json:"createTime"`
}
