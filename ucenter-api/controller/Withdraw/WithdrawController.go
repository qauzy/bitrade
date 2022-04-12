package Withdraw

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/constant/PageModel"
	"bitrade/core/constant/WithdrawStatus"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/MessageResult"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/math"
	"github.com/shopspring/decimal"
	"time"
)

func (this *WithdrawController) AddAddress(ctx *gin.Context, address string, unit string, remark string, code string, aims string, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	hasText(address, this.SourceService.GetMessage("MISSING_COIN_ADDRESS"))
	hasText(unit, this.SourceService.GetMessage("MISSING_COIN_TYPE"))
	hasText(code, this.SourceService.GetMessage("MISSING_VERIFICATION_CODE"))
	hasText(aims, this.SourceService.GetMessage("MISSING_PHONE_OR_EMAIL"))
	var coin = this.CoinService.FindByUnit(unit)
	var memberAddress = this.MemberAddressService.FindByMemberIdAndCoinAndAddress(user.GetId(), coin, address, CommonStatus.NORMAL)
	if memberAddress != nil && memberAddress.Size() > 0 {
		return error("该地址已经存在，请确认地址")
	}
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() != nil && aims.Equals(member.GetMobilePhone()) {
		var info = this.RedisUtil.Get(SysConstant.PHONE_ADD_ADDRESS_PREFIX + member.GetMobilePhone())
		if info == nil {
			return error(this.SourceService.GetMessage("VERIFICATION_CODE_NOT_EXISTS"))
		}
		if !info.ToString().Equals(code) {
			return error(this.SourceService.GetMessage("VERIFICATION_CODE_INCORRECT"))
		} else {
			this.RedisUtil.Delete(SysConstant.PHONE_ADD_ADDRESS_PREFIX + member.GetMobilePhone())
		}
	} else if member.GetEmail() != nil && aims.Equals(member.GetEmail()) {
		var info = this.RedisUtil.Get(SysConstant.ADD_ADDRESS_CODE_PREFIX + member.GetEmail())
		if info == nil {
			return error(this.SourceService.GetMessage("VERIFICATION_CODE_NOT_EXISTS"))
		}
		if !info.ToString().Equals(code) {
			return error(this.SourceService.GetMessage("VERIFICATION_CODE_INCORRECT"))
		} else {
			this.RedisUtil.Delete(SysConstant.ADD_ADDRESS_CODE_PREFIX + member.GetEmail())
		}
	} else {
		return error(this.SourceService.GetMessage("ADD_ADDRESS_FAILED"))
	}
	var result = this.MemberAddressService.AddMemberAddress(user.GetId(), address, unit, remark)
	if result.GetCode() == 0 {
		result.SetMessage(this.SourceService.GetMessage("ADD_ADDRESS_SUCCESS"))
	} else if result.GetCode() == 500 {
		result.SetMessage(this.SourceService.GetMessage("ADD_ADDRESS_FAILED"))
	} else if result.GetCode() == 600 {
		result.SetMessage(this.SourceService.GetMessage("COIN_NOT_SUPPORT"))
	}
	return result
}
func (this *WithdrawController) DeleteAddress(ctx *gin.Context, id int64, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var result = this.MemberAddressService.DeleteMemberAddress(user.GetId(), id)
	if result.GetCode() == 0 {
		result.SetMessage(this.SourceService.GetMessage("DELETE_ADDRESS_SUCCESS"))
	} else {
		result.SetMessage(this.SourceService.GetMessage("DELETE_ADDRESS_FAILED"))
	}
	return result
}
func (this *WithdrawController) AddressPage(ctx *gin.Context, user *transform.AuthMember, pageNo int, pageSize int, unit string) (result *MessageResult.MessageResult) {
	var page = this.MemberAddressService.PageQuery(pageNo, pageSize, user.GetId(), unit)
	var scanMemberAddresses = page.Map(func(x interface {
	}) {
		ScanMemberAddress.ToScanMemberAddress(x)
	})
	var result = MessageResult.Success()
	result.SetData(scanMemberAddresses)
	return result
}
func (this *WithdrawController) QueryWithdraw(ctx *gin.Context) (result *MessageResult.MessageResult) {
	var list = this.CoinService.FindAllCanWithDraw()
	var list1 = arraylist.New[string]()
	list.Stream().ForEach(func(x interface {
	}) {
		list1.Add(x.GetUnit())
	})
	var result = MessageResult.Success()
	result.SetData(list1)
	return result
}
func (this *WithdrawController) QueryAllCoin(ctx *gin.Context) (result *MessageResult.MessageResult) {
	var list = this.CoinService.FindAllByStatus(CommonStatus.NORMAL)
	var list1 = arraylist.New[string]()
	list.Stream().ForEach(func(x interface {
	}) {
		list1.Add(x.GetUnit())
	})
	var result = MessageResult.Success()
	result.SetData(list1)
	return result
}
func (this *WithdrawController) QueryWithdrawCoin(ctx *gin.Context, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var list = this.CoinService.FindAllCanWithDraw()
	var list1 = this.MemberWalletService.FindAllByMemberId(user.GetId())
	var id = user.GetId()
	var list2 = list1.Stream().Filter(func(x interface {
	}) {
		list.Contains(x.GetCoin())
	}).Map(func(x interface {
	}) {
		new(WithdrawWalletInfo).SetBalance(x.GetBalance()).SetWithdrawScale(x.GetCoin().GetWithdrawScale()).SetMaxTxFee(x.GetCoin().GetMaxTxFee()).SetMinTxFee(x.GetCoin().GetMinTxFee()).SetMinAmount(x.GetCoin().GetMinWithdrawAmount()).SetMaxAmount(x.GetCoin().GetMaxWithdrawAmount()).SetName(x.GetCoin().GetName()).SetNameCn(x.GetCoin().GetNameCn()).SetThreshold(x.GetCoin().GetWithdrawThreshold()).SetUnit(x.GetCoin().GetUnit()).SetCanAutoWithdraw(x.GetCoin().GetCanAutoWithdraw()).SetAddresses(this.MemberAddressService.QueryAddress(id, x.GetCoin().GetName()))
	}).Collect(Collectors.ToList())
	var result = MessageResult.Success()
	result.SetData(list2)
	return result
}
func (this *WithdrawController) Withdraw(ctx *gin.Context, user *transform.AuthMember, unit string, address string, amount *math.BigDecimal, fee *math.BigDecimal, remark string, jyPassword string, code string, googleCode string) (result *MessageResult.MessageResult, err error) {
	hasText(jyPassword, this.SourceService.GetMessage("MISSING_JYPASSWORD"))
	hasText(unit, this.SourceService.GetMessage("MISSING_COIN_TYPE"))
	var coin = this.CoinService.FindByUnit(unit)
	amount.SetScale(coin.GetWithdrawScale(), BigDecimal.ROUND_DOWN)
	notNull(coin, this.SourceService.GetMessage("COIN_ILLEGAL"))
	isTrue(coin.GetStatus().Equals(CommonStatus.NORMAL) && coin.GetCanWithdraw().Equals(BooleanEnum.IS_TRUE), this.SourceService.GetMessage("COIN_NOT_SUPPORT"))
	isTrue(compare(fee, BigDecimal(String.ValueOf(coin.GetMinTxFee()))), this.SourceService.GetMessage("CHARGE_MIN")+coin.GetMinTxFee())
	isTrue(compare(BigDecimal(String.ValueOf(coin.GetMaxTxFee())), fee), this.SourceService.GetMessage("CHARGE_MAX")+coin.GetMaxTxFee())
	isTrue(compare(coin.GetMaxWithdrawAmount(), amount), this.SourceService.GetMessage("WITHDRAW_MAX")+coin.GetMaxWithdrawAmount())
	isTrue(compare(amount, coin.GetMinWithdrawAmount()), this.SourceService.GetMessage("WITHDRAW_MIN")+coin.GetMinWithdrawAmount())
	var memberWallet = this.MemberWalletService.FindByCoinAndMemberId(coin, user.GetId())
	isTrue(compare(memberWallet.GetBalance(), amount), this.SourceService.GetMessage("INSUFFICIENT_BALANCE"))
	//        isTrue(memberAddressService.findByMemberIdAndAddress(user.getId(), address).size() > 0, sourceService.getMessage("WRONG_ADDRESS"));
	isTrue(memberWallet.GetIsLock() == BooleanEnum.IS_FALSE, "钱包已锁定")
	var member = this.MemberService.FindOne(user.GetId())
	//是否完成kyc二级认证
	isTrue(member.GetKycStatus() == 4, "请先完成视频认证")
	var mbPassword = member.GetJyPassword()
	isTrue(member.GetMemberLevel() != MemberLevelEnum.GENERAL, "请先进行实名认证!")
	if mbPassword == "" {
		this.SourceService.GetMessage("NO_SET_JYPASSWORD")
	}
	if member.GetGoogleState() == 1 {
		//谷歌验证
		if org.Apache.Commons.Lang.StringUtils.IsNotEmpty(googleCode) {
			var googleCodes = Long.ParseLong(googleCode)
			var t = System.CurrentTimeMillis()
			var ga = new(util.GoogleAuthenticatorUtil)
			var r = ga.Check_code(member.GetGoogleKey(), googleCodes, t)
			if !r {
				return MessageResult.Error("谷歌验证失败")
			}
		} else {
			return MessageResult.Error("请输入谷歌验证码")
		}
	}
	if !code.Equals(this.RedisUtil.Get(SysConstant.PHONE_WITHDRAW_MONEY_CODE_PREFIX + member.GetMobilePhone())) {
		return error(this.SourceService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	} else {
		this.RedisUtil.Delete(SysConstant.PHONE_WITHDRAW_MONEY_CODE_PREFIX + member.GetMobilePhone())
	}
	if Md5.Md5Digest(jyPassword+member.GetSalt()).ToLowerCase().Equals(mbPassword) == false {
		this.SourceService.GetMessage("ERROR_JYPASSWORD")
	}
	var result = this.MemberWalletService.FreezeBalance(memberWallet, amount)
	if result.GetCode() != 0 {
		return InformationExpiredException("Information Expired")
	}
	//判断该用户当日提币笔数与当日提币数量
	var grade = this.GradeService.FindOne(member.GetMemberGradeId())
	var count = this.RedisUtil.Get(SysConstant.CUSTOMER_DAY_WITHDRAW_TOTAL_COUNT + user.GetId())
	var countLong int64
	if count == nil {
		countLong = 0
	} else {
		countLong = Long.ParseLong(count.ToString())
	}
	if countLong > grade.GetDayWithdrawCount() {
		return error("超过当前等级最大提币次数")
	}
	var coverUsdAmount = this.RedisUtil.Get(SysConstant.CUSTOMER_DAY_WITHDRAW_COVER_USD_AMOUNT)
	var coverUsdAmountBigDecimal *math.BigDecimal
	if coverUsdAmount == nil {
		coverUsdAmountBigDecimal = decimal.Zero
	} else {
		coverUsdAmountBigDecimal = coverUsdAmount.(math.BigDecimal)
	}
	if coverUsdAmountBigDecimal.CompareTo(grade.GetWithdrawCoinAmount()) == 1 {
		return error("超过当前等级最大提币数量")
	}
	var expireTime = DateUtil.CalculateCurrentTime2SecondDaySec()
	//设置提币笔数 与折合数量
	var rate = this.CoinExchangeFactory.Get(unit)
	coverUsdAmountBigDecimal = coverUsdAmountBigDecimal.Add(rate.GetUsdRate().Multiply(amount))
	log.Infof("该用户提币次数=%v,提币折合USD数量=%v", coverUsdAmountBigDecimal)
	//判断用户等级最大提币数
	isTrue(compare(grade.GetWithdrawCoinAmount(), coverUsdAmountBigDecimal), "超过等级最大提币数")
	countLong += 1
	this.RedisUtil.Set(SysConstant.CUSTOMER_DAY_WITHDRAW_TOTAL_COUNT+user.GetId(), countLong, expireTime)
	this.RedisUtil.Set(SysConstant.CUSTOMER_DAY_WITHDRAW_COVER_USD_AMOUNT+user.GetId(), coverUsdAmountBigDecimal, expireTime)
	var withdrawApply = new(entity.WithdrawRecord)
	withdrawApply.SetCoin(coin)
	withdrawApply.SetFee(fee)
	withdrawApply.SetArrivedAmount(sub(amount, fee))
	withdrawApply.SetMemberId(user.GetId())
	withdrawApply.SetTotalAmount(amount)
	withdrawApply.SetAddress(address)
	withdrawApply.SetRemark(remark)
	withdrawApply.SetCanAutoWithdraw(coin.GetCanAutoWithdraw())
	//提币数量低于或等于阈值并且该币种支持自动提币
	if amount.CompareTo(coin.GetWithdrawThreshold()) <= 0 && coin.GetCanAutoWithdraw().Equals(BooleanEnum.IS_TRUE) {
		var withAmountSum = this.SumDailyWithdraw(coin)
		//如果币种设置了单日最大提币量，并且当天已申请的数量（包括待审核、待放币、成功、转账中状态的所有记录）加上当前提币量大于每日最大提币量
		// 进入人工审核
		if coin.GetMaxDailyWithdrawRate() != nil && coin.GetMaxDailyWithdrawRate().CompareTo(decimal.Zero) > 0 && coin.GetMaxDailyWithdrawRate().CompareTo(BigDecimal(withAmountSum).Add(amount)) < 0 {
			withdrawApply.SetStatus(WithdrawStatus.PROCESSING)
			withdrawApply.SetIsAuto(BooleanEnum.IS_FALSE)
			if this.WithdrawApplyService.Save(withdrawApply) != nil {
				return MessageResult.Success(this.SourceService.GetMessage("APPLY_AUDIT"))
			} else {
				return InformationExpiredException("Information Expired")
			}
		} else {
			withdrawApply.SetStatus(WithdrawStatus.WAITING)
			withdrawApply.SetIsAuto(BooleanEnum.IS_TRUE)
			withdrawApply.SetDealTime(withdrawApply.GetCreateTime())
			var withdrawRecord = this.WithdrawApplyService.Save(withdrawApply)
			var json = fastjson.NewJSONObject()
			json.Put("uid", user.GetId())
			//提币总数量
			json.Put("totalAmount", amount)
			//手续费
			json.Put("fee", fee)
			//预计到账数量
			json.Put("arriveAmount", sub(amount, fee))
			//币种
			json.Put("coin", coin)
			//提币地址
			json.Put("address", address)
			//提币记录id
			json.Put("withdrawId", withdrawRecord.GetId())
			json.Put("remark", remark)
			this.KafkaTemplate.Send("withdraw", coin.GetUnit(), json.ToJSONString())
			return MessageResult.Success(this.SourceService.GetMessage("APPLY_SUCCESS"))
		}
	} else {
		withdrawApply.SetStatus(WithdrawStatus.PROCESSING)
		withdrawApply.SetIsAuto(BooleanEnum.IS_FALSE)
		if this.WithdrawApplyService.Save(withdrawApply) != nil {
			return MessageResult.Success(this.SourceService.GetMessage("APPLY_AUDIT"))
		} else {
			return InformationExpiredException("Information Expired")
		}
	}
}
func (this *WithdrawController) PageWithdraw(ctx *gin.Context, user *transform.AuthMember, pageModel *PageModel.PageModel, unit string) (result *MessageResult.MessageResult) {
	var mr = MessageResult(0, "success")
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if !StringUtils.IsEmpty(unit) {
		booleanExpressions.Add(QWithdrawRecord.WithdrawRecord.Coin.Unit.Eq(unit))
	}
	booleanExpressions.Add(QWithdrawRecord.WithdrawRecord.MemberId.Eq(user.GetId()))
	var predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var records = this.WithdrawApplyService.FindAll(predicate, pageModel)
	records.Map(func(x interface {
	}) {
		ScanWithdrawRecord.ToScanWithdrawRecord(x)
	})
	mr.SetData(records)
	return mr
}
func (this *WithdrawController) TodayWithdrawSum(ctx *gin.Context, user *transform.AuthMember, symbol string) (result *MessageResult.MessageResult) {
	if StringUtils.IsEmpty(symbol) {
		return error("symbol is not null")
	}
	var coin = this.CoinService.FindByUnit(symbol)
	if coin == nil {
		return error("coin has not found")
	}
	var withAmountSum = this.SumDailyWithdraw(coin)
	var result = MessageResult.Success()
	result.SetData(withAmountSum)
	return result
}
func (this *WithdrawController) SumDailyWithdraw(ctx *gin.Context, coin *entity.Coin) (result float64) {
	var endTime = time.Now()
	var calendar = Calendar.GetInstance()
	calendar.SetTime(endTime)
	calendar.Set(Calendar.HOUR_OF_DAY, 0)
	calendar.Set(Calendar.MINUTE, 0)
	calendar.Set(Calendar.SECOND, 0)
	calendar.Set(Calendar.MILLISECOND, 0)
	calendar.Add(Calendar.DATE, -1)
	var startTime = calendar.GetTime()
	var withAmountSum = this.WithdrawApplyService.CountWithdrawAmountByTimeAndMemberIdAndCoin(startTime, endTime, coin)
	if withAmountSum == nil {
		withAmountSum = 0
	}
	return withAmountSum
}
func NewWithdrawController(memberAddressService *service.MemberAddressService, redisUtil *util.RedisUtil, memberService *service.MemberService, coinService *service.CoinService, memberWalletService *service.MemberWalletService, withdrawApplyService *service.WithdrawRecordService, kafkaTemplate core.KafkaTemplate[string, string], sourceService *service.LocaleMessageSourceService, gradeService *service.MemberGradeService, coinExchangeFactory *system.CoinExchangeFactory) (ret *WithdrawController) {
	ret = new(WithdrawController)
	ret.MemberAddressService = memberAddressService
	ret.RedisUtil = redisUtil
	ret.MemberService = memberService
	ret.CoinService = coinService
	ret.MemberWalletService = memberWalletService
	ret.WithdrawApplyService = withdrawApplyService
	ret.KafkaTemplate = kafkaTemplate
	ret.SourceService = sourceService
	ret.GradeService = gradeService
	ret.CoinExchangeFactory = coinExchangeFactory
	return
}

type WithdrawController struct {
	MemberAddressService *service.MemberAddressService
	RedisUtil            *util.RedisUtil
	MemberService        *service.MemberService
	CoinService          *service.CoinService
	MemberWalletService  *service.MemberWalletService
	WithdrawApplyService *service.WithdrawRecordService
	KafkaTemplate        core.KafkaTemplate[string, string]
	SourceService        *service.LocaleMessageSourceService
	GradeService         *service.MemberGradeService
	CoinExchangeFactory  *system.CoinExchangeFactory
}
