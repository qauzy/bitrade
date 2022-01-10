package entity

import "github.com/qauzy/math"

func (this *TransferAddressInfo) SetInfo(info []map[string]interface {
}) (result *TransferAddressInfo) {
	this.Info = info
	return this
}
func (this *TransferAddressInfo) GetInfo() (info []map[string]interface {
}) {
	return this.Info
}
func (this *TransferAddressInfo) SetBalance(balance math.BigDecimal) (result *TransferAddressInfo) {
	this.Balance = balance
	return this
}
func (this *TransferAddressInfo) GetBalance() (balance math.BigDecimal) {
	return this.Balance
}

type TransferAddressInfo struct {
	Info []map[string]interface {
	}
	Balance math.BigDecimal
}
