package dao

import (
	"bitrade/core/constant/OrderStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/qauzy/util/maps/hashmap"
	"time"
)

type OrderDao interface {
	GetOrderByMemberIdAndOrderSn(memberId int64, orderSn string) (result arraylist.List[entity.Order], err error)
	GetOrderByMemberId(memberId int64) (result arraylist.List[entity.Order], err error)
	GetOrderByMemberIdAndStatus(memberId int64, status *OrderStatus.OrderStatus) (result arraylist.List[entity.Order], err error)
	GetOrderByMemberIdAndStatusAndOrderSn(memberId int64, status *OrderStatus.OrderStatus, orderSn string) (result arraylist.List[entity.Order], err error)
	FindByAdvertiseId(advertiseId int64) (result arraylist.List[entity.Order], err error)
	GetOrderByOrderSn(v2 string) (result *entity.Order, err error)
	UpdatePayOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error)
	CancelOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error)
	ReleaseOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error)
	FindAllExpiredOrder(date time.Time) (result arraylist.List[entity.Order], err error)
	FingAllProcessingOrder(myId int64, unPay *OrderStatus.OrderStatus, paid *OrderStatus.OrderStatus, appeal *OrderStatus.OrderStatus) (result arraylist.List[entity.Order], err error)
	UpdateAppealOrder(status *OrderStatus.OrderStatus, orderSn string) (result int, err error)
	CountByCreateTimeBetween(startTime time.Time, endTime time.Time) (result int, err error)
	CountByStatusAndCreateTimeBetween(status *OrderStatus.OrderStatus, startTime time.Time, endTime time.Time) (result int, err error)
	CountByStatus(status *OrderStatus.OrderStatus) (result int, err error)
	GetOtcTurnoverAmount(date string) (result arraylist.List[[]interface {
	}], err error)
	GetBusinessStatistics(memberId int64) (result *hashmap.Map[string, interface {
	}], err error)
	CancelNopaymentOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error)
	CountOrdersByMemberIdAndCreateTime(startTime time.Time, endTime time.Time) (result arraylist.List[[]interface {
	}], err error)
	Save(m *entity.Order) (result *entity.Order, err error)
	FindById(id int64) (result *entity.Order, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Order], err error)
}
type orderDao struct {
	*db.DB
}

func NewOrderDao(db *db.DB) (dao OrderDao) {
	dao = &orderDao{db}
	return
}
func (this *orderDao) GetOrderByMemberIdAndOrderSn(memberId int64, orderSn string) (result arraylist.List[entity.Order], err error) {
	return
}
func (this *orderDao) GetOrderByMemberId(memberId int64) (result arraylist.List[entity.Order], err error) {
	return
}
func (this *orderDao) GetOrderByMemberIdAndStatus(memberId int64, status *OrderStatus.OrderStatus) (result arraylist.List[entity.Order], err error) {
	return
}
func (this *orderDao) GetOrderByMemberIdAndStatusAndOrderSn(memberId int64, status *OrderStatus.OrderStatus, orderSn string) (result arraylist.List[entity.Order], err error) {
	return
}
func (this *orderDao) FindByAdvertiseId(advertiseId int64) (result arraylist.List[entity.Order], err error) {
	err = this.DBRead().Where("advertise_id = ?", advertiseId).First(&result).Error
	return
}
func (this *orderDao) GetOrderByOrderSn(v2 string) (result *entity.Order, err error) {
	return
}
func (this *orderDao) UpdatePayOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update Order a set a.payTime=?,a.status=? where a.status=1 and a.orderSn=?", date, status, orderSn)
	err = eng.Error
	return
}
func (this *orderDao) CancelOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update Order a set a.cancelTime=?,a.status=? where (a.status=1 or a.status=2 or a.status=4) and a.orderSn=?", date, status, orderSn)
	err = eng.Error
	return
}
func (this *orderDao) ReleaseOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update Order a set a.releaseTime=?,a.status=? where  (a.status=2 or a.status=4) and a.orderSn=?", date, status, orderSn)
	err = eng.Error
	return
}
func (this *orderDao) FindAllExpiredOrder(date time.Time) (result arraylist.List[entity.Order], err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("select a from Order a where timestampdiff(MINUTE,a.createTime,?)>=a.timeLimit and a.status=1", date)
	err = eng.Error
	return
}
func (this *orderDao) FingAllProcessingOrder(myId int64, unPay *OrderStatus.OrderStatus, paid *OrderStatus.OrderStatus, appeal *OrderStatus.OrderStatus) (result arraylist.List[entity.Order], err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("select a from Order a where (a.memberId=? or a.customerId=?) and (a.status=? or a.status=? or a.status=?)", myId, unPay, paid, appeal)
	err = eng.Error
	return
}
func (this *orderDao) UpdateAppealOrder(status *OrderStatus.OrderStatus, orderSn string) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update Order a set a.status=? where a.status=2 and a.orderSn=?", status, orderSn)
	err = eng.Error
	return
}
func (this *orderDao) CountByCreateTimeBetween(startTime time.Time, endTime time.Time) (result int, err error) {
	return
}
func (this *orderDao) CountByStatusAndCreateTimeBetween(status *OrderStatus.OrderStatus, startTime time.Time, endTime time.Time) (result int, err error) {
	return
}
func (this *orderDao) CountByStatus(status *OrderStatus.OrderStatus) (result int, err error) {
	return
}
func (this *orderDao) GetOtcTurnoverAmount(date string) (result arraylist.List[[]interface {
}], err error) {
	eng := this.DBWrite().Table("otc_order as b, otc_coin as a").Select("a.unit as unit, date_format(b.release_time, '%Y-%m-%d'), sum(b.number) as amount, sum(b.commission) as fee, sum(money)").Where("a.id = b.coin_id and b.`status` = 3 and date_format(b.release_time, '%Y-%m-%d') = ?", date).Group(" group by a.unit").Find(&result)
	err = eng.Error
	return
}
func (this *orderDao) GetBusinessStatistics(memberId int64) (result *hashmap.Map[string, interface {
}], err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("select sum(b.commission) as fee,sum(b.money) as money from Order b  where  b.status = 3 and b.memberId = ?", memberId)
	err = eng.Error
	return
}
func (this *orderDao) CancelNopaymentOrder(date time.Time, status *OrderStatus.OrderStatus, orderSn string) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update Order a set a.cancelTime=?,a.status=? where a.status=1 and a.orderSn=?", date, status, orderSn)
	err = eng.Error
	return
}
func (this *orderDao) CountOrdersByMemberIdAndCreateTime(startTime time.Time, endTime time.Time) (result arraylist.List[[]interface {
}], err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("select o.memberId,count(o.memberId) as c from Order o where o.createTime between ? and ? GROUP BY o.memberId", startTime, endTime)
	err = eng.Error
	return
}
func (this *orderDao) Save(m *entity.Order) (result *entity.Order, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *orderDao) FindById(id int64) (result *entity.Order, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *orderDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Order{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *orderDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Order], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
