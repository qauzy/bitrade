package service

func (this *OtcWalletService) FindByMemberId(memberId int64) (result arraylist.List[entity.OtcWallet], err error) {
	return this.OtcWalletDao.FindOtcWalletByMemberId(memberId)
}
func (this *OtcWalletService) FindByOtcCoinAndMemberId(memberId int64, coin *entity.OtcCoin) (result *entity.OtcWallet, err error) {
	var coin1 = this.CoinDao.FindByUnit(coin.GetUnit())
	return this.OtcWalletDao.FindByMemberIdAndCoin(memberId, coin1)
}
func (this *OtcWalletService) FindByCoinAndMember(memberId int64, coin *entity.Coin) (result *entity.OtcWallet, err error) {
	return this.OtcWalletDao.FindByMemberIdAndCoin(memberId, coin)
}
func (this *OtcWalletService) FindByCoinUnitAndMemberId(memberId int64, coinUnit string) (result *entity.OtcWallet, err error) {
	var coin = this.CoinDao.FindByUnit(coinUnit)
	return this.OtcWalletDao.FindByMemberIdAndCoin(memberId, coin)
}
func (this *OtcWalletService) Save(otcWallet *entity.OtcWallet) (result *entity.OtcWallet, err error) {
	return this.OtcWalletDao.Save(otcWallet)
}
func (this *OtcWalletService) AddWallet(id int64, amount *math.BigDecimal) (result int, err error) {
	return this.OtcWalletDao.AddWallet(id, amount)
}
func (this *OtcWalletService) Coin2Otc(memberWallet *entity.MemberWallet, otcWallet *entity.OtcWallet, amount *math.BigDecimal) (result int, err error) {
	//扣减币币账户
	this.MemberWalletService.DecreaseBalance(memberWallet.GetId(), amount)
	//增加法币账户
	var addResult = this.OtcWalletDao.AddWallet(otcWallet.GetId(), amount)
	if addResult == 0 {
		return 0
	}
	//增加划转记录
	var transaction = new(entity.MemberTransaction)
	transaction.SetAmount(amount)
	transaction.SetSymbol(memberWallet.GetCoin().GetName())
	transaction.SetAddress("")
	transaction.SetMemberId(memberWallet.GetMemberId())
	transaction.SetType(TransactionType.COIN_TWO_OTC)
	transaction.SetFee(decimal.Zero)
	transaction.SetCreateTime(time.Now())
	this.TransactionService.Save(transaction)
	return 1
}
func (this *OtcWalletService) Otc2Coin(memberWallet *entity.MemberWallet, otcWallet *entity.OtcWallet, amount *math.BigDecimal) (result int, err error) {
	//扣减法币账户
	var subResult = this.OtcWalletDao.SubWallet(otcWallet.GetId(), amount)
	if subResult == 0 {
		return 0
	}
	//增加币币账户
	this.MemberWalletService.IncreaseBalance(memberWallet.GetId(), amount)
	//增加划转记录
	var transaction = new(entity.MemberTransaction)
	transaction.SetAmount(amount)
	transaction.SetSymbol(memberWallet.GetCoin().GetName())
	transaction.SetAddress("")
	transaction.SetMemberId(memberWallet.GetMemberId())
	transaction.SetType(TransactionType.OTC_TWO_COIN)
	transaction.SetFee(decimal.Zero)
	transaction.SetCreateTime(time.Now())
	this.TransactionService.Save(transaction)
	return 1
}
func (this *OtcWalletService) ThawBalance(otcWallet *entity.OtcWallet, amount *math.BigDecimal) (result *util.MessageResult, err error) {
	var ret = this.OtcWalletDao.ThawBalance(otcWallet.GetId(), amount)
	if ret > 0 {
		return MessageResult.Success()
	} else {
		return MessageResult.Error("Information Expired")
	}
}
func (this *OtcWalletService) DecreaseFrozen(walletId int64, amount *math.BigDecimal) (err error) {
	this.OtcWalletDao.DecreaseFrozen(walletId, amount)
}
func (this *OtcWalletService) Transfer(order *entity.Order, ret int) (err error) {
	if ret == 1 {
		var customerWallet = this.FindByOtcCoinAndMemberId(order.GetCustomerId(), order.GetCoin())
		//卖方付出手续费
		var is = this.OtcWalletDao.DecreaseFrozen(customerWallet.GetId(), BigDecimalUtils.Add(order.GetNumber(), order.GetCommission()))
		if is > 0 {
			var memberWallet = this.FindByOtcCoinAndMemberId(order.GetMemberId(), order.GetCoin())
			//买房得到完整的币
			var a = this.OtcWalletDao.IncreaseBalance(memberWallet.GetId(), order.GetNumber())
			if a <= 0 {
				return InformationExpiredException("Information Expired")
			}
		} else {
			return InformationExpiredException("Information Expired")
		}
	} else {
		var customerWallet = this.FindByOtcCoinAndMemberId(order.GetMemberId(), order.GetCoin())
		//卖方付出手续费
		var is = this.OtcWalletDao.DecreaseFrozen(customerWallet.GetId(), BigDecimalUtils.Add(order.GetNumber(), order.GetCommission()))
		if is > 0 {
			//买方得到完整数量
			var memberWallet = this.FindByOtcCoinAndMemberId(order.GetCustomerId(), order.GetCoin())
			var a = this.OtcWalletDao.IncreaseBalance(memberWallet.GetId(), order.GetNumber())
			if a <= 0 {
				return InformationExpiredException("Information Expired")
			}
		} else {
			return InformationExpiredException("Information Expired")
		}
	}
}
func (this *OtcWalletService) FreezeBalance(otcWallet *entity.OtcWallet, amount *math.BigDecimal) (result *util.MessageResult, err error) {
	var ret = this.OtcWalletDao.FreezeBalance(otcWallet.GetId(), amount)
	if ret > 0 {
		return MessageResult.Success()
	} else {
		return MessageResult.Error("Information Expired")
	}
}
func (this *OtcWalletService) IncreaseBalance(walletId int64, amount *math.BigDecimal) (err error) {
	this.OtcWalletDao.IncreaseBalance(walletId, amount)
}
func (this *OtcWalletService) TransferAdmin(order *entity.Order, ret int) (err error) {
	if ret == 1 || ret == 4 {
		this.TrancerDetail(order, order.GetCustomerId(), order.GetMemberId())
	} else {
		this.TrancerDetail(order, order.GetMemberId(), order.GetCustomerId())
	}
}
func (this *OtcWalletService) TrancerDetail(order *entity.Order, sellerId int64, buyerId int64) (err error) {
	var customerWallet = this.FindByOtcCoinAndMemberId(sellerId, order.GetCoin())
	//卖币者，买币者要处理的金额
	var sellerAmount, buyerAmount *math.BigDecimal
	if order.GetMemberId() == sellerId {
		//广告商卖币
		sellerAmount = BigDecimalUtils.Add(order.GetNumber(), order.GetCommission())
		buyerAmount = order.GetNumber()
	} else {
		//客户卖币
		sellerAmount = BigDecimalUtils.Add(order.GetNumber(), order.GetCommission())
		buyerAmount = order.GetNumber()
	}
	var is = this.OtcWalletDao.DecreaseFrozen(customerWallet.GetId(), sellerAmount)
	if is > 0 {
		var memberWallet = this.FindByOtcCoinAndMemberId(buyerId, order.GetCoin())
		var a = this.OtcWalletDao.IncreaseBalance(memberWallet.GetId(), buyerAmount)
		if a <= 0 {
			return InformationExpiredException("Information Expired")
		}
	} else {
		return InformationExpiredException("Information Expired")
	}
}
func (this *OtcWalletService) LockWallet(uid int64, unit string) (result bool, err error) {
	var wallet = this.FindByCoinUnitAndMemberId(uid, unit)
	if wallet != nil && wallet.GetIsLock() == 0 {
		wallet.SetIsLock(1)
		return true
	} else {
		return false
	}
}
func (this *OtcWalletService) UnlockWallet(uid int64, unit string) (result bool, err error) {
	var wallet = this.FindByCoinUnitAndMemberId(uid, unit)
	if wallet != nil && wallet.GetIsLock() == 1 {
		wallet.SetIsLock(0)
		return true
	} else {
		return false
	}
}
func (this *OtcWalletService) JoinFind(predicates arraylist.List[types.Predicate], qMember QMember, qMemberWallet QOtcWallet, pageModel *PageModel.PageModel) (result domain.Page[dto.OtcWalletDTO], err error) {
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	predicates.Add(qMember.Id.Eq(qMemberWallet.MemberId))
	var query = queryFactory.Select(Projections.Fields(dto.OtcWalletDTO, qMemberWallet.Id.As("id"), qMemberWallet.MemberId.As("memberId"), qMember.Username, qMember.RealName.As("realName"), qMember.Email, qMember.MobilePhone.As("mobilePhone"), qMemberWallet.Balance, qMemberWallet.Coin.Unit, qMemberWallet.FrozenBalance.As("frozenBalance"), qMemberWallet.Balance.Add(qMemberWallet.FrozenBalance).As("allBalance"))).From(QMember.Member, QOtcWallet.OtcWallet).Where(predicates.ToArray(make([]Predicate, 0))).OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var content = query.Offset((pageModel.GetPageNo() - 1) * pageModel.GetPageSize()).Limit(pageModel.GetPageSize()).Fetch()
	var total = query.FetchCount()
	return PageImpl(content, pageModel.GetPageable(), total)
}
func NewOtcWalletService(otcWalletDao *dao.OtcWalletDao, memberWalletService MemberWalletService, transactionService MemberTransactionService, coinDao *dao.CoinDao) (ret *OtcWalletService) {
	ret = new(OtcWalletService)
	ret.OtcWalletDao = otcWalletDao
	ret.MemberWalletService = memberWalletService
	ret.TransactionService = transactionService
	ret.CoinDao = coinDao
	return
}

type OtcWalletService struct {
	OtcWalletDao        dao.OtcWalletDao
	MemberWalletService *MemberWalletService
	TransactionService  *MemberTransactionService
	CoinDao             dao.CoinDao
	Base.BaseService
}
