package constant

type AdvertiseLevel int

const (
	ADVERTISELEVEL_ORDINARY AdvertiseLevel = iota
	ADVERTISELEVEL_EXCELLENT
)

func (this AdvertiseLevel) String() string {
	switch this {
	case ADVERTISELEVEL_ORDINARY:
		return "普通"
	case ADVERTISELEVEL_EXCELLENT:
		return "优质"
	}
	return ""
}
