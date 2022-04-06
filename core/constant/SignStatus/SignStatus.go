package SignStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	UNDERWAY = SignStatus{0}
	FINISH   = SignStatus{1}
)

func (this *SignStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *SignStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *SignStatus) Scan(v interface{}) error {
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
func (this *SignStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *SignStatus) UnmarshalJSON(data []byte) (err error) {
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

type SignStatus struct {
	ordinal int
}
