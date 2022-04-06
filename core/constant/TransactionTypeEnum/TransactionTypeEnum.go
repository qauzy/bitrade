package TransactionTypeEnum

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	OTC_NUM       = TransactionTypeEnum{"法币成交量", 1}
	OTC_MONEY     = TransactionTypeEnum{"法币成交额", 2}
	EXCHANGE      = TransactionTypeEnum{"币币交易手续费统计", 3}
	EXCHANGE_COIN = TransactionTypeEnum{"币币交易成交量统计", 4}
	EXCHANGE_BASE = TransactionTypeEnum{"币币交易成交额统计", 5}
	RECHARGE      = TransactionTypeEnum{"充币", 6}
	WITHDRAW      = TransactionTypeEnum{"提币", 7}
)

func (this *TransactionTypeEnum) Ordinal() (result int) {
	return this.ordinal
}
func (this *TransactionTypeEnum) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *TransactionTypeEnum) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "法币成交量"
		case 2:
			this.CnName = "法币成交额"
		case 3:
			this.CnName = "币币交易手续费统计"
		case 4:
			this.CnName = "币币交易成交量统计"
		case 5:
			this.CnName = "币币交易成交额统计"
		case 6:
			this.CnName = "充币"
		case 7:
			this.CnName = "提币"
		}
	default:
		this = nil
	}
	return nil
}
func (this *TransactionTypeEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *TransactionTypeEnum) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "法币成交量"
	case 2:
		this.CnName = "法币成交额"
	case 3:
		this.CnName = "币币交易手续费统计"
	case 4:
		this.CnName = "币币交易成交量统计"
	case 5:
		this.CnName = "币币交易成交额统计"
	case 6:
		this.CnName = "充币"
	case 7:
		this.CnName = "提币"
	}
	return
}
func (this *TransactionTypeEnum) GetOrdinal() (result int) {
	return this.Ordinal()
}

type TransactionTypeEnum struct {
	CnName  string
	ordinal int
}
