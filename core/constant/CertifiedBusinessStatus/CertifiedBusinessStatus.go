package CertifiedBusinessStatus

type CertifiedBusinessStatus int

const (
	NOT_CERTIFIED CertifiedBusinessStatus = iota
	AUDITING
	VERIFIED
	FAILED
	DEPOSIT_LESS
	CANCEL_AUTH
	RETURN_FAILED
	RETURN_SUCCESS
)

func (this CertifiedBusinessStatus) String() string {
	switch this {
	case NOT_CERTIFIED:
		return "未认证"
	case AUDITING:
		return "认证-待审核"
	case VERIFIED:
		return "认证-审核成功"
	case FAILED:
		return "认证-审核失败"
	case DEPOSIT_LESS:
		return "保证金不足"
	case CANCEL_AUTH:
		return "退保-待审核"
	case RETURN_FAILED:
		return "退保-审核失败"
	case RETURN_SUCCESS:
		return "退保-审核成功"
	}
	return ""
}
