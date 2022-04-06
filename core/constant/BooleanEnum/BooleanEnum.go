package BooleanEnum

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	IS_FALSE = BooleanEnum{false, "否", 1}
	IS_TRUE  = BooleanEnum{true, "是", 2}
)

func (this *BooleanEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *BooleanEnum) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *BooleanEnum) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "否"
		case 2:
			this.CnName = "是"
		}
	default:
		this = nil
	}
	return nil
}
func (this *BooleanEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *BooleanEnum) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "否"
	case 2:
		this.CnName = "是"
	}
	return
}
func (this *BooleanEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type BooleanEnum struct {
	Is      bool
	CnName  string
	ordinal int
}
