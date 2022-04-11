package entity

import (
	"bitrade/core/constant/Platform"
	"github.com/qauzy/chocolate/xtime"
)

func (this *AppRevision) SetId(Id int64) (result *AppRevision) {
	this.Id = Id
	return this
}
func (this *AppRevision) GetId() (Id int64) {
	return this.Id
}
func (this *AppRevision) SetPublishTime(PublishTime xtime.Xtime) (result *AppRevision) {
	this.PublishTime = PublishTime
	return this
}
func (this *AppRevision) GetPublishTime() (PublishTime xtime.Xtime) {
	return this.PublishTime
}
func (this *AppRevision) SetRemark(Remark string) (result *AppRevision) {
	this.Remark = Remark
	return this
}
func (this *AppRevision) GetRemark() (Remark string) {
	return this.Remark
}
func (this *AppRevision) SetVersion(Version string) (result *AppRevision) {
	this.Version = Version
	return this
}
func (this *AppRevision) GetVersion() (Version string) {
	return this.Version
}
func (this *AppRevision) SetDownloadUrl(DownloadUrl string) (result *AppRevision) {
	this.DownloadUrl = DownloadUrl
	return this
}
func (this *AppRevision) GetDownloadUrl() (DownloadUrl string) {
	return this.DownloadUrl
}
func (this *AppRevision) SetPlatform(Platform Platform.Platform) (result *AppRevision) {
	this.Platform = Platform
	return this
}
func (this *AppRevision) GetPlatform() (Platform Platform.Platform) {
	return this.Platform
}
func NewAppRevision(id int64, publishTime xtime.Xtime, remark string, version string, downloadUrl string, platform Platform.Platform) (ret *AppRevision) {
	ret = new(AppRevision)
	ret.Id = id
	ret.PublishTime = publishTime
	ret.Remark = remark
	ret.Version = version
	ret.DownloadUrl = downloadUrl
	ret.Platform = platform
	return
}

type AppRevision struct {
	Id          int64             `gorm:"column:id" json:"id"`
	PublishTime xtime.Xtime       `gorm:"column:publish_time" json:"publishTime"`
	Remark      string            `gorm:"column:remark" json:"remark"`
	Version     string            `gorm:"column:version" json:"version"`
	DownloadUrl string            `gorm:"column:download_url" json:"downloadUrl"`
	Platform    Platform.Platform `gorm:"column:platform" json:"platform"`
}
