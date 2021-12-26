package dao

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/PromotionRewardType"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type RewardPromotionSettingDao interface {
	FindByStatusAndType(booleanEnum BooleanEnum.BooleanEnum, oType PromotionRewardType.PromotionRewardType) (result entity.RewardPromotionSetting, err error)
	Save(m *entity.RewardPromotionSetting) (result *entity.RewardPromotionSetting, err error)
	FindById(id int64) (result *entity.RewardPromotionSetting, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.RewardPromotionSetting, err error)
}
type rewardPromotionSettingDao struct {
	*db.DB
}

func NewRewardPromotionSettingDao(db *db.DB) (dao RewardPromotionSettingDao) {
	dao = &rewardPromotionSettingDao{db}
	return
}
func (this *rewardPromotionSettingDao) FindByStatusAndType(booleanEnum BooleanEnum.BooleanEnum, oType PromotionRewardType.PromotionRewardType) (result entity.RewardPromotionSetting, err error) {
	err = this.DBRead().Where("status = ?", booleanEnum).Where("type = ?", oType).First(&result).Error
	return
}
func (this *rewardPromotionSettingDao) Save(m *entity.RewardPromotionSetting) (result *entity.RewardPromotionSetting, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *rewardPromotionSettingDao) FindById(id int64) (result *entity.RewardPromotionSetting, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *rewardPromotionSettingDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.RewardPromotionSetting{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *rewardPromotionSettingDao) FindAll(qp *types.QueryParam) (result []*entity.RewardPromotionSetting, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
