package BooleanEnum

type BooleanEnum int

const (
	IS_FALSE BooleanEnum = iota
	IS_TRUE
)

func (this BooleanEnum) String() string {
	switch this {
	case IS_FALSE:
		return "否"
	case IS_TRUE:
		return "是"
	}
	return ""
}
