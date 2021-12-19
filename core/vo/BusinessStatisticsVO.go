package vo

import "github.com/qauzy/math"

func (this *BusinessStatisticsVO) SetAdvertisementNum(advertisementNum int64) (result *BusinessStatisticsVO) {
	this.AdvertisementNum = advertisementNum
	return this
}
func (this *BusinessStatisticsVO) GetAdvertisementNum() (advertisementNum int64) {
	return this.AdvertisementNum
}
func (this *BusinessStatisticsVO) SetMoney(money math.BigDecimal) (result *BusinessStatisticsVO) {
	this.Money = money
	return this
}
func (this *BusinessStatisticsVO) GetMoney() (money math.BigDecimal) {
	return this.Money
}
func (this *BusinessStatisticsVO) SetAmount(amount math.BigDecimal) (result *BusinessStatisticsVO) {
	this.Amount = amount
	return this
}
func (this *BusinessStatisticsVO) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *BusinessStatisticsVO) SetComplainantNum(complainantNum int64) (result *BusinessStatisticsVO) {
	this.ComplainantNum = complainantNum
	return this
}
func (this *BusinessStatisticsVO) GetComplainantNum() (complainantNum int64) {
	return this.ComplainantNum
}
func (this *BusinessStatisticsVO) SetDefendantNum(defendantNum int64) (result *BusinessStatisticsVO) {
	this.DefendantNum = defendantNum
	return this
}
func (this *BusinessStatisticsVO) GetDefendantNum() (defendantNum int64) {
	return this.DefendantNum
}

type BusinessStatisticsVO struct {
	AdvertisementNum int64
	Money            math.BigDecimal
	Amount           math.BigDecimal
	ComplainantNum   int64
	DefendantNum     int64
}
