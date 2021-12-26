package entity

func (this *GiftRecord) SetId(id int64) (result *GiftRecord) {
	this.Id = id
	return this
}
func (this *GiftRecord) GetId() (id int64) {
	return this.Id
}
func (this *GiftRecord) SetUserId(userId int64) (result *GiftRecord) {
	this.UserId = userId
	return this
}
func (this *GiftRecord) GetUserId() (userId int64) {
	return this.UserId
}
func (this *GiftRecord) SetUserName(userName string) (result *GiftRecord) {
	this.UserName = userName
	return this
}
func (this *GiftRecord) GetUserName() (userName string) {
	return this.UserName
}
func (this *GiftRecord) SetUserMobile(userMobile string) (result *GiftRecord) {
	this.UserMobile = userMobile
	return this
}
func (this *GiftRecord) GetUserMobile() (userMobile string) {
	return this.UserMobile
}
func (this *GiftRecord) SetGiftName(giftName string) (result *GiftRecord) {
	this.GiftName = giftName
	return this
}
func (this *GiftRecord) GetGiftName() (giftName string) {
	return this.GiftName
}
func (this *GiftRecord) SetGiftCoin(giftCoin string) (result *GiftRecord) {
	this.GiftCoin = giftCoin
	return this
}
func (this *GiftRecord) GetGiftCoin() (giftCoin string) {
	return this.GiftCoin
}
func (this *GiftRecord) SetGiftAmount(giftAmount math.BigDecimal) (result *GiftRecord) {
	this.GiftAmount = giftAmount
	return this
}
func (this *GiftRecord) GetGiftAmount() (giftAmount math.BigDecimal) {
	return this.GiftAmount
}
func (this *GiftRecord) SetCreateTime(createTime time.Time) (result *GiftRecord) {
	this.CreateTime = createTime
	return this
}
func (this *GiftRecord) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}

type GiftRecord struct {
	Id         int64
	UserId     int64
	UserName   string
	UserMobile string
	GiftName   string
	GiftCoin   string
	GiftAmount math.BigDecimal
	CreateTime time.Time
}
