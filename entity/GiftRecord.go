package entity

var SerialVersionUID int64 = 1

func (this *GiftRecord) SetId(id int64) {
	this.Id = id
}
func (this *GiftRecord) GetId() (id int64) {
	return this.Id
}
func (this *GiftRecord) SetUserId(userId int64) {
	this.UserId = userId
}
func (this *GiftRecord) GetUserId() (userId int64) {
	return this.UserId
}
func (this *GiftRecord) SetUserName(userName string) {
	this.UserName = userName
}
func (this *GiftRecord) GetUserName() (userName string) {
	return this.UserName
}
func (this *GiftRecord) SetUserMobile(userMobile string) {
	this.UserMobile = userMobile
}
func (this *GiftRecord) GetUserMobile() (userMobile string) {
	return this.UserMobile
}
func (this *GiftRecord) SetGiftName(giftName string) {
	this.GiftName = giftName
}
func (this *GiftRecord) GetGiftName() (giftName string) {
	return this.GiftName
}
func (this *GiftRecord) SetGiftCoin(giftCoin string) {
	this.GiftCoin = giftCoin
}
func (this *GiftRecord) GetGiftCoin() (giftCoin string) {
	return this.GiftCoin
}
func (this *GiftRecord) SetGiftAmount(giftAmount BigDecimal) {
	this.GiftAmount = giftAmount
}
func (this *GiftRecord) GetGiftAmount() (giftAmount BigDecimal) {
	return this.GiftAmount
}
func (this *GiftRecord) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *GiftRecord) GetCreateTime() (createTime Date) {
	return this.CreateTime
}

type GiftRecord struct {
	Id         int64
	UserId     int64
	UserName   string
	UserMobile string
	GiftName   string
	GiftCoin   string
	GiftAmount BigDecimal
	CreateTime Date
}
