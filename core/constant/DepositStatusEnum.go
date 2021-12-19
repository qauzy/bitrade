package constant

type DepositStatusEnum int

const (
	DEPOSITSTATUSENUM_PAY DepositStatusEnum = iota
	DEPOSITSTATUSENUM_GET_BACK
)

func (this DepositStatusEnum) String() string {
	switch this {
	case DEPOSITSTATUSENUM_PAY:
		return "缴纳"
	case DEPOSITSTATUSENUM_GET_BACK:
		return "索回"
	}
	return ""
}
