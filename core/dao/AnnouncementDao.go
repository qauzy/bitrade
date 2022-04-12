package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type AnnouncementDao interface {
	FindMaxSort() (result int, err error)
	Save(m *entity.Announcement) (result *entity.Announcement, err error)
	FindById(id int64) (result *entity.Announcement, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Announcement], err error)
}
type announcementDao struct {
	*db.DB
}

func NewAnnouncementDao(db *db.DB) (dao AnnouncementDao) {
	dao = &announcementDao{db}
	return
}
func (this *announcementDao) FindMaxSort() (result int, err error) {
	eng := this.DBWrite().Table("Announcement as s").Select("max(s.sort)").Find(&result)
	err = eng.Error
	return
}
func (this *announcementDao) Save(m *entity.Announcement) (result *entity.Announcement, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *announcementDao) FindById(id int64) (result *entity.Announcement, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *announcementDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Announcement{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *announcementDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Announcement], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
