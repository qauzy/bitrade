package PromotionLevel

type PromotionLevel int

const (
	ONE PromotionLevel = iota
	TWO
	THREE
)

func (this PromotionLevel) String() string {
	switch this {
	case ONE:
		return "一级"
	case TWO:
		return "二级"
	case THREE:
		return "三级"
	}
	return ""
}
