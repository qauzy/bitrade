package SysAdvertiseLocation

var (
	APP_SHUFFLING     = SysAdvertiseLocation{"app首页轮播", 0}
	PC_SHUFFLING      = SysAdvertiseLocation{"pc首页轮播", 1}
	PC_CLASSIFICATION = SysAdvertiseLocation{"pc分类广告", 2}
)

func (this *SysAdvertiseLocation) Ordinal() (result int) {
	return this.ordinal
}
func (this *SysAdvertiseLocation) GetOrdinal() (result int) {
	return this.Ordinal()
}

type SysAdvertiseLocation struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
