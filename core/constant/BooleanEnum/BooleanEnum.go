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

func (this *BooleanEnum) SetIs(Is bool) (result *BooleanEnum) {
	this.Is = Is
	return this
}
func (this *BooleanEnum) GetIs() (Is bool) {
	return this.Is
}
func (this *BooleanEnum) SetNameCn(NameCn string) (result *BooleanEnum) {
	this.NameCn = NameCn
	return this
}
func (this *BooleanEnum) GetNameCn() (NameCn string) {
	return this.NameCn
}
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
			this.NameCn = "否"
		case 2:
			this.NameCn = "是"
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
		this.NameCn = "否"
	case 2:
		this.NameCn = "是"
	}
	return
}
func Values() (result []BooleanEnum) {
	return []BooleanEnum{IS_FALSE, IS_TRUE}
}
func (this *BooleanEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type BooleanEnum struct {
	Is      bool
	NameCn  string
	ordinal int
}
