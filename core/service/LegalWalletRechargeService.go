package service

func (this *LegalWalletRechargeService) SetDao(legalWalletRechargeDao *dao.LegalWalletRechargeDao) (err error) {
	super.SetDao(this.LegalWalletRechargeDao)
}
func (this *LegalWalletRechargeService) FindAllByMemberIdAndState(memberId int64, state *LegalWalletState.LegalWalletState, pageNo int, pageSize int) (result domain.Page[entity.LegalWalletRecharge], err error) {
	var predicate *types.Predicate
	if state != nil {
		predicate = QLegalWalletRecharge.LegalWalletRecharge.Member.Id.Eq(memberId).And(QLegalWalletRecharge.LegalWalletRecharge.State.Eq(state))
	}
	var sort = Sort.By(Sort.Order(Sort.Direction.DESC, "id"))
	var pageable = PageRequest.Of(pageNo-1, pageSize, sort)
	return this.LegalWalletRechargeDao.FindAll(predicate, pageable)
}
func (this *LegalWalletRechargeService) FindOneByIdAndMemberId(id int64, memberId int64) (result *entity.LegalWalletRecharge, err error) {
	var predicate = QLegalWalletRecharge.LegalWalletRecharge.Id.Eq(id).And(QLegalWalletRecharge.LegalWalletRecharge.Member.Id.Eq(memberId))
	return this.LegalWalletRechargeDao.FindOne(predicate).Get()
}
func (this *LegalWalletRechargeService) FindOne(id int64) (result *entity.LegalWalletRecharge, err error) {
	return this.LegalWalletRechargeDao.FindById(id).Get()
}
func (this *LegalWalletRechargeService) Save(legalWalletRecharge *entity.LegalWalletRecharge) (result *entity.LegalWalletRecharge, err error) {
	return this.LegalWalletRechargeDao.Save(legalWalletRecharge)
}
func (this *LegalWalletRechargeService) NoPass(legalWalletRecharge *entity.LegalWalletRecharge) (err error) {
	legalWalletRecharge.SetState(LegalWalletState.DEFEATED)
	//标记失败
	this.LegalWalletRechargeDao.Save(legalWalletRecharge)
}
func (this *LegalWalletRechargeService) Pass(wallet *entity.MemberWallet, legalWalletRecharge *entity.LegalWalletRecharge) (err error) {
	wallet.SetBalance(BigDecimalUtils.Add(wallet.GetBalance(), legalWalletRecharge.GetAmount()))
	//充值到账
	legalWalletRecharge.SetState(LegalWalletState.COMPLETE)
	//标记完成
}
func NewLegalWalletRechargeService(legalWalletRechargeDao *dao.LegalWalletRechargeDao) (ret *LegalWalletRechargeService) {
	ret = new(LegalWalletRechargeService)
	ret.LegalWalletRechargeDao = legalWalletRechargeDao
	return
}

type LegalWalletRechargeService struct {
	LegalWalletRechargeDao dao.LegalWalletRechargeDao
	Base.TopBaseService[entity.LegalWalletRecharge, dao.LegalWalletRechargeDao]
}
