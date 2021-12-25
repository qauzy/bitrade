package MemberLevelEnum

type MemberLevelEnum int

const (
	GENERAL MemberLevelEnum = iota
	REALNAME
	IDENTIFICATION
)

func (this MemberLevelEnum) String() string {
	switch this {
	case GENERAL:
		return "普通"
	case REALNAME:
		return "实名"
	case IDENTIFICATION:
		return "认证商家"
	}
	return ""
}
