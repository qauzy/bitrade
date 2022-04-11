package AdvertiseType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	BUY  = AdvertiseType{"购买", 1}
	SELL = AdvertiseType{"出售", 2}
)

func (this *AdvertiseType) SetCnName(CnName string) (result *AdvertiseType) {
	this.CnName = CnName
	return this
}
func (this *AdvertiseType) GetCnName() (CnName string) {
	return this.CnName
}
func (this *AdvertiseType) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *AdvertiseType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "购买"
		case 2:
			this.CnName = "出售"
		}
	default:
		this = nil
	}
	return nil
}
func (this *AdvertiseType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *AdvertiseType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "购买"
	case 2:
		this.CnName = "出售"
	}
	return
}
func Values() (result []AdvertiseType) {
	return []AdvertiseType{BUY, SELL}
}
func (this *AdvertiseType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseType struct {
	CnName  string
	ordinal int
}
