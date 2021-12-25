package MemberRegisterType

type MemberRegisterType int

const (
	BACKSTAGE MemberRegisterType = iota
	AUTONOMOUSLY
	AUTONOMOUSLY_RECOMMEND
)

func (this MemberRegisterType) String() string {
	switch this {
	case BACKSTAGE:
		return "后台"
	case AUTONOMOUSLY:
		return "自主"
	case AUTONOMOUSLY_RECOMMEND:
		return "自主(推荐)"
	}
	return ""
}
