package DepositStatusEnum

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	PAY      = DepositStatusEnum{"缴纳", 1}
	GET_BACK = DepositStatusEnum{"索回", 2}
)

func (this *DepositStatusEnum) SetCnName(CnName string) (result *DepositStatusEnum) {
	this.CnName = CnName
	return this
}
func (this *DepositStatusEnum) GetCnName() (CnName string) {
	return this.CnName
}
func (this *DepositStatusEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *DepositStatusEnum) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *DepositStatusEnum) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "缴纳"
		case 2:
			this.CnName = "索回"
		}
	default:
		this = nil
	}
	return nil
}
func (this *DepositStatusEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *DepositStatusEnum) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "缴纳"
	case 2:
		this.CnName = "索回"
	}
	return
}
func Values() (result []DepositStatusEnum) {
	return []DepositStatusEnum{PAY, GET_BACK}
}
func (this *DepositStatusEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type DepositStatusEnum struct {
	CnName  string
	ordinal int
}
