package entity

func (this *AppRevision) SetId(id int64) {
	this.Id = id
}
func (this *AppRevision) GetId() (id int64) {
	return this.Id
}
func (this *AppRevision) SetPublishTime(publishTime Date) {
	this.PublishTime = publishTime
}
func (this *AppRevision) GetPublishTime() (publishTime Date) {
	return this.PublishTime
}
func (this *AppRevision) SetRemark(remark string) {
	this.Remark = remark
}
func (this *AppRevision) GetRemark() (remark string) {
	return this.Remark
}
func (this *AppRevision) SetVersion(version string) {
	this.Version = version
}
func (this *AppRevision) GetVersion() (version string) {
	return this.Version
}
func (this *AppRevision) SetDownloadUrl(downloadUrl string) {
	this.DownloadUrl = downloadUrl
}
func (this *AppRevision) GetDownloadUrl() (downloadUrl string) {
	return this.DownloadUrl
}
func (this *AppRevision) SetPlatform(platform Platform) {
	this.Platform = platform
}
func (this *AppRevision) GetPlatform() (platform Platform) {
	return this.Platform
}

type AppRevision struct {
	Id          int64
	PublishTime Date
	Remark      string
	Version     string
	DownloadUrl string
	Platform    Platform
}
