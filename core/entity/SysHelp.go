package entity

func (this *SysHelp) SetId(id int64) (result *SysHelp) {
	this.Id = id
	return this
}
func (this *SysHelp) GetId() (id int64) {
	return this.Id
}
func (this *SysHelp) SetTitle(title string) (result *SysHelp) {
	this.Title = title
	return this
}
func (this *SysHelp) GetTitle() (title string) {
	return this.Title
}
func (this *SysHelp) SetSysHelpClassification(sysHelpClassification constant.SysHelpClassification) (result *SysHelp) {
	this.SysHelpClassification = sysHelpClassification
	return this
}
func (this *SysHelp) GetSysHelpClassification() (sysHelpClassification constant.SysHelpClassification) {
	return this.SysHelpClassification
}
func (this *SysHelp) SetImgUrl(imgUrl string) (result *SysHelp) {
	this.ImgUrl = imgUrl
	return this
}
func (this *SysHelp) GetImgUrl() (imgUrl string) {
	return this.ImgUrl
}
func (this *SysHelp) SetCreateTime(createTime time.Time) (result *SysHelp) {
	this.CreateTime = createTime
	return this
}
func (this *SysHelp) GetCreateTime() (createTime time.Time) {
	return this.CreateTime
}
func (this *SysHelp) SetStatus(status constant.CommonStatus) (result *SysHelp) {
	this.Status = status
	return this
}
func (this *SysHelp) GetStatus() (status constant.CommonStatus) {
	return this.Status
}
func (this *SysHelp) SetContent(content string) (result *SysHelp) {
	this.Content = content
	return this
}
func (this *SysHelp) GetContent() (content string) {
	return this.Content
}
func (this *SysHelp) SetAuthor(author string) (result *SysHelp) {
	this.Author = author
	return this
}
func (this *SysHelp) GetAuthor() (author string) {
	return this.Author
}
func (this *SysHelp) SetSort(sort int) (result *SysHelp) {
	this.Sort = sort
	return this
}
func (this *SysHelp) GetSort() (sort int) {
	return this.Sort
}
func (this *SysHelp) SetIsTop(isTop string) (result *SysHelp) {
	this.IsTop = isTop
	return this
}
func (this *SysHelp) GetIsTop() (isTop string) {
	return this.IsTop
}

type SysHelp struct {
	Id                    int64
	Title                 string
	SysHelpClassification constant.SysHelpClassification
	ImgUrl                string
	CreateTime            time.Time
	Status                constant.CommonStatus
	Content               string
	Author                string
	Sort                  int
	IsTop                 string
}
