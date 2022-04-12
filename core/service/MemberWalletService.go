package service

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/PageModel"
	"bitrade/core/constant/TransactionType"
	"bitrade/core/dao"
	"bitrade/core/dto"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"bitrade/core/util"
	"bitrade/core/vo"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
	"github.com/shopspring/decimal"
)

var Limit int = 1000

func (this *MemberWalletService) Save(wallet *entity.MemberWallet) (result *entity.MemberWallet, err error) {
	return this.MemberWalletDao.SaveAndFlush(wallet)
}
func (this *MemberWalletService) FindByOtcCoinAndMemberId(coin *entity.OtcCoin, memberId int64) (result *entity.MemberWallet, err error) {
	var coin1 = this.CoinDao.FindByUnit(coin.GetUnit())
	return this.MemberWalletDao.FindByCoinAndMemberId(coin1, memberId)
}
func (this *MemberWalletService) Recharge(wallet *entity.MemberWallet, amount *math.BigDecimal) (result *util.MessageResult, err error) {
	if wallet == nil {
		return MessageResult(500, "wallet cannot be null")
	}
	if amount.CompareTo(decimal.Zero) <= 0 {
		return MessageResult(500, "amount must large then 0")
	}
	var result = this.MemberWalletDao.IncreaseBalance(wallet.GetId(), amount)
	if result > 0 {
		var transaction = new(entity.MemberTransaction)
		transaction.SetAmount(amount)
		transaction.SetSymbol(wallet.GetCoin().GetUnit())
		transaction.SetAddress(wallet.GetAddress())
		transaction.SetMemberId(wallet.GetMemberId())
		transaction.SetType(TransactionType.RECHARGE)
		transaction.SetFee(decimal.Zero)
		this.TransactionService.Save(transaction)
		//增加记录
		return MessageResult(0, "success")
	} else {
		return MessageResult(500, "recharge failed")
	}
}
func (this *MemberWalletService) Recharge(coin *entity.Coin, address string, amount *math.BigDecimal, txid string) (result *util.MessageResult, err error) {
	var wallet = this.FindByCoinAndAddress(coin, address)
	if wallet == nil {
		return MessageResult(500, "wallet cannot be null")
	}
	if amount.CompareTo(decimal.Zero) <= 0 {
		return MessageResult(500, "amount must large then 0")
	}
	var deposit = new(entity.MemberDeposit)
	deposit.SetAddress(address)
	deposit.SetAmount(amount)
	deposit.SetMemberId(wallet.GetMemberId())
	deposit.SetTxid(txid)
	deposit.SetUnit(wallet.GetCoin().GetUnit())
	this.DepositDao.Save(deposit)
	wallet.SetBalance(wallet.GetBalance().Add(amount))
	var transaction = new(entity.MemberTransaction)
	transaction.SetAmount(amount)
	transaction.SetSymbol(wallet.GetCoin().GetUnit())
	transaction.SetAddress(wallet.GetAddress())
	transaction.SetMemberId(wallet.GetMemberId())
	transaction.SetType(TransactionType.RECHARGE)
	transaction.SetFee(decimal.Zero)
	this.TransactionService.Save(transaction)
	var messageResult = MessageResult(0, "success")
	messageResult.SetData(wallet.GetMemberId())
	return messageResult
}
func (this *MemberWalletService) FindByCoinAndAddress(coin *entity.Coin, address string) (result *entity.MemberWallet, err error) {
	return this.MemberWalletDao.FindByCoinAndAddress(coin, address)
}
func (this *MemberWalletService) FindByCoinAndMember(coin *entity.Coin, member *entity.Member) (result *entity.MemberWallet, err error) {
	return this.MemberWalletDao.FindByCoinAndMemberId(coin, member.GetId())
}
func (this *MemberWalletService) FindByCoinUnitAndMemberId(coinUnit string, memberId int64) (result *entity.MemberWallet, err error) {
	var coin = this.CoinDao.FindByUnit(coinUnit)
	return this.MemberWalletDao.FindByCoinAndMemberId(coin, memberId)
}
func (this *MemberWalletService) FindByCoinAndMemberId(coin *entity.Coin, memberId int64) (result *entity.MemberWallet, err error) {
	return this.MemberWalletDao.FindByCoinAndMemberId(coin, memberId)
}
func (this *MemberWalletService) FindAllByMemberId(member *entity.Member) (result arraylist.List[entity.MemberWallet], err error) {
	return this.MemberWalletDao.FindAllByMemberId(member.GetId())
}
func (this *MemberWalletService) FindAllByMemberId(memberId int64) (result arraylist.List[entity.MemberWallet], err error) {
	return this.MemberWalletDao.FindAllByMemberId(memberId)
}
func (this *MemberWalletService) FreezeBalance(memberWallet *entity.MemberWallet, amount *math.BigDecimal) (result *util.MessageResult, err error) {
	var ret = this.MemberWalletDao.FreezeBalance(memberWallet.GetId(), amount)
	if ret > 0 {
		return MessageResult.Success()
	} else {
		return MessageResult.Error("Information Expired")
	}
}
func (this *MemberWalletService) ThawBalance(memberWallet *entity.MemberWallet, amount *math.BigDecimal) (result *util.MessageResult, err error) {
	var ret = this.MemberWalletDao.ThawBalance(memberWallet.GetId(), amount)
	if ret > 0 {
		return MessageResult.Success()
	} else {
		return MessageResult.Error("Information Expired")
	}
}
func (this *MemberWalletService) Transfer(order *entity.Order, ret int) (err error) {
	if ret == 1 {
		var customerWallet = this.FindByOtcCoinAndMemberId(order.GetCoin(), order.GetCustomerId())
		//卖方付出手续费
		var is = this.MemberWalletDao.DecreaseFrozen(customerWallet.GetId(), BigDecimalUtils.Add(order.GetNumber(), order.GetCommission()))
		if is > 0 {
			var memberWallet = this.FindByOtcCoinAndMemberId(order.GetCoin(), order.GetMemberId())
			//买房得到完整的币
			var a = this.MemberWalletDao.IncreaseBalance(memberWallet.GetId(), order.GetNumber())
			if a <= 0 {
				return InformationExpiredException("Information Expired")
			}
		} else {
			return InformationExpiredException("Information Expired")
		}
	} else {
		var customerWallet = this.FindByOtcCoinAndMemberId(order.GetCoin(), order.GetMemberId())
		//卖方付出手续费
		var is = this.MemberWalletDao.DecreaseFrozen(customerWallet.GetId(), BigDecimalUtils.Add(order.GetNumber(), order.GetCommission()))
		if is > 0 {
			//买方得到完整数量
			var memberWallet = this.FindByOtcCoinAndMemberId(order.GetCoin(), order.GetCustomerId())
			var a = this.MemberWalletDao.IncreaseBalance(memberWallet.GetId(), order.GetNumber())
			if a <= 0 {
				return InformationExpiredException("Information Expired")
			}
		} else {
			return InformationExpiredException("Information Expired")
		}
	}
}
func (this *MemberWalletService) TransferAdmin(order *entity.Order, ret int) (err error) {
	if ret == 1 || ret == 4 {
		this.TrancerDetail(order, order.GetCustomerId(), order.GetMemberId())
	} else {
		this.TrancerDetail(order, order.GetMemberId(), order.GetCustomerId())
	}
}
func (this *MemberWalletService) TrancerDetail(order *entity.Order, sellerId int64, buyerId int64) (err error) {
	var customerWallet = this.FindByOtcCoinAndMemberId(order.GetCoin(), sellerId)
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
	var is = this.MemberWalletDao.DecreaseFrozen(customerWallet.GetId(), sellerAmount)
	if is > 0 {
		var memberWallet = this.FindByOtcCoinAndMemberId(order.GetCoin(), buyerId)
		var a = this.MemberWalletDao.IncreaseBalance(memberWallet.GetId(), buyerAmount)
		if a <= 0 {
			return InformationExpiredException("Information Expired")
		}
	} else {
		return InformationExpiredException("Information Expired")
	}
}
func (this *MemberWalletService) DeductBalance(memberWallet *entity.MemberWallet, amount *math.BigDecimal) (result int, err error) {
	return this.MemberWalletDao.DecreaseBalance(memberWallet.GetId(), amount)
}
func (this *MemberWalletService) FindAll() (result arraylist.List[entity.MemberWallet], err error) {
	return this.MemberWalletDao.FindAll()
}
func (this *MemberWalletService) FindAllByCoin(coin *entity.Coin) (result arraylist.List[entity.MemberWallet], err error) {
	return this.MemberWalletDao.FindAllByCoin(coin)
}
func (this *MemberWalletService) PageCount(coin *entity.Coin) (result int64, err error) {
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("coin", coin, true))
	return this.MemberWalletDao.Count(specification)
}
func (this *MemberWalletService) PageByCoin(coin *entity.Coin, pageNo int, pageSize int) (result domain.Page[entity.MemberWallet], err error) {
	var orders = Sort.By(Sort.Order(Sort.Direction.ASC, "id"))
	var pageRequest = PageRequest.Of(pageNo, pageSize, orders)
	var specification = new(pagination.Criteria)
	specification.Add(Restrictions.Eq("coin", coin, true))
	return this.MemberWalletDao.FindAll(specification, pageRequest)
}
func (this *MemberWalletService) LockWallet(uid int64, unit string) (result bool, err error) {
	var wallet = this.FindByCoinUnitAndMemberId(unit, uid)
	if wallet != nil && wallet.GetIsLock() == BooleanEnum.IS_FALSE {
		wallet.SetIsLock(BooleanEnum.IS_TRUE)
		return true
	} else {
		return false
	}
}
func (this *MemberWalletService) UnlockWallet(uid int64, unit string) (result bool, err error) {
	var wallet = this.FindByCoinUnitAndMemberId(unit, uid)
	if wallet != nil && wallet.GetIsLock() == BooleanEnum.IS_TRUE {
		wallet.SetIsLock(BooleanEnum.IS_FALSE)
		return true
	} else {
		return false
	}
}
func (this *MemberWalletService) FindOneByCoinNameAndMemberId(coinName string, memberId int64) (result *entity.MemberWallet, err error) {
	var and = QMemberWallet.MemberWallet.Coin.Name.Eq(coinName).And(QMemberWallet.MemberWallet.MemberId.Eq(memberId))
	return this.MemberWalletDao.FindOne(and).Get()
}
func (this *MemberWalletService) JoinFind(predicates arraylist.List[Predicate], qMember *entity.QMember, qMemberWallet *entity.QMemberWallet, pageModel *PageModel.PageModel) (result domain.Page[dto.MemberWalletDTO], err error) {
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	predicates.Add(qMember.Id.Eq(qMemberWallet.MemberId))
	var query = queryFactory.Select(Projections.Fields(dto.MemberWalletDTO, qMemberWallet.Id.As("id"), qMemberWallet.MemberId.As("memberId"), qMember.Username, qMember.RealName.As("realName"), qMember.Email, qMember.MobilePhone.As("mobilePhone"), qMemberWallet.Balance, qMemberWallet.Address, qMemberWallet.Coin.Unit, qMemberWallet.FrozenBalance.As("frozenBalance"), qMemberWallet.Balance.Add(qMemberWallet.FrozenBalance).As("allBalance"))).From(QMember.Member, QMemberWallet.MemberWallet).Where(predicates.ToArray(make([]Predicate, 0))).OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var content = query.Offset((pageModel.GetPageNo() - 1) * pageModel.GetPageSize()).Limit(pageModel.GetPageSize()).Fetch()
	var total = query.FetchCount()
	return PageImpl(content, pageModel.GetPageable(), total)
}
func (this *MemberWalletService) GetAllBalance(coinName string) (result *math.BigDecimal, err error) {
	return this.MemberWalletDao.GetWalletAllBalance(coinName)
}
func (this *MemberWalletService) FindDeposit(address string, txid string) (result *entity.MemberDeposit, err error) {
	return this.DepositDao.FindByAddressAndTxid(address, txid)
}
func (this *MemberWalletService) DecreaseFrozen(walletId int64, amount *math.BigDecimal) (err error) {
	this.MemberWalletDao.DecreaseFrozen(walletId, amount)
}
func (this *MemberWalletService) DecreaseBalance(walletId int64, amount *math.BigDecimal) (err error) {
	this.MemberWalletDao.DecreaseBalance(walletId, amount)
}
func (this *MemberWalletService) IncreaseBalance(walletId int64, amount *math.BigDecimal) (err error) {
	this.MemberWalletDao.IncreaseBalance(walletId, amount)
}
func (this *MemberWalletService) CreateGiftTable(times int64, coinId string) (result int, err error) {
	return this.MemberWalletDao.CreateGiftTable(times, coinId)
}
func (this *MemberWalletService) FindGiftTable(times int64) (result arraylist.List[entity.MemberWallet], err error) {
	return this.MemberWalletDao.FindGiftTable(times)
}
func (this *MemberWalletService) SumGiftTable(times int64) (result *math.BigDecimal, err error) {
	return this.MemberWalletDao.SumGiftTable(times)
}
func (this *MemberWalletService) HandleAirdrop(xmlVOS arraylist.List[vo.ImportXmlVO], airdrop int64) (result *util.MessageResult, err error) {
	var memberTransactionList = arraylist.New[entity.MemberTransaction]()
	var memberIdList = arraylist.New[int64]()
	for _, xmlVO := range xmlVOS {
		memberIdList.Add(xmlVO.GetMemberId())
	}
	//根据币种和memberId查询钱包，要求一个上传文件里都是空投同一个币种
	var coin = this.CoinDao.FindByUnit(xmlVOS.Get(0).GetCoinUnit())
	var memberWalletList = this.MemberWalletDao.FindALLByCoinIdAndMemberIdList(coin, memberIdList)
	if memberWalletList == nil || memberWalletList.Size() == 0 || memberWalletList.Size() != xmlVOS.Size() {
		return MessageResult.Error("钱包不存在或钱包数量不足")
	}
	//录入memberTransaction表，新增一个手动空投类型
	for _, xmlVO := range xmlVOS {
		for _, memberWallet := range memberWalletList {
			if xmlVO.GetMemberId().Equals(memberWallet.GetMemberId()) {
				memberWallet.SetBalance(memberWallet.GetBalance().Add(BigDecimal(xmlVO.GetAmount())))
			}
		}
		var transaction = new(entity.MemberTransaction)
		transaction.SetAmount(BigDecimal(xmlVO.GetAmount()))
		transaction.SetSymbol(xmlVO.GetCoinUnit())
		transaction.SetAddress("")
		transaction.SetMemberId(xmlVO.GetMemberId())
		transaction.SetType(TransactionType.MANUAL_AIRDROP)
		transaction.SetFee(decimal.Zero)
		transaction.SetAirdropId(airdrop)
		memberTransactionList.Add(transaction)
	}
	this.TransactionService.Save(memberTransactionList)
	this.MemberWalletDao.SaveAll(memberWalletList)
	return MessageResult.Success()
}
func NewMemberWalletService(memberWalletDao *dao.MemberWalletDao, coinDao *dao.CoinDao, transactionService MemberTransactionService, depositDao *dao.MemberDepositDao) (ret *MemberWalletService) {
	ret = new(MemberWalletService)
	ret.MemberWalletDao = memberWalletDao
	ret.CoinDao = coinDao
	ret.TransactionService = transactionService
	ret.DepositDao = depositDao
	return
}

type MemberWalletService struct {
	MemberWalletDao    dao.MemberWalletDao
	CoinDao            dao.CoinDao
	TransactionService *MemberTransactionService
	DepositDao         dao.MemberDepositDao
	Base.BaseService
}
