package entity

func (this *BindBank) SetRealName(realName string) (result *BindBank) {
	this.RealName = realName
	return this
}
func (this *BindBank) GetRealName() (realName string) {
	return this.RealName
}
func (this *BindBank) SetBank(bank string) (result *BindBank) {
	this.Bank = bank
	return this
}
func (this *BindBank) GetBank() (bank string) {
	return this.Bank
}
func (this *BindBank) SetBranch(branch string) (result *BindBank) {
	this.Branch = branch
	return this
}
func (this *BindBank) GetBranch() (branch string) {
	return this.Branch
}
func (this *BindBank) SetCardNo(cardNo string) (result *BindBank) {
	this.CardNo = cardNo
	return this
}
func (this *BindBank) GetCardNo() (cardNo string) {
	return this.CardNo
}
func (this *BindBank) SetJyPassword(jyPassword string) (result *BindBank) {
	this.JyPassword = jyPassword
	return this
}
func (this *BindBank) GetJyPassword() (jyPassword string) {
	return this.JyPassword
}

type BindBank struct {
	RealName   string
	Bank       string
	Branch     string
	CardNo     string
	JyPassword string
}
