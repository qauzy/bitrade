package service

func (this *MemberTransactionService) QueryWhereOrPage(booleanExpressionList arraylist.List[dsl.BooleanExpression], pageNo int, pageSize int) (result pagination.PageResult[entity.MemberTransaction], err error) {
	var list arraylist.List[entity.MemberTransaction]
	var jpaQuery = queryFactory.SelectFrom(QMemberTransaction.MemberTransaction)
	if booleanExpressionList != nil {
		jpaQuery.Where(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *MemberTransactionService) Save(transaction *entity.MemberTransaction) (result *entity.MemberTransaction, err error) {
	return this.TransactionDao.SaveAndFlush(transaction)
}
func (this *MemberTransactionService) FindAll() (result arraylist.List[entity.MemberTransaction], err error) {
	return this.TransactionDao.FindAll()
}
func (this *MemberTransactionService) FindOne(id int64) (result *entity.MemberTransaction, err error) {
	return this.TransactionDao.FindById(id).Get()
}
func (this *MemberTransactionService) FindAllByWhere(startTime xtime.Xtime, endTime xtime.Xtime, oType *TransactionType.TransactionType, memberId int64) (result arraylist.List[interface {
}], err error) {
	var qMemberTransaction = QMemberTransaction.MemberTransaction
	var booleanExpressionList = arraylist.New[interface {
	}]()
	if startTime != nil {
		booleanExpressionList.Add(qMemberTransaction.CreateTime.Gt(startTime))
	}
	if endTime != nil {
		booleanExpressionList.Add(qMemberTransaction.CreateTime.Lt(endTime))
	}
	if oType != nil {
		booleanExpressionList.Add(qMemberTransaction.Type.Eq(oType))
	}
	if memberId != nil {
		booleanExpressionList.Add(qMemberTransaction.MemberId.Eq(memberId))
	}
	return queryFactory.SelectFrom(qMemberTransaction).Where(booleanExpressionList.ToArray(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))).Fetch()
}
func (this *MemberTransactionService) QueryByMember(uid int64, pageNo int, pageSize int, oType *TransactionType.TransactionType, unit string) (result domain.Page[entity.MemberTransaction], err error) {
	//排序方式 (需要倒序 这样    Criteria.sort("id","createTime.desc") ) //参数实体类为字段名
	var orders = Criteria.SortStatic("createTime.desc")
	//分页参数
	var pageRequest = PageRequest.Of(pageNo-1, pageSize, orders)
	//查询条件
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("memberId", uid, false))
	specification.Add(Restrictions.Eq("type", oType, false))
	if unit != nil {
		specification.Add(Restrictions.Eq("symbol", unit, false))
	}
	return this.TransactionDao.FindAll(specification, pageRequest)
}
func (this *MemberTransactionService) QueryByMember(uid int64, pageNo int, pageSize int, oType *TransactionType.TransactionType, startDate string, endDate string, symbol string) (result domain.Page[entity.MemberTransaction], err error) {
	//排序方式 (需要倒序 这样    Criteria.sort("id","createTime.desc") ) //参数实体类为字段名
	var orders = Criteria.SortStatic("createTime.desc")
	//分页参数
	var pageRequest = PageRequest.Of(pageNo-1, pageSize, orders)
	//查询条件
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("memberId", uid, false))
	if oType != nil {
		specification.Add(Restrictions.Eq("type", oType, false))
	}
	if StringUtils.IsNotBlank(startDate) && StringUtils.IsNotBlank(endDate) {
		specification.Add(Restrictions.Gte("createTime", DateUtil.YYYY_MM_DD_MM_HH_SS.Parse(startDate), false))
		specification.Add(Restrictions.Lte("createTime", DateUtil.YYYY_MM_DD_MM_HH_SS.Parse(endDate), false))
	}
	if StringUtils.IsNotBlank(symbol) {
		specification.Add(Restrictions.Eq("symbol", symbol, false))
	}
	return this.TransactionDao.FindAll(specification, pageRequest)
}
func (this *MemberTransactionService) FindTransactionSum(uid int64, oType *TransactionType.TransactionType) (result *math.BigDecimal, err error) {
	var types = arraylist.New[TransactionType.TransactionType]()
	types.Add(oType)
	var result = this.TransactionDao.FindTransactionSum(uid, types)
	if result == nil || !result.ContainsKey("amount") {
		return decimal.Zero
	} else {
		return BigDecimal(result.Get("amount").ToString())
	}
}
func (this *MemberTransactionService) FindMatchTransaction(uid int64, symbol string) (result arraylist.List[entity.MemberTransaction], err error) {
	var orders = Criteria.SortStatic("createTime.asc")
	//查询条件
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("memberId", uid, false))
	specification.Add(Restrictions.Eq("flag", 0, false))
	specification.Add(Restrictions.Eq("symbol", symbol, false))
	specification.Add(Restrictions.Gt("amount", 0, false))
	var types = arraylist.New[TransactionType.TransactionType]()
	types.Add(TransactionType.RECHARGE)
	types.Add(TransactionType.EXCHANGE)
	types.Add(TransactionType.ADMIN_RECHARGE)
	specification.Add(Restrictions.In("type", types, false))
	var transactions = this.TransactionDao.FindAll(specification, orders)
	return transactions
}
func (this *MemberTransactionService) MatchWallet(uid int64, symbol string, amount *math.BigDecimal) (err error) {
	var transactions = this.FindMatchTransaction(uid, symbol)
	var deltaAmount = decimal.Zero
	var gccWallet = this.WalletService.FindByCoinUnitAndMemberId("GCC", uid)
	var gcxWallet = this.WalletService.FindByCoinUnitAndMemberId("GCX", uid)
	for _, transaction := range transactions {
		if amount.CompareTo(deltaAmount) > 0 {
			var amt *math.BigDecimal
			if amount.Subtract(deltaAmount).CompareTo(transaction.GetAmount()) > 0 {
				amt = transaction.GetAmount()
			} else {
				amt = amount.Subtract(deltaAmount)
			}
			deltaAmount = deltaAmount.Add(amt)
			transaction.SetFlag(1)
		} else {
			break
		}
	}
	gccWallet.SetBalance(gccWallet.GetBalance().Subtract(deltaAmount))
	gcxWallet.SetBalance(gcxWallet.GetBalance().Add(deltaAmount))
	var transaction = new(entity.MemberTransaction)
	transaction.SetAmount(deltaAmount)
	transaction.SetSymbol(gcxWallet.GetCoin().GetUnit())
	transaction.SetAddress(gcxWallet.GetAddress())
	transaction.SetMemberId(gcxWallet.GetMemberId())
	transaction.SetType(TransactionType.MATCH)
	transaction.SetFee(decimal.Zero)
	//保存配对记录
	this.Save(transaction)
	if gccWallet.GetBalance().CompareTo(decimal.Zero) < 0 {
		gccWallet.SetBalance(decimal.Zero)
	}
}
func (this *MemberTransactionService) IsOverMatchLimit(day string, limit float64) (result bool, err error) {
	var totalAmount *math.BigDecimal
	var date1 = DateUtil.YYYY_MM_DD_MM_HH_SS.Parse(day + " 00:00:00")
	var date2 = DateUtil.YYYY_MM_DD_MM_HH_SS.Parse(day + " 23:59:59")
	var result = this.TransactionDao.FindMatchTransactionSum("GCX", TransactionType.MATCH, date1, date2)
	if result != nil && result.ContainsKey("amount") {
		totalAmount = BigDecimal(result.Get("amount").ToString())
	} else {
		totalAmount = decimal.Zero
	}
	log.Info("totalAmount:" + totalAmount)
	return totalAmount.DoubleValue() >= limit
}
func (this *MemberTransactionService) FindMemberDailyMatch(uid int64, day string) (result *math.BigDecimal, err error) {
	var date1 = DateUtil.YYYY_MM_DD_MM_HH_SS.Parse(day + " 00:00:00")
	var date2 = DateUtil.YYYY_MM_DD_MM_HH_SS.Parse(day + " 23:59:59")
	var result = this.TransactionDao.FindMatchTransactionSum(uid, "GCX", TransactionType.MATCH, date1, date2)
	if result != nil && result.ContainsKey("amount") {
		return BigDecimal(result.Get("amount").ToString())
	} else {
		return decimal.Zero
	}
}
func (this *MemberTransactionService) JoinFind(predicates arraylist.List[types.Predicate], pageModel *PageModel.PageModel) (result domain.Page[vo.MemberTransactionVO], err error) {
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	var query = queryFactory.Select(Projections.Fields(vo.MemberTransactionVO, QMemberTransaction.MemberTransaction.Address, QMemberTransaction.MemberTransaction.Amount, QMemberTransaction.MemberTransaction.CreateTime.As("createTime"), QMemberTransaction.MemberTransaction.Fee, QMemberTransaction.MemberTransaction.Flag, QMemberTransaction.MemberTransaction.Id.As("id"), QMemberTransaction.MemberTransaction.Symbol, QMemberTransaction.MemberTransaction.Type, QMember.Member.Username.As("memberUsername"), QMember.Member.MobilePhone.As("phone"), QMember.Member.Email, QMember.Member.RealName.As("memberRealName"), QMember.Member.Id.As("memberId"))).From(QMemberTransaction.MemberTransaction, QMember.Member)
	predicates.Add(QMemberTransaction.MemberTransaction.MemberId.Eq(QMember.Member.Id))
	query.Where(predicates.ToArray(make([]BooleanExpression, 0)))
	query.OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var list = query.Offset((pageModel.GetPageNo() - 1) * pageModel.GetPageSize()).Limit(pageModel.GetPageSize()).Fetch()
	var total = query.FetchCount()
	return PageImpl(list, pageModel.GetPageable(), total)
}
func (this *MemberTransactionService) Save(list arraylist.List[entity.MemberTransaction]) (err error) {
	this.TransactionDao.SaveAll(list)
}
func (this *MemberTransactionService) FindAllByCreateTime(beginDate string, endDate string) (result arraylist.List[entity.MemberTransaction], err error) {
	return this.TransactionDao.FindAllByCreateTime(beginDate, endDate)
}
func NewMemberTransactionService(transactionDao *dao.MemberTransactionDao, walletService MemberWalletService) (ret *MemberTransactionService) {
	ret = new(MemberTransactionService)
	ret.TransactionDao = transactionDao
	ret.WalletService = walletService
	return
}

type MemberTransactionService struct {
	TransactionDao dao.MemberTransactionDao
	WalletService  *MemberWalletService
	Base.BaseService
}
