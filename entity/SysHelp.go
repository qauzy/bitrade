package entity

func (this *SysHelp) SetId(id int64) {
	this.Id = id
}
func (this *SysHelp) GetId() (id int64) {
	return this.Id
}
func (this *SysHelp) SetTitle(title string) {
	this.Title = title
}
func (this *SysHelp) GetTitle() (title string) {
	return this.Title
}
func (this *SysHelp) SetSysHelpClassification(sysHelpClassification SysHelpClassification) {
	this.SysHelpClassification = sysHelpClassification
}
func (this *SysHelp) GetSysHelpClassification() (sysHelpClassification SysHelpClassification) {
	return this.SysHelpClassification
}
func (this *SysHelp) SetImgUrl(imgUrl string) {
	this.ImgUrl = imgUrl
}
func (this *SysHelp) GetImgUrl() (imgUrl string) {
	return this.ImgUrl
}
func (this *SysHelp) SetCreateTime(createTime Date) {
	this.CreateTime = createTime
}
func (this *SysHelp) GetCreateTime() (createTime Date) {
	return this.CreateTime
}
func (this *SysHelp) SetStatus(status CommonStatus) {
	this.Status = status
}
func (this *SysHelp) GetStatus() (status CommonStatus) {
	return this.Status
}
func (this *SysHelp) SetContent(content string) {
	this.Content = content
}
func (this *SysHelp) GetContent() (content string) {
	return this.Content
}
func (this *SysHelp) SetAuthor(author string) {
	this.Author = author
}
func (this *SysHelp) GetAuthor() (author string) {
	return this.Author
}
func (this *SysHelp) SetSort(sort int) {
	this.Sort = sort
}
func (this *SysHelp) GetSort() (sort int) {
	return this.Sort
}
func (this *SysHelp) SetIsTop(isTop string) {
	this.IsTop = isTop
}
func (this *SysHelp) GetIsTop() (isTop string) {
	return this.IsTop
}

type SysHelp struct {
	Id                    int64
	Title                 string
	SysHelpClassification SysHelpClassification
	ImgUrl                string
	CreateTime            Date
	Status                CommonStatus
	Content               string
	Author                string
	Sort                  int
	IsTop                 string
}
