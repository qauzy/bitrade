package CredentialsType

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	CARDED          = CredentialsType{"身份证", 1}
	PASSPORT        = CredentialsType{"护照", 2}
	DRIVING_LICENSE = CredentialsType{"驾照", 3}
)

func (this *CredentialsType) Ordinal() (result int) {
	return this.ordinal
}
func (this *CredentialsType) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *CredentialsType) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.CnName = "身份证"
		case 2:
			this.CnName = "护照"
		case 3:
			this.CnName = "驾照"
		}
	default:
		this = nil
	}
	return nil
}
func (this *CredentialsType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *CredentialsType) UnmarshalJSON(data []byte) (err error) {
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
		this.CnName = "身份证"
	case 2:
		this.CnName = "护照"
	case 3:
		this.CnName = "驾照"
	}
	return
}
func (this *CredentialsType) GetOrdinal() (result int) {
	return this.Ordinal()
}
func (this *CredentialsType) GetAll() (result []CredentialsType) {
	return []CredentialsType{CARDED, PASSPORT, DRIVING_LICENSE}
}
func GetByValue(value int) (result CredentialsType) {
	for _, credentialsType := range []CredentialsType{CARDED, PASSPORT, DRIVING_LICENSE} {
		if credentialsType.GetOrdinal() == value {
			return credentialsType
		}
	}
	return
}

type CredentialsType struct {
	CnName  string
	ordinal int
}
