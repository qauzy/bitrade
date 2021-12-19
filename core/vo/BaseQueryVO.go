package vo

func (this *BaseQueryVO) SetPageNum(pageNum int64) (result *BaseQueryVO) {
	this.PageNum = pageNum
	return this
}
func (this *BaseQueryVO) GetPageNum() (pageNum int64) {
	return this.PageNum
}
func (this *BaseQueryVO) SetPageSize(pageSize int64) (result *BaseQueryVO) {
	this.PageSize = pageSize
	return this
}
func (this *BaseQueryVO) GetPageSize() (pageSize int64) {
	return this.PageSize
}

type BaseQueryVO struct {
	PageNum  int64
	PageSize int64
}
