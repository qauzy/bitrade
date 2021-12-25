package AdvertiseStatus

type AdvertiseStatus int

const (
	HANG AdvertiseStatus = iota
	PAYMENT
	TURNOFF
)

func (this AdvertiseStatus) String() string {
	switch this {
	case HANG:
		return "已挂单"
	case PAYMENT:
		return "待付款"
	case TURNOFF:
		return "已关闭"
	}
	return ""
}
