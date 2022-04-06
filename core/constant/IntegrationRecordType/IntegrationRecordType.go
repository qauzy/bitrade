package IntegrationRecordType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	PROMOTION_GIVING      = IntegrationRecordType{"推广", 1}
	LEGAL_RECHARGE_GIVING = IntegrationRecordType{"法币充值赠送", 2}
	COIN_RECHARGE_GIVING  = IntegrationRecordType{"币币充值赠送", 3}
)

func (this *IntegrationRecordType) Ordinal() (result int) {
	return this.ordinal
}
func (this *IntegrationRecordType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *IntegrationRecordType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "推广"
		case 2:
			this.CnName = "法币充值赠送"
		case 3:
			this.CnName = "币币充值赠送"
		}
	default:
		this = nil
	}
	return nil
}
func (this *IntegrationRecordType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *IntegrationRecordType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "推广"
	case 2:
		this.CnName = "法币充值赠送"
	case 3:
		this.CnName = "币币充值赠送"
	}
	return
}
func (this *IntegrationRecordType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type IntegrationRecordType struct {
	CnName  string
	ordinal int
}
