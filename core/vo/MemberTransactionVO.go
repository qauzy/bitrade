package vo

func (this *MemberTransactionVO) SetMemberUsername(memberUsername string) (result *MemberTransactionVO) {
	this.MemberUsername = memberUsername
	return this
}
func (this *MemberTransactionVO) GetMemberUsername() (memberUsername string) {
	return this.MemberUsername
}
func (this *MemberTransactionVO) SetMemberRealName(memberRealName string) (result *MemberTransactionVO) {
	this.MemberRealName = memberRealName
	return this
}
func (this *MemberTransactionVO) GetMemberRealName() (memberRealName string) {
	return this.MemberRealName
}
func (this *MemberTransactionVO) SetPhone(phone string) (result *MemberTransactionVO) {
	this.Phone = phone
	return this
}
func (this *MemberTransactionVO) GetPhone() (phone string) {
	return this.Phone
}
func (this *MemberTransactionVO) SetEmail(email string) (result *MemberTransactionVO) {
	this.Email = email
	return this
}
func (this *MemberTransactionVO) GetEmail() (email string) {
	return this.Email
}

type MemberTransactionVO struct {
	MemberUsername string
	MemberRealName string
	Phone          string
	Email          string
}
