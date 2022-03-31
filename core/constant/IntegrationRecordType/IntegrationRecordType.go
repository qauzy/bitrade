package IntegrationRecordType

var (
	PROMOTION_GIVING      = IntegrationRecordType{"推广", 0}
	LEGAL_RECHARGE_GIVING = IntegrationRecordType{"法币充值赠送", 1}
	COIN_RECHARGE_GIVING  = IntegrationRecordType{"币币充值赠送", 2}
)

func (this *IntegrationRecordType) Ordinal() (result int) {
	return this.ordinal
}
func (this *IntegrationRecordType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type IntegrationRecordType struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
