package transform

func (this *SpecialPage) SetContext(context []E) {
	this.Context = context
}
func (this *SpecialPage) GetContext() (context []E) {
	return this.Context
}
func (this *SpecialPage) SetCurrentPage(currentPage int) {
	this.CurrentPage = currentPage
}
func (this *SpecialPage) GetCurrentPage() (currentPage int) {
	return this.CurrentPage
}
func (this *SpecialPage) SetTotalPage(totalPage int) {
	this.TotalPage = totalPage
}
func (this *SpecialPage) GetTotalPage() (totalPage int) {
	return this.TotalPage
}
func (this *SpecialPage) SetPageNumber(pageNumber int) {
	this.PageNumber = pageNumber
}
func (this *SpecialPage) GetPageNumber() (pageNumber int) {
	return this.PageNumber
}
func (this *SpecialPage) SetTotalElement(totalElement int) {
	this.TotalElement = totalElement
}
func (this *SpecialPage) GetTotalElement() (totalElement int) {
	return this.TotalElement
}

type SpecialPage struct {
	Context      []E
	CurrentPage  int
	TotalPage    int
	PageNumber   int
	TotalElement int
}
