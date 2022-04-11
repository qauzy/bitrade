package entity

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/math"
)

func (this *BusinessAuthApplyDetailVO) SetId(Id int64) (result *BusinessAuthApplyDetailVO) {
	this.Id = Id
	return this
}
func (this *BusinessAuthApplyDetailVO) GetId() (Id int64) {
	return this.Id
}
func (this *BusinessAuthApplyDetailVO) SetInfo(Info *fastjson.JSONObject) (result *BusinessAuthApplyDetailVO) {
	this.Info = Info
	return this
}
func (this *BusinessAuthApplyDetailVO) GetInfo() (Info *fastjson.JSONObject) {
	return this.Info
}
func (this *BusinessAuthApplyDetailVO) SetStatus(Status CertifiedBusinessStatus.CertifiedBusinessStatus) (result *BusinessAuthApplyDetailVO) {
	this.Status = Status
	return this
}
func (this *BusinessAuthApplyDetailVO) GetStatus() (Status CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.Status
}
func (this *BusinessAuthApplyDetailVO) SetCheckTime(CheckTime xtime.Xtime) (result *BusinessAuthApplyDetailVO) {
	this.CheckTime = CheckTime
	return this
}
func (this *BusinessAuthApplyDetailVO) GetCheckTime() (CheckTime xtime.Xtime) {
	return this.CheckTime
}
func (this *BusinessAuthApplyDetailVO) SetRealName(RealName string) (result *BusinessAuthApplyDetailVO) {
	this.RealName = RealName
	return this
}
func (this *BusinessAuthApplyDetailVO) GetRealName() (RealName string) {
	return this.RealName
}
func (this *BusinessAuthApplyDetailVO) SetDetail(Detail string) (result *BusinessAuthApplyDetailVO) {
	this.Detail = Detail
	return this
}
func (this *BusinessAuthApplyDetailVO) GetDetail() (Detail string) {
	return this.Detail
}
func (this *BusinessAuthApplyDetailVO) SetAmount(Amount math.BigDecimal) (result *BusinessAuthApplyDetailVO) {
	this.Amount = Amount
	return this
}
func (this *BusinessAuthApplyDetailVO) GetAmount() (Amount math.BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthApplyDetailVO) SetAuthInfo(AuthInfo string) (result *BusinessAuthApplyDetailVO) {
	this.AuthInfo = AuthInfo
	return this
}
func (this *BusinessAuthApplyDetailVO) GetAuthInfo() (AuthInfo string) {
	return this.AuthInfo
}
func NewBusinessAuthApplyDetailVO(id int64, info *fastjson.JSONObject, status CertifiedBusinessStatus.CertifiedBusinessStatus, checkTime xtime.Xtime, realName string, detail string, amount math.BigDecimal, authInfo string) (ret *BusinessAuthApplyDetailVO) {
	ret = new(BusinessAuthApplyDetailVO)
	ret.Id = id
	ret.Info = info
	ret.Status = status
	ret.CheckTime = checkTime
	ret.RealName = realName
	ret.Detail = detail
	ret.Amount = amount
	ret.AuthInfo = authInfo
	return
}

type BusinessAuthApplyDetailVO struct {
	Id        int64                                           `gorm:"column:id" json:"id"`
	Info      *fastjson.JSONObject                            `gorm:"column:info" json:"info"`
	Status    CertifiedBusinessStatus.CertifiedBusinessStatus `gorm:"column:status" json:"status"`
	CheckTime xtime.Xtime                                     `gorm:"column:check_time" json:"checkTime"`
	RealName  string                                          `gorm:"column:real_name" json:"realName"`
	Detail    string                                          `gorm:"column:detail" json:"detail"`
	Amount    math.BigDecimal                                 `gorm:"column:amount" json:"amount"`
	AuthInfo  string                                          `gorm:"column:auth_info" json:"authInfo"`
}
