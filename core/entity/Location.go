package entity

func (this *Location) SetCountry(country string) (result *Location) {
	this.Country = country
	return this
}
func (this *Location) GetCountry() (country string) {
	return this.Country
}
func (this *Location) SetProvince(province string) (result *Location) {
	this.Province = province
	return this
}
func (this *Location) GetProvince() (province string) {
	return this.Province
}
func (this *Location) SetCity(city string) (result *Location) {
	this.City = city
	return this
}
func (this *Location) GetCity() (city string) {
	return this.City
}
func (this *Location) SetDistrict(district string) (result *Location) {
	this.District = district
	return this
}
func (this *Location) GetDistrict() (district string) {
	return this.District
}

type Location struct {
	Country  string
	Province string
	City     string
	District string
}
