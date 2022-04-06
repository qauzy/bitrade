package MemberRegisterType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	BACKSTAGE              = MemberRegisterType{"后台", 1}
	AUTONOMOUSLY           = MemberRegisterType{"自主", 2}
	AUTONOMOUSLY_RECOMMEND = MemberRegisterType{"自主(推荐)", 3}
)

func (this *MemberRegisterType) Ordinal() (result int) {
	return this.ordinal
}
func (this *MemberRegisterType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *MemberRegisterType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "后台"
		case 2:
			this.CnName = "自主"
		case 3:
			this.CnName = "自主(推荐)"
		}
	default:
		this = nil
	}
	return nil
}
func (this *MemberRegisterType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *MemberRegisterType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "后台"
	case 2:
		this.CnName = "自主"
	case 3:
		this.CnName = "自主(推荐)"
	}
	return
}
func (this *MemberRegisterType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type MemberRegisterType struct {
	CnName  string
	ordinal int
}
