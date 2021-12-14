package transform

func (this *MemberAdvertiseDetail) SetId(id int64) {
	this.Id = id
}
func (this *MemberAdvertiseDetail) GetId() (id int64) {
	return this.Id
}
func (this *MemberAdvertiseDetail) SetCoinId(coinId int64) {
	this.CoinId = coinId
}
func (this *MemberAdvertiseDetail) GetCoinId() (coinId int64) {
	return this.CoinId
}
func (this *MemberAdvertiseDetail) SetCoinName(coinName string) {
	this.CoinName = coinName
}
func (this *MemberAdvertiseDetail) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *MemberAdvertiseDetail) SetCoinNameCn(coinNameCn string) {
	this.CoinNameCn = coinNameCn
}
func (this *MemberAdvertiseDetail) GetCoinNameCn() (coinNameCn string) {
	return this.CoinNameCn
}
func (this *MemberAdvertiseDetail) SetCoinUnit(coinUnit string) {
	this.CoinUnit = coinUnit
}
func (this *MemberAdvertiseDetail) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *MemberAdvertiseDetail) SetCountry(country Country) {
	this.Country = country
}
func (this *MemberAdvertiseDetail) GetCountry() (country Country) {
	return this.Country
}
func (this *MemberAdvertiseDetail) SetPriceType(priceType PriceType) {
	this.PriceType = priceType
}
func (this *MemberAdvertiseDetail) GetPriceType() (priceType PriceType) {
	return this.PriceType
}
func (this *MemberAdvertiseDetail) SetPrice(price BigDecimal) {
	this.Price = price
}
func (this *MemberAdvertiseDetail) GetPrice() (price BigDecimal) {
	return this.Price
}
func (this *MemberAdvertiseDetail) SetAdvertiseType(advertiseType AdvertiseType) {
	this.AdvertiseType = advertiseType
}
func (this *MemberAdvertiseDetail) GetAdvertiseType() (advertiseType AdvertiseType) {
	return this.AdvertiseType
}
func (this *MemberAdvertiseDetail) SetMinLimit(minLimit BigDecimal) {
	this.MinLimit = minLimit
}
func (this *MemberAdvertiseDetail) GetMinLimit() (minLimit BigDecimal) {
	return this.MinLimit
}
func (this *MemberAdvertiseDetail) SetMaxLimit(maxLimit BigDecimal) {
	this.MaxLimit = maxLimit
}
func (this *MemberAdvertiseDetail) GetMaxLimit() (maxLimit BigDecimal) {
	return this.MaxLimit
}
func (this *MemberAdvertiseDetail) SetRemark(remark string) {
	this.Remark = remark
}
func (this *MemberAdvertiseDetail) GetRemark() (remark string) {
	return this.Remark
}
func (this *MemberAdvertiseDetail) SetTimeLimit(timeLimit int64) {
	this.TimeLimit = timeLimit
}
func (this *MemberAdvertiseDetail) GetTimeLimit() (timeLimit int64) {
	return this.TimeLimit
}
func (this *MemberAdvertiseDetail) SetPremiseRate(premiseRate BigDecimal) {
	this.PremiseRate = premiseRate
}
func (this *MemberAdvertiseDetail) GetPremiseRate() (premiseRate BigDecimal) {
	return this.PremiseRate
}
func (this *MemberAdvertiseDetail) SetPayMode(payMode string) {
	this.PayMode = payMode
}
func (this *MemberAdvertiseDetail) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *MemberAdvertiseDetail) SetStatus(status AdvertiseControlStatus) {
	this.Status = status
}
func (this *MemberAdvertiseDetail) GetStatus() (status AdvertiseControlStatus) {
	return this.Status
}
func (this *MemberAdvertiseDetail) SetNumber(number BigDecimal) {
	this.Number = number
}
func (this *MemberAdvertiseDetail) GetNumber() (number BigDecimal) {
	return this.Number
}
func (this *MemberAdvertiseDetail) SetMarketPrice(marketPrice BigDecimal) {
	this.MarketPrice = marketPrice
}
func (this *MemberAdvertiseDetail) GetMarketPrice() (marketPrice BigDecimal) {
	return this.MarketPrice
}
func (this *MemberAdvertiseDetail) SetAuto(auto BooleanEnum) {
	this.Auto = auto
}
func (this *MemberAdvertiseDetail) GetAuto() (auto BooleanEnum) {
	return this.Auto
}
func (this *MemberAdvertiseDetail) SetAutoword(autoword string) {
	this.Autoword = autoword
}
func (this *MemberAdvertiseDetail) GetAutoword() (autoword string) {
	return this.Autoword
}
func ToMemberAdvertiseDetail(advertise Advertise) (result MemberAdvertiseDetail) {
	return MemberAdvertiseDetail.Builder().Id(advertise.GetId()).AdvertiseType(advertise.GetAdvertiseType()).CoinId(advertise.GetCoin().GetId()).CoinName(advertise.GetCoin().GetName()).CoinNameCn(advertise.GetCoin().GetNameCn()).CoinUnit(advertise.GetCoin().GetUnit()).Country(advertise.GetCountry()).Auto(advertise.GetAuto()).MaxLimit(advertise.GetMaxLimit()).MinLimit(advertise.GetMinLimit()).Number(advertise.GetNumber()).PayMode(advertise.GetPayMode()).PremiseRate(advertise.GetPremiseRate()).Price(advertise.GetPrice()).PriceType(advertise.GetPriceType()).Remark(advertise.GetRemark()).Status(advertise.GetStatus()).TimeLimit(advertise.GetTimeLimit()).Autoword(advertise.GetAutoword()).Build()
}

type MemberAdvertiseDetail struct {
	Id            int64
	CoinId        int64
	CoinName      string
	CoinNameCn    string
	CoinUnit      string
	Country       Country
	PriceType     PriceType
	Price         BigDecimal
	AdvertiseType AdvertiseType
	MinLimit      BigDecimal
	MaxLimit      BigDecimal
	Remark        string
	TimeLimit     int64
	PremiseRate   BigDecimal
	PayMode       string
	Status        AdvertiseControlStatus
	Number        BigDecimal
	MarketPrice   BigDecimal
	Auto          BooleanEnum
	Autoword      string
}
