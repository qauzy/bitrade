package entity

func (this *DataDictionary) SetId(id int64) {
	this.Id = id
}
func (this *DataDictionary) GetId() (id int64) {
	return this.Id
}
func (this *DataDictionary) SetBond(bond string) {
	this.Bond = bond
}
func (this *DataDictionary) GetBond() (bond string) {
	return this.Bond
}
func (this *DataDictionary) SetValue(value string) {
	this.Value = value
}
func (this *DataDictionary) GetValue() (value string) {
	return this.Value
}
func (this *DataDictionary) SetComment(comment string) {
	this.Comment = comment
}
func (this *DataDictionary) GetComment() (comment string) {
	return this.Comment
}
func (this *DataDictionary) SetCreationTime(creationTime Date) {
	this.CreationTime = creationTime
}
func (this *DataDictionary) GetCreationTime() (creationTime Date) {
	return this.CreationTime
}
func (this *DataDictionary) SetUpdateTime(updateTime Date) {
	this.UpdateTime = updateTime
}
func (this *DataDictionary) GetUpdateTime() (updateTime Date) {
	return this.UpdateTime
}
func (this *DataDictionary) SetImgUrl(imgUrl string) {
	this.ImgUrl = imgUrl
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
	CreationTime Date
	UpdateTime   Date
	ImgUrl       string
}
