package PromotionLevel

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	ONE   = PromotionLevel{"一级", 1}
	TWO   = PromotionLevel{"二级", 2}
	THREE = PromotionLevel{"三级", 3}
)

func (this *PromotionLevel) Ordinal() (result int) {
	return this.ordinal
}
func (this *PromotionLevel) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *PromotionLevel) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "一级"
		case 2:
			this.CnName = "二级"
		case 3:
			this.CnName = "三级"
		}
	default:
		this = nil
	}
	return nil
}
func (this *PromotionLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *PromotionLevel) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "一级"
	case 2:
		this.CnName = "二级"
	case 3:
		this.CnName = "三级"
	}
	return
}
func (this *PromotionLevel) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PromotionLevel struct {
	CnName  string
	ordinal int
}
