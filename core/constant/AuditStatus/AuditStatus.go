package AuditStatus

type AuditStatus int

const (
	AUDIT_ING AuditStatus = iota
	AUDIT_DEFEATED
	AUDIT_SUCCESS
)

func (this AuditStatus) String() string {
	switch this {
	case AUDIT_ING:
		return "待审核"
	case AUDIT_DEFEATED:
		return "审核失败"
	case AUDIT_SUCCESS:
		return "审核成功"
	}
	return ""
}
