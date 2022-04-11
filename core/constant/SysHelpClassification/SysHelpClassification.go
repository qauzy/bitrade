package SysHelpClassification

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	HELP        = SysHelpClassification{"新手入门", 1}
	FAQ         = SysHelpClassification{"常见问题", 2}
	RECHARGE    = SysHelpClassification{"充值指南", 3}
	TRANSACTION = SysHelpClassification{"交易指南", 4}
	QR_CODE     = SysHelpClassification{"APP二维码", 5}
)

func (this *SysHelpClassification) SetCnName(CnName string) (result *SysHelpClassification) {
	this.CnName = CnName
	return this
}
func (this *SysHelpClassification) GetCnName() (CnName string) {
	return this.CnName
}
func (this *SysHelpClassification) Ordinal() (result int) {
	return this.ordinal
}
func (this *SysHelpClassification) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *SysHelpClassification) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "新手入门"
		case 2:
			this.CnName = "常见问题"
		case 3:
			this.CnName = "充值指南"
		case 4:
			this.CnName = "交易指南"
		case 5:
			this.CnName = "APP二维码"
		}
	default:
		this = nil
	}
	return nil
}
func (this *SysHelpClassification) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *SysHelpClassification) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "新手入门"
	case 2:
		this.CnName = "常见问题"
	case 3:
		this.CnName = "充值指南"
	case 4:
		this.CnName = "交易指南"
	case 5:
		this.CnName = "APP二维码"
	}
	return
}
func Values() (result []SysHelpClassification) {
	return []SysHelpClassification{HELP, FAQ, RECHARGE, TRANSACTION, QR_CODE}
}
func (this *SysHelpClassification) GetOrdinal() (result int) {
	return this.Ordinal()
}

type SysHelpClassification struct {
	CnName  string
	ordinal int
}
