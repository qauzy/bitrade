package TransactionTypeEnum

var (
	OTC_NUM       = TransactionTypeEnum{"法币成交量", 0}
	OTC_MONEY     = TransactionTypeEnum{"法币成交额", 1}
	EXCHANGE      = TransactionTypeEnum{"币币交易手续费统计", 2}
	EXCHANGE_COIN = TransactionTypeEnum{"币币交易成交量统计", 3}
	EXCHANGE_BASE = TransactionTypeEnum{"币币交易成交额统计", 4}
	RECHARGE      = TransactionTypeEnum{"充币", 5}
	WITHDRAW      = TransactionTypeEnum{"提币", 6}
)

func (this *TransactionTypeEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *TransactionTypeEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type TransactionTypeEnum struct {
	CnName  string `gorm:"column:cn_name" json:"cnName"`
	ordinal int
}
