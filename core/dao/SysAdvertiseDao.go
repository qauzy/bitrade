package dao

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/SysAdvertiseLocation"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type SysAdvertiseDao interface {
	FindAllByStatusAndSysAdvertiseLocationOrderBySortDesc(status *CommonStatus.CommonStatus, sysAdvertiseLocation *SysAdvertiseLocation.SysAdvertiseLocation) (result arraylist.List[entity.SysAdvertise], err error)
	FindMaxSort() (result int, err error)
	Save(m *entity.SysAdvertise) (result *entity.SysAdvertise, err error)
	FindById(id int64) (result *entity.SysAdvertise, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.SysAdvertise], err error)
}
type sysAdvertiseDao struct {
	*db.DB
}

func NewSysAdvertiseDao(db *db.DB) (dao SysAdvertiseDao) {
	dao = &sysAdvertiseDao{db}
	return
}
func (this *sysAdvertiseDao) FindAllByStatusAndSysAdvertiseLocationOrderBySortDesc(status *CommonStatus.CommonStatus, sysAdvertiseLocation *SysAdvertiseLocation.SysAdvertiseLocation) (result arraylist.List[entity.SysAdvertise], err error) {
	err = this.DBRead().Where("sort_desc = ?", status).Find(&result).Error
	return
}
func (this *sysAdvertiseDao) FindMaxSort() (result int, err error) {
	eng := this.DBWrite().Table("SysAdvertise as s").Select("max(s.sort)").Find(&result)
	err = eng.Error
	return
}
func (this *sysAdvertiseDao) Save(m *entity.SysAdvertise) (result *entity.SysAdvertise, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *sysAdvertiseDao) FindById(id int64) (result *entity.SysAdvertise, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *sysAdvertiseDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.SysAdvertise{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *sysAdvertiseDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.SysAdvertise], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
