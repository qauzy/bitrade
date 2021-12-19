package vo

import (
	"bitrade/core/constant"
	"github.com/qauzy/math"
	"time"
)

func (this *WithdrawRecordVO) SetId(id int64) (result *WithdrawRecordVO) {
	this.Id = id
	return this
}
func (this *WithdrawRecordVO) GetId() (id int64) {
	return this.Id
}
func (this *WithdrawRecordVO) SetMemberId(memberId int64) (result *WithdrawRecordVO) {
	this.MemberId = memberId
	return this
}
func (this *WithdrawRecordVO) GetMemberId() (memberId int64) {
	return this.MemberId
}
func (this *WithdrawRecordVO) SetMemberUsername(memberUsername string) (result *WithdrawRecordVO) {
	this.MemberUsername = memberUsername
	return this
}
func (this *WithdrawRecordVO) GetMemberUsername() (memberUsername string) {
	return this.MemberUsername
}
func (this *WithdrawRecordVO) SetMemberRealName(memberRealName string) (result *WithdrawRecordVO) {
	this.MemberRealName = memberRealName
	return this
}
func (this *WithdrawRecordVO) GetMemberRealName() (memberRealName string) {
	return this.MemberRealName
}
func (this *WithdrawRecordVO) SetPhone(phone string) (result *WithdrawRecordVO) {
	this.Phone = phone
	return this
}
func (this *WithdrawRecordVO) GetPhone() (phone string) {
	return this.Phone
}
func (this *WithdrawRecordVO) SetEmail(email string) (result *WithdrawRecordVO) {
	this.Email = email
	return this
}
func (this *WithdrawRecordVO) GetEmail() (email string) {
	return this.Email
}
func (this *WithdrawRecordVO) SetDealTime(dealTime time.Time) (result *WithdrawRecordVO) {
	this.DealTime = dealTime
	return this
}
func (this *WithdrawRecordVO) GetDealTime() (dealTime time.Time) {
	return this.DealTime
}
func (this *WithdrawRecordVO) SetUnit(unit string) (result *WithdrawRecordVO) {
	this.Unit = unit
	return this
}
func (this *WithdrawRecordVO) GetUnit() (unit string) {
	return this.Unit
}
func (this *WithdrawRecordVO) SetTotalAmount(totalAmount math.BigDecimal) (result *WithdrawRecordVO) {
	this.TotalAmount = totalAmount
	return this
}
func (this *WithdrawRecordVO) GetTotalAmount() (totalAmount math.BigDecimal) {
	return this.TotalAmount
}
func (this *WithdrawRecordVO) SetFee(fee math.BigDecimal) (result *WithdrawRecordVO) {
	this.Fee = fee
	return this
}
func (this *WithdrawRecordVO) GetFee() (fee math.BigDecimal) {
	return this.Fee
}
func (this *WithdrawRecordVO) SetArrivedAmount(arrivedAmount math.BigDecimal) (result *WithdrawRecordVO) {
	this.ArrivedAmount = arrivedAmount
	return this
}
func (this *WithdrawRecordVO) GetArrivedAmount() (arrivedAmount math.BigDecimal) {
	return this.ArrivedAmount
}
func (this *WithdrawRecordVO) SetTransactionNumber(transactionNumber string) (result *WithdrawRecordVO) {
	this.TransactionNumber = transactionNumber
	return this
}
func (this *WithdrawRecordVO) GetTransactionNumber() (transactionNumber string) {
	return this.TransactionNumber
}
func (this *WithdrawRecordVO) SetCreateTime(createTime time.Time) (result *WithdrawRecordVO) {
	this.CreateTime = createTime
	return this
}
func (this *WithdrawRecordVO) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *WithdrawRecordVO) SetAddress(address string) (result *WithdrawRecordVO) {
	this.Address = address
	return this
}
func (this *WithdrawRecordVO) GetAddress() (address string) {
	return this.Address
}
func (this *WithdrawRecordVO) SetStatus(status constant.WithdrawStatus) (result *WithdrawRecordVO) {
	this.Status = status
	return this
}
func (this *WithdrawRecordVO) GetStatus() (status constant.WithdrawStatus) {
	return this.Status
}
func (this *WithdrawRecordVO) SetRemark(remark string) (result *WithdrawRecordVO) {
	this.Remark = remark
	return this
}
func (this *WithdrawRecordVO) GetRemark() (remark string) {
	return this.Remark
}
func (this *WithdrawRecordVO) SetIsAuto(isAuto constant.BooleanEnum) (result *WithdrawRecordVO) {
	this.IsAuto = isAuto
	return this
}
func (this *WithdrawRecordVO) GetIsAuto() (isAuto constant.BooleanEnum) {
	return this.IsAuto
}

type WithdrawRecordVO struct {
	Id                int64
	MemberId          int64
	MemberUsername    string
	MemberRealName    string
	Phone             string
	Email             string
	DealTime          time.Time
	Unit              string
	TotalAmount       math.BigDecimal
	Fee               math.BigDecimal
	ArrivedAmount     math.BigDecimal
	TransactionNumber string
	CreateTime        time.Time
	Address           string
	Status            constant.WithdrawStatus
	Remark            string
	IsAuto            constant.BooleanEnum
}
