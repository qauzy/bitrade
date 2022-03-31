
package constant

func (this *PageModel) SetSort() {
	if this.Property == nil || this.Property.Size() == 0 {
		var list = arraylist.New[string]()
		list.Add("id")
		var directions = arraylist.New[>]()
		directions.Add(Sort.Direction.DESC)
		this.Property = list
		this.Direction = directions
	}
}
func (this *PageModel) GetSort() (result *domain.Sort) {
	var orders *arraylist.List[>]
	this.SetSort()
	if this.Direction.Size() == this.Property.Size() {
		orders = arraylist.New[>]()
		var length int = this.Direction.Size()
		for i := 0; i < length; i += 1 {
			orders.Add(Sort.Order(this.Direction.Get(i), this.Property.Get(i)))
		}
	}
	return Sort.By(orders)
}
func (this *PageModel) GetPageable() (result *domain.Pageable) {
	var sort *domain.Sort = this.GetSort()
	if sort == nil {
		return PageRequest.Of(this.PageNo-1, this.PageSize)
	}
	return PageRequest.Of(this.PageNo-1, this.PageSize, sort)
}
func (this *PageModel) DirectoryToOrder(Direction direction) (result *types.Order) {
	if this.Direction.IsAscending() {
		return Order.ASC
	} else {
		return Order.DESC
	}
}
func (this *PageModel) ToOrders(list *arraylist.List[>]) (result *arraylist.List[types.Order]) {
	var orders = arraylist.New[types.Order]()
	for _, this.Direction := range list {
		orders.Add(this.DirectoryToOrder(this.Direction))
	}
	return orders
}
func (this *PageModel) GetOrderSpecifiers() (result *arraylist.List[types.OrderSpecifier]) {
	var orderSpecifiers = arraylist.New[types.OrderSpecifier]()
	this.SetSort()
	if this.GetProperty() != nil {
		for i := 0; i < this.GetProperty().Size(); i += 1 {
			var path *types.Path = ExpressionUtils.Path(types.Path, this.GetProperty().Get(i))
			var orderSpecifier *types.OrderSpecifier = OrderSpecifier(this.ToOrders(this.GetDirection()).Get(i), path)
			orderSpecifiers.Add(orderSpecifier)
		}
	}
	return orderSpecifiers
}

type PageModel struct {
	PageNo    int                     `gorm:"column:page_no" json:"pageNo"`
	PageSize  int                     `gorm:"column:page_size" json:"pageSize"`
	Direction *arraylist.List[>]      `gorm:"column:direction" json:"direction"`
	Property  *arraylist.List[string] `gorm:"column:property" json:"property"`
}

