package SysAdvertiseLocation

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	APP_SHUFFLING     = SysAdvertiseLocation{"app首页轮播", 1}
	PC_SHUFFLING      = SysAdvertiseLocation{"pc首页轮播", 2}
	PC_CLASSIFICATION = SysAdvertiseLocation{"pc分类广告", 3}
)

func (this *SysAdvertiseLocation) Ordinal() (result int) {
	return this.ordinal
}
func (this *SysAdvertiseLocation) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *SysAdvertiseLocation) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "app首页轮播"
		case 2:
			this.CnName = "pc首页轮播"
		case 3:
			this.CnName = "pc分类广告"
		}
	default:
		this = nil
	}
	return nil
}
func (this *SysAdvertiseLocation) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *SysAdvertiseLocation) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "app首页轮播"
	case 2:
		this.CnName = "pc首页轮播"
	case 3:
		this.CnName = "pc分类广告"
	}
	return
}
func (this *SysAdvertiseLocation) GetOrdinal() (result int) {
	return this.Ordinal()
}

type SysAdvertiseLocation struct {
	CnName  string
	ordinal int
}
