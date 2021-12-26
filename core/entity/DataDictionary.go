package entity

func (this *DataDictionary) SetId(id int64) (result *DataDictionary) {
	this.Id = id
	return this
}
func (this *DataDictionary) GetId() (id int64) {
	return this.Id
}
func (this *DataDictionary) SetBond(bond string) (result *DataDictionary) {
	this.Bond = bond
	return this
}
func (this *DataDictionary) GetBond() (bond string) {
	return this.Bond
}
func (this *DataDictionary) SetValue(value string) (result *DataDictionary) {
	this.Value = value
	return this
}
func (this *DataDictionary) GetValue() (value string) {
	return this.Value
}
func (this *DataDictionary) SetComment(comment string) (result *DataDictionary) {
	this.Comment = comment
	return this
}
func (this *DataDictionary) GetComment() (comment string) {
	return this.Comment
}
func (this *DataDictionary) SetCreationTime(creationTime time.Time) (result *DataDictionary) {
	this.CreationTime = creationTime
	return this
}
func (this *DataDictionary) GetCreationTime() (creationTime time.Time) {
	return this.CreationTime
}
func (this *DataDictionary) SetUpdateTime(updateTime time.Time) (result *DataDictionary) {
	this.UpdateTime = updateTime
	return this
}
func (this *DataDictionary) GetUpdateTime() (updateTime time.Time) {
	return this.UpdateTime
}
func (this *DataDictionary) SetImgUrl(imgUrl string) (result *DataDictionary) {
	this.ImgUrl = imgUrl
	return this
}
func (this *DataDictionary) GetImgUrl() (imgUrl string) {
	return this.ImgUrl
}
func NewDataDictionary() (this *DataDictionary) {
	this = new(DataDictionary)
	return
}
func NewDataDictionary(bond string, value string, comment string) (this *DataDictionary) {
	this = new(DataDictionary)
	this.Bond = bond
	this.Value = value
	this.Comment = comment
	return
}

type DataDictionary struct {
	Id           int64
	Bond         string
	Value        string
	Comment      string
	CreationTime time.Time
	UpdateTime   time.Time
	ImgUrl       string
}
