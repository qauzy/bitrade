package ActivityRewardType

var (
	REGISTER    = ActivityRewardType{"注册奖励", 0}
	TRANSACTION = ActivityRewardType{"交易奖励", 1}
	RECHARGE    = ActivityRewardType{"充值奖励", 2}
)

func (this *ActivityRewardType) Ordinal() (result int) {
	return this.ordinal
}
func (this *ActivityRewardType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type ActivityRewardType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
