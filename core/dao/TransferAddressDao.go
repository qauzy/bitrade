package dao

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type TransferAddressDao interface {
	FindAllByStatusAndCoin(status *CommonStatus.CommonStatus, coin *entity.Coin) (result arraylist.List[entity.TransferAddress], err error)
	FindByAddressAndCoin(address string, coin *entity.Coin) (result *entity.TransferAddress, err error)
	Save(m *entity.TransferAddress) (result *entity.TransferAddress, err error)
	FindById(id int64) (result *entity.TransferAddress, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.TransferAddress], err error)
}
type transferAddressDao struct {
	*db.DB
}

func NewTransferAddressDao(db *db.DB) (dao TransferAddressDao) {
	dao = &transferAddressDao{db}
	return
}
func (this *transferAddressDao) FindAllByStatusAndCoin(status *CommonStatus.CommonStatus, coin *entity.Coin) (result arraylist.List[entity.TransferAddress], err error) {
	err = this.DBRead().Where("status = ?", status).Where("coin = ?", coin).Find(&result).Error
	return
}
func (this *transferAddressDao) FindByAddressAndCoin(address string, coin *entity.Coin) (result *entity.TransferAddress, err error) {
	err = this.DBRead().Where("address = ?", address).Where("coin = ?", coin).First(&result).Error
	return
}
func (this *transferAddressDao) Save(m *entity.TransferAddress) (result *entity.TransferAddress, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *transferAddressDao) FindById(id int64) (result *entity.TransferAddress, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *transferAddressDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.TransferAddress{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *transferAddressDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.TransferAddress], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
