package service

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/PageModel"
	"bitrade/core/constant/TransactionType"
	"bitrade/core/constant/WithdrawStatus"
	"bitrade/core/dao"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"bitrade/core/vo"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/chocolate/xtime"
	"time"
)

func (this *WithdrawRecordService) Save(withdrawApply *entity.WithdrawRecord) (result *entity.WithdrawRecord, err error) {
	return this.WithdrawApplyDao.Save(withdrawApply)
}
func (this *WithdrawRecordService) FindAll() (result arraylist.List[entity.WithdrawRecord], err error) {
	return this.WithdrawApplyDao.FindAll()
}
func (this *WithdrawRecordService) FindOne(id int64) (result *entity.WithdrawRecord, err error) {
	return this.WithdrawApplyDao.FindById(id).Get()
}
func (this *WithdrawRecordService) Query(predicateList arraylist.List[Predicate], pageNo int, pageSize int) (result pagination.PageResult[entity.WithdrawRecord], err error) {
	var list arraylist.List[entity.WithdrawRecord]
	var jpaQuery = queryFactory.SelectFrom(withdrawRecord)
	if predicateList != nil {
		jpaQuery.Where(predicateList.ToArray(make([]Predicate, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *WithdrawRecordService) Test() (err error) {
	//查询字段
	var expressions = arraylist.New[Expression]()
	expressions.Add(QWithdrawRecord.WithdrawRecord.MemberId.As("memberId"))
	//查询表
	var entityPaths = arraylist.New[EntityPath]()
	entityPaths.Add(QWithdrawRecord.WithdrawRecord)
	entityPaths.Add(QMember.Member)
	//查询条件
	var predicates = arraylist.New[Predicate]()
	predicates.Add(QWithdrawRecord.WithdrawRecord.MemberId.Eq(QMember.Member.Id))
	//排序
	var orderSpecifierList = arraylist.New[OrderSpecifier]()
	orderSpecifierList.Add(QWithdrawRecord.WithdrawRecord.Id.Desc())
	var pageListMapResult = super.QueryDslForPageListResult(expressions, entityPaths, predicates, orderSpecifierList, 1, 10)
	this.Logger.Info(pageListMapResult.ToString())
}
func (this *WithdrawRecordService) Audit(ids []int64, status *WithdrawStatus.WithdrawStatus) (err error) {
	var withdrawRecord *entity.WithdrawRecord
	for _, id := range ids {
		//20 4.70000000 0 2018-02-27 17:47:37  0.30000000 0 28 0 5.00000000   GalaxyChain
		withdrawRecord = this.WithdrawApplyDao.FindById(id).Get()
		//确认提现申请存在
		notNull(withdrawRecord, "不存在")
		//确认订单状态是审核中
		isTrue(withdrawRecord.GetStatus() == PROCESSING, "id为"+id+"不是审核状态的提现")
		//确认提现类型不是自动提现
		isTrue(withdrawRecord.GetIsAuto() == IS_FALSE, "id为"+id+"不是人工审核提现")
		//审核
		if status == FAIL {
			//审核不通过
			var wallet = this.WalletService.FindByCoinAndMemberId(withdrawRecord.GetCoin(), withdrawRecord.GetMemberId())
			notNull(wallet, "wallet null!")
			wallet.SetBalance(wallet.GetBalance().Add(withdrawRecord.GetTotalAmount()))
			wallet.SetFrozenBalance(wallet.GetFrozenBalance().Subtract(withdrawRecord.GetTotalAmount()))
			this.WalletService.Save(wallet)
		}
		withdrawRecord.SetStatus(status)
		withdrawRecord.SetDealTime(time.Now())
		this.WithdrawApplyDao.Save(withdrawRecord)
	}
}
func (this *WithdrawRecordService) WithdrawSuccess(withdrawId int64, txid string) (err error) {
	var record = this.FindOne(withdrawId)
	if record != nil {
		record.SetTransactionNumber(txid)
		record.SetStatus(WithdrawStatus.SUCCESS)
		var wallet = this.WalletService.FindByCoinUnitAndMemberId(record.GetCoin().GetUnit(), record.GetMemberId())
		if wallet != nil {
			wallet.SetFrozenBalance(wallet.GetFrozenBalance().Subtract(record.GetTotalAmount()))
			var transaction = new(entity.MemberTransaction)
			transaction.SetAmount(record.GetTotalAmount().Negate())
			transaction.SetSymbol(wallet.GetCoin().GetUnit())
			transaction.SetAddress(wallet.GetAddress())
			transaction.SetMemberId(wallet.GetMemberId())
			transaction.SetType(TransactionType.WITHDRAW)
			transaction.SetFee(record.GetFee().Negate())
			this.TransactionService.Save(transaction)
		}
	}
}
func (this *WithdrawRecordService) WithdrawTransfering(withdrawId int64) (err error) {
	var record = this.FindOne(withdrawId)
	if record != nil {
		record.SetStatus(WithdrawStatus.TRANSFER)
	}
}
func (this *WithdrawRecordService) WithdrawFail(withdrawId int64) (err error) {
	var record = this.FindOne(withdrawId)
	if record == nil || record.GetStatus() != WithdrawStatus.PROCESSING {
		return
	}
	var wallet = this.WalletService.FindByCoinAndAddress(record.GetCoin(), record.GetAddress())
	if wallet != nil {
		wallet.SetBalance(wallet.GetBalance().Add(record.GetTotalAmount()))
		wallet.SetFrozenBalance(wallet.GetFrozenBalance().Subtract(record.GetTotalAmount()))
		record.SetStatus(WithdrawStatus.FAIL)
	}
}
func (this *WithdrawRecordService) AutoWithdrawFail(withdrawId int64) (err error) {
	var record = this.FindOne(withdrawId)
	if record == nil || record.GetStatus() != WithdrawStatus.WAITING {
		return
	}
	this.Logger.Info("================  自动转币失败，转为人工处理  ========================")
	this.Logger.Info("================ setIsAuto : " + BooleanEnum.IS_FALSE.GetNameCn() + "  ========================")
	this.Logger.Info("================ setStatus : " + WithdrawStatus.PROCESSING.GetCnName() + "  ========================")
	record.SetIsAuto(BooleanEnum.IS_FALSE)
	record.SetStatus(WithdrawStatus.PROCESSING)
	this.WithdrawApplyDao.Save(record)
}
func (this *WithdrawRecordService) FindAllByMemberId(memberId int64, page int, pageSize int) (result domain.Page[entity.WithdrawRecord], err error) {
	var orders = Criteria.SortStatic("id.desc")
	var pageRequest = PageRequest.Of(page, pageSize, orders)
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("memberId", memberId, false))
	return this.WithdrawApplyDao.FindAll(specification, pageRequest)
}
func (this *WithdrawRecordService) JoinFind(predicates arraylist.List[Predicate], pageModel *PageModel.PageModel) (result domain.Page[vo.WithdrawRecordVO], err error) {
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	var Query = queryFactory.Select(Projections.Fields(vo.WithdrawRecordVO, QWithdrawRecord.WithdrawRecord.Id.As("id"), QWithdrawRecord.WithdrawRecord.MemberId.As("memberId"), QWithdrawRecord.WithdrawRecord.Coin.Unit, QMember.Member.Username.As("memberUsername"), QMember.Member.RealName.As("memberRealName"), QMember.Member.MobilePhone.As("phone"), QMember.Member.Email, QWithdrawRecord.WithdrawRecord.DealTime.As("dealTime"), QWithdrawRecord.WithdrawRecord.TotalAmount.As("totalAmount"), QWithdrawRecord.WithdrawRecord.ArrivedAmount.As("arrivedAmount"), QWithdrawRecord.WithdrawRecord.Status, QWithdrawRecord.WithdrawRecord.IsAuto.As("isAuto"), QWithdrawRecord.WithdrawRecord.Address, QWithdrawRecord.WithdrawRecord.CreateTime.As("createTime"), QWithdrawRecord.WithdrawRecord.Fee, QWithdrawRecord.WithdrawRecord.TransactionNumber.As("transactionNumber"), QWithdrawRecord.WithdrawRecord.Remark)).From(QWithdrawRecord.WithdrawRecord, QMember.Member).Where(predicates.ToArray(make([]BooleanExpression, 0)))
	this.Query.OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var list = this.Query.Offset((pageModel.GetPageNo() - 1) * pageModel.GetPageSize()).Limit(pageModel.GetPageSize()).Fetch()
	var total = this.Query.FetchCount()
	return PageImpl(list, pageModel.GetPageable(), total)
}
func (this *WithdrawRecordService) CountAuditing() (result int64, err error) {
	return this.WithdrawApplyDao.CountAllByStatus(WithdrawStatus.PROCESSING)
}
func (this *WithdrawRecordService) FindByIds(ids []int64) (result arraylist.List[entity.WithdrawRecord], err error) {
	return this.WithdrawApplyDao.FindByIds(ids)
}
func (this *WithdrawRecordService) CountWithdrawAmountByTimeAndMemberIdAndCoin(startTime xtime.Xtime, endTime xtime.Xtime, coin *entity.Coin) (result float64, err error) {
	return this.WithdrawApplyDao.CountWithdrawAmountByTimeAndMemberIdAndCoin(startTime, endTime, coin)
}
func (this *WithdrawRecordService) FindAll(predicate Predicate, pageModel *PageModel.PageModel) (result domain.Page[entity.WithdrawRecord], err error) {
	return this.WithdrawApplyDao.FindAll(predicate, pageModel.GetPageable())
}
func NewWithdrawRecordService(logger *slf4j.Logger, withdrawApplyDao *dao.WithdrawRecordDao, walletService MemberWalletService, transactionService MemberTransactionService) (ret *WithdrawRecordService) {
	ret = new(WithdrawRecordService)
	ret.Logger = logger
	ret.WithdrawApplyDao = withdrawApplyDao
	ret.WalletService = walletService
	ret.TransactionService = transactionService
	return
}

type WithdrawRecordService struct {
	Logger             slf4j.Logger
	WithdrawApplyDao   dao.WithdrawRecordDao
	WalletService      *MemberWalletService
	TransactionService *MemberTransactionService
	Base.BaseService
}
