package constant

type OrderStatus int

const (
	ORDERSTATUS_CANCELLED OrderStatus = iota
	ORDERSTATUS_NONPAYMENT
	ORDERSTATUS_PAID
	ORDERSTATUS_COMPLETED
	ORDERSTATUS_APPEAL
)

func (this OrderStatus) String() string {
	switch this {
	case ORDERSTATUS_CANCELLED:
		return "已取消"
	case ORDERSTATUS_NONPAYMENT:
		return "未付款"
	case ORDERSTATUS_PAID:
		return "已付款"
	case ORDERSTATUS_COMPLETED:
		return "已完成"
	case ORDERSTATUS_APPEAL:
		return "申诉中"
	}
	return ""
}
