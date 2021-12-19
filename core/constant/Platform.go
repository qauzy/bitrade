package constant

type Platform int

const (
	PLATFORM_ANDROID Platform = iota
	PLATFORM_IOS
)

func (this Platform) String() string {
	switch this {
	case PLATFORM_ANDROID:
		return "安卓"
	case PLATFORM_IOS:
		return "苹果"
	}
	return ""
}
