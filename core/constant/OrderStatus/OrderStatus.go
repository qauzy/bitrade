package OrderStatus

type OrderStatus int

const (
	CANCELLED OrderStatus = iota
	NONPAYMENT
	PAID
	COMPLETED
	APPEAL
)

func (this OrderStatus) String() string {
	switch this {
	case CANCELLED:
		return "已取消"
	case NONPAYMENT:
		return "未付款"
	case PAID:
		return "已付款"
	case COMPLETED:
		return "已完成"
	case APPEAL:
		return "申诉中"
	}
	return ""
}
