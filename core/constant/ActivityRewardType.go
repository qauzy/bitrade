package constant

type ActivityRewardType int

const (
	ACTIVITYREWARDTYPE_REGISTER ActivityRewardType = iota
	ACTIVITYREWARDTYPE_TRANSACTION
	ACTIVITYREWARDTYPE_RECHARGE
)

func (this ActivityRewardType) String() string {
	switch this {
	case ACTIVITYREWARDTYPE_REGISTER:
		return "注册奖励"
	case ACTIVITYREWARDTYPE_TRANSACTION:
		return "交易奖励"
	case ACTIVITYREWARDTYPE_RECHARGE:
		return "充值奖励"
	}
	return ""
}
