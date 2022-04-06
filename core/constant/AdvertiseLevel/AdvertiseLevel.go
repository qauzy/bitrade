package AdvertiseLevel

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	ORDINARY  = AdvertiseLevel{"普通", 1}
	EXCELLENT = AdvertiseLevel{"优质", 2}
)

func (this *AdvertiseLevel) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdvertiseLevel) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *AdvertiseLevel) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "普通"
		case 2:
			this.CnName = "优质"
		}
	default:
		this = nil
	}
	return nil
}
func (this *AdvertiseLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *AdvertiseLevel) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "普通"
	case 2:
		this.CnName = "优质"
	}
	return
}
func (this *AdvertiseLevel) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AdvertiseLevel struct {
	CnName  string
	ordinal int
}
