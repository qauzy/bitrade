package PriceType

type PriceType int

const (
	REGULAR PriceType = iota
	MUTATIVE
)

func (this PriceType) String() string {
	switch this {
	case REGULAR:
		return "固定的"
	case MUTATIVE:
		return "变化的"
	}
	return ""
}
