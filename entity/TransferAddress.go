package entity

func (this *TransferAddress) SetId(id int64) {
	this.Id = id
}
func (this *TransferAddress) GetId() (id int64) {
	return this.Id
}
func (this *TransferAddress) SetCoin(coin Coin) {
	this.Coin = coin
}
func (this *TransferAddress) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *TransferAddress) SetAddress(address string) {
	this.Address = address
}
func (this *TransferAddress) GetAddress() (address string) {
	return this.Address
}
func (this *TransferAddress) SetTransferFee(transferFee BigDecimal) {
	this.TransferFee = transferFee
}
func (this *TransferAddress) GetTransferFee() (transferFee BigDecimal) {
	return this.TransferFee
}
func (this *TransferAddress) SetMinAmount(minAmount BigDecimal) {
	this.MinAmount = minAmount
}
func (this *TransferAddress) GetMinAmount() (minAmount BigDecimal) {
	return this.MinAmount
}
func (this *TransferAddress) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *TransferAddress) GetStatus() (status CommonStatus) {
	return this.Status
}

type TransferAddress struct {
	Id          int64
	Coin        Coin
	Address     string
	TransferFee BigDecimal
	MinAmount   BigDecimal
	Status      CommonStatus
}
