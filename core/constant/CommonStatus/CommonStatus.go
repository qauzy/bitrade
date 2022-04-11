package CommonStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	NORMAL  = CommonStatus{"正常", 1}
	ILLEGAL = CommonStatus{"非法", 2}
)

func (this *CommonStatus) SetCnName(CnName string) (result *CommonStatus) {
	this.CnName = CnName
	return this
}
func (this *CommonStatus) GetCnName() (CnName string) {
	return this.CnName
}
func (this *CommonStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *CommonStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *CommonStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "正常"
		case 2:
			this.CnName = "非法"
		}
	default:
		this = nil
	}
	return nil
}
func (this *CommonStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *CommonStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "正常"
	case 2:
		this.CnName = "非法"
	}
	return
}
func Values() (result []CommonStatus) {
	return []CommonStatus{NORMAL, ILLEGAL}
}
func (this *CommonStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type CommonStatus struct {
	CnName  string
	ordinal int
}
