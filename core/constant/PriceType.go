package constant

type PriceType int

const (
	PRICETYPE_REGULAR PriceType = iota
	PRICETYPE_MUTATIVE
)

func (this PriceType) String() string {
	switch this {
	case PRICETYPE_REGULAR:
		return "固定的"
	case PRICETYPE_MUTATIVE:
		return "变化的"
	}
	return ""
}
