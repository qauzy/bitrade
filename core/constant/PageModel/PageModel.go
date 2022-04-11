
package PageModel

func (this *PageModel) SetPageNo(PageNo int) (result *PageModel) {
	this.PageNo = PageNo
	return this
}
func (this *PageModel) GetPageNo() (PageNo int) {
	return this.PageNo
}
func (this *PageModel) SetPageSize(PageSize int) (result *PageModel) {
	this.PageSize = PageSize
	return this
}
func (this *PageModel) GetPageSize() (PageSize int) {
	return this.PageSize
}
func (this *PageModel) SetDirection(Direction arraylist.List[>]) (result *PageModel) {
	this.Direction = Direction
	return this
}
func (this *PageModel) GetDirection() (Direction arraylist.List[>]) {
	return this.Direction
}
func (this *PageModel) SetProperty(Property arraylist.List[string]) (result *PageModel) {
	this.Property = Property
	return this
}
func (this *PageModel) GetProperty() (Property arraylist.List[string]) {
	return this.Property
}
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
	var orders arraylist.List[>]
	this.SetSort()
	if this.Direction.Size() == this.Property.Size() {
		orders = arraylist.New[>]()
		var length = this.Direction.Size()
		for i := 0; i < length; i += 1 {
			orders.Add(Sort.Order(this.Direction.Get(i), this.Property.Get(i)))
		}
	}
	return Sort.By(orders)
}
func (this *PageModel) GetPageable() (result *domain.Pageable) {
	var sort = this.GetSort()
	if sort == nil {
		return PageRequest.Of(this.PageNo-1, this.PageSize)
	}
	return PageRequest.Of(this.PageNo-1, this.PageSize, sort)
}
func (this *PageModel) DirectoryToOrder(direction direction) (result *types.Order) {
	if this.Direction.IsAscending() {
		return Order.ASC
	} else {
		return Order.DESC
	}
}
func (this *PageModel) ToOrders(list arraylist.List[>]) (result arraylist.List[types.Order]) {
	var orders = arraylist.New[types.Order]()
	for _, this.Direction := range list {
		orders.Add(this.DirectoryToOrder(this.Direction))
	}
	return orders
}
func (this *PageModel) GetOrderSpecifiers() (result arraylist.List[types.OrderSpecifier]) {
	var orderSpecifiers = arraylist.New[types.OrderSpecifier]()
	this.SetSort()
	if this.GetProperty() != nil {
		for i := 0; i < this.GetProperty().Size(); i += 1 {
			var path = ExpressionUtils.Path(types.Path, this.GetProperty().Get(i))
			var orderSpecifier = OrderSpecifier(this.ToOrders(this.GetDirection()).Get(i), path)
			orderSpecifiers.Add(orderSpecifier)
		}
	}
	return orderSpecifiers
}
func NewPageModel(pageNo int, pageSize int, direction arraylist.List[>], property arraylist.List[string]) (ret *PageModel) {
	ret = new(PageModel)
	ret.PageNo = pageNo
	ret.PageSize = pageSize
	ret.Direction = direction
	ret.Property = property
	return
}

type PageModel struct {
	PageNo    int
	PageSize  int
	Direction arraylist.List[>]
	Property  arraylist.List[string]
}

