package AdvertiseControlStatus

var (
	PUT_ON_SHELVES  = AdvertiseControlStatus{"上架", 0}
	PUT_OFF_SHELVES = AdvertiseControlStatus{"下架", 1}
	TURNOFF         = AdvertiseControlStatus{"已关闭", 2}
)

func (this *AdvertiseControlStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseControlStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseControlStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
