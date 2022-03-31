package SysHelpClassification

var (
	HELP        = SysHelpClassification{"新手入门", 0}
	FAQ         = SysHelpClassification{"常见问题", 1}
	RECHARGE    = SysHelpClassification{"充值指南", 2}
	TRANSACTION = SysHelpClassification{"交易指南", 3}
	QR_CODE     = SysHelpClassification{"APP二维码", 4}
)

func (this *SysHelpClassification) Ordinal() (result int) {
	return this.ordinal
}
func (this *SysHelpClassification) GetOrdinal() (result int) {
	return this.Ordinal()
}

type SysHelpClassification struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
