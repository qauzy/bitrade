package entity

func (this *Location) SetCountry(country string) {
	this.Country = country
}
func (this *Location) GetCountry() (country string) {
	return this.Country
}
func (this *Location) SetProvince(province string) {
	this.Province = province
}
func (this *Location) GetProvince() (province string) {
	return this.Province
}
func (this *Location) SetCity(city string) {
	this.City = city
}
func (this *Location) GetCity() (city string) {
	return this.City
}
func (this *Location) SetDistrict(district string) {
	this.District = district
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
