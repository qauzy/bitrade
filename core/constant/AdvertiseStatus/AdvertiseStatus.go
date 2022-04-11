package AdvertiseStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	HANG    = AdvertiseStatus{"已挂单", 1}
	PAYMENT = AdvertiseStatus{"待付款", 2}
	TURNOFF = AdvertiseStatus{"已关闭", 3}
)

func (this *AdvertiseStatus) SetCnName(CnName string) (result *AdvertiseStatus) {
	this.CnName = CnName
	return this
}
func (this *AdvertiseStatus) GetCnName() (CnName string) {
	return this.CnName
}
func (this *AdvertiseStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *AdvertiseStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "已挂单"
		case 2:
			this.CnName = "待付款"
		case 3:
			this.CnName = "已关闭"
		}
	default:
		this = nil
	}
	return nil
}
func (this *AdvertiseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *AdvertiseStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "已挂单"
	case 2:
		this.CnName = "待付款"
	case 3:
		this.CnName = "已关闭"
	}
	return
}
func Values() (result []AdvertiseStatus) {
	return []AdvertiseStatus{HANG, PAYMENT, TURNOFF}
}
func (this *AdvertiseStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseStatus struct {
	CnName  string
	ordinal int
}
