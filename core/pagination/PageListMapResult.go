package pagination

import (
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/chocolate/maps/hashmap"
)

func (this *PageListMapResult) SetContent(Content *arraylist.List[*hashmap.Map[string, interface {
}]]) (result *PageListMapResult) {
	this.Content = Content
	return this
}
func (this *PageListMapResult) GetContent() (Content *arraylist.List[*hashmap.Map[string, interface{}]]) {
	return this.Content
}
func (this *PageListMapResult) SetNumber(Number int) (result *PageListMapResult) {
	this.Number = Number
	return this
}
func (this *PageListMapResult) GetNumber() (Number int) {
	return this.Number
}
func (this *PageListMapResult) SetSize(Size int) (result *PageListMapResult) {
	this.Size = Size
	return this
}
func (this *PageListMapResult) GetSize() (Size int) {
	return this.Size
}
func (this *PageListMapResult) SetTotalElements(TotalElements int64) (result *PageListMapResult) {
	this.TotalElements = TotalElements
	return this
}
func (this *PageListMapResult) GetTotalElements() (TotalElements int64) {
	return this.TotalElements
}
func NewPageListMapResult(Content *arraylist.List[*hashmap.Map[string, interface{}]], Number int, Size int, TotalElements int64) (this *PageListMapResult) {
	this = new(PageListMapResult)
	this.Content = this.Content
	this.Number = this.Number - 1
	this.Size = this.Size
	this.TotalElements = this.TotalElements
	return
}

type PageListMapResult struct {
	Content       *arraylist.List[*hashmap.Map[string, interface{}]] `gorm:"column:content" json:"content"`
	Number        int                                                `gorm:"column:number" json:"number"`
	Size          int                                                `gorm:"column:size" json:"size"`
	TotalElements int64                                              `gorm:"column:total_elements" json:"totalElements"`
}
