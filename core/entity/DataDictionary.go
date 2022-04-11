package entity

import "github.com/qauzy/chocolate/xtime"

func (this *DataDictionary) SetId(Id int64) (result *DataDictionary) {
	this.Id = Id
	return this
}
func (this *DataDictionary) GetId() (Id int64) {
	return this.Id
}
func (this *DataDictionary) SetBond(Bond string) (result *DataDictionary) {
	this.Bond = Bond
	return this
}
func (this *DataDictionary) GetBond() (Bond string) {
	return this.Bond
}
func (this *DataDictionary) SetValue(Value string) (result *DataDictionary) {
	this.Value = Value
	return this
}
func (this *DataDictionary) GetValue() (Value string) {
	return this.Value
}
func (this *DataDictionary) SetComment(Comment string) (result *DataDictionary) {
	this.Comment = Comment
	return this
}
func (this *DataDictionary) GetComment() (Comment string) {
	return this.Comment
}
func (this *DataDictionary) SetCreationTime(CreationTime xtime.Xtime) (result *DataDictionary) {
	this.CreationTime = CreationTime
	return this
}
func (this *DataDictionary) GetCreationTime() (CreationTime xtime.Xtime) {
	return this.CreationTime
}
func (this *DataDictionary) SetUpdateTime(UpdateTime xtime.Xtime) (result *DataDictionary) {
	this.UpdateTime = UpdateTime
	return this
}
func (this *DataDictionary) GetUpdateTime() (UpdateTime xtime.Xtime) {
	return this.UpdateTime
}
func (this *DataDictionary) SetImgUrl(ImgUrl string) (result *DataDictionary) {
	this.ImgUrl = ImgUrl
	return this
}
func (this *DataDictionary) GetImgUrl() (ImgUrl string) {
	return this.ImgUrl
}
func NewDataDictionary() (this *DataDictionary) {
	this = new(DataDictionary)
	return
}
func NewDataDictionaryV3(bond string, value string, comment string) (this *DataDictionary) {
	this = new(DataDictionary)
	this.Bond = this.Bond
	this.Value = this.Value
	this.Comment = this.Comment
	return
}
func NewDataDictionaryEx(id int64, bond string, value string, comment string, creationTime xtime.Xtime, updateTime xtime.Xtime, imgUrl string) (ret *DataDictionary) {
	ret = new(DataDictionary)
	ret.Id = id
	ret.Bond = bond
	ret.Value = value
	ret.Comment = comment
	ret.CreationTime = creationTime
	ret.UpdateTime = updateTime
	ret.ImgUrl = imgUrl
	return
}

type DataDictionary struct {
	Id           int64       `gorm:"column:id" json:"id"`
	Bond         string      `gorm:"column:bond" json:"bond"`
	Value        string      `gorm:"column:value" json:"value"`
	Comment      string      `gorm:"column:comment" json:"comment"`
	CreationTime xtime.Xtime `gorm:"column:creation_time" json:"creationTime"`
	UpdateTime   xtime.Xtime `gorm:"column:update_time" json:"updateTime"`
	ImgUrl       string      `gorm:"column:img_url" json:"imgUrl"`
}
