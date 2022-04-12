package service

func (this *RewardWalletService) FindRewardWalletByMemberId(memberId int64) (result *entity.RewardWallet, err error) {
	return this.RewardWalletDao.FindRewardWalletByMemberId(memberId)
}
func (this *RewardWalletService) Save(wallet *entity.RewardWallet) (result *entity.RewardWallet, err error) {
	return this.RewardWalletDao.Save(wallet)
}
func (this *RewardWalletService) FindRewardWalletByMemberIdAndCoinUnit(memberId int64, unit string) (result *entity.RewardWallet, err error) {
	return this.RewardWalletDao.FindRewardWalletByMemberIdAndCoinUnit(memberId, unit)
}
func (this *RewardWalletService) FindAllByMemberId(memberId int64) (result arraylist.List[entity.RewardWallet], err error) {
	return this.RewardWalletDao.FindAllByMemberId(memberId)
}
func (this *RewardWalletService) FindAll() (result arraylist.List[entity.RewardWallet], err error) {
	return this.RewardWalletDao.FindAll()
}
func NewRewardWalletService(rewardWalletDao *dao.RewardWalletDao) (ret *RewardWalletService) {
	ret = new(RewardWalletService)
	ret.RewardWalletDao = rewardWalletDao
	return
}

type RewardWalletService struct {
	RewardWalletDao dao.RewardWalletDao
	Base.BaseService
}
