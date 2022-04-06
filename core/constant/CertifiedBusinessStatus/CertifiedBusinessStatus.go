package CertifiedBusinessStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	NOT_CERTIFIED  = CertifiedBusinessStatus{"未认证", 1}
	AUDITING       = CertifiedBusinessStatus{"认证-待审核", 2}
	VERIFIED       = CertifiedBusinessStatus{"认证-审核成功", 3}
	FAILED         = CertifiedBusinessStatus{"认证-审核失败", 4}
	DEPOSIT_LESS   = CertifiedBusinessStatus{"保证金不足", 5}
	CANCEL_AUTH    = CertifiedBusinessStatus{"退保-待审核", 6}
	RETURN_FAILED  = CertifiedBusinessStatus{"退保-审核失败", 7}
	RETURN_SUCCESS = CertifiedBusinessStatus{"退保-审核成功", 8}
)

func (this *CertifiedBusinessStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *CertifiedBusinessStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *CertifiedBusinessStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "未认证"
		case 2:
			this.CnName = "认证-待审核"
		case 3:
			this.CnName = "认证-审核成功"
		case 4:
			this.CnName = "认证-审核失败"
		case 5:
			this.CnName = "保证金不足"
		case 6:
			this.CnName = "退保-待审核"
		case 7:
			this.CnName = "退保-审核失败"
		case 8:
			this.CnName = "退保-审核成功"
		}
	default:
		this = nil
	}
	return nil
}
func (this *CertifiedBusinessStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *CertifiedBusinessStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "认证-待审核"
	case 3:
		this.CnName = "认证-审核成功"
	case 4:
		this.CnName = "认证-审核失败"
	case 5:
		this.CnName = "保证金不足"
	case 6:
		this.CnName = "退保-待审核"
	case 7:
		this.CnName = "退保-审核失败"
	case 8:
		this.CnName = "退保-审核成功"
	}
	return
}
func (this *CertifiedBusinessStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type CertifiedBusinessStatus struct {
	CnName  string
	ordinal int
}
