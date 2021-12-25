package ActivityRewardType

type ActivityRewardType int

const (
	REGISTER ActivityRewardType = iota
	TRANSACTION
	RECHARGE
)

func (this ActivityRewardType) String() string {
	switch this {
	case REGISTER:
		return "注册奖励"
	case TRANSACTION:
		return "交易奖励"
	case RECHARGE:
		return "充值奖励"
	}
	return ""
}
