package constant

import "go/types"

func (this *PageModel) SetPageNo(pageNo int64) (result *PageModel) {
	this.PageNo = pageNo
	return this
}
func (this *PageModel) GetPageNo() (pageNo int64) {
	return this.PageNo
}
func (this *PageModel) SetPageSize(pageSize int64) (result *PageModel) {
	this.PageSize = pageSize
	return this
}
func (this *PageModel) GetPageSize() (pageSize int64) {
	return this.PageSize
}
func (this *PageModel) SetDirection(direction []Sort.Direction) (result *PageModel) {
	this.Direction = direction
	return this
}
func (this *PageModel) GetDirection() (direction []Sort.Direction) {
	return this.Direction
}
func (this *PageModel) SetProperty(property []string) (result *PageModel) {
	this.Property = property
	return this
}
func (this *PageModel) GetProperty() (property []string) {
	return this.Property
}
func (this *PageModel) SetSort() {
	if this.Property == nil || this.Property.Size() == 0 {
		var list []string = make([]string, 0)
		list.Add("id")
		var directions []Sort.Direction = make([]Sort.Direction, 0)
		directions.Add(Sort.Direction.DESC)
		this.Property = list
		this.Direction = directions
	}
}
func (this *PageModel) GetSort() (result domain.Sort) {
	var orders []Sort.Order = nil
	this.SetSort()
	if this.Direction.Size() == this.Property.Size() {
		orders = make([]Sort.Order, 0)
		var length int = this.Direction.Size()
		for
		var i = 0
		i < length
		i += 1
		{
			orders.Add(Sort.Order(this.Direction.Get(i), this.Property.Get(i)))
		}
	}
	return Sort.By(orders)
}
func (this *PageModel) GetPageable() (result domain.Pageable) {
	var sort domain.Sort = this.GetSort()
	if sort == nil {
		return PageRequest.Of(this.PageNo-1, this.PageSize)
	}
	return PageRequest.Of(this.PageNo-1, this.PageSize, sort)
}
func (this *PageModel) DirectoryToOrder(direction Sort.Direction) (result types.Order) {
	return func() {
		if direction.IsAscending() {
			return Order.ASC
		} else {
			return Order.DESC
		}
	}
}
func (this *PageModel) ToOrders(list []Sort.Direction) (result []types.Order) {
	var orders []types.Order = make([]types.Order, 0)
	for _, this.Direction := range list{
		orders.Add(this.DirectoryToOrder(this.Direction)),
	}
	return orders
}
func (this *PageModel) GetOrderSpecifiers() (result []types.OrderSpecifier) {
	var orderSpecifiers []types.OrderSpecifier = make([]types.OrderSpecifier, 0)
	this.SetSort()
	if this.GetProperty() != nil {
		for
		var i int = 0
		i < this.GetProperty().Size()
		i += 1
		{
			var path types.Path = ExpressionUtils.Path(class, this.GetProperty().Get(i))
			var orderSpecifier types.OrderSpecifier = OrderSpecifier(this.ToOrders(this.GetDirection()).Get(i), path)
			orderSpecifiers.Add(orderSpecifier)
		}
	}
	return orderSpecifiers
}

type PageModel struct {
	PageNo    int64
	PageSize  int64
	Direction []Sort.Direction
	Property  []string
}
