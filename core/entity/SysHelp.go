package entity

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/SysHelpClassification"
	"github.com/qauzy/chocolate/xtime"
)

func (this *SysHelp) SetId(Id int64) (result *SysHelp) {
	this.Id = Id
	return this
}
func (this *SysHelp) GetId() (Id int64) {
	return this.Id
}
func (this *SysHelp) SetTitle(Title string) (result *SysHelp) {
	this.Title = Title
	return this
}
func (this *SysHelp) GetTitle() (Title string) {
	return this.Title
}
func (this *SysHelp) SetSysHelpClassification(SysHelpClassification SysHelpClassification.SysHelpClassification) (result *SysHelp) {
	this.SysHelpClassification = SysHelpClassification
	return this
}
func (this *SysHelp) GetSysHelpClassification() (SysHelpClassification SysHelpClassification.SysHelpClassification) {
	return this.SysHelpClassification
}
func (this *SysHelp) SetImgUrl(ImgUrl string) (result *SysHelp) {
	this.ImgUrl = ImgUrl
	return this
}
func (this *SysHelp) GetImgUrl() (ImgUrl string) {
	return this.ImgUrl
}
func (this *SysHelp) SetCreateTime(CreateTime xtime.Xtime) (result *SysHelp) {
	this.CreateTime = CreateTime
	return this
}
func (this *SysHelp) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *SysHelp) SetStatus(Status CommonStatus.CommonStatus) (result *SysHelp) {
	this.Status = Status
	return this
}
func (this *SysHelp) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *SysHelp) SetContent(Content string) (result *SysHelp) {
	this.Content = Content
	return this
}
func (this *SysHelp) GetContent() (Content string) {
	return this.Content
}
func (this *SysHelp) SetAuthor(Author string) (result *SysHelp) {
	this.Author = Author
	return this
}
func (this *SysHelp) GetAuthor() (Author string) {
	return this.Author
}
func (this *SysHelp) SetSort(Sort int) (result *SysHelp) {
	this.Sort = Sort
	return this
}
func (this *SysHelp) GetSort() (Sort int) {
	return this.Sort
}
func (this *SysHelp) SetIsTop(IsTop string) (result *SysHelp) {
	this.IsTop = IsTop
	return this
}
func (this *SysHelp) GetIsTop() (IsTop string) {
	return this.IsTop
}
func NewSysHelp(id int64, title string, sysHelpClassification SysHelpClassification.SysHelpClassification, imgUrl string, createTime xtime.Xtime, status CommonStatus.CommonStatus, content string, author string, sort int, isTop string) (ret *SysHelp) {
	ret = new(SysHelp)
	ret.Id = id
	ret.Title = title
	ret.SysHelpClassification = sysHelpClassification
	ret.ImgUrl = imgUrl
	ret.CreateTime = createTime
	ret.Status = status
	ret.Content = content
	ret.Author = author
	ret.Sort = sort
	ret.IsTop = isTop
	return
}

type SysHelp struct {
	Id                    int64                                       `gorm:"column:id" json:"id"`
	Title                 string                                      `gorm:"column:title" json:"title"`
	SysHelpClassification SysHelpClassification.SysHelpClassification `gorm:"column:sys_help_classification" json:"sysHelpClassification"`
	ImgUrl                string                                      `gorm:"column:img_url" json:"imgUrl"`
	CreateTime            xtime.Xtime                                 `gorm:"column:create_time" json:"createTime"`
	Status                CommonStatus.CommonStatus                   `gorm:"column:status" json:"status"`
	Content               string                                      `gorm:"column:content" json:"content"`
	Author                string                                      `gorm:"column:author" json:"author"`
	Sort                  int                                         `gorm:"column:sort" json:"sort"`
	IsTop                 string                                      `gorm:"column:is_top" json:"isTop"`
}
