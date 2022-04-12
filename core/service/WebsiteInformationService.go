package service

import (
	"bitrade/core/dao"
	"bitrade/core/entity"
)

func (this *WebsiteInformationService) FetchOne() (result *entity.WebsiteInformation, err error) {
	return this.WebsiteInformationDao.FetchOne()
}
func (this *WebsiteInformationService) Save(websiteInformation *entity.WebsiteInformation) (result *entity.WebsiteInformation, err error) {
	return this.WebsiteInformationDao.Save(websiteInformation)
}
func NewWebsiteInformationService(websiteInformationDao dao.WebsiteInformationDao) (ret *WebsiteInformationService) {
	ret = new(WebsiteInformationService)
	ret.WebsiteInformationDao = websiteInformationDao
	return
}

type WebsiteInformationService struct {
	WebsiteInformationDao dao.WebsiteInformationDao
}
