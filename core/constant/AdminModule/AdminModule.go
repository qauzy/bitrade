package AdminModule

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

var (
	CMS       = AdminModule{"CMS", 1}
	COMMON    = AdminModule{"COMMON", 2}
	EXCHANGE  = AdminModule{"EXCHANGE", 3}
	FINANCE   = AdminModule{"FINANCE", 4}
	MEMBER    = AdminModule{"MEMBER", 5}
	OTC       = AdminModule{"OTC", 6}
	SYSTEM    = AdminModule{"SYSTEM", 7}
	PROMOTION = AdminModule{"PROMOTION", 8}
	INDEX     = AdminModule{"INDEX", 9}
	IEO       = AdminModule{"IEO", 10}
	GIFT      = AdminModule{"GIFT", 11}
	MARGIN    = AdminModule{"MARGIN", 12}
)

func (this *AdminModule) SetTitle(Title string) (result *AdminModule) {
	this.Title = Title
	return this
}
func (this *AdminModule) GetTitle() (Title string) {
	return this.Title
}
func (this *AdminModule) Ordinal() (result int) {
	return this.ordinal
}
func (this *AdminModule) Value() (driver.Value, error) {
	return this.Ordinal(), nil
}
func (this *AdminModule) Scan(v interface{}) error {
	switch vt := v.(type) {
	case int:
		this.ordinal = vt
		switch vt {
		case 1:
			this.Title = "CMS"
		case 2:
			this.Title = "COMMON"
		case 3:
			this.Title = "EXCHANGE"
		case 4:
			this.Title = "FINANCE"
		case 5:
			this.Title = "MEMBER"
		case 6:
			this.Title = "OTC"
		case 7:
			this.Title = "SYSTEM"
		case 8:
			this.Title = "PROMOTION"
		case 9:
			this.Title = "INDEX"
		case 10:
			this.Title = "IEO"
		case 11:
			this.Title = "GIFT"
		case 12:
			this.Title = "MARGIN"
		}
	default:
		this = nil
	}
	return nil
}
func (this *AdminModule) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", this.ordinal)), nil
}
func (this *AdminModule) UnmarshalJSON(data []byte) (err error) {
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
		this.Title = "CMS"
	case 2:
		this.Title = "COMMON"
	case 3:
		this.Title = "EXCHANGE"
	case 4:
		this.Title = "FINANCE"
	case 5:
		this.Title = "MEMBER"
	case 6:
		this.Title = "OTC"
	case 7:
		this.Title = "SYSTEM"
	case 8:
		this.Title = "PROMOTION"
	case 9:
		this.Title = "INDEX"
	case 10:
		this.Title = "IEO"
	case 11:
		this.Title = "GIFT"
	case 12:
		this.Title = "MARGIN"
	}
	return
}
func Values() (result []AdminModule) {
	return []AdminModule{CMS, COMMON, EXCHANGE, FINANCE, MEMBER, OTC, SYSTEM, PROMOTION, INDEX, IEO, GIFT, MARGIN}
}

type AdminModule struct {
	Title   string
	ordinal int
}
