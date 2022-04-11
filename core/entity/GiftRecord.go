package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *GiftRecord) SetId(Id int64) (result *GiftRecord) {
	this.Id = Id
	return this
}
func (this *GiftRecord) GetId() (Id int64) {
	return this.Id
}
func (this *GiftRecord) SetUserId(UserId int64) (result *GiftRecord) {
	this.UserId = UserId
	return this
}
func (this *GiftRecord) GetUserId() (UserId int64) {
	return this.UserId
}
func (this *GiftRecord) SetUserName(UserName string) (result *GiftRecord) {
	this.UserName = UserName
	return this
}
func (this *GiftRecord) GetUserName() (UserName string) {
	return this.UserName
}
func (this *GiftRecord) SetUserMobile(UserMobile string) (result *GiftRecord) {
	this.UserMobile = UserMobile
	return this
}
func (this *GiftRecord) GetUserMobile() (UserMobile string) {
	return this.UserMobile
}
func (this *GiftRecord) SetGiftName(GiftName string) (result *GiftRecord) {
	this.GiftName = GiftName
	return this
}
func (this *GiftRecord) GetGiftName() (GiftName string) {
	return this.GiftName
}
func (this *GiftRecord) SetGiftCoin(GiftCoin string) (result *GiftRecord) {
	this.GiftCoin = GiftCoin
	return this
}
func (this *GiftRecord) GetGiftCoin() (GiftCoin string) {
	return this.GiftCoin
}
func (this *GiftRecord) SetGiftAmount(GiftAmount math.BigDecimal) (result *GiftRecord) {
	this.GiftAmount = GiftAmount
	return this
}
func (this *GiftRecord) GetGiftAmount() (GiftAmount math.BigDecimal) {
	return this.GiftAmount
}
func (this *GiftRecord) SetCreateTime(CreateTime xtime.Xtime) (result *GiftRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *GiftRecord) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func NewGiftRecord(id int64, userId int64, userName string, userMobile string, giftName string, giftCoin string, giftAmount math.BigDecimal, createTime xtime.Xtime) (ret *GiftRecord) {
	ret = new(GiftRecord)
	ret.Id = id
	ret.UserId = userId
	ret.UserName = userName
	ret.UserMobile = userMobile
	ret.GiftName = giftName
	ret.GiftCoin = giftCoin
	ret.GiftAmount = giftAmount
	ret.CreateTime = createTime
	return
}

type GiftRecord struct {
	Id         int64           `gorm:"column:id" json:"id"`
	UserId     int64           `gorm:"column:user_id" json:"userId"`
	UserName   string          `gorm:"column:user_name" json:"userName"`
	UserMobile string          `gorm:"column:user_mobile" json:"userMobile"`
	GiftName   string          `gorm:"column:gift_name" json:"giftName"`
	GiftCoin   string          `gorm:"column:gift_coin" json:"giftCoin"`
	GiftAmount math.BigDecimal `gorm:"column:gift_amount" json:"giftAmount"`
	CreateTime xtime.Xtime     `gorm:"column:create_time" json:"createTime"`
}
