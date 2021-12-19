package constant

type AppealStatus int

const (
	APPEALSTATUS_NOT_PROCESSED AppealStatus = iota
	APPEALSTATUS_PROCESSED
)

func (this AppealStatus) String() string {
	switch this {
	case APPEALSTATUS_NOT_PROCESSED:
		return "未处理"
	case APPEALSTATUS_PROCESSED:
		return "已处理"
	}
	return ""
}
