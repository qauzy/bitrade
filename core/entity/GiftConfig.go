package entity

func (this *GiftConfig) SetId(id int64) (result *GiftConfig) {
	this.Id = id
	return this
}
func (this *GiftConfig) GetId() (id int64) {
	return this.Id
}
func (this *GiftConfig) SetGiftName(giftName string) (result *GiftConfig) {
	this.GiftName = giftName
	return this
}
func (this *GiftConfig) GetGiftName() (giftName string) {
	return this.GiftName
}
func (this *GiftConfig) SetGiftCoin(giftCoin string) (result *GiftConfig) {
	this.GiftCoin = giftCoin
	return this
}
func (this *GiftConfig) GetGiftCoin() (giftCoin string) {
	return this.GiftCoin
}
func (this *GiftConfig) SetAmount(amount math.BigDecimal) (result *GiftConfig) {
	this.Amount = amount
	return this
}
func (this *GiftConfig) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *GiftConfig) SetHaveCoin(haveCoin string) (result *GiftConfig) {
	this.HaveCoin = haveCoin
	return this
}
func (this *GiftConfig) GetHaveCoin() (haveCoin string) {
	return this.HaveCoin
}
func (this *GiftConfig) SetHaveAmount(haveAmount math.BigDecimal) (result *GiftConfig) {
	this.HaveAmount = haveAmount
	return this
}
func (this *GiftConfig) GetHaveAmount() (haveAmount math.BigDecimal) {
	return this.HaveAmount
}
func (this *GiftConfig) SetCreateTime(createTime time.Time) (result *GiftConfig) {
	this.CreateTime = createTime
	return this
}
func (this *GiftConfig) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}

type GiftConfig struct {
	Id         int64
	GiftName   string
	GiftCoin   string
	Amount     math.BigDecimal
	HaveCoin   string
	HaveAmount math.BigDecimal
	CreateTime time.Time
}
