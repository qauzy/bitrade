package constant

type RealNameStatus int

const (
	REALNAMESTATUS_NOT_CERTIFIED RealNameStatus = iota
	REALNAMESTATUS_AUDITING
	REALNAMESTATUS_VERIFIED
)

func (this RealNameStatus) String() string {
	switch this {
	case REALNAMESTATUS_NOT_CERTIFIED:
		return "未认证"
	case REALNAMESTATUS_AUDITING:
		return "审核中"
	case REALNAMESTATUS_VERIFIED:
		return "已认证"
	}
	return ""
}
