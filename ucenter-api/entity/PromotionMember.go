package entity

import (
	"bitrade/core/constant/PromotionLevel"
	"time"
)

func (this *PromotionMember) SetCreateTime(CreateTime time.Time) (result *PromotionMember) {
	this.CreateTime = CreateTime
	return this
}
func (this *PromotionMember) GetCreateTime() (CreateTime time.Time) {
	return this.CreateTime
}
func (this *PromotionMember) SetUsername(Username string) (result *PromotionMember) {
	this.Username = Username
	return this
}
func (this *PromotionMember) GetUsername() (Username string) {
	return this.Username
}
func (this *PromotionMember) SetLevel(Level *PromotionLevel.PromotionLevel) (result *PromotionMember) {
	this.Level = Level
	return this
}
func (this *PromotionMember) GetLevel() (Level *PromotionLevel.PromotionLevel) {
	return this.Level
}

type PromotionMember struct {
	CreateTime time.Time                      `gorm:"column:create_time" json:"createTime"`
	Username   string                         `gorm:"column:username" json:"username"`
	Level      *PromotionLevel.PromotionLevel `gorm:"column:level" json:"level"`
}
