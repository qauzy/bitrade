package constant

type PromotionLevel int

const (
	PROMOTIONLEVEL_ONE PromotionLevel = iota
	PROMOTIONLEVEL_TWO
	PROMOTIONLEVEL_THREE
)

func (this PromotionLevel) String() string {
	switch this {
	case PROMOTIONLEVEL_ONE:
		return "一级"
	case PROMOTIONLEVEL_TWO:
		return "二级"
	case PROMOTIONLEVEL_THREE:
		return "三级"
	}
	return ""
}
