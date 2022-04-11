package entity

import (
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
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
func (this *PromotionRewardRecord) SetCreateTime(createTime xtime.Xtime) (result *PromotionRewardRecord) {
	this.CreateTime = createTime
	return this
}
func (this *PromotionRewardRecord) GetCreateTime() (createTime xtime.Xtime) {
	return this.CreateTime
}

type PromotionRewardRecord struct {
	Symbol     string
	Remark     string
	Amount     math.BigDecimal
	CreateTime xtime.Xtime
}
