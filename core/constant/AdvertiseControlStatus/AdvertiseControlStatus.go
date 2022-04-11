package AdvertiseControlStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	PUT_ON_SHELVES  = AdvertiseControlStatus{"上架", 1}
	PUT_OFF_SHELVES = AdvertiseControlStatus{"下架", 2}
	TURNOFF         = AdvertiseControlStatus{"已关闭", 3}
)

func (this *AdvertiseControlStatus) SetCnName(CnName string) (result *AdvertiseControlStatus) {
	this.CnName = CnName
	return this
}
func (this *AdvertiseControlStatus) GetCnName() (CnName string) {
	return this.CnName
}
func (this *AdvertiseControlStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseControlStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *AdvertiseControlStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "上架"
		case 2:
			this.CnName = "下架"
		case 3:
			this.CnName = "已关闭"
		}
	default:
		this = nil
	}
	return nil
}
func (this *AdvertiseControlStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *AdvertiseControlStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "上架"
	case 2:
		this.CnName = "下架"
	case 3:
		this.CnName = "已关闭"
	}
	return
}
func Values() (result []AdvertiseControlStatus) {
	return []AdvertiseControlStatus{PUT_ON_SHELVES, PUT_OFF_SHELVES, TURNOFF}
}
func (this *AdvertiseControlStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseControlStatus struct {
	CnName  string
	ordinal int
}
