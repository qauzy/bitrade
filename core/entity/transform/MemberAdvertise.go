package transform

func (this *MemberAdvertise) SetId(id int64) (result *MemberAdvertise) {
	this.Id = id
	return this
}
func (this *MemberAdvertise) GetId() (id int64) {
	return this.Id
}
func (this *MemberAdvertise) SetAdvertiseType(advertiseType AdvertiseType.AdvertiseType) (result *MemberAdvertise) {
	this.AdvertiseType = advertiseType
	return this
}
func (this *MemberAdvertise) GetAdvertiseType() (advertiseType AdvertiseType.AdvertiseType) {
	return this.AdvertiseType
}
func (this *MemberAdvertise) SetMinLimit(minLimit math.BigDecimal) (result *MemberAdvertise) {
	this.MinLimit = minLimit
	return this
}
func (this *MemberAdvertise) GetMinLimit() (minLimit math.BigDecimal) {
	return this.MinLimit
}
func (this *MemberAdvertise) SetMaxLimit(maxLimit math.BigDecimal) (result *MemberAdvertise) {
	this.MaxLimit = maxLimit
	return this
}
func (this *MemberAdvertise) GetMaxLimit() (maxLimit math.BigDecimal) {
	return this.MaxLimit
}
func (this *MemberAdvertise) SetStatus(status AdvertiseControlStatus.AdvertiseControlStatus) (result *MemberAdvertise) {
	this.Status = status
	return this
}
func (this *MemberAdvertise) GetStatus() (status AdvertiseControlStatus.AdvertiseControlStatus) {
	return this.Status
}
func (this *MemberAdvertise) SetRemainAmount(remainAmount math.BigDecimal) (result *MemberAdvertise) {
	this.RemainAmount = remainAmount
	return this
}
func (this *MemberAdvertise) GetRemainAmount() (remainAmount math.BigDecimal) {
	return this.RemainAmount
}
func (this *MemberAdvertise) SetCoinUnit(coinUnit string) (result *MemberAdvertise) {
	this.CoinUnit = coinUnit
	return this
}
func (this *MemberAdvertise) GetCoinUnit() (coinUnit string) {
	return this.CoinUnit
}
func (this *MemberAdvertise) SetCreateTime(createTime time.Time) (result *MemberAdvertise) {
	this.CreateTime = createTime
	return this
}
func (this *MemberAdvertise) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *MemberAdvertise) SetCountry(country *entity.Country) (result *MemberAdvertise) {
	this.Country = country
	return this
}
func (this *MemberAdvertise) GetCountry() (country *entity.Country) {
	return this.Country
}
func ToMemberAdvertise(x *entity.Advertise) (result *MemberAdvertise) {
	return new(MemberAdvertise).SetId(x.GetId()).SetAdvertiseType(x.GetAdvertiseType()).SetCoinUnit(x.GetCoin().GetUnit()).SetCreateTime(x.GetCreateTime()).SetMinLimit(x.GetMinLimit()).SetMaxLimit(x.GetMaxLimit()).SetRemainAmount(x.GetRemainAmount()).SetStatus(x.GetStatus()).SetCountry(x.GetCountry())
}

type MemberAdvertise struct {
	Id            int64
	AdvertiseType AdvertiseType.AdvertiseType
	MinLimit      math.BigDecimal
	MaxLimit      math.BigDecimal
	Status        AdvertiseControlStatus.AdvertiseControlStatus
	RemainAmount  math.BigDecimal
	CoinUnit      string
	CreateTime    time.Time
	Country       *entity.Country
}
