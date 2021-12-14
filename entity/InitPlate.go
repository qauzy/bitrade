package entity

func (this *InitPlate) SetId(id int64) {
	this.Id = id
}
func (this *InitPlate) GetId() (id int64) {
	return this.Id
}
func (this *InitPlate) SetSymbol(symbol string) {
	this.Symbol = symbol
}
func (this *InitPlate) GetSymbol() (symbol string) {
	return this.Symbol
}
func (this *InitPlate) SetFinalPrice(finalPrice string) {
	this.FinalPrice = finalPrice
}
func (this *InitPlate) GetFinalPrice() (finalPrice string) {
	return this.FinalPrice
}
func (this *InitPlate) SetInitPrice(initPrice string) {
	this.InitPrice = initPrice
}
func (this *InitPlate) GetInitPrice() (initPrice string) {
	return this.InitPrice
}
func (this *InitPlate) SetRelativeTime(relativeTime string) {
	this.RelativeTime = relativeTime
}
func (this *InitPlate) GetRelativeTime() (relativeTime string) {
	return this.RelativeTime
}
func (this *InitPlate) SetInterferenceFactor(interferenceFactor string) {
	this.InterferenceFactor = interferenceFactor
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
