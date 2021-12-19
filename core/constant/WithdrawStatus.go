package constant

type WithdrawStatus int

const (
	WITHDRAWSTATUS_PROCESSING WithdrawStatus = iota
	WITHDRAWSTATUS_WAITING
	WITHDRAWSTATUS_FAIL
	WITHDRAWSTATUS_SUCCESS
	WITHDRAWSTATUS_TRANSFER
)

func (this WithdrawStatus) String() string {
	switch this {
	case WITHDRAWSTATUS_PROCESSING:
		return "审核中"
	case WITHDRAWSTATUS_WAITING:
		return "等待放币"
	case WITHDRAWSTATUS_FAIL:
		return "失败"
	case WITHDRAWSTATUS_SUCCESS:
		return "成功"
	case WITHDRAWSTATUS_TRANSFER:
		return "转账中"
	}
	return ""
}
