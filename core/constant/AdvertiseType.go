package constant

type AdvertiseType int

const (
	ADVERTISETYPE_BUY AdvertiseType = iota
	ADVERTISETYPE_SELL
)

func (this AdvertiseType) String() string {
	switch this {
	case ADVERTISETYPE_BUY:
		return "购买"
	case ADVERTISETYPE_SELL:
		return "出售"
	}
	return ""
}
