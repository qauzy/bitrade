package service

func (this *ReleaseBalanceService) Save(releaseBalance *entity.ReleaseBalance) (result *entity.ReleaseBalance, err error) {
	return this.ReleaseBalanceDao.Save(releaseBalance)
}
func (this *ReleaseBalanceService) UpdateReleaseBalance(releaseBalanceVO *vo.ReleaseBalanceVO) (result *util.MessageResult, err error) {
	var id = releaseBalanceVO.GetId()
	var releaseBalance *entity.ReleaseBalance
	for i := 0; i < id.Size(); i += 1 {
		releaseBalance = this.ReleaseBalanceDao.FindById(id.Get(i)).Get()
		if !"1".Equals(releaseBalance.GetReleaseState()) {
			var coinId = releaseBalance.GetCoinUnit()
			var memberId = releaseBalance.GetMemberId()
			// 根据传过来的  用户id 去找钱包
			var memberWallet = this.MemberWalletDao.FindCoinIdAndMemberId(coinId, memberId)
			// 获取释放余额.
			var rbMoney = releaseBalance.GetReleaseBalance()
			var mwMoney = memberWallet.GetReleaseBalance()
			if mwMoney.CompareTo(rbMoney) > -1 {
				//a大于等于b
				// 将获取的释放余额  加 到钱包余额里面
				var num = this.MemberWalletDao.ReleaseReisterGiving(rbMoney, memberWallet.GetId())
				//更改审核状态 1 - 已审核
				if num == 1 {
					releaseBalance.SetReleaseState("1")
					releaseBalance.SetReleaseTime(time.Now())
					this.ReleaseBalanceDao.Save(releaseBalance)
					//会员交易记录，包括充值、提现、转账等
					var memberTransaction = new(entity.MemberTransaction)
					memberTransaction.SetFee(decimal.Zero)
					memberTransaction.SetAmount(rbMoney)
					memberTransaction.SetSymbol(coinId)
					memberTransaction.SetType(TransactionType.ACTIVITY_AWARD)
					memberTransaction.SetMemberId(id.Get(i))
					this.MemberTransactionDao.Save(memberTransaction)
				}
			}
		}
	}
	return MessageResult.Success("审核已通过")
}
func (this *ReleaseBalanceService) FindByReleaseBalanceState(releaseBalanceVO *vo.ReleaseBalanceVO) (result domain.Page[entity.ReleaseBalance], err error) {
	// 查询条件
	var releaseBalance = new(pagination.Criteria)
	// 条件查询
	releaseBalance.Add(Restrictions.Eq("releaseState", releaseBalanceVO.GetReleaseState(), false))
	// 根据注册时间倒序
	var sorts = Criteria.SortStatic("registerTime.desc")
	// 分页
	var pageRequest = PageRequest.Of(releaseBalanceVO.GetPageNo()-1, releaseBalanceVO.GetPageSize(), sorts)
	// 返回
	return this.ReleaseBalanceDao.FindAll(releaseBalance, pageRequest)
}
func (this *ReleaseBalanceService) ConditionQueryAll(releaseBalanceVO *vo.ReleaseBalanceVO) (result domain.Page[entity.ReleaseBalance], err error) {
	var releaseBalance = new(pagination.Criteria)
	// 条件查询
	if StringUtils.IsNotBlank(releaseBalanceVO.GetMemberName()) {
		releaseBalance.Add(Restrictions.Eq("userName", releaseBalanceVO.GetMemberName(), false))
	}
	if StringUtils.IsNotBlank(releaseBalanceVO.GetPhone()) {
		releaseBalance.Add(Restrictions.Eq("phone", releaseBalanceVO.GetPhone(), false))
	}
	if StringUtils.IsNotBlank(releaseBalanceVO.GetReleaseState()) {
		releaseBalance.Add(Restrictions.Eq("releaseState", releaseBalanceVO.GetReleaseState(), false))
	}
	if releaseBalanceVO.GetStartTime() != nil {
		releaseBalance.Add(Restrictions.Gte("registerTime", releaseBalanceVO.GetStartTime(), false))
	}
	if releaseBalanceVO.GetEndTime() != nil {
		releaseBalance.Add(Restrictions.Lte("registerTime", releaseBalanceVO.GetEndTime(), false))
	}
	// 根据注册时间倒序
	var sorts = Criteria.SortStatic("registerTime.desc")
	var pageRequest = PageRequest.Of(releaseBalanceVO.GetPageNo()-1, releaseBalanceVO.GetPageSize(), sorts)
	return this.ReleaseBalanceDao.FindAll(releaseBalance, pageRequest)
}
func NewReleaseBalanceService(releaseBalanceDao *dao.ReleaseBalanceDao, memberWalletDao *dao.MemberWalletDao, memberTransactionDao *dao.MemberTransactionDao) (ret *ReleaseBalanceService) {
	ret = new(ReleaseBalanceService)
	ret.ReleaseBalanceDao = releaseBalanceDao
	ret.MemberWalletDao = memberWalletDao
	ret.MemberTransactionDao = memberTransactionDao
	return
}

type ReleaseBalanceService struct {
	ReleaseBalanceDao    dao.ReleaseBalanceDao
	MemberWalletDao      dao.MemberWalletDao
	MemberTransactionDao dao.MemberTransactionDao
	Base.BaseService
}
