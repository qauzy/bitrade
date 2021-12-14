package transform

func (this *ScanAdvertise) SetMemberName(memberName string) {
	this.MemberName = memberName
}
func (this *ScanAdvertise) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *ScanAdvertise) SetAvatar(avatar string) {
	this.Avatar = avatar
}
func (this *ScanAdvertise) GetAvatar() (avatar string) {
	return this.Avatar
}
func (this *ScanAdvertise) SetAdvertiseId(advertiseId int64) {
	this.AdvertiseId = advertiseId
}
func (this *ScanAdvertise) GetAdvertiseId() (advertiseId int64) {
	return this.AdvertiseId
}
func (this *ScanAdvertise) SetTransactions(transactions int) {
	this.Transactions = transactions
}
func (this *ScanAdvertise) GetTransactions() (transactions int) {
	return this.Transactions
}
func (this *ScanAdvertise) SetPrice(price BigDecimal) {
	this.Price = price
}
func (this *ScanAdvertise) GetPrice() (price BigDecimal) {
	return this.Price
}
func (this *ScanAdvertise) SetMinLimit(minLimit BigDecimal) {
	this.MinLimit = minLimit
}
func (this *ScanAdvertise) GetMinLimit() (minLimit BigDecimal) {
	return this.MinLimit
}
func (this *ScanAdvertise) SetMaxLimit(maxLimit BigDecimal) {
	this.MaxLimit = maxLimit
}
func (this *ScanAdvertise) GetMaxLimit() (maxLimit BigDecimal) {
	return this.MaxLimit
}
func (this *ScanAdvertise) SetRemainAmount(remainAmount BigDecimal) {
	this.RemainAmount = remainAmount
}
func (this *ScanAdvertise) GetRemainAmount() (remainAmount BigDecimal) {
	return this.RemainAmount
}
func (this *ScanAdvertise) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *ScanAdvertise) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *ScanAdvertise) SetPayMode(payMode string) {
	this.PayMode = payMode
}
func (this *ScanAdvertise) GetPayMode() (payMode string) {
	return this.PayMode
}
func (this *ScanAdvertise) SetCoinId(coinId int64) {
	this.CoinId = coinId
}
func (this *ScanAdvertise) GetCoinId() (coinId int64) {
	return this.CoinId
}
func (this *ScanAdvertise) SetUnit(unit string) {
	this.Unit = unit
}
func (this *ScanAdvertise) GetUnit() (unit string) {
	return this.Unit
}
func (this *ScanAdvertise) SetCoinName(coinName string) {
	this.CoinName = coinName
}
func (this *ScanAdvertise) GetCoinName() (coinName string) {
	return this.CoinName
}
func (this *ScanAdvertise) SetCoinNameCn(coinNameCn string) {
	this.CoinNameCn = coinNameCn
}
func (this *ScanAdvertise) GetCoinNameCn() (coinNameCn string) {
	return this.CoinNameCn
}
func (this *ScanAdvertise) SetLevel(level int) {
	this.Level = level
}
func (this *ScanAdvertise) GetLevel() (level int) {
	return this.Level
}
func (this *ScanAdvertise) SetAdvertiseType(advertiseType AdvertiseType) {
	this.AdvertiseType = advertiseType
}
func (this *ScanAdvertise) GetAdvertiseType() (advertiseType AdvertiseType) {
	return this.AdvertiseType
}
func (this *ScanAdvertise) SetAdvType(advType int) {
	this.AdvType = advType
}
func (this *ScanAdvertise) GetAdvType() (advType int) {
	return this.AdvType
}
func (this *ScanAdvertise) SetPremiseRate(premiseRate BigDecimal) {
	this.PremiseRate = premiseRate
}
func (this *ScanAdvertise) GetPremiseRate() (premiseRate BigDecimal) {
	return this.PremiseRate
}

type ScanAdvertise struct {
	MemberName    string
	Avatar        string
	AdvertiseId   int64
	Transactions  int
	Price         BigDecimal
	MinLimit      BigDecimal
	MaxLimit      BigDecimal
	RemainAmount  BigDecimal
	CreateTime    Date
	PayMode       string
	CoinId        int64
	Unit          string
	CoinName      string
	CoinNameCn    string
	Level         int
	AdvertiseType AdvertiseType
	AdvType       int
	PremiseRate   BigDecimal
}
