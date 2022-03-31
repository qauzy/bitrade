package BooleanEnum

var (
	IS_FALSE = BooleanEnum{false, "否", 0}
	IS_TRUE  = BooleanEnum{true, "是", 1}
)

func (this *BooleanEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *BooleanEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type BooleanEnum struct {
	Is      bool   `gorm:"column:is" json:"is"`
	NameCn  string `gorm:"column:name_cn" json:"nameCn"`
	ordinal int
}
