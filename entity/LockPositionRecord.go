package entity

func (this *LockPositionRecord) SetId(id int64) {
	this.Id = id
}
func (this *LockPositionRecord) GetId() (id int64) {
	return this.Id
}
func (this *LockPositionRecord) SetMemberId(memberId int64) {
	this.MemberId = memberId
}
func (this *LockPositionRecord) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *LockPositionRecord) SetMemberName(memberName string) {
	this.MemberName = memberName
}
func (this *LockPositionRecord) GetMemberName() (memberName string) {
	return this.MemberName
}
func (this *LockPositionRecord) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *LockPositionRecord) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *LockPositionRecord) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *LockPositionRecord) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *LockPositionRecord) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *LockPositionRecord) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *LockPositionRecord) SetUnlockTime(unlockTime Date) {
	this.UnlockTime = unlockTime
}
func (this *LockPositionRecord) GetUnlockTime() (unlockTime Date) {
	return this.UnlockTime
}
func (this *LockPositionRecord) SetReason(reason string) {
	this.Reason = reason
}
func (this *LockPositionRecord) GetReason() (reason string) {
	return this.Reason
}
func (this *LockPositionRecord) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *LockPositionRecord) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *LockPositionRecord) SetWalletId(walletId int64) {
	this.WalletId = walletId
}
func (this *LockPositionRecord) GetWalletId() (walletId int64) {
	return this.WalletId
}

type LockPositionRecord struct {
	Id         int64
	MemberId   int64
	MemberName string
	Coin       Coin
	CreateTime Date
	Status     CommonStatus
	UnlockTime Date
	Reason     string
	Amount     BigDecimal
	WalletId   int64
}
