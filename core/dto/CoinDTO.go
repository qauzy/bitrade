package dto

func (this *CoinDTO) SetName(Name string) (result *CoinDTO) {
	this.Name = Name
	return this
}
func (this *CoinDTO) GetName() (Name string) {
	return this.Name
}
func (this *CoinDTO) SetUnit(Unit string) (result *CoinDTO) {
	this.Unit = Unit
	return this
}
func (this *CoinDTO) GetUnit() (Unit string) {
	return this.Unit
}

type CoinDTO struct {
	Name string `gorm:"column:name" json:"name"`
	Unit string `gorm:"column:unit" json:"unit"`
}
