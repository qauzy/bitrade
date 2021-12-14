package entity

func (this *SysAdvertise) SetSerialNumber(serialNumber string) {
	this.SerialNumber = serialNumber
}
func (this *SysAdvertise) GetSerialNumber() (serialNumber string) {
	return this.SerialNumber
}
func (this *SysAdvertise) SetName(name string) {
	this.Name = name
}
func (this *SysAdvertise) GetName() (name string) {
	return this.Name
}
func (this *SysAdvertise) SetSysAdvertiseLocation(sysAdvertiseLocation SysAdvertiseLocation) {
	this.SysAdvertiseLocation = sysAdvertiseLocation
}
func (this *SysAdvertise) GetSysAdvertiseLocation() (sysAdvertiseLocation SysAdvertiseLocation) {
	return this.SysAdvertiseLocation
}
func (this *SysAdvertise) SetStartTime(startTime string) {
	this.StartTime = startTime
}
func (this *SysAdvertise) GetStartTime() (startTime string) {
	return this.StartTime
}
func (this *SysAdvertise) SetEndTime(endTime string) {
	this.EndTime = endTime
}
func (this *SysAdvertise) GetEndTime() (endTime string) {
	return this.EndTime
}
func (this *SysAdvertise) SetUrl(url string) {
	this.Url = url
}
func (this *SysAdvertise) GetUrl() (url string) {
	return this.Url
}
func (this *SysAdvertise) SetLinkUrl(linkUrl string) {
	this.LinkUrl = linkUrl
}
func (this *SysAdvertise) GetLinkUrl() (linkUrl string) {
	return this.LinkUrl
}
func (this *SysAdvertise) SetRemark(remark string) {
	this.Remark = remark
}
func (this *SysAdvertise) GetRemark() (remark string) {
	return this.Remark
}
func (this *SysAdvertise) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *SysAdvertise) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *SysAdvertise) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *SysAdvertise) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *SysAdvertise) SetContent(content string) {
	this.Content = content
}
func (this *SysAdvertise) GetContent() (content string) {
	return this.Content
}
func (this *SysAdvertise) SetAuthor(author string) {
	this.Author = author
}
func (this *SysAdvertise) GetAuthor() (author string) {
	return this.Author
}
func (this *SysAdvertise) SetSort(sort int) {
	this.Sort = sort
}
func (this *SysAdvertise) GetSort() (sort int) {
	return this.Sort
}

type SysAdvertise struct {
	SerialNumber         string
	Name                 string
	SysAdvertiseLocation SysAdvertiseLocation
	StartTime            string
	EndTime              string
	Url                  string
	LinkUrl              string
	Remark               string
	Status               CommonStatus
	CreateTime           Date
	Content              string
	Author               string
	Sort                 int
}
