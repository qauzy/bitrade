package DepositStatusEnum

type DepositStatusEnum int

const (
	PAY DepositStatusEnum = iota
	GET_BACK
)

func (this DepositStatusEnum) String() string {
	switch this {
	case PAY:
		return "缴纳"
	case GET_BACK:
		return "索回"
	}
	return ""
}
