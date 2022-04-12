package service

func (this *RewardPromotionSettingService) SetDao(dao *dao.RewardPromotionSettingDao) (err error) {
	super.SetDao(dao)
}
func (this *RewardPromotionSettingService) FindByType(oType *PromotionRewardType.PromotionRewardType) (result *entity.RewardPromotionSetting, err error) {
	return dao.FindByStatusAndType(BooleanEnum.IS_TRUE, oType)
}
func (this *RewardPromotionSettingService) Save(setting *entity.RewardPromotionSetting) (result *entity.RewardPromotionSetting, err error) {
	return dao.Save(setting)
}
func (this *RewardPromotionSettingService) Deletes(ids []int64) (err error) {
	for _, id := range ids {
		delete(id)
	}
}
func NewRewardPromotionSettingService() (ret *RewardPromotionSettingService) {
	ret = new(RewardPromotionSettingService)
	return
}

type RewardPromotionSettingService struct {
	Base.TopBaseService[entity.RewardPromotionSetting, dao.RewardPromotionSettingDao]
}
