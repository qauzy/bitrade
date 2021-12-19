package entity

func (this *MemberGrade) SetId(id int64) (result *MemberGrade) {
	this.Id = id
	return this
}
func (this *MemberGrade) GetId() (id int64) {
	return this.Id
}
func (this *MemberGrade) SetGradeName(gradeName string) (result *MemberGrade) {
	this.GradeName = gradeName
	return this
}
func (this *MemberGrade) GetGradeName() (gradeName string) {
	return this.GradeName
}
func (this *MemberGrade) SetGradeCode(gradeCode string) (result *MemberGrade) {
	this.GradeCode = gradeCode
	return this
}
func (this *MemberGrade) GetGradeCode() (gradeCode string) {
	return this.GradeCode
}
func (this *MemberGrade) SetWithdrawCoinAmount(withdrawCoinAmount math.BigDecimal) (result *MemberGrade) {
	this.WithdrawCoinAmount = withdrawCoinAmount
	return this
}
func (this *MemberGrade) GetWithdrawCoinAmount() (withdrawCoinAmount math.BigDecimal) {
	return this.WithdrawCoinAmount
}
func (this *MemberGrade) SetDayWithdrawCount(dayWithdrawCount int64) (result *MemberGrade) {
	this.DayWithdrawCount = dayWithdrawCount
	return this
}
func (this *MemberGrade) GetDayWithdrawCount() (dayWithdrawCount int64) {
	return this.DayWithdrawCount
}
func (this *MemberGrade) SetExchangeFeeRate(exchangeFeeRate math.BigDecimal) (result *MemberGrade) {
	this.ExchangeFeeRate = exchangeFeeRate
	return this
}
func (this *MemberGrade) GetExchangeFeeRate() (exchangeFeeRate math.BigDecimal) {
	return this.ExchangeFeeRate
}
func (this *MemberGrade) SetGradeBound(gradeBound int64) (result *MemberGrade) {
	this.GradeBound = gradeBound
	return this
}
func (this *MemberGrade) GetGradeBound() (gradeBound int64) {
	return this.GradeBound
}

type MemberGrade struct {
	Id                 int64
	GradeName          string
	GradeCode          string
	WithdrawCoinAmount math.BigDecimal
	DayWithdrawCount   int64
	ExchangeFeeRate    math.BigDecimal
	GradeBound         int64
}
