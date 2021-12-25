package AppealStatus

type AppealStatus int

const (
	NOT_PROCESSED AppealStatus = iota
	PROCESSED
)

func (this AppealStatus) String() string {
	switch this {
	case NOT_PROCESSED:
		return "未处理"
	case PROCESSED:
		return "已处理"
	}
	return ""
}
