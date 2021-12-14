package entity

func (this *Advertise) SetId(id int64) {
	this.Id = id
}
func (this *Advertise) GetId() (id int64) {
	return this.Id
}
func (this *Advertise) SetPrice(price BigDecimal) {
	this.Price = price
}
func (this *Advertise) GetPrice() (price BigDecimal) {
	return this.Price
}
func (this *Advertise) SetAdvertiseType(advertiseType AdvertiseType) {
	this.AdvertiseType = advertiseType
}
func (this *Advertise) GetAdvertiseType() (advertiseType AdvertiseType) {
	return this.AdvertiseType
}
func (this *Advertise) SetMember(member Member) {
	this.Member = member
}
func (this *Advertise) GetMember() (member Member) {
	return this.Member
}
func (this *Advertise) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *Advertise) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *Advertise) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *Advertise) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}
func (this *Advertise) SetCoin(coin OtcCoin) {
	this.Coin = coin
}
func (this *Advertise) GetCoin() (coin OtcCoin) {
	return this.Coin
}
func (this *Advertise) SetCountry(country Country) {
	this.Country = country
}
func (this *Advertise) GetCountry() (country Country) {
	return this.Country
}
func (this *Advertise) SetMinLimit(minLimit BigDecimal) {
	this.MinLimit = minLimit
}
func (this *Advertise) GetMinLimit() (minLimit BigDecimal) {
	return this.MinLimit
}
func (this *Advertise) SetMaxLimit(maxLimit BigDecimal) {
	this.MaxLimit = maxLimit
}
func (this *Advertise) GetMaxLimit() (maxLimit BigDecimal) {
	return this.MaxLimit
}
func (this *Advertise) SetRemark(remark string) {
	this.Remark = remark
}
func (this *Advertise) GetRemark() (remark string) {
	return this.Remark
}
func (this *Advertise) SetTimeLimit(timeLimit int64) {
	this.TimeLimit = timeLimit
}
func (this *Advertise) GetTimeLimit() (timeLimit int64) {
	return this.TimeLimit
}
func (this *Advertise) SetPremiseRate(premiseRate BigDecimal) {
	this.PremiseRate = premiseRate
}
func (this *Advertise) GetPremiseRate() (premiseRate BigDecimal) {
	return this.PremiseRate
}
func (this *Advertise) SetLevel(level AdvertiseLevel) {
	this.Level = level
}
func (this *Advertise) GetLevel() (level AdvertiseLevel) {
	return this.Level
}
func (this *Advertise) SetPayMode(payMode string) {
	this.PayMode = payMode
}
func (this *Advertise) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *Advertise) SetVersion(version int64) {
	this.Version = version
}
func (this *Advertise) GetVersion() (version int64) {
	return this.Version
}
func (this *Advertise) SetStatus(status AdvertiseControlStatus) {
	this.Status = status
}
func (this *Advertise) GetStatus() (status AdvertiseControlStatus) {
	return this.Status
}
func (this *Advertise) SetPriceType(priceType PriceType) {
	this.PriceType = priceType
}
func (this *Advertise) GetPriceType() (priceType PriceType) {
	return this.PriceType
}
func (this *Advertise) SetNumber(number BigDecimal) {
	this.Number = number
}
func (this *Advertise) GetNumber() (number BigDecimal) {
	return this.Number
}
func (this *Advertise) SetDealAmount(dealAmount BigDecimal) {
	this.DealAmount = dealAmount
}
func (this *Advertise) GetDealAmount() (dealAmount BigDecimal) {
	return this.DealAmount
}
func (this *Advertise) SetRemainAmount(remainAmount BigDecimal) {
	this.RemainAmount = remainAmount
}
func (this *Advertise) GetRemainAmount() (remainAmount BigDecimal) {
	return this.RemainAmount
}
func (this *Advertise) SetAuto(auto BooleanEnum) {
	this.Auto = auto
}
func (this *Advertise) GetAuto() (auto BooleanEnum) {
	return this.Auto
}
func (this *Advertise) SetMarketPrice(marketPrice BigDecimal) {
	this.MarketPrice = marketPrice
}
func (this *Advertise) GetMarketPrice() (marketPrice BigDecimal) {
	return this.MarketPrice
}
func (this *Advertise) SetAutoword(autoword string) {
	this.Autoword = autoword
}
func (this *Advertise) GetAutoword() (autoword string) {
	return this.Autoword
}
func (this *Advertise) SetLimitMoney(limitMoney string) {
	this.LimitMoney = limitMoney
}
func (this *Advertise) GetLimitMoney() (limitMoney string) {
	return this.LimitMoney
}
func (this *Advertise) SetUsername(username string) {
	this.Username = username
}
func (this *Advertise) GetUsername() (username string) {
	return this.Username
}
func (this *Advertise) SetCoinUnit(coinUnit string) {
	this.CoinUnit = coinUnit
}
func (this *Advertise) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}

type Advertise struct {
	Id            int64
	Price         BigDecimal
	AdvertiseType AdvertiseType
	Member        Member
	CreateTime    Date
	UpdateTime    Date
	Coin          OtcCoin
	Country       Country
	MinLimit      BigDecimal
	MaxLimit      BigDecimal
	Remark        string
	TimeLimit     int64
	PremiseRate   BigDecimal
	Level         AdvertiseLevel
	PayMode       string
	Version       int64
	Status        AdvertiseControlStatus
	PriceType     PriceType
	Number        BigDecimal
	DealAmount    BigDecimal
	RemainAmount  BigDecimal
	Auto          BooleanEnum
	MarketPrice   BigDecimal
	Autoword      string
	LimitMoney    string
	Username      string
	CoinUnit      string
}
