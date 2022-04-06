package MemberLevelEnum

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	GENERAL        = MemberLevelEnum{"普通", 1}
	REALNAME       = MemberLevelEnum{"实名", 2}
	IDENTIFICATION = MemberLevelEnum{"认证商家", 3}
)

func (this *MemberLevelEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *MemberLevelEnum) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *MemberLevelEnum) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "普通"
		case 2:
			this.CnName = "实名"
		case 3:
			this.CnName = "认证商家"
		}
	default:
		this = nil
	}
	return nil
}
func (this *MemberLevelEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *MemberLevelEnum) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "实名"
	case 3:
		this.CnName = "认证商家"
	}
	return
}
func (this *MemberLevelEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type MemberLevelEnum struct {
	CnName  string
	ordinal int
}
