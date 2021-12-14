package entity

func (this *Country) SetZhName(zhName string) {
	this.ZhName = zhName
}
func (this *Country) GetZhName() (zhName string) {
	return this.ZhName
}
func (this *Country) SetEnName(enName string) {
	this.EnName = enName
}
func (this *Country) GetEnName() (enName string) {
	return this.EnName
}
func (this *Country) SetAreaCode(areaCode string) {
	this.AreaCode = areaCode
}
func (this *Country) GetAreaCode() (areaCode string) {
	return this.AreaCode
}
func (this *Country) SetLanguage(language string) {
	this.Language = language
}
func (this *Country) GetLanguage() (language string) {
	return this.Language
}
func (this *Country) SetLocalCurrency(localCurrency string) {
	this.LocalCurrency = localCurrency
}
func (this *Country) GetLocalCurrency() (localCurrency string) {
	return this.LocalCurrency
}
func (this *Country) SetSort(sort int) {
	this.Sort = sort
}
func (this *Country) GetSort() (sort int) {
	return this.Sort
}

type Country struct {
	ZhName        string
	EnName        string
	AreaCode      string
	Language      string
	LocalCurrency string
	Sort          int
}
