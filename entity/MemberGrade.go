package entity

func (this *MemberGrade) SetId(id int64) {
	this.Id = id
}
func (this *MemberGrade) GetId() (id int64) {
	return this.Id
}
func (this *MemberGrade) SetGradeName(gradeName string) {
	this.GradeName = gradeName
}
func (this *MemberGrade) GetGradeName() (gradeName string) {
	return this.GradeName
}
func (this *MemberGrade) SetGradeCode(gradeCode string) {
	this.GradeCode = gradeCode
}
func (this *MemberGrade) GetGradeCode() (gradeCode string) {
	return this.GradeCode
}
func (this *MemberGrade) SetWithdrawCoinAmount(withdrawCoinAmount BigDecimal) {
	this.WithdrawCoinAmount = withdrawCoinAmount
}
func (this *MemberGrade) GetWithdrawCoinAmount() (withdrawCoinAmount BigDecimal) {
	return this.WithdrawCoinAmount
}
func (this *MemberGrade) SetDayWithdrawCount(dayWithdrawCount int64) {
	this.DayWithdrawCount = dayWithdrawCount
}
func (this *MemberGrade) GetDayWithdrawCount() (dayWithdrawCount int64) {
	return this.DayWithdrawCount
}
func (this *MemberGrade) SetExchangeFeeRate(exchangeFeeRate BigDecimal) {
	this.ExchangeFeeRate = exchangeFeeRate
}
func (this *MemberGrade) GetExchangeFeeRate() (exchangeFeeRate BigDecimal) {
	return this.ExchangeFeeRate
}
func (this *MemberGrade) SetGradeBound(gradeBound int64) {
	this.GradeBound = gradeBound
}
func (this *MemberGrade) GetGradeBound() (gradeBound int64) {
	return this.GradeBound
}

type MemberGrade struct {
	Id                 int64
	GradeName          string
	GradeCode          string
	WithdrawCoinAmount BigDecimal
	DayWithdrawCount   int64
	ExchangeFeeRate    BigDecimal
	GradeBound         int64
}
