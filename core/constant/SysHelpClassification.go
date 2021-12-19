package constant

type SysHelpClassification int

const (
	SYSHELPCLASSIFICATION_HELP SysHelpClassification = iota
	SYSHELPCLASSIFICATION_FAQ
	SYSHELPCLASSIFICATION_RECHARGE
	SYSHELPCLASSIFICATION_TRANSACTION
	SYSHELPCLASSIFICATION_QR_CODE
)

func (this SysHelpClassification) String() string {
	switch this {
	case SYSHELPCLASSIFICATION_HELP:
		return "新手入门"
	case SYSHELPCLASSIFICATION_FAQ:
		return "常见问题"
	case SYSHELPCLASSIFICATION_RECHARGE:
		return "充值指南"
	case SYSHELPCLASSIFICATION_TRANSACTION:
		return "交易指南"
	case SYSHELPCLASSIFICATION_QR_CODE:
		return "APP二维码"
	}
	return ""
}
