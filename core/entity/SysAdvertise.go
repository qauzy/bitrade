package entity

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/SysAdvertiseLocation"
	"time"
)

func (this *SysAdvertise) SetSerialNumber(serialNumber string) (result *SysAdvertise) {
	this.SerialNumber = serialNumber
	return this
}
func (this *SysAdvertise) GetSerialNumber() (serialNumber string) {
	return this.SerialNumber
}
func (this *SysAdvertise) SetName(name string) (result *SysAdvertise) {
	this.Name = name
	return this
}
func (this *SysAdvertise) GetName() (name string) {
	return this.Name
}
func (this *SysAdvertise) SetSysAdvertiseLocation(sysAdvertiseLocation SysAdvertiseLocation.SysAdvertiseLocation) (result *SysAdvertise) {
	this.SysAdvertiseLocation = sysAdvertiseLocation
	return this
}
func (this *SysAdvertise) GetSysAdvertiseLocation() (sysAdvertiseLocation SysAdvertiseLocation.SysAdvertiseLocation) {
	return this.SysAdvertiseLocation
}
func (this *SysAdvertise) SetStartTime(startTime string) (result *SysAdvertise) {
	this.StartTime = startTime
	return this
}
func (this *SysAdvertise) GetStartTime() (startTime string) {
	return this.StartTime
}
func (this *SysAdvertise) SetEndTime(endTime string) (result *SysAdvertise) {
	this.EndTime = endTime
	return this
}
func (this *SysAdvertise) GetEndTime() (endTime string) {
	return this.EndTime
}
func (this *SysAdvertise) SetUrl(url string) (result *SysAdvertise) {
	this.Url = url
	return this
}
func (this *SysAdvertise) GetUrl() (url string) {
	return this.Url
}
func (this *SysAdvertise) SetLinkUrl(linkUrl string) (result *SysAdvertise) {
	this.LinkUrl = linkUrl
	return this
}
func (this *SysAdvertise) GetLinkUrl() (linkUrl string) {
	return this.LinkUrl
}
func (this *SysAdvertise) SetRemark(remark string) (result *SysAdvertise) {
	this.Remark = remark
	return this
}
func (this *SysAdvertise) GetRemark() (remark string) {
	return this.Remark
}
func (this *SysAdvertise) SetStatus(status CommonStatus.CommonStatus) (result *SysAdvertise) {
	this.Status = status
	return this
}
func (this *SysAdvertise) GetStatus() (status CommonStatus.CommonStatus) {
	return this.Status
}
func (this *SysAdvertise) SetCreateTime(createTime time.Time) (result *SysAdvertise) {
	this.CreateTime = createTime
	return this
}
func (this *SysAdvertise) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *SysAdvertise) SetContent(content string) (result *SysAdvertise) {
	this.Content = content
	return this
}
func (this *SysAdvertise) GetContent() (content string) {
	return this.Content
}
func (this *SysAdvertise) SetAuthor(author string) (result *SysAdvertise) {
	this.Author = author
	return this
}
func (this *SysAdvertise) GetAuthor() (author string) {
	return this.Author
}
func (this *SysAdvertise) SetSort(sort int) (result *SysAdvertise) {
	this.Sort = sort
	return this
}
func (this *SysAdvertise) GetSort() (sort int) {
	return this.Sort
}

type SysAdvertise struct {
	SerialNumber         string
	Name                 string
	SysAdvertiseLocation SysAdvertiseLocation.SysAdvertiseLocation
	StartTime            string
	EndTime              string
	Url                  string
	LinkUrl              string
	Remark               string
	Status               CommonStatus.CommonStatus
	CreateTime           time.Time
	Content              string
	Author               string
	Sort                 int
}
