package entity

func (this *ExchangeTurnoverStatistics) SetDate(date time.Time) (result *ExchangeTurnoverStatistics) {
	this.Date = date
	return this
}
func (this *ExchangeTurnoverStatistics) GetDate() (date time.Time) {
	return this.Date
}
func (this *ExchangeTurnoverStatistics) SetYear(year int) (result *ExchangeTurnoverStatistics) {
	this.Year = year
	return this
}
func (this *ExchangeTurnoverStatistics) GetYear() (year int) {
	return this.Year
}
func (this *ExchangeTurnoverStatistics) SetMonth(month int) (result *ExchangeTurnoverStatistics) {
	this.Month = month
	return this
}
func (this *ExchangeTurnoverStatistics) GetMonth() (month int) {
	return this.Month
}
func (this *ExchangeTurnoverStatistics) SetDay(day int) (result *ExchangeTurnoverStatistics) {
	this.Day = day
	return this
}
func (this *ExchangeTurnoverStatistics) GetDay() (day int) {
	return this.Day
}
func (this *ExchangeTurnoverStatistics) SetBaseSymbol(baseSymbol string) (result *ExchangeTurnoverStatistics) {
	this.BaseSymbol = baseSymbol
	return this
}
func (this *ExchangeTurnoverStatistics) GetBaseSymbol() (baseSymbol string) {
	return this.BaseSymbol
}
func (this *ExchangeTurnoverStatistics) SetCoinSymbol(coinSymbol string) (result *ExchangeTurnoverStatistics) {
	this.CoinSymbol = coinSymbol
	return this
}
func (this *ExchangeTurnoverStatistics) GetCoinSymbol() (coinSymbol string) {
	return this.CoinSymbol
}
func (this *ExchangeTurnoverStatistics) SetAmount(amount math.BigDecimal) (result *ExchangeTurnoverStatistics) {
	this.Amount = amount
	return this
}
func (this *ExchangeTurnoverStatistics) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *ExchangeTurnoverStatistics) SetMoney(money math.BigDecimal) (result *ExchangeTurnoverStatistics) {
	this.Money = money
	return this
}
func (this *ExchangeTurnoverStatistics) GetMoney() (money math.BigDecimal) {
	return this.Money
}

type ExchangeTurnoverStatistics struct {
	Date       time.Time
	Year       int
	Month      int
	Day        int
	BaseSymbol string
	CoinSymbol string
	Amount     math.BigDecimal
	Money      math.BigDecimal
}
