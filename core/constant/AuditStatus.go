package constant

type AuditStatus int

const (
	AUDITSTATUS_AUDIT_ING AuditStatus = iota
	AUDITSTATUS_AUDIT_DEFEATED
	AUDITSTATUS_AUDIT_SUCCESS
)

func (this AuditStatus) String() string {
	switch this {
	case AUDITSTATUS_AUDIT_ING:
		return "待审核"
	case AUDITSTATUS_AUDIT_DEFEATED:
		return "审核失败"
	case AUDITSTATUS_AUDIT_SUCCESS:
		return "审核成功"
	}
	return ""
}
