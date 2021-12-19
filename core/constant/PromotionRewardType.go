package constant

type PromotionRewardType int

const (
	PROMOTIONREWARDTYPE_REGISTER PromotionRewardType = iota
	PROMOTIONREWARDTYPE_TRANSACTION
	PROMOTIONREWARDTYPE_EXCHANGE_TRANSACTION
)

func (this PromotionRewardType) String() string {
	switch this {
	case PROMOTIONREWARDTYPE_REGISTER:
		return "推广注册"
	case PROMOTIONREWARDTYPE_TRANSACTION:
		return "法币推广交易"
	case PROMOTIONREWARDTYPE_EXCHANGE_TRANSACTION:
		return "币币推广交易"
	}
	return ""
}
