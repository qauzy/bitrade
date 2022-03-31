package RewardRecordType

var (
	PROMOTION = RewardRecordType{"推广", 0}
	ACTIVITY  = RewardRecordType{"活动", 1}
	DIVIDEND  = RewardRecordType{"分红", 2}
)

func (this *RewardRecordType) Ordinal() (result int) {
	return this.ordinal
}
func (this *RewardRecordType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type RewardRecordType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
