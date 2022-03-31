package AuditStatus

var (
	AUDIT_ING      = AuditStatus{"待审核", 0}
	AUDIT_DEFEATED = AuditStatus{"审核失败", 1}
	AUDIT_SUCCESS  = AuditStatus{"审核成功", 2}
)

func (this *AuditStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AuditStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AuditStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
