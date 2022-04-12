package service

func (this *RewardRecordService) Save(rewardRecord *entity.RewardRecord) (result *entity.RewardRecord, err error) {
	return this.RewardRecordDao.Save(rewardRecord)
}
func (this *RewardRecordService) QueryRewardPromotionList(member *entity.Member) (result arraylist.List[entity.RewardRecord], err error) {
	return this.RewardRecordDao.FindAllByMemberAndType(member, RewardRecordType.PROMOTION)
}
func (this *RewardRecordService) GetAllPromotionReward(memberId int64, oType *RewardRecordType.RewardRecordType) (result *hashmap.Map[string, math.BigDecimal], err error) {
	var list = this.RewardRecordDao.GetAllPromotionReward(memberId, oType.GetOrdinal())
	var oMap = hashset.New[math.BigDecimal]()
	for _, array := range list {
		oMap.Put(array[0].ToString(), array[1].(math.BigDecimal))
	}
	return oMap
}
func (this *RewardRecordService) FindAll(predicate *types.Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.RewardRecord], err error) {
	return this.RewardRecordDao.FindAll(predicate, pageModel.GetPageable())
}
func (this *RewardRecordService) FindCount(predicate *types.Predicate) (result int64, err error) {
	return this.RewardRecordDao.Count(predicate)
}
func NewRewardRecordService(rewardRecordDao *dao.RewardRecordDao) (ret *RewardRecordService) {
	ret = new(RewardRecordService)
	ret.RewardRecordDao = rewardRecordDao
	return
}

type RewardRecordService struct {
	RewardRecordDao dao.RewardRecordDao
	Base.BaseService
}
