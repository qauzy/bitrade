package WithdrawStatus

type WithdrawStatus int

const (
	PROCESSING WithdrawStatus = iota
	WAITING
	FAIL
	SUCCESS
	TRANSFER
)

func (this WithdrawStatus) String() string {
	switch this {
	case PROCESSING:
		return "审核中"
	case WAITING:
		return "等待放币"
	case FAIL:
		return "失败"
	case SUCCESS:
		return "成功"
	case TRANSFER:
		return "转账中"
	}
	return ""
}
