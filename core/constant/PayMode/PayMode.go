package PayMode

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	ALIPAY = PayMode{"支付宝", 1}
	WECHAT = PayMode{"微信", 2}
	BANK   = PayMode{"银联", 3}
	PAYPAL = PayMode{"Paypal", 4}
)

func (this *PayMode) SetCnName(CnName string) (result *PayMode) {
	this.CnName = CnName
	return this
}
func (this *PayMode) GetCnName() (CnName string) {
	return this.CnName
}
func (this *PayMode) Ordinal() (result int) {
	return this.ordinal
}
func (this *PayMode) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *PayMode) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "支付宝"
		case 2:
			this.CnName = "微信"
		case 3:
			this.CnName = "银联"
		case 4:
			this.CnName = "Paypal"
		}
	default:
		this = nil
	}
	return nil
}
func (this *PayMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *PayMode) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "支付宝"
	case 2:
		this.CnName = "微信"
	case 3:
		this.CnName = "银联"
	case 4:
		this.CnName = "Paypal"
	}
	return
}
func Values() (result []PayMode) {
	return []PayMode{ALIPAY, WECHAT, BANK, PAYPAL}
}
func (this *PayMode) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PayMode struct {
	CnName  string
	ordinal int
}
