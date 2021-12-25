package IntegrationRecordType

type IntegrationRecordType int

const (
	PROMOTION_GIVING IntegrationRecordType = iota
	LEGAL_RECHARGE_GIVING
	COIN_RECHARGE_GIVING
)

func (this IntegrationRecordType) String() string {
	switch this {
	case PROMOTION_GIVING:
		return "推广"
	case LEGAL_RECHARGE_GIVING:
		return "法币充值赠送"
	case COIN_RECHARGE_GIVING:
		return "币币充值赠送"
	}
	return ""
}
