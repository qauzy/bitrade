package RewardRecordType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	PROMOTION = RewardRecordType{"推广", 1}
	ACTIVITY  = RewardRecordType{"活动", 2}
	DIVIDEND  = RewardRecordType{"分红", 3}
)

func (this *RewardRecordType) SetCnName(CnName string) (result *RewardRecordType) {
	this.CnName = CnName
	return this
}
func (this *RewardRecordType) GetCnName() (CnName string) {
	return this.CnName
}
func (this *RewardRecordType) Ordinal() (result int) {
	return this.ordinal
}
func (this *RewardRecordType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *RewardRecordType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "推广"
		case 2:
			this.CnName = "活动"
		case 3:
			this.CnName = "分红"
		}
	default:
		this = nil
	}
	return nil
}
func (this *RewardRecordType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *RewardRecordType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "活动"
	case 3:
		this.CnName = "分红"
	}
	return
}
func Values() (result []RewardRecordType) {
	return []RewardRecordType{PROMOTION, ACTIVITY, DIVIDEND}
}
func (this *RewardRecordType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type RewardRecordType struct {
	CnName  string
	ordinal int
}
