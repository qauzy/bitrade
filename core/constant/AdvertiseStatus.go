package constant

type AdvertiseStatus int

const (
	ADVERTISESTATUS_HANG AdvertiseStatus = iota
	ADVERTISESTATUS_PAYMENT
	ADVERTISESTATUS_TURNOFF
)

func (this AdvertiseStatus) String() string {
	switch this {
	case ADVERTISESTATUS_HANG:
		return "已挂单"
	case ADVERTISESTATUS_PAYMENT:
		return "待付款"
	case ADVERTISESTATUS_TURNOFF:
		return "已关闭"
	}
	return ""
}
