package entity

func (this *InitPlate) SetId(Id int64) (result *InitPlate) {
	this.Id = Id
	return this
}
func (this *InitPlate) GetId() (Id int64) {
	return this.Id
}
func (this *InitPlate) SetSymbol(Symbol string) (result *InitPlate) {
	this.Symbol = Symbol
	return this
}
func (this *InitPlate) GetSymbol() (Symbol string) {
	return this.Symbol
}
func (this *InitPlate) SetFinalPrice(FinalPrice string) (result *InitPlate) {
	this.FinalPrice = FinalPrice
	return this
}
func (this *InitPlate) GetFinalPrice() (FinalPrice string) {
	return this.FinalPrice
}
func (this *InitPlate) SetInitPrice(InitPrice string) (result *InitPlate) {
	this.InitPrice = InitPrice
	return this
}
func (this *InitPlate) GetInitPrice() (InitPrice string) {
	return this.InitPrice
}
func (this *InitPlate) SetRelativeTime(RelativeTime string) (result *InitPlate) {
	this.RelativeTime = RelativeTime
	return this
}
func (this *InitPlate) GetRelativeTime() (RelativeTime string) {
	return this.RelativeTime
}
func (this *InitPlate) SetInterferenceFactor(InterferenceFactor string) (result *InitPlate) {
	this.InterferenceFactor = InterferenceFactor
	return this
}
func (this *InitPlate) GetInterferenceFactor() (InterferenceFactor string) {
	return this.InterferenceFactor
}
func NewInitPlate(id int64, symbol string, finalPrice string, initPrice string, relativeTime string, interferenceFactor string) (ret *InitPlate) {
	ret = new(InitPlate)
	ret.Id = id
	ret.Symbol = symbol
	ret.FinalPrice = finalPrice
	ret.InitPrice = initPrice
	ret.RelativeTime = relativeTime
	ret.InterferenceFactor = interferenceFactor
	return
}

type InitPlate struct {
	Id                 int64  `gorm:"column:id" json:"id"`
	Symbol             string `gorm:"column:symbol" json:"symbol"`
	FinalPrice         string `gorm:"column:final_price" json:"finalPrice"`
	InitPrice          string `gorm:"column:init_price" json:"initPrice"`
	RelativeTime       string `gorm:"column:relative_time" json:"relativeTime"`
	InterferenceFactor string `gorm:"column:interference_factor" json:"interferenceFactor"`
}
