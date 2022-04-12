package service

func (this *CountryService) GetAllCountry() (result arraylist.List[entity.Country], err error) {
	return this.CountryDao.FindAllOrderBySort()
}
func (this *CountryService) FindOne(zhName string) (result *entity.Country, err error) {
	return this.CountryDao.FindByZhName(zhName)
}
func (this *CountryService) FindByLegalCurrency(legalCurrency string) (result arraylist.List[entity.Country], err error) {
	return this.CountryDao.FindByLocalCurrency(legalCurrency)
}
func NewCountryService(countryDao *dao.CountryDao) (ret *CountryService) {
	ret = new(CountryService)
	ret.CountryDao = countryDao
	return
}

type CountryService struct {
	CountryDao dao.CountryDao
}
