package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type WebsiteInformationDao interface {
	Save(m *entity.WebsiteInformation) (result *entity.WebsiteInformation, err error)
	FindById(id int64) (result *entity.WebsiteInformation, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.WebsiteInformation], err error)
}
type websiteInformationDao struct {
	*db.DB
}

func NewWebsiteInformationDao(db *db.DB) (dao WebsiteInformationDao) {
	dao = &websiteInformationDao{db}
	return
}
func (this *websiteInformationDao) Save(m *entity.WebsiteInformation) (result *entity.WebsiteInformation, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *websiteInformationDao) FindById(id int64) (result *entity.WebsiteInformation, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *websiteInformationDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.WebsiteInformation{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *websiteInformationDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.WebsiteInformation], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
