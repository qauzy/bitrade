package entity

func (this *TransferAddress) SetId(id int64) (result *TransferAddress) {
	this.Id = id
	return this
}
func (this *TransferAddress) GetId() (id int64) {
	return this.Id
}
func (this *TransferAddress) SetCoin(coin Coin) (result *TransferAddress) {
	this.Coin = coin
	return this
}
func (this *TransferAddress) GetCoin() (coin Coin) {
	return this.Coin
}
func (this *TransferAddress) SetAddress(address string) (result *TransferAddress) {
	this.Address = address
	return this
}
func (this *TransferAddress) GetAddress() (address string) {
	return this.Address
}
func (this *TransferAddress) SetTransferFee(transferFee math.BigDecimal) (result *TransferAddress) {
	this.TransferFee = transferFee
	return this
}
func (this *TransferAddress) GetTransferFee() (transferFee math.BigDecimal) {
	return this.TransferFee
}
func (this *TransferAddress) SetMinAmount(minAmount math.BigDecimal) (result *TransferAddress) {
	this.MinAmount = minAmount
	return this
}
func (this *TransferAddress) GetMinAmount() (minAmount math.BigDecimal) {
	return this.MinAmount
}
func (this *TransferAddress) SetStatus(status constant.CommonStatus) (result *TransferAddress) {
	this.Status = status
	return this
}
func (this *TransferAddress) GetStatus() (status constant.CommonStatus) {
	return this.Status
}

type TransferAddress struct {
	Id          int64
	Coin        Coin
	Address     string
	TransferFee math.BigDecimal
	MinAmount   math.BigDecimal
	Status      constant.CommonStatus
}
