package BusinessStatus

type BusinessStatus int

const (
	ZERO BusinessStatus = iota
	PORTION
	ALL
)

func (this BusinessStatus) String() string {
	switch this {
	case ZERO:
		return "为成交"
	case PORTION:
		return "部分成交"
	case ALL:
		return "全部成交"
	}
	return ""
}
