package LegalWalletState

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	APPLYING = LegalWalletState{"申请中", 1}
	COMPLETE = LegalWalletState{"完成", 2}
	DEFEATED = LegalWalletState{"失败", 3}
)

func (this *LegalWalletState) SetCnName(CnName string) (result *LegalWalletState) {
	this.CnName = CnName
	return this
}
func (this *LegalWalletState) GetCnName() (CnName string) {
	return this.CnName
}
func (this *LegalWalletState) Ordinal() (result int) {
	return this.ordinal
}
func (this *LegalWalletState) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *LegalWalletState) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "申请中"
		case 2:
			this.CnName = "完成"
		case 3:
			this.CnName = "失败"
		}
	default:
		this = nil
	}
	return nil
}
func (this *LegalWalletState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *LegalWalletState) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "申请中"
	case 2:
		this.CnName = "完成"
	case 3:
		this.CnName = "失败"
	}
	return
}
func Values() (result []LegalWalletState) {
	return []LegalWalletState{APPLYING, COMPLETE, DEFEATED}
}
func NewLegalWalletState(cnName string) (this *LegalWalletState) {
	this = new(LegalWalletState)
	this.CnName = this.CnName
	return
}

type LegalWalletState struct {
	CnName  string
	ordinal int
}
