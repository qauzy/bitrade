package entity

import (
	"github.com/qauzy/math"
	"time"
)

func (this *PromotionRewardRecord) SetSymbol(symbol string) (result *PromotionRewardRecord) {
	this.Symbol = symbol
	return this
}
func (this *PromotionRewardRecord) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *PromotionRewardRecord) SetRemark(remark string) (result *PromotionRewardRecord) {
	this.Remark = remark
	return this
}
func (this *PromotionRewardRecord) GetRemark() (remark string) {
	return this.Remark
}
func (this *PromotionRewardRecord) SetAmount(amount math.BigDecimal) (result *PromotionRewardRecord) {
	this.Amount = amount
	return this
}
func (this *PromotionRewardRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *PromotionRewardRecord) SetCreateTime(createTime time.Time) (result *PromotionRewardRecord) {
	this.CreateTime = createTime
	return this
}
func (this *PromotionRewardRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}

type PromotionRewardRecord struct {
	Symbol     string
	Remark     string
	Amount     math.BigDecimal
	CreateTime time.Time
}
