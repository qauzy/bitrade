package dto

func (this *Pagenation) SetPageParam(pageParam PageParam) (result *Pagenation) {
	this.PageParam = pageParam
	return this
}
func (this *Pagenation) GetPageParam() (pageParam PageParam) {
	return this.PageParam
}
func (this *Pagenation) SetTotalCount(totalCount int64) (result *Pagenation) {
	this.TotalCount = totalCount
	return this
}
func (this *Pagenation) GetTotalCount() (totalCount int64) {
	return this.TotalCount
}
func (this *Pagenation) SetTotalPage(totalPage int64) (result *Pagenation) {
	this.TotalPage = totalPage
	return this
}
func (this *Pagenation) GetTotalPage() (totalPage int64) {
	return this.TotalPage
}
func (this *Pagenation) SetList(list []T) (result *Pagenation) {
	this.List = list
	return this
}
func (this *Pagenation) GetList() (list []T) {
	return this.List
}
func NewPagenation(pageNo int, pageSize int) (this *Pagenation) {
	this = new(Pagenation)
	this.PageParam.SetPageNo(pageNo)
	this.PageParam.SetPageSize(pageSize)
	return
}
func NewPagenation(pageParam PageParam) (this *Pagenation) {
	this = new(Pagenation)
	if pageParam.GetOrders().Size() < 1 {
		pageParam.GetOrders().Add("id")
	}
	this.PageParam = pageParam
	return
}
func (this *Pagenation) SetData(list []interface {
}, totalCount int64, totalPage int64) (result Pagenation) {
	this.List = list
	this.TotalCount = totalCount
	this.TotalPage = totalPage
	return this
}

type Pagenation struct {
	PageParam  PageParam
	TotalCount int64
	TotalPage  int64
	List       []T
}
