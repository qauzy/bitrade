package vo

import (
	"github.com/qauzy/math"
	"time"
)

func (this *MemberDepositVO) SetId(id int64) (result *MemberDepositVO) {
	this.Id = id
	return this
}
func (this *MemberDepositVO) GetId() (id int64) {
	return this.Id
}
func (this *MemberDepositVO) SetMemberId(memberId int64) (result *MemberDepositVO) {
	this.MemberId = memberId
	return this
}
func (this *MemberDepositVO) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *MemberDepositVO) SetUsername(username string) (result *MemberDepositVO) {
	this.Username = username
	return this
}
func (this *MemberDepositVO) GetUsername() (username string) {
	return this.Username
}
func (this *MemberDepositVO) SetUnit(unit string) (result *MemberDepositVO) {
	this.Unit = unit
	return this
}
func (this *MemberDepositVO) GetUnit() (unit string) {
	return this.Unit
}
func (this *MemberDepositVO) SetAddress(address string) (result *MemberDepositVO) {
	this.Address = address
	return this
}
func (this *MemberDepositVO) GetAddress() (address string) {
	return this.Address
}
func (this *MemberDepositVO) SetAmount(amount math.BigDecimal) (result *MemberDepositVO) {
	this.Amount = amount
	return this
}
func (this *MemberDepositVO) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *MemberDepositVO) SetCreateTime(createTime time.Time) (result *MemberDepositVO) {
	this.CreateTime = createTime
	return this
}
func (this *MemberDepositVO) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}

type MemberDepositVO struct {
	Id         int64
	MemberId   int64
	Username   string
	Unit       string
	Address    string
	Amount     math.BigDecimal
	CreateTime time.Time
}
