package vo

import (
	"bitrade/core/constant"
	"github.com/qauzy/math"
	"time"
)

func (this *OtcOrderVO) SetId(id int64) (result *OtcOrderVO) {
	this.Id = id
	return this
}
func (this *OtcOrderVO) GetId() (id int64) {
	return this.Id
}
func (this *OtcOrderVO) SetOrderSn(orderSn string) (result *OtcOrderVO) {
	this.OrderSn = orderSn
	return this
}
func (this *OtcOrderVO) GetOrderSn() (orderSn string) {
	return this.OrderSn
}
func (this *OtcOrderVO) SetCreateTime(createTime time.Time) (result *OtcOrderVO) {
	this.CreateTime = createTime
	return this
}
func (this *OtcOrderVO) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *OtcOrderVO) SetMemberName(memberName string) (result *OtcOrderVO) {
	this.MemberName = memberName
	return this
}
func (this *OtcOrderVO) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *OtcOrderVO) SetCustomerName(customerName string) (result *OtcOrderVO) {
	this.CustomerName = customerName
	return this
}
func (this *OtcOrderVO) GetCustomerName() (customerName string) {
	return this.CustomerName
}
func (this *OtcOrderVO) SetUnit(unit string) (result *OtcOrderVO) {
	this.Unit = unit
	return this
}
func (this *OtcOrderVO) GetUnit() (unit string) {
	return this.Unit
}
func (this *OtcOrderVO) SetAdvertiseType(advertiseType constant.AdvertiseType) (result *OtcOrderVO) {
	this.AdvertiseType = advertiseType
	return this
}
func (this *OtcOrderVO) GetAdvertiseType() (advertiseType constant.AdvertiseType) {
	return this.AdvertiseType
}
func (this *OtcOrderVO) SetMoney(money math.BigDecimal) (result *OtcOrderVO) {
	this.Money = money
	return this
}
func (this *OtcOrderVO) GetMoney() (money math.BigDecimal) {
	return this.Money
}
func (this *OtcOrderVO) SetNumber(number math.BigDecimal) (result *OtcOrderVO) {
	this.Number = number
	return this
}
func (this *OtcOrderVO) GetNumber() (number math.BigDecimal) {
	return this.Number
}
func (this *OtcOrderVO) SetFee(fee math.BigDecimal) (result *OtcOrderVO) {
	this.Fee = fee
	return this
}
func (this *OtcOrderVO) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *OtcOrderVO) SetPayMode(payMode string) (result *OtcOrderVO) {
	this.PayMode = payMode
	return this
}
func (this *OtcOrderVO) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *OtcOrderVO) SetStatus(status constant.OrderStatus) (result *OtcOrderVO) {
	this.Status = status
	return this
}
func (this *OtcOrderVO) GetStatus() (status constant.OrderStatus) {
	return this.Status
}
func (this *OtcOrderVO) SetCancelTime(cancelTime time.Time) (result *OtcOrderVO) {
	this.CancelTime = cancelTime
	return this
}
func (this *OtcOrderVO) GetCancelTime() (cancelTime time.Time) {
	return this.CancelTime
}
func (this *OtcOrderVO) SetReleaseTime(releaseTime time.Time) (result *OtcOrderVO) {
	this.ReleaseTime = releaseTime
	return this
}
func (this *OtcOrderVO) GetReleaseTime() (releaseTime time.Time) {
	return this.ReleaseTime
}
func (this *OtcOrderVO) SetPayTime(payTime time.Time) (result *OtcOrderVO) {
	this.PayTime = payTime
	return this
}
func (this *OtcOrderVO) GetPayTime() (payTime time.Time) {
	return this.PayTime
}

type OtcOrderVO struct {
	Id            int64
	OrderSn       string
	CreateTime    time.Time
	MemberName    string
	CustomerName  string
	Unit          string
	AdvertiseType constant.AdvertiseType
	Money         math.BigDecimal
	Number        math.BigDecimal
	Fee           math.BigDecimal
	PayMode       string
	Status        constant.OrderStatus
	CancelTime    time.Time
	ReleaseTime   time.Time
	PayTime       time.Time
}
