package AppealStatus

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	NOT_PROCESSED = AppealStatus{"未处理", 1}
	PROCESSED     = AppealStatus{"已处理", 2}
)

func (this *AppealStatus) Ordinal() (result int) {
	return this.ordinal
}
func (this *AppealStatus) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *AppealStatus) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "未处理"
		case 2:
			this.CnName = "已处理"
		}
	default:
		this = nil
	}
	return nil
}
func (this *AppealStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *AppealStatus) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "未处理"
	case 2:
		this.CnName = "已处理"
	}
	return
}
func (this *AppealStatus) GetOrdinal() (result int) {
	return this.Ordinal()
}

type AppealStatus struct {
	CnName  string
	ordinal int
}
