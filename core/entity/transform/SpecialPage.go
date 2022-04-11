package transform

import "github.com/qauzy/chocolate/lists/arraylist"

func (this *SpecialPage[E]) SetContext(Context arraylist.List[E]) (result *SpecialPage[E]) {
	this.Context = Context
	return this
}
func (this *SpecialPage[E]) GetContext() (Context arraylist.List[E]) {
	return this.Context
}
func (this *SpecialPage[E]) SetCurrentPage(CurrentPage int) (result *SpecialPage[E]) {
	this.CurrentPage = CurrentPage
	return this
}
func (this *SpecialPage[E]) GetCurrentPage() (CurrentPage int) {
	return this.CurrentPage
}
func (this *SpecialPage[E]) SetTotalPage(TotalPage int) (result *SpecialPage[E]) {
	this.TotalPage = TotalPage
	return this
}
func (this *SpecialPage[E]) GetTotalPage() (TotalPage int) {
	return this.TotalPage
}
func (this *SpecialPage[E]) SetPageNumber(PageNumber int) (result *SpecialPage[E]) {
	this.PageNumber = PageNumber
	return this
}
func (this *SpecialPage[E]) GetPageNumber() (PageNumber int) {
	return this.PageNumber
}
func (this *SpecialPage[E]) SetTotalElement(TotalElement int) (result *SpecialPage[E]) {
	this.TotalElement = TotalElement
	return this
}
func (this *SpecialPage[E]) GetTotalElement() (TotalElement int) {
	return this.TotalElement
}
func NewSpecialPage[E any](context arraylist.List[E], currentPage int, totalPage int, pageNumber int, totalElement int) (ret *SpecialPage[E]) {
	ret = new(SpecialPage[E])
	ret.Context = context
	ret.CurrentPage = currentPage
	ret.TotalPage = totalPage
	ret.PageNumber = pageNumber
	ret.TotalElement = totalElement
	return
}

type SpecialPage[E any] struct {
	Context      arraylist.List[E] `gorm:"column:context" json:"context"`
	CurrentPage  int               `gorm:"column:current_page" json:"currentPage"`
	TotalPage    int               `gorm:"column:total_page" json:"totalPage"`
	PageNumber   int               `gorm:"column:page_number" json:"pageNumber"`
	TotalElement int               `gorm:"column:total_element" json:"totalElement"`
}
