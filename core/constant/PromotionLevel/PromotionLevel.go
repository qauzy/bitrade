package PromotionLevel

var (
	ONE   = PromotionLevel{"一级", 0}
	TWO   = PromotionLevel{"二级", 1}
	THREE = PromotionLevel{"三级", 2}
)

func (this *PromotionLevel) Ordinal() (result int) {
	return this.ordinal
}
func (this *PromotionLevel) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PromotionLevel struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
