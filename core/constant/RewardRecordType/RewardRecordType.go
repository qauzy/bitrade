package RewardRecordType

type RewardRecordType int

const (
	PROMOTION RewardRecordType = iota
	ACTIVITY
	DIVIDEND
)

func (this RewardRecordType) String() string {
	switch this {
	case PROMOTION:
		return "推广"
	case ACTIVITY:
		return "活动"
	case DIVIDEND:
		return "分红"
	}
	return ""
}
