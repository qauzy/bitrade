package entity

import (
	"bitrade/core/constant/PromotionLevel"
	"time"
)

func (this *PromotionMember) SetCreateTime(createTime time.Time) (result *PromotionMember) {
	this.CreateTime = createTime
	return this
}
func (this *PromotionMember) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *PromotionMember) SetUsername(username string) (result *PromotionMember) {
	this.Username = username
	return this
}
func (this *PromotionMember) GetUsername() (username string) {
	return this.Username
}
func (this *PromotionMember) SetLevel(level PromotionLevel.PromotionLevel) (result *PromotionMember) {
	this.Level = level
	return this
}
func (this *PromotionMember) GetLevel() (level PromotionLevel.PromotionLevel) {
	return this.Level
}

type PromotionMember struct {
	CreateTime time.Time
	Username   string
	Level      PromotionLevel.PromotionLevel
}
