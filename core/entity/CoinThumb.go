package entity

import (
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
)

func (this *CoinThumb) SetSymbol(Symbol string) (result *CoinThumb) {
	this.Symbol = Symbol
	return this
}
func (this *CoinThumb) GetSymbol() (Symbol string) {
	return this.Symbol
}
func (this *CoinThumb) SetOpen(Open math.BigDecimal) (result *CoinThumb) {
	this.Open = Open
	return this
}
func (this *CoinThumb) GetOpen() (Open math.BigDecimal) {
	return this.Open
}
func (this *CoinThumb) SetHigh(High math.BigDecimal) (result *CoinThumb) {
	this.High = High
	return this
}
func (this *CoinThumb) GetHigh() (High math.BigDecimal) {
	return this.High
}
func (this *CoinThumb) SetLow(Low math.BigDecimal) (result *CoinThumb) {
	this.Low = Low
	return this
}
func (this *CoinThumb) GetLow() (Low math.BigDecimal) {
	return this.Low
}
func (this *CoinThumb) SetClose(Close math.BigDecimal) (result *CoinThumb) {
	this.Close = Close
	return this
}
func (this *CoinThumb) GetClose() (Close math.BigDecimal) {
	return this.Close
}
func (this *CoinThumb) SetChg(Chg math.BigDecimal) (result *CoinThumb) {
	this.Chg = Chg
	return this
}
func (this *CoinThumb) GetChg() (Chg math.BigDecimal) {
	return this.Chg
}
func (this *CoinThumb) SetChange(Change math.BigDecimal) (result *CoinThumb) {
	this.Change = Change
	return this
}
func (this *CoinThumb) GetChange() (Change math.BigDecimal) {
	return this.Change
}
func (this *CoinThumb) SetVolume(Volume math.BigDecimal) (result *CoinThumb) {
	this.Volume = Volume
	return this
}
func (this *CoinThumb) GetVolume() (Volume math.BigDecimal) {
	return this.Volume
}
func (this *CoinThumb) SetTurnover(Turnover math.BigDecimal) (result *CoinThumb) {
	this.Turnover = Turnover
	return this
}
func (this *CoinThumb) GetTurnover() (Turnover math.BigDecimal) {
	return this.Turnover
}
func (this *CoinThumb) SetLastDayClose(LastDayClose math.BigDecimal) (result *CoinThumb) {
	this.LastDayClose = LastDayClose
	return this
}
func (this *CoinThumb) GetLastDayClose() (LastDayClose math.BigDecimal) {
	return this.LastDayClose
}
func (this *CoinThumb) SetUsdRate(UsdRate math.BigDecimal) (result *CoinThumb) {
	this.UsdRate = UsdRate
	return this
}
func (this *CoinThumb) GetUsdRate() (UsdRate math.BigDecimal) {
	return this.UsdRate
}
func (this *CoinThumb) SetBaseUsdRate(BaseUsdRate math.BigDecimal) (result *CoinThumb) {
	this.BaseUsdRate = BaseUsdRate
	return this
}
func (this *CoinThumb) GetBaseUsdRate() (BaseUsdRate math.BigDecimal) {
	return this.BaseUsdRate
}
func (this *CoinThumb) SetCloseStr(CloseStr string) (result *CoinThumb) {
	this.CloseStr = CloseStr
	return this
}
func (this *CoinThumb) GetCloseStr() (CloseStr string) {
	return this.CloseStr
}
func (this *CoinThumb) SetTrend(Trend arraylist.List[math.BigDecimal]) (result *CoinThumb) {
	this.Trend = Trend
	return this
}
func (this *CoinThumb) GetTrend() (Trend arraylist.List[math.BigDecimal]) {
	return this.Trend
}
func (this *CoinThumb) SetProportion(Proportion math.BigDecimal) (result *CoinThumb) {
	this.Proportion = Proportion
	return this
}
func (this *CoinThumb) GetProportion() (Proportion math.BigDecimal) {
	return this.Proportion
}
func (this *CoinThumb) SetCnyPrice(CnyPrice string) (result *CoinThumb) {
	this.CnyPrice = CnyPrice
	return this
}
func (this *CoinThumb) GetCnyPrice() (CnyPrice string) {
	return this.CnyPrice
}
func NewCoinThumb(symbol string, open math.BigDecimal, high math.BigDecimal, low math.BigDecimal, close math.BigDecimal, chg math.BigDecimal, change math.BigDecimal, volume math.BigDecimal, turnover math.BigDecimal, lastDayClose math.BigDecimal, usdRate math.BigDecimal, baseUsdRate math.BigDecimal, closeStr string, trend arraylist.List[math.BigDecimal], proportion math.BigDecimal, cnyPrice string) (ret *CoinThumb) {
	ret = new(CoinThumb)
	ret.Symbol = symbol
	ret.Open = open
	ret.High = high
	ret.Low = low
	ret.Close = close
	ret.Chg = chg
	ret.Change = change
	ret.Volume = volume
	ret.Turnover = turnover
	ret.LastDayClose = lastDayClose
	ret.UsdRate = usdRate
	ret.BaseUsdRate = baseUsdRate
	ret.CloseStr = closeStr
	ret.Trend = trend
	ret.Proportion = proportion
	ret.CnyPrice = cnyPrice
	return
}

type CoinThumb struct {
	Symbol       string                          `gorm:"column:symbol" json:"symbol"`
	Open         math.BigDecimal                 `gorm:"column:open" json:"open"`
	High         math.BigDecimal                 `gorm:"column:high" json:"high"`
	Low          math.BigDecimal                 `gorm:"column:low" json:"low"`
	Close        math.BigDecimal                 `gorm:"column:close" json:"close"`
	Chg          math.BigDecimal                 `gorm:"column:chg" json:"chg"`
	Change       math.BigDecimal                 `gorm:"column:change" json:"change"`
	Volume       math.BigDecimal                 `gorm:"column:volume" json:"volume"`
	Turnover     math.BigDecimal                 `gorm:"column:turnover" json:"turnover"`
	LastDayClose math.BigDecimal                 `gorm:"column:last_day_close" json:"lastDayClose"`
	UsdRate      math.BigDecimal                 `gorm:"column:usd_rate" json:"usdRate"`
	BaseUsdRate  math.BigDecimal                 `gorm:"column:base_usd_rate" json:"baseUsdRate"`
	CloseStr     string                          `gorm:"column:close_str" json:"closeStr"`
	Trend        arraylist.List[math.BigDecimal] `gorm:"column:trend" json:"trend"`
	Proportion   math.BigDecimal                 `gorm:"column:proportion" json:"proportion"`
	CnyPrice     string                          `gorm:"column:cny_price" json:"cnyPrice"`
}
