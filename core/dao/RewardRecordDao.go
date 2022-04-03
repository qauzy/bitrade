package dao

import (
	"bitrade/core/constant/RewardRecordType"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"github.com/qauzy/util/lists/arraylist"
)

type RewardRecordDao interface {
	FindAllByMemberAndType(member *entity.Member, oType *RewardRecordType.RewardRecordType) (result arraylist.List[entity.RewardRecord], err error)
	GetAllPromotionReward(memberId int64, oType int) (result arraylist.List[[]interface {
	}], err error)
	Save(m *entity.RewardRecord) (result *entity.RewardRecord, err error)
	FindById(id int64) (result *entity.RewardRecord, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result arraylist.List[*entity.RewardRecord], err error)
}
type rewardRecordDao struct {
	*db.DB
}

func NewRewardRecordDao(db *db.DB) (dao RewardRecordDao) {
	dao = &rewardRecordDao{db}
	return
}
func (this *rewardRecordDao) FindAllByMemberAndType(member *entity.Member, oType *RewardRecordType.RewardRecordType) (result arraylist.List[entity.RewardRecord], err error) {
	err = this.DBRead().Where("member = ?", member).Where("type = ?", oType).Find(&result).Error
	return
}
func (this *rewardRecordDao) GetAllPromotionReward(memberId int64, oType int) (result arraylist.List[[]interface {
}], err error) {
	eng := this.DBWrite().Table("reward_record").Select("coin_id, sum(amount)").Where("member_id = ? and type = ?", memberId, oType).Group(" group by coin_id").Find(&result)
	err = eng.Error
	return
}
func (this *rewardRecordDao) Save(m *entity.RewardRecord) (result *entity.RewardRecord, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *rewardRecordDao) FindById(id int64) (result *entity.RewardRecord, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *rewardRecordDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.RewardRecord{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *rewardRecordDao) FindAll(qp *types.QueryParam) (result arraylist.List[*entity.RewardRecord], err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
