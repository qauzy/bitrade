package constant

type CommonStatus int

const (
	COMMONSTATUS_NORMAL CommonStatus = iota
	COMMONSTATUS_ILLEGAL
)

func (this CommonStatus) String() string {
	switch this {
	case COMMONSTATUS_NORMAL:
		return "正常"
	case COMMONSTATUS_ILLEGAL:
		return "非法"
	}
	return ""
}
