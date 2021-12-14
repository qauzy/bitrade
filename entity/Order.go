package entity

func (this *Order) SetId(id int64) {
	this.Id = id
}
func (this *Order) GetId() (id int64) {
	return this.Id
}
func (this *Order) SetOrderSn(orderSn string) {
	this.OrderSn = orderSn
}
func (this *Order) GetOrderSn() (orderSn string) {
	return this.OrderSn
}
func (this *Order) SetReferenceNumber(referenceNumber string) {
	this.ReferenceNumber = referenceNumber
}
func (this *Order) GetReferenceNumber() (referenceNumber string) {
	return this.ReferenceNumber
}
func (this *Order) SetAdvertiseType(advertiseType AdvertiseType) {
	this.AdvertiseType = advertiseType
}
func (this *Order) GetAdvertiseType() (advertiseType AdvertiseType) {
	return this.AdvertiseType
}
func (this *Order) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *Order) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *Order) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *Order) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *Order) SetMemberName(memberName string) {
	this.MemberName = memberName
}
func (this *Order) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *Order) SetMemberRealName(memberRealName string) {
	this.MemberRealName = memberRealName
}
func (this *Order) GetMemberRealName() (memberRealName string) {
	return this.MemberRealName
}
func (this *Order) SetCustomerId(customerId int64) {
	this.CustomerId = customerId
}
func (this *Order) GetCustomerId() (customerId int64) {
	return this.CustomerId
}
func (this *Order) SetCustomerName(customerName string) {
	this.CustomerName = customerName
}
func (this *Order) GetCustomerName() (customerName string) {
	return this.CustomerName
}
func (this *Order) SetCustomerRealName(customerRealName string) {
	this.CustomerRealName = customerRealName
}
func (this *Order) GetCustomerRealName() (customerRealName string) {
	return this.CustomerRealName
}
func (this *Order) SetCoin(coin OtcCoin) {
	this.Coin = coin
}
func (this *Order) GetCoin() (coin OtcCoin) {
	return this.Coin
}
func (this *Order) SetPrice(price BigDecimal) {
	this.Price = price
}
func (this *Order) GetPrice() (price BigDecimal) {
	return this.Price
}
func (this *Order) SetMaxLimit(maxLimit BigDecimal) {
	this.MaxLimit = maxLimit
}
func (this *Order) GetMaxLimit() (maxLimit BigDecimal) {
	return this.MaxLimit
}
func (this *Order) SetCountry(country string) {
	this.Country = country
}
func (this *Order) GetCountry() (country string) {
	return this.Country
}
func (this *Order) SetMinLimit(minLimit BigDecimal) {
	this.MinLimit = minLimit
}
func (this *Order) GetMinLimit() (minLimit BigDecimal) {
	return this.MinLimit
}
func (this *Order) SetRemark(remark string) {
	this.Remark = remark
}
func (this *Order) GetRemark() (remark string) {
	return this.Remark
}
func (this *Order) SetTimeLimit(timeLimit int64) {
	this.TimeLimit = timeLimit
}
func (this *Order) GetTimeLimit() (timeLimit int64) {
	return this.TimeLimit
}
func (this *Order) SetMoney(money BigDecimal) {
	this.Money = money
}
func (this *Order) GetMoney() (money BigDecimal) {
	return this.Money
}
func (this *Order) SetNumber(number BigDecimal) {
	this.Number = number
}
func (this *Order) GetNumber() (number BigDecimal) {
	return this.Number
}
func (this *Order) SetCommission(commission BigDecimal) {
	this.Commission = commission
}
func (this *Order) GetCommission() (commission BigDecimal) {
	return this.Commission
}
func (this *Order) SetStatus(status OrderStatus) {
	this.Status = status
}
func (this *Order) GetStatus() (status OrderStatus) {
	return this.Status
}
func (this *Order) SetPayTime(payTime Date) {
	this.PayTime = payTime
}
func (this *Order) GetPayTime() (payTime Date) {
	return this.PayTime
}
func (this *Order) SetPayMode(payMode string) {
	this.PayMode = payMode
}
func (this *Order) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *Order) SetAdvertiseId(advertiseId int64) {
	this.AdvertiseId = advertiseId
}
func (this *Order) GetAdvertiseId() (advertiseId int64) {
	return this.AdvertiseId
}
func (this *Order) SetCancelTime(cancelTime Date) {
	this.CancelTime = cancelTime
}
func (this *Order) GetCancelTime() (cancelTime Date) {
	return this.CancelTime
}
func (this *Order) SetReleaseTime(releaseTime Date) {
	this.ReleaseTime = releaseTime
}
func (this *Order) GetReleaseTime() (releaseTime Date) {
	return this.ReleaseTime
}
func (this *Order) SetAlipay(alipay Alipay) {
	this.Alipay = alipay
}
func (this *Order) GetAlipay() (alipay Alipay) {
	return this.Alipay
}
func (this *Order) SetBankInfo(bankInfo BankInfo) {
	this.BankInfo = bankInfo
}
func (this *Order) GetBankInfo() (bankInfo BankInfo) {
	return this.BankInfo
}
func (this *Order) SetWechatPay(wechatPay WechatPay) {
	this.WechatPay = wechatPay
}
func (this *Order) GetWechatPay() (wechatPay WechatPay) {
	return this.WechatPay
}
func (this *Order) SetVersion(version int64) {
	this.Version = version
}
func (this *Order) GetVersion() (version int64) {
	return this.Version
}

type Order struct {
	Id               int64
	OrderSn          string
	ReferenceNumber  string
	AdvertiseType    AdvertiseType
	CreateTime       Date
	MemberId         int64
	MemberName       string
	MemberRealName   string
	CustomerId       int64
	CustomerName     string
	CustomerRealName string
	Coin             OtcCoin
	Price            BigDecimal
	MaxLimit         BigDecimal
	Country          string
	MinLimit         BigDecimal
	Remark           string
	TimeLimit        int64
	Money            BigDecimal
	Number           BigDecimal
	Commission       BigDecimal
	Status           OrderStatus
	PayTime          Date
	PayMode          string
	AdvertiseId      int64
	CancelTime       Date
	ReleaseTime      Date
	Alipay           Alipay
	BankInfo         BankInfo
	WechatPay        WechatPay
	Version          int64
}
