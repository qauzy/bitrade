package entity

import (
	"bitrade/core/constant/ActivityRewardType"
	"bitrade/core/constant/BooleanEnum"
	"github.com/qauzy/chocolate/xtime"
)

func (this *RewardActivitySetting) SetId(Id int64) (result *RewardActivitySetting) {
	this.Id = Id
	return this
}
func (this *RewardActivitySetting) GetId() (Id int64) {
	return this.Id
}
func (this *RewardActivitySetting) SetCoin(Coin *Coin) (result *RewardActivitySetting) {
	this.Coin = Coin
	return this
}
func (this *RewardActivitySetting) GetCoin() (Coin *Coin) {
	return this.Coin
}
func (this *RewardActivitySetting) SetStatus(Status BooleanEnum.BooleanEnum) (result *RewardActivitySetting) {
	this.Status = Status
	return this
}
func (this *RewardActivitySetting) GetStatus() (Status BooleanEnum.BooleanEnum) {
	return this.Status
}
func (this *RewardActivitySetting) SetType(Type ActivityRewardType.ActivityRewardType) (result *RewardActivitySetting) {
	this.Type = Type
	return this
}
func (this *RewardActivitySetting) GetType() (Type ActivityRewardType.ActivityRewardType) {
	return this.Type
}
func (this *RewardActivitySetting) SetInfo(Info string) (result *RewardActivitySetting) {
	this.Info = Info
	return this
}
func (this *RewardActivitySetting) GetInfo() (Info string) {
	return this.Info
}
func (this *RewardActivitySetting) SetUpdateTime(UpdateTime xtime.Xtime) (result *RewardActivitySetting) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *RewardActivitySetting) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func (this *RewardActivitySetting) SetAdmin(Admin *Admin) (result *RewardActivitySetting) {
	this.Admin = Admin
	return this
}
func (this *RewardActivitySetting) GetAdmin() (Admin *Admin) {
	return this.Admin
}
func NewRewardActivitySetting(id int64, coin *Coin, status BooleanEnum.BooleanEnum, oType ActivityRewardType.ActivityRewardType, info string, updateTime xtime.Xtime, admin *Admin) (ret *RewardActivitySetting) {
	ret = new(RewardActivitySetting)
	ret.Id = id
	ret.Coin = coin
	ret.Status = status
	ret.Type = oType
	ret.Info = info
	ret.UpdateTime = updateTime
	ret.Admin = admin
	return
}

type RewardActivitySetting struct {
	Id         int64                                 `gorm:"column:id" json:"id"`
	Coin       *Coin                                 `gorm:"column:coin" json:"coin"`
	Status     BooleanEnum.BooleanEnum               `gorm:"column:status" json:"status"`
	Type       ActivityRewardType.ActivityRewardType `gorm:"column:type" json:"type"`
	Info       string                                `gorm:"column:info" json:"info"`
	UpdateTime xtime.Xtime                           `gorm:"column:update_time" json:"updateTime"`
	Admin      *Admin                                `gorm:"column:admin" json:"admin"`
}
