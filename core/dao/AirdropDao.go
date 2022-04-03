package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type AirdropDao interface {
	Save(m *entity.Airdrop) (result *entity.Airdrop, err error)
	FindById(id int64) (result *entity.Airdrop, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Airdrop], err error)
}
type airdropDao struct {
	*db.DB
}

func NewAirdropDao(db *db.DB) (dao AirdropDao) {
	dao = &airdropDao{db}
	return
}
func (this *airdropDao) Save(m *entity.Airdrop) (result *entity.Airdrop, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *airdropDao) FindById(id int64) (result *entity.Airdrop, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *airdropDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.Airdrop{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *airdropDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.Airdrop], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
