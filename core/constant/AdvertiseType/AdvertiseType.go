package AdvertiseType

type AdvertiseType int

const (
	BUY AdvertiseType = iota
	SELL
)

func (this AdvertiseType) String() string {
	switch this {
	case BUY:
		return "购买"
	case SELL:
		return "出售"
	}
	return ""
}
