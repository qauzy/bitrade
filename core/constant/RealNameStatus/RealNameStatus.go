package RealNameStatus

type RealNameStatus int

const (
	NOT_CERTIFIED RealNameStatus = iota
	AUDITING
	VERIFIED
)

func (this RealNameStatus) String() string {
	switch this {
	case NOT_CERTIFIED:
		return "未认证"
	case AUDITING:
		return "审核中"
	case VERIFIED:
		return "已认证"
	}
	return ""
}
