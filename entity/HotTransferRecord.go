package entity

func (this *HotTransferRecord) SetId(id int64) {
	this.Id = id
}
func (this *HotTransferRecord) GetId() (id int64) {
	return this.Id
}
func (this *HotTransferRecord) SetAdminId(adminId int64) {
	this.AdminId = adminId
}
func (this *HotTransferRecord) GetAdminId() (adminId int64) {
	return this.AdminId
}
func (this *HotTransferRecord) SetTransferTime(transferTime Date) {
	this.TransferTime = transferTime
}
func (this *HotTransferRecord) GetTransferTime() (transferTime Date) {
	return this.TransferTime
}
func (this *HotTransferRecord) SetAmount(amount BigDecimal) {
	this.Amount = amount
}
func (this *HotTransferRecord) GetAmount() (amount BigDecimal) {
	return this.Amount
}
func (this *HotTransferRecord) SetBalance(balance BigDecimal) {
	this.Balance = balance
}
func (this *HotTransferRecord) GetBalance() (balance BigDecimal) {
	return this.Balance
}
func (this *HotTransferRecord) SetMinerFee(minerFee BigDecimal) {
	this.MinerFee = minerFee
}
func (this *HotTransferRecord) GetMinerFee() (minerFee BigDecimal) {
	return this.MinerFee
}
func (this *HotTransferRecord) SetUnit(unit string) {
	this.Unit = unit
}
func (this *HotTransferRecord) GetUnit() (unit string) {
	return this.Unit
}
func (this *HotTransferRecord) SetAdminName(adminName string) {
	this.AdminName = adminName
}
func (this *HotTransferRecord) GetAdminName() (adminName string) {
	return this.AdminName
}
func (this *HotTransferRecord) SetColdAddress(coldAddress string) {
	this.ColdAddress = coldAddress
}
func (this *HotTransferRecord) GetColdAddress() (coldAddress string) {
	return this.ColdAddress
}
func (this *HotTransferRecord) SetTransactionNumber(transactionNumber string) {
	this.TransactionNumber = transactionNumber
}
func (this *HotTransferRecord) GetTransactionNumber() (transactionNumber string) {
	return this.TransactionNumber
}
func NewHotTransferRecord() (this *HotTransferRecord) {
	this = new(HotTransferRecord)
	return
}

type HotTransferRecord struct {
	Id                int64
	AdminId           int64
	TransferTime      Date
	Amount            BigDecimal
	Balance           BigDecimal
	MinerFee          BigDecimal
	Unit              string
	AdminName         string
	ColdAddress       string
	TransactionNumber string
}
