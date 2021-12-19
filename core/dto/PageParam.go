package dto

func (this *PageParam) SetPageNo(pageNo int) (result *PageParam) {
	this.PageNo = pageNo
	return this
}
func (this *PageParam) GetPageNo() (pageNo int) {
	return this.PageNo
}
func (this *PageParam) SetPageSize(pageSize int) (result *PageParam) {
	this.PageSize = pageSize
	return this
}
func (this *PageParam) GetPageSize() (pageSize int) {
	return this.PageSize
}
func (this *PageParam) SetOrders(orders []string) (result *PageParam) {
	this.Orders = orders
	return this
}
func (this *PageParam) GetOrders() (orders []string) {
	return this.Orders
}
func (this *PageParam) SetDirection(direction Sort.Direction) (result *PageParam) {
	this.Direction = direction
	return this
}
func (this *PageParam) GetDirection() (direction Sort.Direction) {
	return this.Direction
}

type PageParam struct {
	PageNo    int
	PageSize  int
	Orders    []string
	Direction Sort.Direction
}
