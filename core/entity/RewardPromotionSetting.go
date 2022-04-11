package entity

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/PromotionRewardType"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *RewardPromotionSetting) SetId(Id int64) (result *RewardPromotionSetting) {
	this.Id = Id
	return this
}
func (this *RewardPromotionSetting) GetId() (Id int64) {
	return this.Id
}
func (this *RewardPromotionSetting) SetCoin(Coin *Coin) (result *RewardPromotionSetting) {
	this.Coin = Coin
	return this
}
func (this *RewardPromotionSetting) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *RewardPromotionSetting) SetStatus(Status BooleanEnum.BooleanEnum) (result *RewardPromotionSetting) {
	this.Status = Status
	return this
}
func (this *RewardPromotionSetting) GetStatus() (Status BooleanEnum.BooleanEnum) {
	return this.Status
}
func (this *RewardPromotionSetting) SetType(Type PromotionRewardType.PromotionRewardType) (result *RewardPromotionSetting) {
	this.Type = Type
	return this
}
func (this *RewardPromotionSetting) GetType() (Type PromotionRewardType.PromotionRewardType) {
	return this.Type
}
func (this *RewardPromotionSetting) SetInfo(Info string) (result *RewardPromotionSetting) {
	this.Info = Info
	return this
}
func (this *RewardPromotionSetting) GetInfo() (Info string) {
	return this.Info
}
func (this *RewardPromotionSetting) SetEffectiveTime(EffectiveTime int) (result *RewardPromotionSetting) {
	this.EffectiveTime = EffectiveTime
	return this
}
func (this *RewardPromotionSetting) GetEffectiveTime() (EffectiveTime int) {
	return this.EffectiveTime
}
func (this *RewardPromotionSetting) SetUpdateTime(UpdateTime xtime.Xtime) (result *RewardPromotionSetting) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *RewardPromotionSetting) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func (this *RewardPromotionSetting) SetAdmin(Admin *Admin) (result *RewardPromotionSetting) {
	this.Admin = Admin
	return this
}
func (this *RewardPromotionSetting) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func (this *RewardPromotionSetting) SetOne(One math.BigDecimal) (result *RewardPromotionSetting) {
	this.One = One
	return this
}
func (this *RewardPromotionSetting) GetOne() (One math.BigDecimal) {
	return this.One
}
func (this *RewardPromotionSetting) SetTwo(Two math.BigDecimal) (result *RewardPromotionSetting) {
	this.Two = Two
	return this
}
func (this *RewardPromotionSetting) GetTwo() (Two math.BigDecimal) {
	return this.Two
}
func NewRewardPromotionSetting(id int64, coin *Coin, status BooleanEnum.BooleanEnum, oType PromotionRewardType.PromotionRewardType, info string, effectiveTime int, updateTime xtime.Xtime, admin *Admin, one math.BigDecimal, two math.BigDecimal) (ret *RewardPromotionSetting) {
	ret = new(RewardPromotionSetting)
	ret.Id = id
	ret.Coin = coin
	ret.Status = status
	ret.Type = oType
	ret.Info = info
	ret.EffectiveTime = effectiveTime
	ret.UpdateTime = updateTime
	ret.Admin = admin
	ret.One = one
	ret.Two = two
	return
}

type RewardPromotionSetting struct {
	Id            int64                                   `gorm:"column:id" json:"id"`
	Coin          *Coin                                   `gorm:"column:coin" json:"coin"`
	Status        BooleanEnum.BooleanEnum                 `gorm:"column:status" json:"status"`
	Type          PromotionRewardType.PromotionRewardType `gorm:"column:type" json:"type"`
	Info          string                                  `gorm:"column:info" json:"info"`
	EffectiveTime int                                     `gorm:"column:effective_time" json:"effectiveTime"`
	UpdateTime    xtime.Xtime                             `gorm:"column:update_time" json:"updateTime"`
	Admin         *Admin                                  `gorm:"column:admin" json:"admin"`
	One           math.BigDecimal                         `gorm:"column:one" json:"one"`
	Two           math.BigDecimal                         `gorm:"column:two" json:"two"`
}
