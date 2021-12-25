package BooleanEnum

type BooleanEnum int

const (
	IS_FALSE BooleanEnum = iota
	IS_TRUE
)

func (this BooleanEnum) String() string {
	switch this {
	case IS_FALSE:
		return false
	case IS_TRUE:
		return true
	}
	return ""
}
