package Symbol

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	USDT = Symbol{1}
	BTC  = Symbol{2}
	ETH  = Symbol{3}
	GCC  = Symbol{4}
	GCX  = Symbol{5}
)

func (this *Symbol) Ordinal() (result int) {
	return this.ordinal
}
func (this *Symbol) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *Symbol) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		}
	default:
		this = nil
	}
	return nil
}
func (this *Symbol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *Symbol) UnmarshalJSON(data []byte) (err error) {
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
	}
	return
}
func Values() (result []Symbol) {
	return []Symbol{USDT, BTC, ETH, GCC, GCX}
}

type Symbol struct {
	ordinal int
}
