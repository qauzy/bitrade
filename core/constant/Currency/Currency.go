package Currency

var BTC = Currency{"比特币", 0}

func (this *Currency) Ordinal() (result int) {
	return this.ordinal
}
func (this *Currency) GetOrdinal() (result int) {
	return this.Ordinal()
}

type Currency struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
