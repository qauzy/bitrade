package entity

func (this *CoinThumb) SetSymbol(symbol string) {
	this.Symbol = symbol
}
func (this *CoinThumb) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *CoinThumb) SetOpen(open BigDecimal) {
	this.Open = open
}
func (this *CoinThumb) GetOpen() (open BigDecimal) {
	return this.Open
}
func (this *CoinThumb) SetHigh(high BigDecimal) {
	this.High = high
}
func (this *CoinThumb) GetHigh() (high BigDecimal) {
	return this.High
}
func (this *CoinThumb) SetLow(low BigDecimal) {
	this.Low = low
}
func (this *CoinThumb) GetLow() (low BigDecimal) {
	return this.Low
}
func (this *CoinThumb) SetClose(close BigDecimal) {
	this.Close = close
}
func (this *CoinThumb) GetClose() (close BigDecimal) {
	return this.Close
}
func (this *CoinThumb) SetChg(chg BigDecimal) {
	this.Chg = chg
}
func (this *CoinThumb) GetChg() (chg BigDecimal) {
	return this.Chg
}
func (this *CoinThumb) SetChange(change BigDecimal) {
	this.Change = change
}
func (this *CoinThumb) GetChange() (change BigDecimal) {
	return this.Change
}
func (this *CoinThumb) SetVolume(volume BigDecimal) {
	this.Volume = volume
}
func (this *CoinThumb) GetVolume() (volume BigDecimal) {
	return this.Volume
}
func (this *CoinThumb) SetTurnover(turnover BigDecimal) {
	this.Turnover = turnover
}
func (this *CoinThumb) GetTurnover() (turnover BigDecimal) {
	return this.Turnover
}
func (this *CoinThumb) SetLastDayClose(lastDayClose BigDecimal) {
	this.LastDayClose = lastDayClose
}
func (this *CoinThumb) GetLastDayClose() (lastDayClose BigDecimal) {
	return this.LastDayClose
}
func (this *CoinThumb) SetUsdRate(usdRate BigDecimal) {
	this.UsdRate = usdRate
}
func (this *CoinThumb) GetUsdRate() (usdRate BigDecimal) {
	return this.UsdRate
}
func (this *CoinThumb) SetBaseUsdRate(baseUsdRate BigDecimal) {
	this.BaseUsdRate = baseUsdRate
}
func (this *CoinThumb) GetBaseUsdRate() (baseUsdRate BigDecimal) {
	return this.BaseUsdRate
}
func (this *CoinThumb) SetCloseStr(closeStr string) {
	this.CloseStr = closeStr
}
func (this *CoinThumb) GetCloseStr() (closeStr string) {
	return this.CloseStr
}
func (this *CoinThumb) SetTrend(trend []BigDecimal) {
	this.Trend = trend
}
func (this *CoinThumb) GetTrend() (trend []BigDecimal) {
	return this.Trend
}
func (this *CoinThumb) SetProportion(proportion BigDecimal) {
	this.Proportion = proportion
}
func (this *CoinThumb) GetProportion() (proportion BigDecimal) {
	return this.Proportion
}
func (this *CoinThumb) SetCnyPrice(cnyPrice string) {
	this.CnyPrice = cnyPrice
}
func (this *CoinThumb) GetCnyPrice() (cnyPrice string) {
	return this.CnyPrice
}

type CoinThumb struct {
	Symbol       string
	Open         BigDecimal
	High         BigDecimal
	Low          BigDecimal
	Close        BigDecimal
	Chg          BigDecimal
	Change       BigDecimal
	Volume       BigDecimal
	Turnover     BigDecimal
	LastDayClose BigDecimal
	UsdRate      BigDecimal
	BaseUsdRate  BigDecimal
	CloseStr     string
	Trend        []BigDecimal
	Proportion   BigDecimal
	CnyPrice     string
}
