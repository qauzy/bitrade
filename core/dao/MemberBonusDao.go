package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/dto"
	"github.com/qauzy/math"
)

type MemberBonusDao interface {
	GetBonusByMemberId(memberId int64) (result []dto.MemberBonusDTO, err error)
	GetBonusAmountByMemberId(memberId int64) (result math.BigDecimal, err error)
	Save(m *entity.MemberBonus) (result *entity.MemberBonus, err error)
	FindById(id int64) (result *entity.MemberBonus, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.MemberBonus, err error)
}
type memberBonusDao struct {
	*db.DB
}

func NewMemberBonusDao(db *db.DB) (dao MemberBonusDao) {
	dao = &memberBonusDao{db}
	return
}
func (this *memberBonusDao) GetBonusByMemberId(memberId int64) (result []dto.MemberBonusDTO, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT * from member_bonus  where member_id=? ORDER BY id DESC ", memberId)
	err = eng.Error
	return
}
func (this *memberBonusDao) GetBonusAmountByMemberId(memberId int64) (result math.BigDecimal, err error) {

	//FIXME 非原生sql，需要处理
	eng := this.DBWrite().Exec("SELECT SUM(mem_bouns) from member_bonus  where member_id=?", memberId)
	err = eng.Error
	return
}
func (this *memberBonusDao) Save(m *entity.MemberBonus) (result *entity.MemberBonus, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *memberBonusDao) FindById(id int64) (result *entity.MemberBonus, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *memberBonusDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.MemberBonus{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *memberBonusDao) FindAll(qp *types.QueryParam) (result []*entity.MemberBonus, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
