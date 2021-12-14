package entity

var SerialVersionUID int64 = 1

func (this *GiftConfig) SetId(id int64) {
	this.Id = id
}
func (this *GiftConfig) GetId() (id int64) {
	return this.Id
}
func (this *GiftConfig) SetGiftName(giftName string) {
	this.GiftName = giftName
}
func (this *GiftConfig) GetGiftName() (giftName string) {
	return this.GiftName
}
func (this *GiftConfig) SetGiftCoin(giftCoin string) {
	this.GiftCoin = giftCoin
}
func (this *GiftConfig) GetGiftCoin() (giftCoin string) {
	return this.GiftCoin
}
func (this *GiftConfig) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *GiftConfig) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *GiftConfig) SetHaveCoin(haveCoin string) {
	this.HaveCoin = haveCoin
}
func (this *GiftConfig) GetHaveCoin() (haveCoin string) {
	return this.HaveCoin
}
func (this *GiftConfig) SetHaveAmount(haveAmount BigDecimal) {
	this.HaveAmount = haveAmount
}
func (this *GiftConfig) GetHaveAmount() (haveAmount BigDecimal) {
	return this.HaveAmount
}
func (this *GiftConfig) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *GiftConfig) GetCreateTime() (createTime Date) {
	return this.CreateTime
}

type GiftConfig struct {
	Id         int64
	GiftName   string
	GiftCoin   string
	Amount     BigDecimal
	HaveCoin   string
	HaveAmount BigDecimal
	CreateTime Date
}
