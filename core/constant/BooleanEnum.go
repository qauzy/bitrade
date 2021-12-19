package constant

type BooleanEnum int

const (
	BOOLEANENUM_IS_FALSE BooleanEnum = iota
	BOOLEANENUM_IS_TRUE
)

func (this BooleanEnum) String() string {
	switch this {
	case BOOLEANENUM_IS_FALSE:
		return "否"
	case BOOLEANENUM_IS_TRUE:
		return "是"
	}
	return ""
}
