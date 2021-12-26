package entity

import (
	"bitrade/core/constant/CertifiedBusinessStatus"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/math"
	"time"
)

func (this *BusinessAuthApplyDetailVO) SetId(id int64) (result *BusinessAuthApplyDetailVO) {
	this.Id = id
	return this
}
func (this *BusinessAuthApplyDetailVO) GetId() (id int64) {
	return this.Id
}
func (this *BusinessAuthApplyDetailVO) SetInfo(info fastjson.JSONObject) (result *BusinessAuthApplyDetailVO) {
	this.Info = info
	return this
}
func (this *BusinessAuthApplyDetailVO) GetInfo() (info fastjson.JSONObject) {
	return this.Info
}
func (this *BusinessAuthApplyDetailVO) SetStatus(status CertifiedBusinessStatus.CertifiedBusinessStatus) (result *BusinessAuthApplyDetailVO) {
	this.Status = status
	return this
}
func (this *BusinessAuthApplyDetailVO) GetStatus() (status CertifiedBusinessStatus.CertifiedBusinessStatus) {
	return this.Status
}
func (this *BusinessAuthApplyDetailVO) SetCheckTime(checkTime time.Time) (result *BusinessAuthApplyDetailVO) {
	this.CheckTime = checkTime
	return this
}
func (this *BusinessAuthApplyDetailVO) GetCheckTime() (checkTime time.Time) {
	return this.CheckTime
}
func (this *BusinessAuthApplyDetailVO) SetRealName(realName string) (result *BusinessAuthApplyDetailVO) {
	this.RealName = realName
	return this
}
func (this *BusinessAuthApplyDetailVO) GetRealName() (realName string) {
	return this.RealName
}
func (this *BusinessAuthApplyDetailVO) SetDetail(detail string) (result *BusinessAuthApplyDetailVO) {
	this.Detail = detail
	return this
}
func (this *BusinessAuthApplyDetailVO) GetDetail() (detail string) {
	return this.Detail
}
func (this *BusinessAuthApplyDetailVO) SetAmount(amount math.BigDecimal) (result *BusinessAuthApplyDetailVO) {
	this.Amount = amount
	return this
}
func (this *BusinessAuthApplyDetailVO) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *BusinessAuthApplyDetailVO) SetAuthInfo(authInfo string) (result *BusinessAuthApplyDetailVO) {
	this.AuthInfo = authInfo
	return this
}
func (this *BusinessAuthApplyDetailVO) GetAuthInfo() (authInfo string) {
	return this.AuthInfo
}

type BusinessAuthApplyDetailVO struct {
	Id        int64
	Info      fastjson.JSONObject
	Status    CertifiedBusinessStatus.CertifiedBusinessStatus
	CheckTime time.Time
	RealName  string
	Detail    string
	Amount    math.BigDecimal
	AuthInfo  string
}
