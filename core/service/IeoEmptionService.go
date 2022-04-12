package service

func (this *IeoEmptionService) GetByPage(ieoEmptionVO *vo.IeoEmptionVO) (result domain.Page[entity.IeoEmption], err error) {
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if StringUtils.IsNotEmpty(ieoEmptionVO.GetStartTime()) {
		booleanExpressions.Add(QIeoEmption.IeoEmption.CreateTime.Goe(DateUtil.StringToDate(ieoEmptionVO.GetStartTime())))
	}
	if StringUtils.IsNotEmpty(ieoEmptionVO.GetEndTime()) {
		booleanExpressions.Add(QIeoEmption.IeoEmption.CreateTime.Loe(DateUtil.StringToDate(ieoEmptionVO.GetEndTime())))
	}
	if StringUtils.IsNotEmpty(ieoEmptionVO.GetStatus()) {
		var date = time.Now()
		var status = ieoEmptionVO.GetStatus()
		if "1".Equals(status) {
			booleanExpressions.Add(QIeoEmption.IeoEmption.StartTime.Gt(date))
		} else if "2".Equals(status) {
			booleanExpressions.Add(QIeoEmption.IeoEmption.StartTime.Lt(date))
			booleanExpressions.Add(QIeoEmption.IeoEmption.EndTime.Goe(date))
		} else if "3".Equals(status) {
			booleanExpressions.Add(QIeoEmption.IeoEmption.EndTime.Lt(date))
		}
	}
	if ieoEmptionVO.GetId() != nil {
		booleanExpressions.Add(QIeoEmption.IeoEmption.Id.Eq(ieoEmptionVO.GetId()))
	}
	var predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var sort = Sort.By(Sort.Direction.DESC, "id")
	var pageable = PageRequest.Of(ieoEmptionVO.GetPageNum()-1, ieoEmptionVO.GetPageSize(), sort)
	return this.IeoEmptionDao.FindAll(predicate, pageable)
}
func (this *IeoEmptionService) Save(ieoEmption *entity.IeoEmption) (result *entity.IeoEmption, err error) {
	return this.IeoEmptionDao.Save(ieoEmption)
}
func (this *IeoEmptionService) FindById(id int64) (result *entity.IeoEmption, err error) {
	return this.IeoEmptionDao.FindById(id).Get()
}
func (this *IeoEmptionService) Del(id int64) (err error) {
	this.IeoEmptionDao.DeleteById(id)
}
func (this *IeoEmptionService) FindbyCondition(id int64, time string) (result *entity.IeoEmption, err error) {
	return this.IeoEmptionDao.FindbyCondition(id, time)
}
func (this *IeoEmptionService) SubAmount(amount *math.BigDecimal, ieoEmption *entity.IeoEmption, userId int64) (result int, err error) {
	var receAmount = amount.Multiply(ieoEmption.GetRatio()).SetScale(4, BigDecimal.ROUND_DOWN)
	//减少库存
	var subResult = this.IeoEmptionDao.SubAmount(ieoEmption.GetId(), receAmount)
	if subResult == 1 {
		var raiseCoin = this.CoinService.FindByUnit(ieoEmption.GetRaiseCoin())
		//扣减用户募集币种
		var raiseWallet = this.MemberWalletDao.FindByCoinAndMemberId(raiseCoin, userId)
		var subWallet = this.MemberWalletDao.DecreaseBalance(raiseWallet.GetId(), amount)
		if subWallet == 1 {
			//增加用户认购币种
			var saleCoin = this.CoinService.FindByUnit(ieoEmption.GetSaleCoin())
			var saleWallet = this.MemberWalletDao.FindByCoinAndMemberId(saleCoin, userId)
			if saleWallet == nil {
				var insertResult = this.MemberService.SaveWallet(ieoEmption.GetSaleCoin(), userId, receAmount)
				if insertResult == 1 {
					return 1
				} else {
					return Exception("增加用户余额异常：" + userId + "，coin:" + ieoEmption.GetSaleCoin())
				}
			} else {
				var addResult = this.MemberWalletDao.IncreaseBalance(saleWallet.GetId(), receAmount)
				if addResult == 1 {
					return 1
				} else {
					return Exception("增加用户余额异常：" + userId + "，coin:" + ieoEmption.GetSaleCoin())
				}
			}
		} else {
			return Exception("扣减用户余额异常：" + userId + "，coin:" + ieoEmption.GetRaiseCoin())
		}
	} else {
		return 0
	}
}
func NewIeoEmptionService(ieoEmptionDao *dao.IeoEmptionDao, memberWalletDao *dao.MemberWalletDao, coinService CoinService, memberService MemberService) (ret *IeoEmptionService) {
	ret = new(IeoEmptionService)
	ret.IeoEmptionDao = ieoEmptionDao
	ret.MemberWalletDao = memberWalletDao
	ret.CoinService = coinService
	ret.MemberService = memberService
	return
}

type IeoEmptionService struct {
	IeoEmptionDao   dao.IeoEmptionDao
	MemberWalletDao dao.MemberWalletDao
	CoinService     *CoinService
	MemberService   *MemberService
	Base.BaseService
}
