package entity

func (this *OtcCoin) SetId(id int64) {
	this.Id = id
}
func (this *OtcCoin) GetId() (id int64) {
	return this.Id
}
func (this *OtcCoin) SetName(name string) {
	this.Name = name
}
func (this *OtcCoin) GetName() (name string) {
	return this.Name
}
func (this *OtcCoin) SetNameCn(nameCn string) {
	this.NameCn = nameCn
}
func (this *OtcCoin) GetNameCn() (nameCn string) {
	return this.NameCn
}
func (this *OtcCoin) SetUnit(unit string) {
	this.Unit = unit
}
func (this *OtcCoin) GetUnit() (unit string) {
	return this.Unit
}
func (this *OtcCoin) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *OtcCoin) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *OtcCoin) SetJyRate(jyRate BigDecimal) {
	this.JyRate = jyRate
}
func (this *OtcCoin) GetJyRate() (jyRate BigDecimal) {
	return this.JyRate
}
func (this *OtcCoin) SetSellMinAmount(sellMinAmount BigDecimal) {
	this.SellMinAmount = sellMinAmount
}
func (this *OtcCoin) GetSellMinAmount() (sellMinAmount BigDecimal) {
	return this.SellMinAmount
}
func (this *OtcCoin) SetBuyMinAmount(buyMinAmount BigDecimal) {
	this.BuyMinAmount = buyMinAmount
}
func (this *OtcCoin) GetBuyMinAmount() (buyMinAmount BigDecimal) {
	return this.BuyMinAmount
}
func (this *OtcCoin) SetSort(sort int) {
	this.Sort = sort
}
func (this *OtcCoin) GetSort() (sort int) {
	return this.Sort
}
func (this *OtcCoin) SetIsPlatformCoin(isPlatformCoin BooleanEnum) {
	this.IsPlatformCoin = isPlatformCoin
}
func (this *OtcCoin) GetIsPlatformCoin() (isPlatformCoin BooleanEnum) {
	return this.IsPlatformCoin
}
func (this *OtcCoin) SetCoinScale(coinScale int64) {
	this.CoinScale = coinScale
}
func (this *OtcCoin) GetCoinScale() (coinScale int64) {
	return this.CoinScale
}
func (this *OtcCoin) SetMaxTradingTime(maxTradingTime int64) {
	this.MaxTradingTime = maxTradingTime
}
func (this *OtcCoin) GetMaxTradingTime() (maxTradingTime int64) {
	return this.MaxTradingTime
}
func (this *OtcCoin) SetMaxVolume(maxVolume int64) {
	this.MaxVolume = maxVolume
}
func (this *OtcCoin) GetMaxVolume() (maxVolume int64) {
	return this.MaxVolume
}

type OtcCoin struct {
	Id             int64
	Name           string
	NameCn         string
	Unit           string
	Status         CommonStatus
	JyRate         BigDecimal
	SellMinAmount  BigDecimal
	BuyMinAmount   BigDecimal
	Sort           int
	IsPlatformCoin BooleanEnum
	CoinScale      int64
	MaxTradingTime int64
	MaxVolume      int64
}
