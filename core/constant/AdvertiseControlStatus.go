package constant

type AdvertiseControlStatus int

const (
	ADVERTISECONTROLSTATUS_PUT_ON_SHELVES AdvertiseControlStatus = iota
	ADVERTISECONTROLSTATUS_PUT_OFF_SHELVES
	ADVERTISECONTROLSTATUS_TURNOFF
)

func (this AdvertiseControlStatus) String() string {
	switch this {
	case ADVERTISECONTROLSTATUS_PUT_ON_SHELVES:
		return "上架"
	case ADVERTISECONTROLSTATUS_PUT_OFF_SHELVES:
		return "下架"
	case ADVERTISECONTROLSTATUS_TURNOFF:
		return "已关闭"
	}
	return ""
}
