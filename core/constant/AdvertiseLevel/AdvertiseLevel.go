package AdvertiseLevel

type AdvertiseLevel int

const (
	ORDINARY AdvertiseLevel = iota
	EXCELLENT
)

func (this AdvertiseLevel) String() string {
	switch this {
	case ORDINARY:
		return "普通"
	case EXCELLENT:
		return "优质"
	}
	return ""
}
