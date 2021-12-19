package constant

type IntegrationRecordType int

const (
	INTEGRATIONRECORDTYPE_PROMOTION_GIVING IntegrationRecordType = iota
	INTEGRATIONRECORDTYPE_LEGAL_RECHARGE_GIVING
	INTEGRATIONRECORDTYPE_COIN_RECHARGE_GIVING
)

func (this IntegrationRecordType) String() string {
	switch this {
	case INTEGRATIONRECORDTYPE_PROMOTION_GIVING:
		return "推广"
	case INTEGRATIONRECORDTYPE_LEGAL_RECHARGE_GIVING:
		return "法币充值赠送"
	case INTEGRATIONRECORDTYPE_COIN_RECHARGE_GIVING:
		return "币币充值赠送"
	}
	return ""
}
