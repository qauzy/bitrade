package CommonStatus

type CommonStatus int

const (
	NORMAL CommonStatus = iota
	ILLEGAL
)

func (this CommonStatus) String() string {
	switch this {
	case NORMAL:
		return "正常"
	case ILLEGAL:
		return "非法"
	}
	return ""
}
