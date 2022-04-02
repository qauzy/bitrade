package dto

import "github.com/qauzy/util/lists/arraylist"

func (this *Pagenation[T]) SetPageParam(PageParam *PageParam) (result *Pagenation[T]) {
	this.PageParam = PageParam
	return this
}
func (this *Pagenation[T]) GetPageParam() (PageParam *PageParam) {
	return this.PageParam
}
func (this *Pagenation[T]) SetTotalCount(TotalCount int64) (result *Pagenation[T]) {
	this.TotalCount = TotalCount
	return this
}
func (this *Pagenation[T]) GetTotalCount() (TotalCount int64) {
	return this.TotalCount
}
func (this *Pagenation[T]) SetTotalPage(TotalPage int64) (result *Pagenation[T]) {
	this.TotalPage = TotalPage
	return this
}
func (this *Pagenation[T]) GetTotalPage() (TotalPage int64) {
	return this.TotalPage
}
func (this *Pagenation[T]) SetList(List *arraylist.List[T]) (result *Pagenation[T]) {
	this.List = List
	return this
}
func (this *Pagenation[T]) GetList() (List *arraylist.List[T]) {
	return this.List
}
func NewPagenation[T any](pageNo int, pageSize int) (this *Pagenation[T]) {
	this = new(Pagenation[T])
	this.PageParam.SetPageNo(pageNo)
	this.PageParam.SetPageSize(pageSize)
	return
}
func NewPagenationEx[T any](PageParam PageParam) (this *Pagenation[T]) {
	this = new(Pagenation[T])
	if this.PageParam.GetOrders().Size() < 1 {
		this.PageParam.GetOrders().Add("id")
	}
	this.PageParam = this.PageParam
	return
}
func (this *Pagenation[T]) SetData(List *arraylist.List[T], TotalCount int64, TotalPage int64) (result *Pagenation[T]) {
	this.List = this.List
	this.TotalCount = this.TotalCount
	this.TotalPage = this.TotalPage
	return this
}

type Pagenation[T any] struct {
	PageParam  *PageParam         `gorm:"column:page_param" json:"pageParam"`
	TotalCount int64              `gorm:"column:total_count" json:"totalCount"`
	TotalPage  int64              `gorm:"column:total_page" json:"totalPage"`
	List       *arraylist.List[T] `gorm:"column:list" json:"list"`
}
