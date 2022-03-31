package PromotionRewardType

var (
	REGISTER             = PromotionRewardType{"推广注册", 0}
	TRANSACTION          = PromotionRewardType{"法币推广交易", 1}
	EXCHANGE_TRANSACTION = PromotionRewardType{"币币推广交易", 2}
)

func (this *PromotionRewardType) Ordinal() (result int) {
	return this.ordinal
}
func (this *PromotionRewardType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PromotionRewardType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
