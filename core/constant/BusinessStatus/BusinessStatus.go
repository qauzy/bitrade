package BusinessStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	ZERO    = BusinessStatus{"为成交", 1}
	PORTION = BusinessStatus{"部分成交", 2}
	ALL     = BusinessStatus{"全部成交", 3}
)

func (this *BusinessStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *BusinessStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *BusinessStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "为成交"
		case 2:
			this.CnName = "部分成交"
		case 3:
			this.CnName = "全部成交"
		}
	default:
		this = nil
	}
	return nil
}
func (this *BusinessStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *BusinessStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "为成交"
	case 2:
		this.CnName = "部分成交"
	case 3:
		this.CnName = "全部成交"
	}
	return
}
func (this *BusinessStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type BusinessStatus struct {
	CnName  string
	ordinal int
}
