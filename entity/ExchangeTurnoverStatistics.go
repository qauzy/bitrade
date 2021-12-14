package entity

func (this *ExchangeTurnoverStatistics) SetDate(date Date) {
	this.Date = date
}
func (this *ExchangeTurnoverStatistics) GetDate() (date Date) {
	return this.Date
}
func (this *ExchangeTurnoverStatistics) SetYear(year int) {
	this.Year = year
}
func (this *ExchangeTurnoverStatistics) GetYear() (year int) {
	return this.Year
}
func (this *ExchangeTurnoverStatistics) SetMonth(month int) {
	this.Month = month
}
func (this *ExchangeTurnoverStatistics) GetMonth() (month int) {
	return this.Month
}
func (this *ExchangeTurnoverStatistics) SetDay(day int) {
	this.Day = day
}
func (this *ExchangeTurnoverStatistics) GetDay() (day int) {
	return this.Day
}
func (this *ExchangeTurnoverStatistics) SetBaseSymbol(baseSymbol string) {
	this.BaseSymbol = baseSymbol
}
func (this *ExchangeTurnoverStatistics) GetBaseSymbol() (baseSymbol string) {
	return this.BaseSymbol
}
func (this *ExchangeTurnoverStatistics) SetCoinSymbol(coinSymbol string) {
	this.CoinSymbol = coinSymbol
}
func (this *ExchangeTurnoverStatistics) GetCoinSymbol() (coinSymbol string) {
	return this.CoinSymbol
}
func (this *ExchangeTurnoverStatistics) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *ExchangeTurnoverStatistics) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *ExchangeTurnoverStatistics) SetMoney(money BigDecimal) {
	this.Money = money
}
func (this *ExchangeTurnoverStatistics) GetMoney() (money BigDecimal) {
	return this.Money
}

type ExchangeTurnoverStatistics struct {
	Date       Date
	Year       int
	Month      int
	Day        int
	BaseSymbol string
	CoinSymbol string
	Amount     BigDecimal
	Money      BigDecimal
}
