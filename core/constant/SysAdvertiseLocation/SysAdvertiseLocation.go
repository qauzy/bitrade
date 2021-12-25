package SysAdvertiseLocation

type SysAdvertiseLocation int

const (
	APP_SHUFFLING SysAdvertiseLocation = iota
	PC_SHUFFLING
	PC_CLASSIFICATION
)

func (this SysAdvertiseLocation) String() string {
	switch this {
	case APP_SHUFFLING:
		return "app首页轮播"
	case PC_SHUFFLING:
		return "pc首页轮播"
	case PC_CLASSIFICATION:
		return "pc分类广告"
	}
	return ""
}
