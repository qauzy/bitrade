package entity

import "github.com/qauzy/math"

func (this *CoinThumb) SetSymbol(symbol string) (result *CoinThumb) {
	this.Symbol = symbol
	return this
}
func (this *CoinThumb) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *CoinThumb) SetOpen(open math.BigDecimal) (result *CoinThumb) {
	this.Open = open
	return this
}
func (this *CoinThumb) GetOpen() (open math.BigDecimal) {
	return this.Open
}
func (this *CoinThumb) SetHigh(high math.BigDecimal) (result *CoinThumb) {
	this.High = high
	return this
}
func (this *CoinThumb) GetHigh() (high math.BigDecimal) {
	return this.High
}
func (this *CoinThumb) SetLow(low math.BigDecimal) (result *CoinThumb) {
	this.Low = low
	return this
}
func (this *CoinThumb) GetLow() (low math.BigDecimal) {
	return this.Low
}
func (this *CoinThumb) SetClose(close math.BigDecimal) (result *CoinThumb) {
	this.Close = close
	return this
}
func (this *CoinThumb) GetClose() (close math.BigDecimal) {
	return this.Close
}
func (this *CoinThumb) SetChg(chg math.BigDecimal) (result *CoinThumb) {
	this.Chg = chg
	return this
}
func (this *CoinThumb) GetChg() (chg math.BigDecimal) {
	return this.Chg
}
func (this *CoinThumb) SetChange(change math.BigDecimal) (result *CoinThumb) {
	this.Change = change
	return this
}
func (this *CoinThumb) GetChange() (change math.BigDecimal) {
	return this.Change
}
func (this *CoinThumb) SetVolume(volume math.BigDecimal) (result *CoinThumb) {
	this.Volume = volume
	return this
}
func (this *CoinThumb) GetVolume() (volume math.BigDecimal) {
	return this.Volume
}
func (this *CoinThumb) SetTurnover(turnover math.BigDecimal) (result *CoinThumb) {
	this.Turnover = turnover
	return this
}
func (this *CoinThumb) GetTurnover() (turnover math.BigDecimal) {
	return this.Turnover
}
func (this *CoinThumb) SetLastDayClose(lastDayClose math.BigDecimal) (result *CoinThumb) {
	this.LastDayClose = lastDayClose
	return this
}
func (this *CoinThumb) GetLastDayClose() (lastDayClose math.BigDecimal) {
	return this.LastDayClose
}
func (this *CoinThumb) SetUsdRate(usdRate math.BigDecimal) (result *CoinThumb) {
	this.UsdRate = usdRate
	return this
}
func (this *CoinThumb) GetUsdRate() (usdRate math.BigDecimal) {
	return this.UsdRate
}
func (this *CoinThumb) SetBaseUsdRate(baseUsdRate math.BigDecimal) (result *CoinThumb) {
	this.BaseUsdRate = baseUsdRate
	return this
}
func (this *CoinThumb) GetBaseUsdRate() (baseUsdRate math.BigDecimal) {
	return this.BaseUsdRate
}
func (this *CoinThumb) SetCloseStr(closeStr string) (result *CoinThumb) {
	this.CloseStr = closeStr
	return this
}
func (this *CoinThumb) GetCloseStr() (closeStr string) {
	return this.CloseStr
}
func (this *CoinThumb) SetTrend(trend []math.BigDecimal) (result *CoinThumb) {
	this.Trend = trend
	return this
}
func (this *CoinThumb) GetTrend() (trend []math.BigDecimal) {
	return this.Trend
}
func (this *CoinThumb) SetProportion(proportion math.BigDecimal) (result *CoinThumb) {
	this.Proportion = proportion
	return this
}
func (this *CoinThumb) GetProportion() (proportion math.BigDecimal) {
	return this.Proportion
}
func (this *CoinThumb) SetCnyPrice(cnyPrice string) (result *CoinThumb) {
	this.CnyPrice = cnyPrice
	return this
}
func (this *CoinThumb) GetCnyPrice() (cnyPrice string) {
	return this.CnyPrice
}

type CoinThumb struct {
	Symbol       string
	Open         math.BigDecimal
	High         math.BigDecimal
	Low          math.BigDecimal
	Close        math.BigDecimal
	Chg          math.BigDecimal
	Change       math.BigDecimal
	Volume       math.BigDecimal
	Turnover     math.BigDecimal
	LastDayClose math.BigDecimal
	UsdRate      math.BigDecimal
	BaseUsdRate  math.BigDecimal
	CloseStr     string
	Trend        []math.BigDecimal
	Proportion   math.BigDecimal
	CnyPrice     string
}
