package TransactionTypeEnum

type TransactionTypeEnum int

const (
	OTC_NUM TransactionTypeEnum = iota
	OTC_MONEY
	EXCHANGE
	EXCHANGE_COIN
	EXCHANGE_BASE
	RECHARGE
	WITHDRAW
)

func (this TransactionTypeEnum) String() string {
	switch this {
	case OTC_NUM:
		return "法币成交量"
	case OTC_MONEY:
		return "法币成交额"
	case EXCHANGE:
		return "币币交易手续费统计"
	case EXCHANGE_COIN:
		return "币币交易成交量统计"
	case EXCHANGE_BASE:
		return "币币交易成交额统计"
	case RECHARGE:
		return "充币"
	case WITHDRAW:
		return "提币"
	}
	return ""
}
