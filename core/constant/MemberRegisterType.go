package constant

type MemberRegisterType int

const (
	MEMBERREGISTERTYPE_BACKSTAGE MemberRegisterType = iota
	MEMBERREGISTERTYPE_AUTONOMOUSLY
	MEMBERREGISTERTYPE_AUTONOMOUSLY_RECOMMEND
)

func (this MemberRegisterType) String() string {
	switch this {
	case MEMBERREGISTERTYPE_BACKSTAGE:
		return "后台"
	case MEMBERREGISTERTYPE_AUTONOMOUSLY:
		return "自主"
	case MEMBERREGISTERTYPE_AUTONOMOUSLY_RECOMMEND:
		return "自主(推荐)"
	}
	return ""
}
