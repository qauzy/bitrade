package service

import (
	"bitrade/core/dao"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *AnnouncementService) Save(announcement *entity.Announcement) (result *entity.Announcement, err error) {
	return this.AnnouncementDao.Save(announcement)
}
func (this *AnnouncementService) FindAll() (result arraylist.List[entity.Announcement], err error) {
	return this.AnnouncementDao.FindAll()
}
func (this *AnnouncementService) FindById(id int64) (result *entity.Announcement, err error) {
	return this.AnnouncementDao.FindById(id)
}
func (this *AnnouncementService) DeleteById(id int64) (err error) {
	this.AnnouncementDao.DeleteById(id)
}
func (this *AnnouncementService) DeleteBatch(ids []int64) (err error) {
	for _, id := range ids {
		var announcement = this.FindById(id)
		if announcement == nil {
			"validate id!"
		}
		this.DeleteById(id)
	}
}
func (this *AnnouncementService) GetMaxSort() (result int, err error) {
	return this.AnnouncementDao.FindMaxSort()
}
func (this *AnnouncementService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.Announcement], err error) {
	return this.AnnouncementDao.FindAll(predicate, pageable)
}
func NewAnnouncementService(announcementDao *dao.AnnouncementDao) (ret *AnnouncementService) {
	ret = new(AnnouncementService)
	ret.AnnouncementDao = announcementDao
	return
}

type AnnouncementService struct {
	AnnouncementDao dao.AnnouncementDao
	Base.BaseService[entity.Announcement]
}
