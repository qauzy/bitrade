package dao

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/SysHelpClassification"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type SysHelpDao interface {
	FindAllBySysHelpClassificationAndStatusNot(sysHelpClassification SysHelpClassification.SysHelpClassification, commonStatus CommonStatus.CommonStatus) (result []entity.SysHelp, err error)
	FindMaxSort() (result int, err error)
	GetCateTop(cate string) (result []entity.SysHelp, err error)
	FindAllByStatusNotAndSort() (result []entity.SysHelp, err error)
	Save(m *entity.SysHelp) (result *entity.SysHelp, err error)
	FindById(id int64) (result *entity.SysHelp, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.SysHelp, err error)
}
type sysHelpDao struct {
	*db.DB
}

func NewSysHelpDao(db *db.DB) (dao SysHelpDao) {
	dao = &sysHelpDao{db}
	return
}
func (this *sysHelpDao) FindAllBySysHelpClassificationAndStatusNot(sysHelpClassification SysHelpClassification.SysHelpClassification, commonStatus CommonStatus.CommonStatus) (result []entity.SysHelp, err error) {
	err = this.DBRead().Where("sys_help_classification = ?", sysHelpClassification).Where("status <> ?", commonStatus).Find(&result).Error
	return
}
func (this *sysHelpDao) FindMaxSort() (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select max(s.sort) from SysHelp s")
	err = eng.Error
	return
}
func (this *sysHelpDao) GetCateTop(cate string) (result []entity.SysHelp, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select * from sys_help WHERE sys_help_classification=? and is_top='0' ", cate)
	err = eng.Error
	return
}
func (this *sysHelpDao) FindAllByStatusNotAndSort() (result []entity.SysHelp, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT * FROM sys_help s WHERE s.status ='0' ORDER BY s.is_top asc , s.sort DESC")
	err = eng.Error
	return
}
func (this *sysHelpDao) Save(m *entity.SysHelp) (result *entity.SysHelp, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *sysHelpDao) FindById(id int64) (result *entity.SysHelp, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *sysHelpDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.SysHelp{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *sysHelpDao) FindAll(qp *types.QueryParam) (result []*entity.SysHelp, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
