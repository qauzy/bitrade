package constant

type BusinessStatus int

const (
	BUSINESSSTATUS_ZERO BusinessStatus = iota
	BUSINESSSTATUS_PORTION
	BUSINESSSTATUS_ALL
)

func (this BusinessStatus) String() string {
	switch this {
	case BUSINESSSTATUS_ZERO:
		return "为成交"
	case BUSINESSSTATUS_PORTION:
		return "部分成交"
	case BUSINESSSTATUS_ALL:
		return "全部成交"
	}
	return ""
}
