package entity

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/SysAdvertiseLocation"
	"github.com/qauzy/chocolate/xtime"
)

func (this *SysAdvertise) SetSerialNumber(SerialNumber string) (result *SysAdvertise) {
	this.SerialNumber = SerialNumber
	return this
}
func (this *SysAdvertise) GetSerialNumber() (SerialNumber string) {
	return this.SerialNumber
}
func (this *SysAdvertise) SetName(Name string) (result *SysAdvertise) {
	this.Name = Name
	return this
}
func (this *SysAdvertise) GetName() (Name string) {
	return this.Name
}
func (this *SysAdvertise) SetSysAdvertiseLocation(SysAdvertiseLocation SysAdvertiseLocation.SysAdvertiseLocation) (result *SysAdvertise) {
	this.SysAdvertiseLocation = SysAdvertiseLocation
	return this
}
func (this *SysAdvertise) GetSysAdvertiseLocation() (SysAdvertiseLocation SysAdvertiseLocation.SysAdvertiseLocation) {
	return this.SysAdvertiseLocation
}
func (this *SysAdvertise) SetStartTime(StartTime string) (result *SysAdvertise) {
	this.StartTime = StartTime
	return this
}
func (this *SysAdvertise) GetStartTime() (StartTime string) {
	return this.StartTime
}
func (this *SysAdvertise) SetEndTime(EndTime string) (result *SysAdvertise) {
	this.EndTime = EndTime
	return this
}
func (this *SysAdvertise) GetEndTime() (EndTime string) {
	return this.EndTime
}
func (this *SysAdvertise) SetUrl(Url string) (result *SysAdvertise) {
	this.Url = Url
	return this
}
func (this *SysAdvertise) GetUrl() (Url string) {
	return this.Url
}
func (this *SysAdvertise) SetLinkUrl(LinkUrl string) (result *SysAdvertise) {
	this.LinkUrl = LinkUrl
	return this
}
func (this *SysAdvertise) GetLinkUrl() (LinkUrl string) {
	return this.LinkUrl
}
func (this *SysAdvertise) SetRemark(Remark string) (result *SysAdvertise) {
	this.Remark = Remark
	return this
}
func (this *SysAdvertise) GetRemark() (Remark string) {
	return this.Remark
}
func (this *SysAdvertise) SetStatus(Status CommonStatus.CommonStatus) (result *SysAdvertise) {
	this.Status = Status
	return this
}
func (this *SysAdvertise) GetStatus() (Status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *SysAdvertise) SetCreateTime(CreateTime xtime.Xtime) (result *SysAdvertise) {
	this.CreateTime = CreateTime
	return this
}
func (this *SysAdvertise) GetCreateTime() (CreateTime xtime.Xtime) {
	return this.CreateTime
}
func (this *SysAdvertise) SetContent(Content string) (result *SysAdvertise) {
	this.Content = Content
	return this
}
func (this *SysAdvertise) GetContent() (Content string) {
	return this.Content
}
func (this *SysAdvertise) SetAuthor(Author string) (result *SysAdvertise) {
	this.Author = Author
	return this
}
func (this *SysAdvertise) GetAuthor() (Author string) {
	return this.Author
}
func (this *SysAdvertise) SetSort(Sort int) (result *SysAdvertise) {
	this.Sort = Sort
	return this
}
func (this *SysAdvertise) GetSort() (Sort int) {
	return this.Sort
}
func NewSysAdvertise(serialNumber string, name string, sysAdvertiseLocation SysAdvertiseLocation.SysAdvertiseLocation, startTime string, endTime string, url string, linkUrl string, remark string, status CommonStatus.CommonStatus, createTime xtime.Xtime, content string, author string, sort int) (ret *SysAdvertise) {
	ret = new(SysAdvertise)
	ret.SerialNumber = serialNumber
	ret.Name = name
	ret.SysAdvertiseLocation = sysAdvertiseLocation
	ret.StartTime = startTime
	ret.EndTime = endTime
	ret.Url = url
	ret.LinkUrl = linkUrl
	ret.Remark = remark
	ret.Status = status
	ret.CreateTime = createTime
	ret.Content = content
	ret.Author = author
	ret.Sort = sort
	return
}

type SysAdvertise struct {
	SerialNumber         string                                    `gorm:"column:serial_number" json:"serialNumber"`
	Name                 string                                    `gorm:"column:name" json:"name"`
	SysAdvertiseLocation SysAdvertiseLocation.SysAdvertiseLocation `gorm:"column:sys_advertise_location" json:"sysAdvertiseLocation"`
	StartTime            string                                    `gorm:"column:start_time" json:"startTime"`
	EndTime              string                                    `gorm:"column:end_time" json:"endTime"`
	Url                  string                                    `gorm:"column:url" json:"url"`
	LinkUrl              string                                    `gorm:"column:link_url" json:"linkUrl"`
	Remark               string                                    `gorm:"column:remark" json:"remark"`
	Status               CommonStatus.CommonStatus                 `gorm:"column:status" json:"status"`
	CreateTime           xtime.Xtime                               `gorm:"column:create_time" json:"createTime"`
	Content              string                                    `gorm:"column:content" json:"content"`
	Author               string                                    `gorm:"column:author" json:"author"`
	Sort                 int                                       `gorm:"column:sort" json:"sort"`
}
