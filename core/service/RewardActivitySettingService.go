package service

func (this *RewardActivitySettingService) SetDao(dao *dao.RewardActivitySettingDao) (err error) {
	this.Dao = dao
}
func (this *RewardActivitySettingService) FindByType(oType *ActivityRewardType.ActivityRewardType) (result *entity.RewardActivitySetting, err error) {
	return dao.FindByStatusAndType(BooleanEnum.IS_TRUE, oType)
}
func (this *RewardActivitySettingService) Save(rewardActivitySetting *entity.RewardActivitySetting) (result *entity.RewardActivitySetting, err error) {
	return dao.Save(rewardActivitySetting)
}
func NewRewardActivitySettingService() (ret *RewardActivitySettingService) {
	ret = new(RewardActivitySettingService)
	return
}

type RewardActivitySettingService struct {
	Base.TopBaseService[entity.RewardActivitySetting, dao.RewardActivitySettingDao]
}
