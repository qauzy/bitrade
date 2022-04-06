package Platform

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	ANDROID = Platform{"安卓", 1}
	IOS     = Platform{"苹果", 2}
)

func (this *Platform) Ordinal() (result int) {
	return this.ordinal
}
func (this *Platform) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *Platform) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "安卓"
		case 2:
			this.CnName = "苹果"
		}
	default:
		this = nil
	}
	return nil
}
func (this *Platform) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *Platform) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "安卓"
	case 2:
		this.CnName = "苹果"
	}
	return
}
func (this *Platform) GetOrdinal() (result int) {
	return this.Ordinal()
}

type Platform struct {
	CnName  string
	ordinal int
}
