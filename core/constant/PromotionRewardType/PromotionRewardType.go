package PromotionRewardType

type PromotionRewardType int

const (
	REGISTER PromotionRewardType = iota
	TRANSACTION
	EXCHANGE_TRANSACTION
)

func (this PromotionRewardType) String() string {
	switch this {
	case REGISTER:
		return "推广注册"
	case TRANSACTION:
		return "法币推广交易"
	case EXCHANGE_TRANSACTION:
		return "币币推广交易"
	}
	return ""
}
