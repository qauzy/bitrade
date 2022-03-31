package Platform

var (
	ANDROID = Platform{"安卓", 0}
	IOS     = Platform{"苹果", 1}
)

func (this *Platform) Ordinal() (result int) {
	return this.ordinal
}
func (this *Platform) GetOrdinal() (result int) {
	return this.Ordinal()
}

type Platform struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
