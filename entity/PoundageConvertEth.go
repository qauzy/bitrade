
package entity

func (this *PoundageConvertEth) SetId(id int64) {
	this.Id = id
}
func (this *PoundageConvertEth) GetId() (id int64) {
	return this.Id
}
func (this *PoundageConvertEth) SetExchangeOrderId(exchangeOrderId string) {
	this.ExchangeOrderId = exchangeOrderId
}
func (this *PoundageConvertEth) GetExchangeOrderId() (exchangeOrderId string) {
	return this.ExchangeOrderId
}
func (this *PoundageConvertEth) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *PoundageConvertEth) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *PoundageConvertEth) SetMineAmount(mineAmount BigDecimal) {
	this.MineAmount = mineAmount
}
func (this *PoundageConvertEth) GetMineAmount() (mineAmount BigDecimal) {
	return this.MineAmount
}
func (this *PoundageConvertEth) SetPoundageAmount(poundageAmount BigDecimal) {
	this.PoundageAmount = poundageAmount
}
func (this *PoundageConvertEth) GetPoundageAmount() (poundageAmount BigDecimal) {
	return this.PoundageAmount
}
func (this *PoundageConvertEth) SetPoundageAmountEth(poundageAmountEth BigDecimal) {
	this.PoundageAmountEth = poundageAmountEth
}
func (this *PoundageConvertEth) GetPoundageAmountEth() (poundageAmountEth BigDecimal) {
	return this.PoundageAmountEth
}
func (this *PoundageConvertEth) SetCoinId(coinId string) {
	this.CoinId = coinId
}
func (this *PoundageConvertEth) GetCoinId() (coinId string) {
	return this.CoinId
}
func (this *PoundageConvertEth) SetTransactionTime(transactionTime Date) {
	this.TransactionTime = transactionTime
}
func (this *PoundageConvertEth) GetTransactionTime() (transactionTime Date) {
	return this.TransactionTime
}
func (this *PoundageConvertEth) SetSymbol(symbol string) {
	this.Symbol = symbol
}
func (this *PoundageConvertEth) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *PoundageConvertEth) SetType(type string) {
	this.Type = type
}
func (this *PoundageConvertEth) GetType() (type string) {
	return this.Type
}
func (this *PoundageConvertEth) SetDirection(direction string) {
	this.Direction = direction
}
func (this *PoundageConvertEth) GetDirection() (direction string) {
	return this.Direction
}
func (this *PoundageConvertEth) SetUsdtRate(usdtRate string) {
	this.UsdtRate = usdtRate
}
func (this *PoundageConvertEth) GetUsdtRate() (usdtRate string) {
	return this.UsdtRate
}
func (this *PoundageConvertEth) SetEthUsdtRate(ethUsdtRate string) {
	this.EthUsdtRate = ethUsdtRate
}
func (this *PoundageConvertEth) GetEthUsdtRate() (ethUsdtRate string) {
	return this.EthUsdtRate
}

type PoundageConvertEth struct {
	Id                int64
	ExchangeOrderId   string
	MemberId          int64
	MineAmount        BigDecimal
	PoundageAmount    BigDecimal
	PoundageAmountEth BigDecimal
	CoinId            string
	TransactionTime   Date
	Symbol            string
	Type              string
	Direction         string
	UsdtRate          string
	EthUsdtRate       string
}

