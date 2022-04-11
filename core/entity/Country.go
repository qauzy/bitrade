package entity

func (this *Country) SetZhName(ZhName string) (result *Country) {
	this.ZhName = ZhName
	return this
}
func (this *Country) GetZhName() (ZhName string) {
	return this.ZhName
}
func (this *Country) SetEnName(EnName string) (result *Country) {
	this.EnName = EnName
	return this
}
func (this *Country) GetEnName() (EnName string) {
	return this.EnName
}
func (this *Country) SetAreaCode(AreaCode string) (result *Country) {
	this.AreaCode = AreaCode
	return this
}
func (this *Country) GetAreaCode() (AreaCode string) {
	return this.AreaCode
}
func (this *Country) SetLanguage(Language string) (result *Country) {
	this.Language = Language
	return this
}
func (this *Country) GetLanguage() (Language string) {
	return this.Language
}
func (this *Country) SetLocalCurrency(LocalCurrency string) (result *Country) {
	this.LocalCurrency = LocalCurrency
	return this
}
func (this *Country) GetLocalCurrency() (LocalCurrency string) {
	return this.LocalCurrency
}
func (this *Country) SetSort(Sort int) (result *Country) {
	this.Sort = Sort
	return this
}
func (this *Country) GetSort() (Sort int) {
	return this.Sort
}
func NewCountry(zhName string, enName string, areaCode string, language string, localCurrency string, sort int) (ret *Country) {
	ret = new(Country)
	ret.ZhName = zhName
	ret.EnName = enName
	ret.AreaCode = areaCode
	ret.Language = language
	ret.LocalCurrency = localCurrency
	ret.Sort = sort
	return
}

type Country struct {
	ZhName        string `gorm:"column:zh_name" json:"zhName"`
	EnName        string `gorm:"column:en_name" json:"enName"`
	AreaCode      string `gorm:"column:area_code" json:"areaCode"`
	Language      string `gorm:"column:language" json:"language"`
	LocalCurrency string `gorm:"column:local_currency" json:"localCurrency"`
	Sort          int    `gorm:"column:sort" json:"sort"`
}
