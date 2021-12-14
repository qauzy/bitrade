package transform

func (this *MemberAdvertise) SetId(id int64) {
	this.Id = id
}
func (this *MemberAdvertise) GetId() (id int64) {
	return this.Id
}
func (this *MemberAdvertise) SetAdvertiseType(advertiseType AdvertiseType) {
	this.AdvertiseType = advertiseType
}
func (this *MemberAdvertise) GetAdvertiseType() (advertiseType AdvertiseType) {
	return this.AdvertiseType
}
func (this *MemberAdvertise) SetMinLimit(minLimit BigDecimal) {
	this.MinLimit = minLimit
}
func (this *MemberAdvertise) GetMinLimit() (minLimit BigDecimal) {
	return this.MinLimit
}
func (this *MemberAdvertise) SetMaxLimit(maxLimit BigDecimal) {
	this.MaxLimit = maxLimit
}
func (this *MemberAdvertise) GetMaxLimit() (maxLimit BigDecimal) {
	return this.MaxLimit
}
func (this *MemberAdvertise) SetStatus(status AdvertiseControlStatus) {
	this.Status = status
}
func (this *MemberAdvertise) GetStatus() (status AdvertiseControlStatus) {
	return this.Status
}
func (this *MemberAdvertise) SetRemainAmount(remainAmount BigDecimal) {
	this.RemainAmount = remainAmount
}
func (this *MemberAdvertise) GetRemainAmount() (remainAmount BigDecimal) {
	return this.RemainAmount
}
func (this *MemberAdvertise) SetCoinUnit(coinUnit string) {
	this.CoinUnit = coinUnit
}
func (this *MemberAdvertise) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *MemberAdvertise) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *MemberAdvertise) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *MemberAdvertise) SetCountry(country Country) {
	this.Country = country
}
func (this *MemberAdvertise) GetCountry() (country Country) {
	return this.Country
}
func ToMemberAdvertise(x Advertise) (result MemberAdvertise) {
	return MemberAdvertise.Builder().Id(x.GetId()).AdvertiseType(x.GetAdvertiseType()).CoinUnit(x.GetCoin().GetUnit()).CreateTime(x.GetCreateTime()).MinLimit(x.GetMinLimit()).MaxLimit(x.GetMaxLimit()).RemainAmount(x.GetRemainAmount()).Status(x.GetStatus()).Country(x.GetCountry()).Build()
}

type MemberAdvertise struct {
	Id            int64
	AdvertiseType AdvertiseType
	MinLimit      BigDecimal
	MaxLimit      BigDecimal
	Status        AdvertiseControlStatus
	RemainAmount  BigDecimal
	CoinUnit      string
	CreateTime    Date
	Country       Country
}
