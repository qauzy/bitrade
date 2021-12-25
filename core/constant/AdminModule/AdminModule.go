package AdminModule

type AdminModule int

const (
	CMS AdminModule = iota
	COMMON
	EXCHANGE
	FINANCE
	MEMBER
	OTC
	SYSTEM
	PROMOTION
	INDEX
	IEO
	GIFT
	MARGIN
)

func (this AdminModule) String() string {
	switch this {
	case CMS:
		return "CMS"
	case COMMON:
		return "COMMON"
	case EXCHANGE:
		return "EXCHANGE"
	case FINANCE:
		return "FINANCE"
	case MEMBER:
		return "MEMBER"
	case OTC:
		return "OTC"
	case SYSTEM:
		return "SYSTEM"
	case PROMOTION:
		return "PROMOTION"
	case INDEX:
		return "INDEX"
	case IEO:
		return "IEO"
	case GIFT:
		return "GIFT"
	case MARGIN:
		return "MARGIN"
	}
	return ""
}
