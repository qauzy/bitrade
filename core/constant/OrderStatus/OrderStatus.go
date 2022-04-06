package OrderStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	CANCELLED  = OrderStatus{"已取消", 1}
	NONPAYMENT = OrderStatus{"未付款", 2}
	PAID       = OrderStatus{"已付款", 3}
	COMPLETED  = OrderStatus{"已完成", 4}
	APPEAL     = OrderStatus{"申诉中", 5}
)

func (this *OrderStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *OrderStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *OrderStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "已取消"
		case 2:
			this.CnName = "未付款"
		case 3:
			this.CnName = "已付款"
		case 4:
			this.CnName = "已完成"
		case 5:
			this.CnName = "申诉中"
		}
	default:
		this = nil
	}
	return nil
}
func (this *OrderStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *OrderStatus) UnmarshalJSON(data []byte) (err error) {
	if data == nil || len(data) == 2 {
		return
	}
	this.ordinal, err = strconv.Atoi(string(data))
	if err != nil {
		return
	}
	switch this.ordinal {
	}
	switch this.ordinal {
	case 1:
		this.CnName = "已取消"
	case 2:
		this.CnName = "未付款"
	case 3:
		this.CnName = "已付款"
	case 4:
		this.CnName = "已完成"
	case 5:
		this.CnName = "申诉中"
	}
	return
}
func (this *OrderStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type OrderStatus struct {
	CnName  string
	ordinal int
}
