package AdvertiseStatus

var (
	HANG    = AdvertiseStatus{"已挂单", 0}
	PAYMENT = AdvertiseStatus{"待付款", 1}
	TURNOFF = AdvertiseStatus{"已关闭", 2}
)

func (this *AdvertiseStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
