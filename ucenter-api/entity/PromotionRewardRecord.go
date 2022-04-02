package entity

import (
	"github.com/qauzy/math"
	"time"
)

func (this *PromotionRewardRecord) SetSymbol(Symbol string) (result *PromotionRewardRecord) {
	this.Symbol = Symbol
	return this
}
func (this *PromotionRewardRecord) GetSymbol() (Symbol string) {
	return this.Symbol
}
func (this *PromotionRewardRecord) SetRemark(Remark string) (result *PromotionRewardRecord) {
	this.Remark = Remark
	return this
}
func (this *PromotionRewardRecord) GetRemark() (Remark string) {
	return this.Remark
}
func (this *PromotionRewardRecord) SetAmount(Amount *math.BigDecimal) (result *PromotionRewardRecord) {
	this.Amount = Amount
	return this
}
func (this *PromotionRewardRecord) GetAmount() (Amount *math.BigDecimal) {
	return this.Amount
}
func (this *PromotionRewardRecord) SetCreateTime(CreateTime time.Time) (result *PromotionRewardRecord) {
	this.CreateTime = CreateTime
	return this
}
func (this *PromotionRewardRecord) GetCreateTime() (CreateTime time.Time) {
	return this.CreateTime
}

type PromotionRewardRecord struct {
	Symbol     string           `gorm:"column:symbol" json:"symbol"`
	Remark     string           `gorm:"column:remark" json:"remark"`
	Amount     *math.BigDecimal `gorm:"column:amount" json:"amount"`
	CreateTime time.Time        `gorm:"column:create_time" json:"createTime"`
}
