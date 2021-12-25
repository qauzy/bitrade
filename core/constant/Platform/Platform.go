package Platform

type Platform int

const (
	ANDROID Platform = iota
	IOS
)

func (this Platform) String() string {
	switch this {
	case ANDROID:
		return "安卓"
	case IOS:
		return "苹果"
	}
	return ""
}
