package entity

import "github.com/qauzy/chocolate/xtime"

func (this *Announcement) SetId(Id int64) (result *Announcement) {
	this.Id = Id
	return this
}
func (this *Announcement) GetId() (Id int64) {
	return this.Id
}
func (this *Announcement) SetTitle(Title string) (result *Announcement) {
	this.Title = Title
	return this
}
func (this *Announcement) GetTitle() (Title string) {
	return this.Title
}
func (this *Announcement) SetContent(Content string) (result *Announcement) {
	this.Content = Content
	return this
}
func (this *Announcement) GetContent() (Content string) {
	return this.Content
}
func (this *Announcement) SetCreateTime(CreateTime xtime.Xtime) (result *Announcement) {
	this.CreateTime = CreateTime
	return this
}
func (this *Announcement) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *Announcement) SetIsShow(IsShow bool) (result *Announcement) {
	this.IsShow = IsShow
	return this
}
func (this *Announcement) GetIsShow() (IsShow bool) {
	return this.IsShow
}
func (this *Announcement) SetImgUrl(ImgUrl string) (result *Announcement) {
	this.ImgUrl = ImgUrl
	return this
}
func (this *Announcement) GetImgUrl() (ImgUrl string) {
	return this.ImgUrl
}
func (this *Announcement) SetSort(Sort int) (result *Announcement) {
	this.Sort = Sort
	return this
}
func (this *Announcement) GetSort() (Sort int) {
	return this.Sort
}
func (this *Announcement) SetIsTop(IsTop string) (result *Announcement) {
	this.IsTop = IsTop
	return this
}
func (this *Announcement) GetIsTop() (IsTop string) {
	return this.IsTop
}
func NewAnnouncement(id int64, title string, content string, createTime xtime.Xtime, isShow bool, imgUrl string, sort int, isTop string) (ret *Announcement) {
	ret = new(Announcement)
	ret.Id = id
	ret.Title = title
	ret.Content = content
	ret.CreateTime = createTime
	ret.IsShow = isShow
	ret.ImgUrl = imgUrl
	ret.Sort = sort
	ret.IsTop = isTop
	return
}

type Announcement struct {
	Id         int64       `gorm:"column:id" json:"id"`
	Title      string      `gorm:"column:title" json:"title"`
	Content    string      `gorm:"column:content" json:"content"`
	CreateTime xtime.Xtime `gorm:"column:create_time" json:"createTime"`
	IsShow     bool        `gorm:"column:is_show" json:"isShow"`
	ImgUrl     string      `gorm:"column:img_url" json:"imgUrl"`
	Sort       int         `gorm:"column:sort" json:"sort"`
	IsTop      string      `gorm:"column:is_top" json:"isTop"`
}
