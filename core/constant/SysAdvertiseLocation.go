package constant

type SysAdvertiseLocation int

const (
	SYSADVERTISELOCATION_APP_SHUFFLING SysAdvertiseLocation = iota
	SYSADVERTISELOCATION_PC_SHUFFLING
	SYSADVERTISELOCATION_PC_CLASSIFICATION
)

func (this SysAdvertiseLocation) String() string {
	switch this {
	case SYSADVERTISELOCATION_APP_SHUFFLING:
		return "app首页轮播"
	case SYSADVERTISELOCATION_PC_SHUFFLING:
		return "pc首页轮播"
	case SYSADVERTISELOCATION_PC_CLASSIFICATION:
		return "pc分类广告"
	}
	return ""
}
