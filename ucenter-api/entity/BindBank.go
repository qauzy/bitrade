package entity

func (this *BindBank) SetRealName(RealName string) (result *BindBank) {
	this.RealName = RealName
	return this
}
func (this *BindBank) GetRealName() (RealName string) {
	return this.RealName
}
func (this *BindBank) SetBank(Bank string) (result *BindBank) {
	this.Bank = Bank
	return this
}
func (this *BindBank) GetBank() (Bank string) {
	return this.Bank
}
func (this *BindBank) SetBranch(Branch string) (result *BindBank) {
	this.Branch = Branch
	return this
}
func (this *BindBank) GetBranch() (Branch string) {
	return this.Branch
}
func (this *BindBank) SetCardNo(CardNo string) (result *BindBank) {
	this.CardNo = CardNo
	return this
}
func (this *BindBank) GetCardNo() (CardNo string) {
	return this.CardNo
}
func (this *BindBank) SetJyPassword(JyPassword string) (result *BindBank) {
	this.JyPassword = JyPassword
	return this
}
func (this *BindBank) GetJyPassword() (JyPassword string) {
	return this.JyPassword
}

type BindBank struct {
	RealName   string `gorm:"column:real_name" json:"realName"`
	Bank       string `gorm:"column:bank" json:"bank"`
	Branch     string `gorm:"column:branch" json:"branch"`
	CardNo     string `gorm:"column:card_no" json:"cardNo"`
	JyPassword string `gorm:"column:jy_password" json:"jyPassword"`
}
