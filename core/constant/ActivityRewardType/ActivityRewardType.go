package ActivityRewardType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	REGISTER    = ActivityRewardType{"注册奖励", 1}
	TRANSACTION = ActivityRewardType{"交易奖励", 2}
	RECHARGE    = ActivityRewardType{"充值奖励", 3}
)

func (this *ActivityRewardType) SetCnName(CnName string) (result *ActivityRewardType) {
	this.CnName = CnName
	return this
}
func (this *ActivityRewardType) GetCnName() (CnName string) {
	return this.CnName
}
func (this *ActivityRewardType) Ordinal() (result int) {
	return this.ordinal
}
func (this *ActivityRewardType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *ActivityRewardType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "注册奖励"
		case 2:
			this.CnName = "交易奖励"
		case 3:
			this.CnName = "充值奖励"
		}
	default:
		this = nil
	}
	return nil
}
func (this *ActivityRewardType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *ActivityRewardType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "注册奖励"
	case 2:
		this.CnName = "交易奖励"
	case 3:
		this.CnName = "充值奖励"
	}
	return
}
func Values() (result []ActivityRewardType) {
	return []ActivityRewardType{REGISTER, TRANSACTION, RECHARGE}
}
func (this *ActivityRewardType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type ActivityRewardType struct {
	CnName  string
	ordinal int
}
