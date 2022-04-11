package entity

func (this *Location) SetCountry(Country string) (result *Location) {
	this.Country = Country
	return this
}
func (this *Location) GetCountry() (Country string) {
	return this.Country
}
func (this *Location) SetProvince(Province string) (result *Location) {
	this.Province = Province
	return this
}
func (this *Location) GetProvince() (Province string) {
	return this.Province
}
func (this *Location) SetCity(City string) (result *Location) {
	this.City = City
	return this
}
func (this *Location) GetCity() (City string) {
	return this.City
}
func (this *Location) SetDistrict(District string) (result *Location) {
	this.District = District
	return this
}
func (this *Location) GetDistrict() (District string) {
	return this.District
}
func NewLocation(country string, province string, city string, district string) (ret *Location) {
	ret = new(Location)
	ret.Country = country
	ret.Province = province
	ret.City = city
	ret.District = district
	return
}

type Location struct {
	Country  string `gorm:"column:country" json:"country"`
	Province string `gorm:"column:province" json:"province"`
	City     string `gorm:"column:city" json:"city"`
	District string `gorm:"column:district" json:"district"`
}
