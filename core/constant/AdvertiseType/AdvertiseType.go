package AdvertiseType

var (
	BUY  = AdvertiseType{"购买", 0}
	SELL = AdvertiseType{"出售", 1}
)

func (this *AdvertiseType) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
