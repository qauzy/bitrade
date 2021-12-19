package vo

import "github.com/qauzy/math"

func (this *OtcWalletVO) SetCoinName(coinName string) (result *OtcWalletVO) {
	this.CoinName = coinName
	return this
}
func (this *OtcWalletVO) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *OtcWalletVO) SetAmount(amount math.BigDecimal) (result *OtcWalletVO) {
	this.Amount = amount
	return this
}
func (this *OtcWalletVO) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *OtcWalletVO) SetDirection(direction string) (result *OtcWalletVO) {
	this.Direction = direction
	return this
}
func (this *OtcWalletVO) GetDirection() (direction string) {
	return this.Direction
}

type OtcWalletVO struct {
	CoinName  string
	Amount    math.BigDecimal
	Direction string
}
