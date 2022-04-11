package RealNameStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	NOT_CERTIFIED = RealNameStatus{"未认证", 1}
	AUDITING      = RealNameStatus{"审核中", 2}
	VERIFIED      = RealNameStatus{"已认证", 3}
)

func (this *RealNameStatus) SetCnName(CnName string) (result *RealNameStatus) {
	this.CnName = CnName
	return this
}
func (this *RealNameStatus) GetCnName() (CnName string) {
	return this.CnName
}
func (this *RealNameStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *RealNameStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *RealNameStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "未认证"
		case 2:
			this.CnName = "审核中"
		case 3:
			this.CnName = "已认证"
		}
	default:
		this = nil
	}
	return nil
}
func (this *RealNameStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *RealNameStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "未认证"
	case 2:
		this.CnName = "审核中"
	case 3:
		this.CnName = "已认证"
	}
	return
}
func Values() (result []RealNameStatus) {
	return []RealNameStatus{NOT_CERTIFIED, AUDITING, VERIFIED}
}
func (this *RealNameStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type RealNameStatus struct {
	CnName  string
	ordinal int
}
