package entity

func (this *HotTransferRecord) SetId(id int64) (result *HotTransferRecord) {
	this.Id = id
	return this
}
func (this *HotTransferRecord) GetId() (id int64) {
	return this.Id
}
func (this *HotTransferRecord) SetAdminId(adminId int64) (result *HotTransferRecord) {
	this.AdminId = adminId
	return this
}
func (this *HotTransferRecord) GetAdminId() (adminId int64) {
	return this.AdminId
}
func (this *HotTransferRecord) SetTransferTime(transferTime time.Time) (result *HotTransferRecord) {
	this.TransferTime = transferTime
	return this
}
func (this *HotTransferRecord) GetTransferTime() (transferTime time.Time) {
	return this.TransferTime
}
func (this *HotTransferRecord) SetAmount(amount math.BigDecimal) (result *HotTransferRecord) {
	this.Amount = amount
	return this
}
func (this *HotTransferRecord) GetAmount() (amount math.BigDecimal) {
	return this.Amount
}
func (this *HotTransferRecord) SetBalance(balance math.BigDecimal) (result *HotTransferRecord) {
	this.Balance = balance
	return this
}
func (this *HotTransferRecord) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}
func (this *HotTransferRecord) SetMinerFee(minerFee math.BigDecimal) (result *HotTransferRecord) {
	this.MinerFee = minerFee
	return this
}
func (this *HotTransferRecord) GetMinerFee() (minerFee math.BigDecimal) {
	return this.MinerFee
}
func (this *HotTransferRecord) SetUnit(unit string) (result *HotTransferRecord) {
	this.Unit = unit
	return this
}
func (this *HotTransferRecord) GetUnit() (unit string) {
	return this.Unit
}
func (this *HotTransferRecord) SetAdminName(adminName string) (result *HotTransferRecord) {
	this.AdminName = adminName
	return this
}
func (this *HotTransferRecord) GetAdminName() (adminName string) {
	return this.AdminName
}
func (this *HotTransferRecord) SetColdAddress(coldAddress string) (result *HotTransferRecord) {
	this.ColdAddress = coldAddress
	return this
}
func (this *HotTransferRecord) GetColdAddress() (coldAddress string) {
	return this.ColdAddress
}
func (this *HotTransferRecord) SetTransactionNumber(transactionNumber string) (result *HotTransferRecord) {
	this.TransactionNumber = transactionNumber
	return this
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
	TransferTime      time.Time
	Amount            math.BigDecimal
	Balance           math.BigDecimal
	MinerFee          math.BigDecimal
	Unit              string
	AdminName         string
	ColdAddress       string
	TransactionNumber string
}
