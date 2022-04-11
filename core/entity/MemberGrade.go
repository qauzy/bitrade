package entity

import "github.com/qauzy/math"

func (this *MemberGrade) SetId(Id int64) (result *MemberGrade) {
	this.Id = Id
	return this
}
func (this *MemberGrade) GetId() (Id int64) {
	return this.Id
}
func (this *MemberGrade) SetGradeName(GradeName string) (result *MemberGrade) {
	this.GradeName = GradeName
	return this
}
func (this *MemberGrade) GetGradeName() (GradeName string) {
	return this.GradeName
}
func (this *MemberGrade) SetGradeCode(GradeCode string) (result *MemberGrade) {
	this.GradeCode = GradeCode
	return this
}
func (this *MemberGrade) GetGradeCode() (GradeCode string) {
	return this.GradeCode
}
func (this *MemberGrade) SetWithdrawCoinAmount(WithdrawCoinAmount math.BigDecimal) (result *MemberGrade) {
	this.WithdrawCoinAmount = WithdrawCoinAmount
	return this
}
func (this *MemberGrade) GetWithdrawCoinAmount() (WithdrawCoinAmount math.BigDecimal) {
	return this.WithdrawCoinAmount
}
func (this *MemberGrade) SetDayWithdrawCount(DayWithdrawCount int) (result *MemberGrade) {
	this.DayWithdrawCount = DayWithdrawCount
	return this
}
func (this *MemberGrade) GetDayWithdrawCount() (DayWithdrawCount int) {
	return this.DayWithdrawCount
}
func (this *MemberGrade) SetExchangeFeeRate(ExchangeFeeRate math.BigDecimal) (result *MemberGrade) {
	this.ExchangeFeeRate = ExchangeFeeRate
	return this
}
func (this *MemberGrade) GetExchangeFeeRate() (ExchangeFeeRate math.BigDecimal) {
	return this.ExchangeFeeRate
}
func (this *MemberGrade) SetGradeBound(GradeBound int) (result *MemberGrade) {
	this.GradeBound = GradeBound
	return this
}
func (this *MemberGrade) GetGradeBound() (GradeBound int) {
	return this.GradeBound
}
func NewMemberGrade(id int64, gradeName string, gradeCode string, withdrawCoinAmount math.BigDecimal, dayWithdrawCount int, exchangeFeeRate math.BigDecimal, gradeBound int) (ret *MemberGrade) {
	ret = new(MemberGrade)
	ret.Id = id
	ret.GradeName = gradeName
	ret.GradeCode = gradeCode
	ret.WithdrawCoinAmount = withdrawCoinAmount
	ret.DayWithdrawCount = dayWithdrawCount
	ret.ExchangeFeeRate = exchangeFeeRate
	ret.GradeBound = gradeBound
	return
}

type MemberGrade struct {
	Id                 int64           `gorm:"column:id" json:"id"`
	GradeName          string          `gorm:"column:grade_name" json:"gradeName"`
	GradeCode          string          `gorm:"column:grade_code" json:"gradeCode"`
	WithdrawCoinAmount math.BigDecimal `gorm:"column:withdraw_coin_amount" json:"withdrawCoinAmount"`
	DayWithdrawCount   int             `gorm:"column:day_withdraw_count" json:"dayWithdrawCount"`
	ExchangeFeeRate    math.BigDecimal `gorm:"column:exchange_fee_rate" json:"exchangeFeeRate"`
	GradeBound         int             `gorm:"column:grade_bound" json:"gradeBound"`
}
