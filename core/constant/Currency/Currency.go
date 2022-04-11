package Currency

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var BTC = Currency{"比特币", 1}

func (this *Currency) SetCnName(CnName string) (result *Currency) {
	this.CnName = CnName
	return this
}
func (this *Currency) GetCnName() (CnName string) {
	return this.CnName
}
func (this *Currency) Ordinal() (result int) {
	return this.ordinal
}
func (this *Currency) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *Currency) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "比特币"
		}
	default:
		this = nil
	}
	return nil
}
func (this *Currency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *Currency) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "比特币"
	}
	return
}
func Values() (result []Currency) {
	return []Currency{BTC}
}
func (this *Currency) GetOrdinal() (result int) {
	return this.Ordinal()
}

type Currency struct {
	CnName  string
	ordinal int
}
