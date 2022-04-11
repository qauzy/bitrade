package entity

import (
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/constant/OrderStatus"
	"github.com/qauzy/chocolate/xtime"
	"github.com/qauzy/math"
)

func (this *Order) SetId(Id int64) (result *Order) {
	this.Id = Id
	return this
}
func (this *Order) GetId() (Id int64) {
	return this.Id
}
func (this *Order) SetOrderSn(OrderSn string) (result *Order) {
	this.OrderSn = OrderSn
	return this
}
func (this *Order) GetOrderSn() (OrderSn string) {
	return this.OrderSn
}
func (this *Order) SetReferenceNumber(ReferenceNumber string) (result *Order) {
	this.ReferenceNumber = ReferenceNumber
	return this
}
func (this *Order) GetReferenceNumber() (ReferenceNumber string) {
	return this.ReferenceNumber
}
func (this *Order) SetAdvertiseType(AdvertiseType AdvertiseType.AdvertiseType) (result *Order) {
	this.AdvertiseType = AdvertiseType
	return this
}
func (this *Order) GetAdvertiseType() (AdvertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *Order) SetCreateTime(CreateTime xtime.Xtime) (result *Order) {
	this.CreateTime = CreateTime
	return this
}
func (this *Order) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *Order) SetMemberId(MemberId int64) (result *Order) {
	this.MemberId = MemberId
	return this
}
func (this *Order) GetMemberId() (MemberId int64) {
	return this.MemberId
}
func (this *Order) SetMemberName(MemberName string) (result *Order) {
	this.MemberName = MemberName
	return this
}
func (this *Order) GetMemberName() (MemberName string) {
	return this.MemberName
}
func (this *Order) SetMemberRealName(MemberRealName string) (result *Order) {
	this.MemberRealName = MemberRealName
	return this
}
func (this *Order) GetMemberRealName() (MemberRealName string) {
	return this.MemberRealName
}
func (this *Order) SetCustomerId(CustomerId int64) (result *Order) {
	this.CustomerId = CustomerId
	return this
}
func (this *Order) GetCustomerId() (CustomerId int64) {
	return this.CustomerId
}
func (this *Order) SetCustomerName(CustomerName string) (result *Order) {
	this.CustomerName = CustomerName
	return this
}
func (this *Order) GetCustomerName() (CustomerName string) {
	return this.CustomerName
}
func (this *Order) SetCustomerRealName(CustomerRealName string) (result *Order) {
	this.CustomerRealName = CustomerRealName
	return this
}
func (this *Order) GetCustomerRealName() (CustomerRealName string) {
	return this.CustomerRealName
}
func (this *Order) SetCoin(Coin *OtcCoin) (result *Order) {
	this.Coin = Coin
	return this
}
func (this *Order) GetCoin() (Coin *OtcCoin) {
	return this.Coin
}
func (this *Order) SetPrice(Price math.BigDecimal) (result *Order) {
	this.Price = Price
	return this
}
func (this *Order) GetPrice() (Price math.BigDecimal) {
	return this.Price
}
func (this *Order) SetMaxLimit(MaxLimit math.BigDecimal) (result *Order) {
	this.MaxLimit = MaxLimit
	return this
}
func (this *Order) GetMaxLimit() (MaxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *Order) SetCountry(Country string) (result *Order) {
	this.Country = Country
	return this
}
func (this *Order) GetCountry() (Country string) {
	return this.Country
}
func (this *Order) SetMinLimit(MinLimit math.BigDecimal) (result *Order) {
	this.MinLimit = MinLimit
	return this
}
func (this *Order) GetMinLimit() (MinLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *Order) SetRemark(Remark string) (result *Order) {
	this.Remark = Remark
	return this
}
func (this *Order) GetRemark() (Remark string) {
	return this.Remark
}
func (this *Order) SetTimeLimit(TimeLimit int) (result *Order) {
	this.TimeLimit = TimeLimit
	return this
}
func (this *Order) GetTimeLimit() (TimeLimit int) {
	return this.TimeLimit
}
func (this *Order) SetMoney(Money math.BigDecimal) (result *Order) {
	this.Money = Money
	return this
}
func (this *Order) GetMoney() (Money math.BigDecimal) {
	return this.Money
}
func (this *Order) SetNumber(Number math.BigDecimal) (result *Order) {
	this.Number = Number
	return this
}
func (this *Order) GetNumber() (Number math.BigDecimal) {
	return this.Number
}
func (this *Order) SetCommission(Commission math.BigDecimal) (result *Order) {
	this.Commission = Commission
	return this
}
func (this *Order) GetCommission() (Commission math.BigDecimal) {
	return this.Commission
}
func (this *Order) SetStatus(Status OrderStatus.OrderStatus) (result *Order) {
	this.Status = Status
	return this
}
func (this *Order) GetStatus() (Status OrderStatus.OrderStatus) {
	return this.Status
}
func (this *Order) SetPayTime(PayTime xtime.Xtime) (result *Order) {
	this.PayTime = PayTime
	return this
}
func (this *Order) GetPayTime() (PayTime xtime.Xtime) {
	return this.PayTime
}
func (this *Order) SetPayMode(PayMode string) (result *Order) {
	this.PayMode = PayMode
	return this
}
func (this *Order) GetPayMode() (PayMode string) {
	return this.PayMode
}
func (this *Order) SetAdvertiseId(AdvertiseId int64) (result *Order) {
	this.AdvertiseId = AdvertiseId
	return this
}
func (this *Order) GetAdvertiseId() (AdvertiseId int64) {
	return this.AdvertiseId
}
func (this *Order) SetCancelTime(CancelTime xtime.Xtime) (result *Order) {
	this.CancelTime = CancelTime
	return this
}
func (this *Order) GetCancelTime() (CancelTime xtime.Xtime) {
	return this.CancelTime
}
func (this *Order) SetReleaseTime(ReleaseTime xtime.Xtime) (result *Order) {
	this.ReleaseTime = ReleaseTime
	return this
}
func (this *Order) GetReleaseTime() (ReleaseTime xtime.Xtime) {
	return this.ReleaseTime
}
func (this *Order) SetAlipay(Alipay *Alipay) (result *Order) {
	this.Alipay = Alipay
	return this
}
func (this *Order) GetAlipay() (Alipay *Alipay) {
	return this.Alipay
}
func (this *Order) SetBankInfo(BankInfo *BankInfo) (result *Order) {
	this.BankInfo = BankInfo
	return this
}
func (this *Order) GetBankInfo() (BankInfo *BankInfo) {
	return this.BankInfo
}
func (this *Order) SetWechatPay(WechatPay *WechatPay) (result *Order) {
	this.WechatPay = WechatPay
	return this
}
func (this *Order) GetWechatPay() (WechatPay *WechatPay) {
	return this.WechatPay
}
func (this *Order) SetVersion(Version int64) (result *Order) {
	this.Version = Version
	return this
}
func (this *Order) GetVersion() (Version int64) {
	return this.Version
}
func NewOrder(id int64, orderSn string, referenceNumber string, advertiseType AdvertiseType.AdvertiseType, createTime xtime.Xtime, memberId int64, memberName string, memberRealName string, customerId int64, customerName string, customerRealName string, coin *OtcCoin, price math.BigDecimal, maxLimit math.BigDecimal, country string, minLimit math.BigDecimal, remark string, timeLimit int, money math.BigDecimal, number math.BigDecimal, commission math.BigDecimal, status OrderStatus.OrderStatus, payTime xtime.Xtime, payMode string, advertiseId int64, cancelTime xtime.Xtime, releaseTime xtime.Xtime, alipay *Alipay, bankInfo *BankInfo, wechatPay *WechatPay, version int64) (ret *Order) {
	ret = new(Order)
	ret.Id = id
	ret.OrderSn = orderSn
	ret.ReferenceNumber = referenceNumber
	ret.AdvertiseType = advertiseType
	ret.CreateTime = createTime
	ret.MemberId = memberId
	ret.MemberName = memberName
	ret.MemberRealName = memberRealName
	ret.CustomerId = customerId
	ret.CustomerName = customerName
	ret.CustomerRealName = customerRealName
	ret.Coin = coin
	ret.Price = price
	ret.MaxLimit = maxLimit
	ret.Country = country
	ret.MinLimit = minLimit
	ret.Remark = remark
	ret.TimeLimit = timeLimit
	ret.Money = money
	ret.Number = number
	ret.Commission = commission
	ret.Status = status
	ret.PayTime = payTime
	ret.PayMode = payMode
	ret.AdvertiseId = advertiseId
	ret.CancelTime = cancelTime
	ret.ReleaseTime = releaseTime
	ret.Alipay = alipay
	ret.BankInfo = bankInfo
	ret.WechatPay = wechatPay
	ret.Version = version
	return
}

