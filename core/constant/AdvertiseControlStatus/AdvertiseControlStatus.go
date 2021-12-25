package AdvertiseControlStatus

type AdvertiseControlStatus int

const (
	PUT_ON_SHELVES AdvertiseControlStatus = iota
	PUT_OFF_SHELVES
	TURNOFF
)

func (this AdvertiseControlStatus) String() string {
	switch this {
	case PUT_ON_SHELVES:
		return "上架"
	case PUT_OFF_SHELVES:
		return "下架"
	case TURNOFF:
		return "已关闭"
	}
	return ""
}
