package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type LegalWalletRechargeDao interface {
	Save(m *entity.LegalWalletRecharge) (result *entity.LegalWalletRecharge, err error)
	FindById(id int64) (result *entity.LegalWalletRecharge, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.LegalWalletRecharge], err error)
}
type legalWalletRechargeDao struct {
	*db.DB
}

func NewLegalWalletRechargeDao(db *db.DB) (dao LegalWalletRechargeDao) {
	dao = &legalWalletRechargeDao{db}
	return
}
func (this *legalWalletRechargeDao) Save(m *entity.LegalWalletRecharge) (result *entity.LegalWalletRecharge, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *legalWalletRechargeDao) FindById(id int64) (result *entity.LegalWalletRecharge, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *legalWalletRechargeDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.LegalWalletRecharge{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *legalWalletRechargeDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.LegalWalletRecharge], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
