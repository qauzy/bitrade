package constant

type RewardRecordType int

const (
	REWARDRECORDTYPE_PROMOTION RewardRecordType = iota
	REWARDRECORDTYPE_ACTIVITY
	REWARDRECORDTYPE_DIVIDEND
)

func (this RewardRecordType) String() string {
	switch this {
	case REWARDRECORDTYPE_PROMOTION:
		return "推广"
	case REWARDRECORDTYPE_ACTIVITY:
		return "活动"
	case REWARDRECORDTYPE_DIVIDEND:
		return "分红"
	}
	return ""
}
