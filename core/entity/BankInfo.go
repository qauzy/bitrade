package entity

func (this *BankInfo) SetBank(bank string) (result *BankInfo) {
	this.Bank = bank
	return this
}
func (this *BankInfo) GetBank() (bank string) {
	return this.Bank
}
func (this *BankInfo) SetBranch(branch string) (result *BankInfo) {
	this.Branch = branch
	return this
}
func (this *BankInfo) GetBranch() (branch string) {
	return this.Branch
}
func (this *BankInfo) SetCardNo(cardNo string) (result *BankInfo) {
	this.CardNo = cardNo
	return this
}
func (this *BankInfo) GetCardNo() (cardNo string) {
	return this.CardNo
}

type BankInfo struct {
	Bank   string
	Branch string
	CardNo string
}
