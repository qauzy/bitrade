package entity

func (this *BankInfo) SetBank(Bank string) (result *BankInfo) {
	this.Bank = Bank
	return this
}
func (this *BankInfo) GetBank() (Bank string) {
	return this.Bank
}
func (this *BankInfo) SetBranch(Branch string) (result *BankInfo) {
	this.Branch = Branch
	return this
}
func (this *BankInfo) GetBranch() (Branch string) {
	return this.Branch
}
func (this *BankInfo) SetCardNo(CardNo string) (result *BankInfo) {
	this.CardNo = CardNo
	return this
}
func (this *BankInfo) GetCardNo() (CardNo string) {
	return this.CardNo
}
func NewBankInfo(bank string, branch string, cardNo string) (ret *BankInfo) {
	ret = new(BankInfo)
	ret.Bank = bank
	ret.Branch = branch
	ret.CardNo = cardNo
	return
}

type BankInfo struct {
	Bank   string `gorm:"column:bank" json:"bank"`
	Branch string `gorm:"column:branch" json:"branch"`
	CardNo string `gorm:"column:card_no" json:"cardNo"`
}
