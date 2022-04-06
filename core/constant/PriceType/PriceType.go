package PriceType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	REGULAR  = PriceType{"固定的", 1}
	MUTATIVE = PriceType{"变化的", 2}
)

func (this *PriceType) Ordinal() (result int) {
	return this.ordinal
}
func (this *PriceType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *PriceType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "固定的"
		case 2:
			this.CnName = "变化的"
		}
	default:
		this = nil
	}
	return nil
}
func (this *PriceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *PriceType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "固定的"
	case 2:
		this.CnName = "变化的"
	}
	return
}
func (this *PriceType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PriceType struct {
	CnName  string
	ordinal int
}
