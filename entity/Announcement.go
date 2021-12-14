package entity

func (this *Announcement) SetId(id int64) {
	this.Id = id
}
func (this *Announcement) GetId() (id int64) {
	return this.Id
}
func (this *Announcement) SetTitle(title string) {
	this.Title = title
}
func (this *Announcement) GetTitle() (title string) {
	return this.Title
}
func (this *Announcement) SetContent(content string) {
	this.Content = content
}
func (this *Announcement) GetContent() (content string) {
	return this.Content
}
func (this *Announcement) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *Announcement) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *Announcement) SetIsShow(isShow bool) {
	this.IsShow = isShow
}
func (this *Announcement) GetIsShow() (isShow bool) {
	return this.IsShow
}
func (this *Announcement) SetImgUrl(imgUrl string) {
	this.ImgUrl = imgUrl
}
func (this *Announcement) GetImgUrl() (imgUrl string) {
	return this.ImgUrl
}
func (this *Announcement) SetSort(sort int) {
	this.Sort = sort
}
func (this *Announcement) GetSort() (sort int) {
	return this.Sort
}
func (this *Announcement) SetIsTop(isTop string) {
	this.IsTop = isTop
}
func (this *Announcement) GetIsTop() (isTop string) {
	return this.IsTop
}

type Announcement struct {
	Id         int64
	Title      string
	Content    string
	CreateTime Date
	IsShow     bool
	ImgUrl     string
	Sort       int
	IsTop      string
}
