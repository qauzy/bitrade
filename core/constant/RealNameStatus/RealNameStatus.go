package RealNameStatus

var (
	NOT_CERTIFIED = RealNameStatus{"未认证", 0}
	AUDITING      = RealNameStatus{"审核中", 1}
	VERIFIED      = RealNameStatus{"已认证", 2}
)

func (this *RealNameStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *RealNameStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type RealNameStatus struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
