package PromotionRewardType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	REGISTER             = PromotionRewardType{"推广注册", 1}
	TRANSACTION          = PromotionRewardType{"法币推广交易", 2}
	EXCHANGE_TRANSACTION = PromotionRewardType{"币币推广交易", 3}
)

func (this *PromotionRewardType) Ordinal() (result int) {
	return this.ordinal
}
func (this *PromotionRewardType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *PromotionRewardType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "推广注册"
		case 2:
			this.CnName = "法币推广交易"
		case 3:
			this.CnName = "币币推广交易"
		}
	default:
		this = nil
	}
	return nil
}
func (this *PromotionRewardType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *PromotionRewardType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "推广注册"
	case 2:
		this.CnName = "法币推广交易"
	case 3:
		this.CnName = "币币推广交易"
	}
	return
}
func (this *PromotionRewardType) GetOrdinal() (result int) {
	return this.Ordinal()
}

type PromotionRewardType struct {
	CnName  string
	ordinal int
}
