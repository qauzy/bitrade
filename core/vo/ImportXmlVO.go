package vo

func (this *ImportXmlVO) SetMemberId(memberId int64) (result *ImportXmlVO) {
	this.MemberId = memberId
	return this
}
func (this *ImportXmlVO) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *ImportXmlVO) SetMemberName(memberName string) (result *ImportXmlVO) {
	this.MemberName = memberName
	return this
}
func (this *ImportXmlVO) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *ImportXmlVO) SetPhone(phone string) (result *ImportXmlVO) {
	this.Phone = phone
	return this
}
func (this *ImportXmlVO) GetPhone() (phone string) {
	return this.Phone
}
func (this *ImportXmlVO) SetCoinName(coinName string) (result *ImportXmlVO) {
	this.CoinName = coinName
	return this
}
func (this *ImportXmlVO) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *ImportXmlVO) SetCoinUnit(coinUnit string) (result *ImportXmlVO) {
	this.CoinUnit = coinUnit
	return this
}
func (this *ImportXmlVO) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *ImportXmlVO) SetAmount(amount float64) (result *ImportXmlVO) {
	this.Amount = amount
	return this
}
func (this *ImportXmlVO) GetAmount() (amount float64) {
	return this.Amount
}

type ImportXmlVO struct {
	MemberId   int64
	MemberName string
	Phone      string
	CoinName   string
	CoinUnit   string
	Amount     float64
}
