package entity

import (
	"github.com/qauzy/math"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/qauzy/util/maps/hashmap"
)

func (this *TransferAddressInfo) SetInfo(Info *arraylist.List[*hashmap.Map[string, interface {
}]]) (result *TransferAddressInfo) {
	this.Info = Info
	return this
}
func (this *TransferAddressInfo) GetInfo() (Info *arraylist.List[*hashmap.Map[string, interface {
}]]) {
	return this.Info
}
func (this *TransferAddressInfo) SetBalance(Balance *math.BigDecimal) (result *TransferAddressInfo) {
	this.Balance = Balance
	return this
}
func (this *TransferAddressInfo) GetBalance() (Balance *math.BigDecimal) {
	return this.Balance
}

type TransferAddressInfo struct {
	Info *arraylist.List[*hashmap.Map[string, interface {
	}]] `gorm:"column:info" json:"info"`
	Balance *math.BigDecimal `gorm:"column:balance" json:"balance"`
}
