package entity

import (
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/constant/OrderStatus"
	"github.com/qauzy/math"
	"time"
)

func (this *Order) SetId(id int64) (result *Order) {
	this.Id = id
	return this
}
func (this *Order) GetId() (id int64) {
	return this.Id
}
func (this *Order) SetOrderSn(orderSn string) (result *Order) {
	this.OrderSn = orderSn
	return this
}
func (this *Order) GetOrderSn() (orderSn string) {
	return this.OrderSn
}
func (this *Order) SetReferenceNumber(referenceNumber string) (result *Order) {
	this.ReferenceNumber = referenceNumber
	return this
}
func (this *Order) GetReferenceNumber() (referenceNumber string) {
	return this.ReferenceNumber
}
func (this *Order) SetAdvertiseType(advertiseType AdvertiseType.AdvertiseType) (result *Order) {
	this.AdvertiseType = advertiseType
	return this
}
func (this *Order) GetAdvertiseType() (advertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *Order) SetCreateTime(createTime time.Time) (result *Order) {
	this.CreateTime = createTime
	return this
}
func (this *Order) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *Order) SetMemberId(memberId int64) (result *Order) {
	this.MemberId = memberId
	return this
}
func (this *Order) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *Order) SetMemberName(memberName string) (result *Order) {
	this.MemberName = memberName
	return this
}
func (this *Order) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *Order) SetMemberRealName(memberRealName string) (result *Order) {
	this.MemberRealName = memberRealName
	return this
}
func (this *Order) GetMemberRealName() (memberRealName string) {
	return this.MemberRealName
}
func (this *Order) SetCustomerId(customerId int64) (result *Order) {
	this.CustomerId = customerId
	return this
}
func (this *Order) GetCustomerId() (customerId int64) {
	return this.CustomerId
}
func (this *Order) SetCustomerName(customerName string) (result *Order) {
	this.CustomerName = customerName
	return this
}
func (this *Order) GetCustomerName() (customerName string) {
	return this.CustomerName
}
func (this *Order) SetCustomerRealName(customerRealName string) (result *Order) {
	this.CustomerRealName = customerRealName
	return this
}
func (this *Order) GetCustomerRealName() (customerRealName string) {
	return this.CustomerRealName
}
func (this *Order) SetCoin(coin *OtcCoin) (result *Order) {
	this.Coin = coin
	return this
}
func (this *Order) GetCoin() (coin *OtcCoin) {
	return this.Coin
}
func (this *Order) SetPrice(price math.BigDecimal) (result *Order) {
	this.Price = price
	return this
}
func (this *Order) GetPrice() (price math.BigDecimal) {
	return this.Price
}
func (this *Order) SetMaxLimit(maxLimit math.BigDecimal) (result *Order) {
	this.MaxLimit = maxLimit
	return this
}
func (this *Order) GetMaxLimit() (maxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *Order) SetCountry(country string) (result *Order) {
	this.Country = country
	return this
}
func (this *Order) GetCountry() (country string) {
	return this.Country
}
func (this *Order) SetMinLimit(minLimit math.BigDecimal) (result *Order) {
	this.MinLimit = minLimit
	return this
}
func (this *Order) GetMinLimit() (minLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *Order) SetRemark(remark string) (result *Order) {
	this.Remark = remark
	return this
}
func (this *Order) GetRemark() (remark string) {
	return this.Remark
}
func (this *Order) SetTimeLimit(timeLimit int64) (result *Order) {
	this.TimeLimit = timeLimit
	return this
}
func (this *Order) GetTimeLimit() (timeLimit int64) {
	return this.TimeLimit
}
func (this *Order) SetMoney(money math.BigDecimal) (result *Order) {
	this.Money = money
	return this
}
func (this *Order) GetMoney() (money math.BigDecimal) {
	return this.Money
}
func (this *Order) SetNumber(number math.BigDecimal) (result *Order) {
	this.Number = number
	return this
}
func (this *Order) GetNumber() (number math.BigDecimal) {
	return this.Number
}
func (this *Order) SetCommission(commission math.BigDecimal) (result *Order) {
	this.Commission = commission
	return this
}
func (this *Order) GetCommission() (commission math.BigDecimal) {
	return this.Commission
}
func (this *Order) SetStatus(status OrderStatus.OrderStatus) (result *Order) {
	this.Status = status
	return this
}
func (this *Order) GetStatus() (status OrderStatus.OrderStatus) {
	return this.Status
}
func (this *Order) SetPayTime(payTime time.Time) (result *Order) {
	this.PayTime = payTime
	return this
}
func (this *Order) GetPayTime() (payTime time.Time) {
	return this.PayTime
}
func (this *Order) SetPayMode(payMode string) (result *Order) {
	this.PayMode = payMode
	return this
}
func (this *Order) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *Order) SetAdvertiseId(advertiseId int64) (result *Order) {
	this.AdvertiseId = advertiseId
	return this
}
func (this *Order) GetAdvertiseId() (advertiseId int64) {
	return this.AdvertiseId
}
func (this *Order) SetCancelTime(cancelTime time.Time) (result *Order) {
	this.CancelTime = cancelTime
	return this
}
func (this *Order) GetCancelTime() (cancelTime time.Time) {
	return this.CancelTime
}
func (this *Order) SetReleaseTime(releaseTime time.Time) (result *Order) {
	this.ReleaseTime = releaseTime
	return this
}
func (this *Order) GetReleaseTime() (releaseTime time.Time) {
	return this.ReleaseTime
}
func (this *Order) SetAlipay(alipay *Alipay) (result *Order) {
	this.Alipay = alipay
	return this
}
func (this *Order) GetAlipay() (alipay *Alipay) {
	return this.Alipay
}
func (this *Order) SetBankInfo(bankInfo *BankInfo) (result *Order) {
	this.BankInfo = bankInfo
	return this
}
func (this *Order) GetBankInfo() (bankInfo *BankInfo) {
	return this.BankInfo
}
func (this *Order) SetWechatPay(wechatPay *WechatPay) (result *Order) {
	this.WechatPay = wechatPay
	return this
}
func (this *Order) GetWechatPay() (wechatPay *WechatPay) {
	return this.WechatPay
}
func (this *Order) SetVersion(version int64) (result *Order) {
	this.Version = version
	return this
}
func (this *Order) GetVersion() (version int64) {
	return this.Version
}

type Order struct {
	Id               int64
	OrderSn          string
	ReferenceNumber  string
	AdvertiseType    AdvertiseType.AdvertiseType
	CreateTime       time.Time
	MemberId         int64
	MemberName       string
	MemberRealName   string
	CustomerId       int64
	CustomerName     string
	CustomerRealName string
	Coin             *OtcCoin
	Price            math.BigDecimal
	MaxLimit         math.BigDecimal
	Country          string
	MinLimit         math.BigDecimal
	Remark           string
	TimeLimit        int64
	Money            math.BigDecimal
	Number           math.BigDecimal
	Commission       math.BigDecimal
	Status           OrderStatus.OrderStatus
	PayTime          time.Time
	PayMode          string
	AdvertiseId      int64
	CancelTime       time.Time
	ReleaseTime      time.Time
	Alipay           *Alipay
	BankInfo         *BankInfo
	WechatPay        *WechatPay
	Version          int64
}