type Order struct {
	Id               int64                       `gorm:"column:id" json:"id"`
	OrderSn          string                      `gorm:"column:order_sn" json:"orderSn"`
	ReferenceNumber  string                      `gorm:"column:reference_number" json:"referenceNumber"`
	AdvertiseType    AdvertiseType.AdvertiseType `gorm:"column:advertise_type" json:"advertiseType"`
	CreateTime       xtime.Xtime                 `gorm:"column:create_time" json:"createTime"`
	MemberId         int64                       `gorm:"column:member_id" json:"memberId"`
	MemberName       string                      `gorm:"column:member_name" json:"memberName"`
	MemberRealName   string                      `gorm:"column:member_real_name" json:"memberRealName"`
	CustomerId       int64                       `gorm:"column:customer_id" json:"customerId"`
	CustomerName     string                      `gorm:"column:customer_name" json:"customerName"`
	CustomerRealName string                      `gorm:"column:customer_real_name" json:"customerRealName"`
	Coin             *OtcCoin                    `gorm:"column:coin" json:"coin"`
	Price            math.BigDecimal             `gorm:"column:price" json:"price"`
	MaxLimit         math.BigDecimal             `gorm:"column:max_limit" json:"maxLimit"`
	Country          string                      `gorm:"column:country" json:"country"`
	MinLimit         math.BigDecimal             `gorm:"column:min_limit" json:"minLimit"`
	Remark           string                      `gorm:"column:remark" json:"remark"`
	TimeLimit        int                         `gorm:"column:time_limit" json:"timeLimit"`
	Money            math.BigDecimal             `gorm:"column:money" json:"money"`
	Number           math.BigDecimal             `gorm:"column:number" json:"number"`
	Commission       math.BigDecimal             `gorm:"column:commission" json:"commission"`
	Status           OrderStatus.OrderStatus     `gorm:"column:status" json:"status"`
	PayTime          xtime.Xtime                 `gorm:"column:pay_time" json:"payTime"`
	PayMode          string                      `gorm:"column:pay_mode" json:"payMode"`
	AdvertiseId      int64                       `gorm:"column:advertise_id" json:"advertiseId"`
	CancelTime       xtime.Xtime                 `gorm:"column:cancel_time" json:"cancelTime"`
	ReleaseTime      xtime.Xtime                 `gorm:"column:release_time" json:"releaseTime"`
	Alipay           *Alipay                     `gorm:"column:alipay" json:"alipay"`
	BankInfo         *BankInfo                   `gorm:"column:bank_info" json:"bankInfo"`
	WechatPay        *WechatPay                  `gorm:"column:wechat_pay" json:"wechatPay"`
	Version          int64                       `gorm:"column:version" json:"version"`
}
