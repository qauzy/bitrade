package AuditStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	AUDIT_ING      = AuditStatus{"待审核", 1}
	AUDIT_DEFEATED = AuditStatus{"审核失败", 2}
	AUDIT_SUCCESS  = AuditStatus{"审核成功", 3}
)

func (this *AuditStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AuditStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *AuditStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "待审核"
		case 2:
			this.CnName = "审核失败"
		case 3:
			this.CnName = "审核成功"
		}
	default:
		this = nil
	}
	return nil
}
func (this *AuditStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *AuditStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "待审核"
	case 2:
		this.CnName = "审核失败"
	case 3:
		this.CnName = "审核成功"
	}
	return
}
func (this *AuditStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AuditStatus struct {
	CnName  string
	ordinal int
}
