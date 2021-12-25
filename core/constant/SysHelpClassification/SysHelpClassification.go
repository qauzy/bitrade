package SysHelpClassification

type SysHelpClassification int

const (
	HELP SysHelpClassification = iota
	FAQ
	RECHARGE
	TRANSACTION
	QR_CODE
)

func (this SysHelpClassification) String() string {
	switch this {
	case HELP:
		return "新手入门"
	case FAQ:
		return "常见问题"
	case RECHARGE:
		return "充值指南"
	case TRANSACTION:
		return "交易指南"
	case QR_CODE:
		return "APP二维码"
	}
	return ""
}
