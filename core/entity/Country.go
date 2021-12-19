package entity

func (this *Country) SetZhName(zhName string) (result *Country) {
	this.ZhName = zhName
	return this
}
func (this *Country) GetZhName() (zhName string) {
	return this.ZhName
}
func (this *Country) SetEnName(enName string) (result *Country) {
	this.EnName = enName
	return this
}
func (this *Country) GetEnName() (enName string) {
	return this.EnName
}
func (this *Country) SetAreaCode(areaCode string) (result *Country) {
	this.AreaCode = areaCode
	return this
}
func (this *Country) GetAreaCode() (areaCode string) {
	return this.AreaCode
}
func (this *Country) SetLanguage(language string) (result *Country) {
	this.Language = language
	return this
}
func (this *Country) GetLanguage() (language string) {
	return this.Language
}
func (this *Country) SetLocalCurrency(localCurrency string) (result *Country) {
	this.LocalCurrency = localCurrency
	return this
}
func (this *Country) GetLocalCurrency() (localCurrency string) {
	return this.LocalCurrency
}
func (this *Country) SetSort(sort int) (result *Country) {
	this.Sort = sort
	return this
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
