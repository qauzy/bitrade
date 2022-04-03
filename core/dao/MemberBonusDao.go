package dao

import (
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/dto"
	"github.com/qauzy/math"
	"github.com/qauzy/util/lists/arraylist"
)

type MemberBonusDao interface {
	GetBonusByMemberId(memberId int64) (result arraylist.List[dto.MemberBonusDTO], err error)
	GetBonusAmountByMemberId(memberId int64) (result *math.BigDecimal, err error)
	Save(m *entity.MemberBonus) (result *entity.MemberBonus, err error)
	FindById(id int64) (result *entity.MemberBonus, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberBonus], err error)
}
type memberBonusDao struct {
	*db.DB
}

func NewMemberBonusDao(db *db.DB) (dao MemberBonusDao) {
	dao = &memberBonusDao{db}
	return
}
func (this *memberBonusDao) GetBonusByMemberId(memberId int64) (result arraylist.List[dto.MemberBonusDTO], err error) {
	eng := this.DBWrite().Table("member_bonus").Select("*").Where("member_id = ?", memberId).Order("id desc").Find(&result)
	err = eng.Error
	return
}
func (this *memberBonusDao) GetBonusAmountByMemberId(memberId int64) (result *math.BigDecimal, err error) {
	eng := this.DBWrite().Table("member_bonus").Select("SUM(mem_bouns)").Where("member_id = ?", memberId).Find(&result)
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
func (this *memberBonusDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.MemberBonus], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
