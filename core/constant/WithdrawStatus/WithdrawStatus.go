package WithdrawStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	PROCESSING = WithdrawStatus{"审核中", 1}
	WAITING    = WithdrawStatus{"等待放币", 2}
	FAIL       = WithdrawStatus{"失败", 3}
	SUCCESS    = WithdrawStatus{"成功", 4}
	TRANSFER   = WithdrawStatus{"转账中", 5}
)

func (this *WithdrawStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *WithdrawStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *WithdrawStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "审核中"
		case 2:
			this.CnName = "等待放币"
		case 3:
			this.CnName = "失败"
		case 4:
			this.CnName = "成功"
		case 5:
			this.CnName = "转账中"
		}
	default:
		this = nil
	}
	return nil
}
func (this *WithdrawStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *WithdrawStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "审核中"
	case 2:
		this.CnName = "等待放币"
	case 3:
		this.CnName = "失败"
	case 4:
		this.CnName = "成功"
	case 5:
		this.CnName = "转账中"
	}
	return
}
func (this *WithdrawStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type WithdrawStatus struct {
	CnName  string
	ordinal int
}
