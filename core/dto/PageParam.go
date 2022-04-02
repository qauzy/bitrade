package dto

import (
	"bitrade/core/constant/Direction"
	"github.com/qauzy/util/lists/arraylist"
)

func (this *PageParam) SetPageNo(PageNo int) (result *PageParam) {
	this.PageNo = PageNo
	return this
}
func (this *PageParam) GetPageNo() (PageNo int) {
	return this.PageNo
}
func (this *PageParam) SetPageSize(PageSize int) (result *PageParam) {
	this.PageSize = PageSize
	return this
}
func (this *PageParam) GetPageSize() (PageSize int) {
	return this.PageSize
}
func (this *PageParam) SetOrders(Orders *arraylist.List[string]) (result *PageParam) {
	this.Orders = Orders
	return this
}
func (this *PageParam) GetOrders() (Orders *arraylist.List[string]) {
	return this.Orders
}
func (this *PageParam) SetDirection(Direction Direction.Direction) (result *PageParam) {
	this.Direction = Direction
	return this
}
func (this *PageParam) GetDirection() (Direction Direction.Direction) {
	return this.Direction
}

type PageParam struct {
	PageNo    int                     `gorm:"column:page_no" json:"pageNo"`
	PageSize  int                     `gorm:"column:page_size" json:"pageSize"`
	Orders    *arraylist.List[string] `gorm:"column:orders" json:"orders"`
	Direction Direction.Direction     `gorm:"column:direction" json:"direction"`
}
