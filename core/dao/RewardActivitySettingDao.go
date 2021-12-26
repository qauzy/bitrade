package dao

import (
	"bitrade/core/constant/ActivityRewardType"
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/dao/db"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
)

type RewardActivitySettingDao interface {
	FindByStatusAndType(booleanEnum BooleanEnum.BooleanEnum, oType ActivityRewardType.ActivityRewardType) (result entity.RewardActivitySetting, err error)
	Save(m *entity.RewardActivitySetting) (result *entity.RewardActivitySetting, err error)
	FindById(id int64) (result *entity.RewardActivitySetting, err error)
	DeleteById(id int64) (count int64, err error)
	FindAll(qp *types.QueryParam) (result []*entity.RewardActivitySetting, err error)
}
type rewardActivitySettingDao struct {
	*db.DB
}

func NewRewardActivitySettingDao(db *db.DB) (dao RewardActivitySettingDao) {
	dao = &rewardActivitySettingDao{db}
	return
}
func (this *rewardActivitySettingDao) FindByStatusAndType(booleanEnum BooleanEnum.BooleanEnum, oType ActivityRewardType.ActivityRewardType) (result entity.RewardActivitySetting, err error) {
	err = this.DBRead().Where("status = ?", booleanEnum).Where("type = ?", oType).First(&result).Error
	return
}
func (this *rewardActivitySettingDao) Save(m *entity.RewardActivitySetting) (result *entity.RewardActivitySetting, err error) {
	err = this.DBWrite().Save(m).Error
	return
}
func (this *rewardActivitySettingDao) FindById(id int64) (result *entity.RewardActivitySetting, err error) {
	err = this.DBRead().Where("id = ?", id).First(&result).Error
	return
}
func (this *rewardActivitySettingDao) DeleteById(id int64) (count int64, err error) {
	d := this.DBRead().Where("id = ?", id).Delete(entity.RewardActivitySetting{})
	err = d.Error
	count = d.RowsAffected
	return
}
func (this *rewardActivitySettingDao) FindAll(qp *types.QueryParam) (result []*entity.RewardActivitySetting, err error) {
	d := this.DBRead()
	if qp != nil {
		d = qp.BuildQuery(d)
	}
	d = d.Find(&result)
	err = d.Error
	return
}
