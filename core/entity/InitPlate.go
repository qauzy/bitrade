package entity

func (this *InitPlate) SetId(id int64) (result *InitPlate) {
	this.Id = id
	return this
}
func (this *InitPlate) GetId() (id int64) {
	return this.Id
}
func (this *InitPlate) SetSymbol(symbol string) (result *InitPlate) {
	this.Symbol = symbol
	return this
}
func (this *InitPlate) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *InitPlate) SetFinalPrice(finalPrice string) (result *InitPlate) {
	this.FinalPrice = finalPrice
	return this
}
func (this *InitPlate) GetFinalPrice() (finalPrice string) {
	return this.FinalPrice
}
func (this *InitPlate) SetInitPrice(initPrice string) (result *InitPlate) {
	this.InitPrice = initPrice
	return this
}
func (this *InitPlate) GetInitPrice() (initPrice string) {
	return this.InitPrice
}
func (this *InitPlate) SetRelativeTime(relativeTime string) (result *InitPlate) {
	this.RelativeTime = relativeTime
	return this
}
func (this *InitPlate) GetRelativeTime() (relativeTime string) {
	return this.RelativeTime
}
func (this *InitPlate) SetInterferenceFactor(interferenceFactor string) (result *InitPlate) {
	this.InterferenceFactor = interferenceFactor
	return this
}
func (this *InitPlate) GetInterferenceFactor() (interferenceFactor string) {
	return this.InterferenceFactor
}

type InitPlate struct {
	Id                 int64
	Symbol             string
	FinalPrice         string
	InitPrice          string
	RelativeTime       string
	InterferenceFactor string
}
