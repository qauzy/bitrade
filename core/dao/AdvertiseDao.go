package dao

import (
	"bitrade/core/constant/AdvertiseControlStatus"
	"bitrade/core/constant/AdvertiseType"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/math"
	"time"
)

type AdvertiseDao interface {
	FindAllByMemberIdAndStatusNot(var1 int64, status AdvertiseControlStatus.AdvertiseControlStatus, var2 domain.Sort) (result []entity.Advertise, err error)
	FindAdvertiseByIdAndMemberIdAndStatusNot(var1 int64, var2 int64, status AdvertiseControlStatus.AdvertiseControlStatus) (result entity.Advertise, err error)
	FindByIdAndMemberId(v1 int64, v2 int64) (result entity.Advertise, err error)
	UpdateAdvertiseStatus(status AdvertiseControlStatus.AdvertiseControlStatus, id int64, mid int64, advertiseControlStatus AdvertiseControlStatus.AdvertiseControlStatus) (result int, err error)
	FindAllByMemberIdAndStatusAndAdvertiseType(var1 int64, status AdvertiseControlStatus.AdvertiseControlStatus, oType AdvertiseType.AdvertiseType) (result []entity.Advertise, err error)
	UpdateAdvertiseAmount(status AdvertiseControlStatus.AdvertiseControlStatus, id int64, amount math.BigDecimal) (result int, err error)
	UpdateAdvertiseDealAmount(id int64, amount math.BigDecimal) (result int, err error)
	PutOffAdvertise(id int64, amount math.BigDecimal) (result int, err error)
	FindAllByMemberIdAndStatus(vari int64, status AdvertiseControlStatus.AdvertiseControlStatus) (result []entity.Advertise, err error)
	AlterStatusBatch(status AdvertiseControlStatus.AdvertiseControlStatus, updateTime time.Time, ids []int) (result int, err error)
	GetAdvertiseNum(member entity.Member) (result int64, err error)
	UpdateAdvertiseDealAmountAndRemainAmount(id int64, amount math.BigDecimal) (result int, err error)
	CountAllByMemberAndStatus(member entity.Member, status AdvertiseControlStatus.AdvertiseControlStatus) (result int, err error)
	Save(m *entity.Advertise) (result *entity.Advertise, err error)
	FindById(id int64) (result *entity.Advertise, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.Advertise, err error)
}
type advertiseDao struct {
	*db.DB
}

func NewAdvertiseDao(db *db.DB) (dao AdvertiseDao) {
	dao = &advertiseDao{db}
	return
}
func (this *advertiseDao) FindAllByMemberIdAndStatusNot(var1 int64, status AdvertiseControlStatus.AdvertiseControlStatus, var2 domain.Sort) (result []entity.Advertise, err error) {
	err = this.DBRead().Where("member_id = ?", var1).Where("status <> ?", status).Find(&result).Error
	return
}
func (this *advertiseDao) FindAdvertiseByIdAndMemberIdAndStatusNot(var1 int64, var2 int64, status AdvertiseControlStatus.AdvertiseControlStatus) (result entity.Advertise, err error) {
	err = this.DBRead().Where("id = ?", var1).Where("member_id = ?", var2).Where("status <> ?", status).First(&result).Error
	return
}
func (this *advertiseDao) FindByIdAndMemberId(v1 int64, v2 int64) (result entity.Advertise, err error) {
	err = this.DBRead().Where("id = ?", v1).Where("member_id = ?", v2).First(&result).Error
	return
}
func (this *advertiseDao) UpdateAdvertiseStatus(status AdvertiseControlStatus.AdvertiseControlStatus, id int64, mid int64, advertiseControlStatus AdvertiseControlStatus.AdvertiseControlStatus) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Advertise a set a.status=? where a.id=? and a.member.id=? and a.status<>?", status, id, mid, advertiseControlStatus)
	err = eng.Error
	return
}
func (this *advertiseDao) FindAllByMemberIdAndStatusAndAdvertiseType(var1 int64, status AdvertiseControlStatus.AdvertiseControlStatus, oType AdvertiseType.AdvertiseType) (result []entity.Advertise, err error) {
	err = this.DBRead().Where("member_id = ?", var1).Where("status = ?", status).Where("advertise_type = ?", oType).Find(&result).Error
	return
}
func (this *advertiseDao) UpdateAdvertiseAmount(status AdvertiseControlStatus.AdvertiseControlStatus, id int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Advertise a set a.remainAmount=a.remainAmount-?,a.dealAmount=a.dealAmount+? where a.remainAmount>=:amount and a.status=:sta and a.id=?", status, id, amount)
	err = eng.Error
	return
}
func (this *advertiseDao) UpdateAdvertiseDealAmount(id int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Advertise a set a.dealAmount=a.dealAmount-? where a.dealAmount>=?  and a.id=?", id, amount)
	err = eng.Error
	return
}
func (this *advertiseDao) PutOffAdvertise(id int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Advertise a set a.status=1,a.remainAmount=0,a.number=0 where a.remainAmount=?  and a.id=? and a.status=0", id, amount)
	err = eng.Error
	return
}
func (this *advertiseDao) FindAllByMemberIdAndStatus(vari int64, status AdvertiseControlStatus.AdvertiseControlStatus) (result []entity.Advertise, err error) {
	err = this.DBRead().Where("member_id = ?", vari).Where("status = ?", status).Find(&result).Error
	return
}
func (this *advertiseDao) AlterStatusBatch(status AdvertiseControlStatus.AdvertiseControlStatus, updateTime time.Time, ids []int) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Advertise a set a.status = ?,a.updateTime=? where a.id in ?", status, updateTime, ids)
	err = eng.Error
	return
}
func (this *advertiseDao) GetAdvertiseNum(member entity.Member) (result int64, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("select count(a.id) from Advertise a where a.member = ?", member)
	err = eng.Error
	return
}
func (this *advertiseDao) UpdateAdvertiseDealAmountAndRemainAmount(id int64, amount math.BigDecimal) (result int, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("update Advertise a set a.dealAmount=a.dealAmount-?,a.remainAmount=a.remainAmount+? where a.dealAmount>=:amount  and a.id=?", id, amount)
	err = eng.Error
	return
}
func (this *advertiseDao) CountAllByMemberAndStatus(member entity.Member, status AdvertiseControlStatus.AdvertiseControlStatus) (result int, err error) {
	return
}
func (this *advertiseDao) Save(m *entity.Advertise) (result *entity.Advertise, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *advertiseDao) FindById(id int64) (result *entity.Advertise, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *advertiseDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Advertise{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *advertiseDao) FindAll(qp *types.QueryParam) (result []*entity.Advertise, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
