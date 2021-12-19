package dto

func (this *CoinDTO) SetName(name string) (result *CoinDTO) {
	this.Name = name
	return this
}
func (this *CoinDTO) GetName() (name string) {
	return this.Name
}
func (this *CoinDTO) SetUnit(unit string) (result *CoinDTO) {
	this.Unit = unit
	return this
}
func (this *CoinDTO) GetUnit() (unit string) {
	return this.Unit
}

type CoinDTO struct {
	Name string
	Unit string
}
