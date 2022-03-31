package CertifiedBusinessStatus

var (
	NOT_CERTIFIED  = CertifiedBusinessStatus{"未认证", 0}
	AUDITING       = CertifiedBusinessStatus{"认证-待审核", 1}
	VERIFIED       = CertifiedBusinessStatus{"认证-审核成功", 2}
	FAILED         = CertifiedBusinessStatus{"认证-审核失败", 3}
	DEPOSIT_LESS   = CertifiedBusinessStatus{"保证金不足", 4}
	CANCEL_AUTH    = CertifiedBusinessStatus{"退保-待审核", 5}
	RETURN_FAILED  = CertifiedBusinessStatus{"退保-审核失败", 6}
	RETURN_SUCCESS = CertifiedBusinessStatus{"退保-审核成功", 7}
)

func (this *CertifiedBusinessStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *CertifiedBusinessStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type CertifiedBusinessStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
