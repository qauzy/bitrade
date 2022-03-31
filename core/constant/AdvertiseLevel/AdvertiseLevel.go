package AdvertiseLevel

var (
	ORDINARY  = AdvertiseLevel{"普通", 0}
	EXCELLENT = AdvertiseLevel{"优质", 1}
)

func (this *AdvertiseLevel) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseLevel) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseLevel struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
