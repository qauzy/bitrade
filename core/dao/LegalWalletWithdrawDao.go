package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/chocolate/lists/arraylist"
)

type LegalWalletWithdrawDao interface {
	Save(m *entity.LegalWalletWithdraw) (result *entity.LegalWalletWithdraw, err error)
	FindById(id int64) (result *entity.LegalWalletWithdraw, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.LegalWalletWithdraw], err error)
}
type legalWalletWithdrawDao struct {
	*db.DB
}

func NewLegalWalletWithdrawDao(db *db.DB) (dao LegalWalletWithdrawDao) {
	dao = &legalWalletWithdrawDao{db}
	return
}
func (this *legalWalletWithdrawDao) Save(m *entity.LegalWalletWithdraw) (result *entity.LegalWalletWithdraw, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *legalWalletWithdrawDao) FindById(id int64) (result *entity.LegalWalletWithdraw, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *legalWalletWithdrawDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.LegalWalletWithdraw{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *legalWalletWithdrawDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.LegalWalletWithdraw], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
