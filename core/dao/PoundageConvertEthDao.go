package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/math"
)

type PoundageConvertEthDao interface {
	GetAmountEthByDate(startDate string, endDate string) (result math.BigDecimal, err error)
	GetAmountEth(startTime string, endTime string) (result math.BigDecimal, err error)
	Save(m *entity.PoundageConvertEth) (result *entity.PoundageConvertEth, err error)
	FindById(id int64) (result *entity.PoundageConvertEth, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.PoundageConvertEth, err error)
}
type poundageConvertEthDao struct {
	*db.DB
}

func NewPoundageConvertEthDao(db *db.DB) (dao PoundageConvertEthDao) {
	dao = &poundageConvertEthDao{db}
	return
}
func (this *poundageConvertEthDao) GetAmountEthByDate(startDate string, endDate string) (result math.BigDecimal, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT SUM(poundage_amount_Eth) poundage_amount_eth FROM poundage_convert_eth WHERE transaction_time>= :startTime AND transaction_time<=:endTime ", startDate, endDate)
	err = eng.Error
	return
}
func (this *poundageConvertEthDao) GetAmountEth(startTime string, endTime string) (result math.BigDecimal, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT SUM(poundage_amount_Eth) poundage_amount_eth FROM poundage_convert_eth WHERE transaction_time>= ? AND transaction_time<=? AND member_id NOT IN (66946,65859,52350,118863,119284,76895)", startTime, endTime)
	err = eng.Error
	return
}
func (this *poundageConvertEthDao) Save(m *entity.PoundageConvertEth) (result *entity.PoundageConvertEth, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *poundageConvertEthDao) FindById(id int64) (result *entity.PoundageConvertEth, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *poundageConvertEthDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.PoundageConvertEth{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *poundageConvertEthDao) FindAll(qp *types.QueryParam) (result []*entity.PoundageConvertEth, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
