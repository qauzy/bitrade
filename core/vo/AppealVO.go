package vo

import (
	"github.com/qauzy/math"
	"time"
)

func (this *AppealVO) SetAppealId(appealId math.BigInteger) (result *AppealVO) {
	this.AppealId = appealId
	return this
}
func (this *AppealVO) GetAppealId() (appealId math.BigInteger) {
	return this.AppealId
}
func (this *AppealVO) SetAdvertiseCreaterUserName(advertiseCreaterUserName string) (result *AppealVO) {
	this.AdvertiseCreaterUserName = advertiseCreaterUserName
	return this
}
func (this *AppealVO) GetAdvertiseCreaterUserName() (advertiseCreaterUserName string) {
	return this.AdvertiseCreaterUserName
}
func (this *AppealVO) SetAdvertiseCreaterName(advertiseCreaterName string) (result *AppealVO) {
	this.AdvertiseCreaterName = advertiseCreaterName
	return this
}
func (this *AppealVO) GetAdvertiseCreaterName() (advertiseCreaterName string) {
	return this.AdvertiseCreaterName
}
func (this *AppealVO) SetCustomerUserName(customerUserName string) (result *AppealVO) {
	this.CustomerUserName = customerUserName
	return this
}
func (this *AppealVO) GetCustomerUserName() (customerUserName string) {
	return this.CustomerUserName
}
func (this *AppealVO) SetCustomerName(customerName string) (result *AppealVO) {
	this.CustomerName = customerName
	return this
}
func (this *AppealVO) GetCustomerName() (customerName string) {
	return this.CustomerName
}
func (this *AppealVO) SetAssociateName(associateName string) (result *AppealVO) {
	this.AssociateName = associateName
	return this
}
func (this *AppealVO) GetAssociateName() (associateName string) {
	return this.AssociateName
}
func (this *AppealVO) SetAssociateUsername(associateUsername string) (result *AppealVO) {
	this.AssociateUsername = associateUsername
	return this
}
func (this *AppealVO) GetAssociateUsername() (associateUsername string) {
	return this.AssociateUsername
}
func (this *AppealVO) SetInitiatorUsername(initiatorUsername string) (result *AppealVO) {
	this.InitiatorUsername = initiatorUsername
	return this
}
func (this *AppealVO) GetInitiatorUsername() (initiatorUsername string) {
	return this.InitiatorUsername
}
func (this *AppealVO) SetInitiatorName(initiatorName string) (result *AppealVO) {
	this.InitiatorName = initiatorName
	return this
}
func (this *AppealVO) GetInitiatorName() (initiatorName string) {
	return this.InitiatorName
}
func (this *AppealVO) SetFee(fee math.BigDecimal) (result *AppealVO) {
	this.Fee = fee
	return this
}
func (this *AppealVO) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *AppealVO) SetNumber(number math.BigDecimal) (result *AppealVO) {
	this.Number = number
	return this
}
func (this *AppealVO) GetNumber() (number math.BigDecimal) {
	return this.Number
}
func (this *AppealVO) SetMoney(money math.BigDecimal) (result *AppealVO) {
	this.Money = money
	return this
}
func (this *AppealVO) GetMoney() (money math.BigDecimal) {
	return this.Money
}
func (this *AppealVO) SetOrderSn(orderSn string) (result *AppealVO) {
	this.OrderSn = orderSn
	return this
}
func (this *AppealVO) GetOrderSn() (orderSn string) {
	return this.OrderSn
}
func (this *AppealVO) SetTransactionTime(transactionTime time.Time) (result *AppealVO) {
	this.TransactionTime = transactionTime
	return this
}
func (this *AppealVO) GetTransactionTime() (transactionTime time.Time) {
	return this.TransactionTime
}
func (this *AppealVO) SetCreateTime(createTime time.Time) (result *AppealVO) {
	this.CreateTime = createTime
	return this
}
func (this *AppealVO) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *AppealVO) SetDealTime(dealTime time.Time) (result *AppealVO) {
	this.DealTime = dealTime
	return this
}
func (this *AppealVO) GetDealTime() (dealTime time.Time) {
	return this.DealTime
}
func (this *AppealVO) SetPayMode(payMode string) (result *AppealVO) {
	this.PayMode = payMode
	return this
}
func (this *AppealVO) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *AppealVO) SetCoinName(coinName string) (result *AppealVO) {
	this.CoinName = coinName
	return this
}
func (this *AppealVO) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *AppealVO) SetOrderStatus(orderStatus int) (result *AppealVO) {
	this.OrderStatus = orderStatus
	return this
}
func (this *AppealVO) GetOrderStatus() (orderStatus int) {
	return this.OrderStatus
}
func (this *AppealVO) SetIsSuccess(isSuccess int) (result *AppealVO) {
	this.IsSuccess = isSuccess
	return this
}
func (this *AppealVO) GetIsSuccess() (isSuccess int) {
	return this.IsSuccess
}
func (this *AppealVO) SetAdvertiseType(advertiseType int) (result *AppealVO) {
	this.AdvertiseType = advertiseType
	return this
}
func (this *AppealVO) GetAdvertiseType() (advertiseType int) {
	return this.AdvertiseType
}
func (this *AppealVO) SetStatus(status int) (result *AppealVO) {
	this.Status = status
	return this
}
func (this *AppealVO) GetStatus() (status int) {
	return this.Status
}
func (this *AppealVO) SetRemark(remark string) (result *AppealVO) {
	this.Remark = remark
	return this
}
func (this *AppealVO) GetRemark() (remark string) {
	return this.Remark
}

type AppealVO struct {
	AppealId                 math.BigInteger
	AdvertiseCreaterUserName string
	AdvertiseCreaterName     string
	CustomerUserName         string
	CustomerName             string
	AssociateName            string
	AssociateUsername        string
	InitiatorUsername        string
	InitiatorName            string
	Fee                      math.BigDecimal
	Number                   math.BigDecimal
	Money                    math.BigDecimal
	OrderSn                  string
	TransactionTime          time.Time
	CreateTime               time.Time
	DealTime                 time.Time
	PayMode                  string
	CoinName                 string
	OrderStatus              int
	IsSuccess                int
	AdvertiseType            int
	Status                   int
	Remark                   string
}
