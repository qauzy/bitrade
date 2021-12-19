package entity

import "time"

func (this *Announcement) SetId(id int64) (result *Announcement) {
	this.Id = id
	return this
}
func (this *Announcement) GetId() (id int64) {
	return this.Id
}
func (this *Announcement) SetTitle(title string) (result *Announcement) {
	this.Title = title
	return this
}
func (this *Announcement) GetTitle() (title string) {
	return this.Title
}
func (this *Announcement) SetContent(content string) (result *Announcement) {
	this.Content = content
	return this
}
func (this *Announcement) GetContent() (content string) {
	return this.Content
}
func (this *Announcement) SetCreateTime(createTime time.Time) (result *Announcement) {
	this.CreateTime = createTime
	return this
}
func (this *Announcement) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *Announcement) SetIsShow(isShow bool) (result *Announcement) {
	this.IsShow = isShow
	return this
}
func (this *Announcement) GetIsShow() (isShow bool) {
	return this.IsShow
}
func (this *Announcement) SetImgUrl(imgUrl string) (result *Announcement) {
	this.ImgUrl = imgUrl
	return this
}
func (this *Announcement) GetImgUrl() (imgUrl string) {
	return this.ImgUrl
}
func (this *Announcement) SetSort(sort int) (result *Announcement) {
	this.Sort = sort
	return this
}
func (this *Announcement) GetSort() (sort int) {
	return this.Sort
}
func (this *Announcement) SetIsTop(isTop string) (result *Announcement) {
	this.IsTop = isTop
	return this
}
func (this *Announcement) GetIsTop() (isTop string) {
	return this.IsTop
}

type Announcement struct {
	Id         int64
	Title      string
	Content    string
	CreateTime time.Time
	IsShow     bool
	ImgUrl     string
	Sort       int
	IsTop      string
}
