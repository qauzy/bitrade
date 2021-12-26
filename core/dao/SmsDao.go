package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/dto"
)

type SmsDao interface {
	FindBySmsStatus() (result dto.SmsDTO, err error)
	Save(m *entity.Sms) (result *entity.Sms, err error)
	FindById(id int64) (result *entity.Sms, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Sms, err error)
}
type smsDao struct {
	*db.DB
}

func NewSmsDao(db *db.DB) (dao SmsDao) {
	dao = &smsDao{db}
	return
}
func (this *smsDao) FindBySmsStatus() (result dto.SmsDTO, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select * from tb_sms where sms_status = '0' ")
	err = eng.Error
	return
}
func (this *smsDao) Save(m *entity.Sms) (result *entity.Sms, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *smsDao) FindById(id int64) (result *entity.Sms, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *smsDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Sms{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *smsDao) FindAll(qp *types.QueryParam) (result []*entity.Sms, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
