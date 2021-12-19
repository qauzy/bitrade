package constant

type MemberLevelEnum int

const (
	MEMBERLEVELENUM_GENERAL MemberLevelEnum = iota
	MEMBERLEVELENUM_REALNAME
	MEMBERLEVELENUM_IDENTIFICATION
)

func (this MemberLevelEnum) String() string {
	switch this {
	case MEMBERLEVELENUM_GENERAL:
		return "普通"
	case MEMBERLEVELENUM_REALNAME:
		return "实名"
	case MEMBERLEVELENUM_IDENTIFICATION:
		return "认证商家"
	}
	return ""
}
