package dao

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
	"time"
)

type MemberAddressDao interface {
	DeleteMemberAddress(date time.Time, id int64, memberId int64) (result int, err error)
	FindAllByMemberIdAndAddressAndStatus(memberId int64, address string, status *CommonStatus.CommonStatus) (result arraylist.List[entity.MemberAddress], err error)
	FindByMemberIdAndCoinAndAddressAndStatus(memberId int64, coin *entity.Coin, address string, status *CommonStatus.CommonStatus) (result arraylist.List[entity.MemberAddress], err error)
	Save(m *entity.MemberAddress) (result *entity.MemberAddress, err error)
	FindById(id int64) (result *entity.MemberAddress, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberAddress], err error)
}
type memberAddressDao struct {
	*db.DB
}

func NewMemberAddressDao(db *db.DB) (dao MemberAddressDao) {
	dao = &memberAddressDao{db}
	return
}
func (this *memberAddressDao) DeleteMemberAddress(date time.Time, id int64, memberId int64) (result int, err error) {

	//FIXME 非原生sql,需要处理
	eng := this.DBWrite().Exec("update MemberAddress a set a.deleteTime=?,a.status=1 where a.status=0 and a.id=? and a.memberId=?", date, id, memberId)
	err = eng.Error
	return
}
func (this *memberAddressDao) FindAllByMemberIdAndAddressAndStatus(memberId int64, address string, status *CommonStatus.CommonStatus) (result arraylist.List[entity.MemberAddress], err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Where("address = ?", address).Where("status = ?", status).Find(&result).Error
	return
}
func (this *memberAddressDao) FindByMemberIdAndCoinAndAddressAndStatus(memberId int64, coin *entity.Coin, address string, status *CommonStatus.CommonStatus) (result arraylist.List[entity.MemberAddress], err error) {
	err = this.DBRead().Where("member_id = ?", memberId).Where("coin = ?", coin).Where("address = ?", address).Where("status = ?", status).First(&result).Error
	return
}
func (this *memberAddressDao) Save(m *entity.MemberAddress) (result *entity.MemberAddress, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberAddressDao) FindById(id int64) (result *entity.MemberAddress, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberAddressDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberAddress{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberAddressDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberAddress], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
