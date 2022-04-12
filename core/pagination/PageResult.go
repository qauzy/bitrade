package pagination

import "github.com/qauzy/chocolate/lists/arraylist"

func (this *PageResult[T]) SetContent(Content *arraylist.List[T]) (result *PageResult[T]) {
	this.Content = Content
	return this
}
func (this *PageResult[T]) GetContent() (Content *arraylist.List[T]) {
	return this.Content
}
func (this *PageResult[T]) SetNumber(Number int) (result *PageResult[T]) {
	this.Number = Number - 1
	return this
}
func (this *PageResult[T]) GetNumber() (Number int) {
	return this.Number
}
func (this *PageResult[T]) SetSize(Size int) (result *PageResult[T]) {
	this.Size = Size
	return this
}
func (this *PageResult[T]) GetSize() (Size int) {
	return this.Size
}
func (this *PageResult[T]) SetTotalElements(TotalElements int64) (result *PageResult[T]) {
	this.TotalElements = TotalElements
	return this
}
func (this *PageResult[T]) GetTotalElements() (TotalElements int64) {
	return this.TotalElements
}

func (this *PageResult[T]) SetTotalPages(TotalPages int) (result *PageResult[T]) {
	this.TotalPages = TotalPages
	return this
}
func (this *PageResult[T]) GetTotalPages() (TotalPages int) {
	return this.TotalPages
}

func NewPageResult[T any](list *arraylist.List[T], totalNumber int64) (this *PageResult[T]) {
	this = new(PageResult[T])
	this.Content = list
	this.TotalElements = totalNumber
	return
}
func NewPageResultEx[T any](Content *arraylist.List[T], Number int, Size int, TotalElements int64) (this *PageResult[T]) {
	this = new(PageResult[T])
	this.Content = Content
	this.Number = Number - 1
	this.Size = Size
	this.TotalElements = TotalElements
	return
}

func NewPageResultV5[T any](Content *arraylist.List[T], Number int, Size int, TotalPages int, TotalElements int64) (this *PageResult[T]) {
	this = new(PageResult[T])
	this.Content = Content
	this.Number = Number - 1
	this.Size = Size
	this.TotalPages = TotalPages
	this.TotalElements = TotalElements
	return
}

type PageResult[T any] struct {
	Content       *arraylist.List[T] `gorm:"column:content" json:"content"`
	Number        int                `gorm:"column:number" json:"number"`
	Size          int                `gorm:"column:size" json:"size"`
	TotalPages    int                `gorm:"column:totalPages" json:"totalPages"`
	TotalElements int64              `gorm:"column:total_elements" json:"totalElements"`
}
