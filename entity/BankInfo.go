package entity

var SerialVersionUID int64 = -5602439827000928628

func (this *BankInfo) SetBank(bank string) {
	this.Bank = bank
}
func (this *BankInfo) GetBank() (bank string) {
	return this.Bank
}
func (this *BankInfo) SetBranch(branch string) {
	this.Branch = branch
}
func (this *BankInfo) GetBranch() (branch string) {
	return this.Branch
}
func (this *BankInfo) SetCardNo(cardNo string) {
	this.CardNo = cardNo
}
func (this *BankInfo) GetCardNo() (cardNo string) {
	return this.CardNo
}

type BankInfo struct {
	Bank   string
	Branch string
	CardNo string
}
