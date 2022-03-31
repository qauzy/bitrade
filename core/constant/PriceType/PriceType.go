package PriceType

var (
	REGULAR  = PriceType{"固定的", 0}
	MUTATIVE = PriceType{"变化的", 1}
)

func (this *PriceType) Ordinal() (result int) {
	return this.ordinal
}
func (this *PriceType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PriceType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
