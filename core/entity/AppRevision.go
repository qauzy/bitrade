package entity

import (
	"bitrade/core/constant/Platform"
	"time"
)

func (this *AppRevision) SetId(id int64) (result *AppRevision) {
	this.Id = id
	return this
}
func (this *AppRevision) GetId() (id int64) {
	return this.Id
}
func (this *AppRevision) SetPublishTime(publishTime time.Time) (result *AppRevision) {
	this.PublishTime = publishTime
	return this
}
func (this *AppRevision) GetPublishTime() (publishTime time.Time) {
	return this.PublishTime
}
func (this *AppRevision) SetRemark(remark string) (result *AppRevision) {
	this.Remark = remark
	return this
}
func (this *AppRevision) GetRemark() (remark string) {
	return this.Remark
}
func (this *AppRevision) SetVersion(version string) (result *AppRevision) {
	this.Version = version
	return this
}
func (this *AppRevision) GetVersion() (version string) {
	return this.Version
}
func (this *AppRevision) SetDownloadUrl(downloadUrl string) (result *AppRevision) {
	this.DownloadUrl = downloadUrl
	return this
}
func (this *AppRevision) GetDownloadUrl() (downloadUrl string) {
	return this.DownloadUrl
}
func (this *AppRevision) SetPlatform(platform Platform.Platform) (result *AppRevision) {
	this.Platform = platform
	return this
}
func (this *AppRevision) GetPlatform() (platform Platform.Platform) {
	return this.Platform
}

type AppRevision struct {
	Id          int64
	PublishTime time.Time
	Remark      string
	Version     string
	DownloadUrl string
	Platform    Platform.Platform
}
