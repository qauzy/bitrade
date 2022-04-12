package Asset

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/PageModel"
	"bitrade/core/constant/TransactionType"
	"bitrade/core/controller"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/fastjson"
	"strings"
)

func (this *AssetController) FindWallet(ctx *gin.Context, member *transform.AuthMember) (result *util.MessageResult) {
	var wallets, err = this.WalletService.FindAllByMemberId(member.GetId())
	wallets.ForEach(func(wallet interface {
	}) {
		var rate = this.CoinExchangeFactory.Get(wallet.GetCoin().GetUnit())
		if rate != nil {
			wallet.GetCoin().SetUsdRate(rate.GetUsdRate())
			wallet.GetCoin().SetCnyRate(rate.GetCnyRate())
			wallet.GetCoin().SetSgdRate(rate.GetSgdRate())
		} else {
			log.Infof("unit = %v , rate = null ", wallet.GetCoin().GetUnit())
		}
	})
	var mr = MessageResult.Success("success")
	mr.SetData(wallets)
	return mr
}
func (this *AssetController) FindTransaction(ctx *gin.Context, member *transform.AuthMember, pageNo int, pageSize int, oType *TransactionType.TransactionType, symbol string, startTime string, endTime string) (result *util.MessageResult, err error) {
	var page = this.TransactionService.QueryByMember(member.GetId(), pageNo, pageSize, oType, startTime, endTime, symbol)
	return this.SuccessWithData(page)
}
func (this *AssetController) FindTransaction(ctx *gin.Context, member *transform.AuthMember, request *http.HttpServletRequest, pageNo int, pageSize int, symbol string) (result *util.MessageResult, err error) {
	var oType *TransactionType.TransactionType
	if StringUtils.IsNotEmpty(request.GetParameter("type")) {
		oType = TransactionType.ValueOfOrdinal(Convert.StrToInt(request.GetParameter("type"), 0))
	}
	var startDate = ""
	var endDate = ""
	if StringUtils.IsNotEmpty(request.GetParameter("dateRange")) {
		var parts = strings.Split(request.GetParameter("dateRange"), "~")
		startDate = parts[0].Trim()
		endDate = parts[1].Trim()
	}
	return success(this.TransactionService.QueryByMember(member.GetId(), pageNo, pageSize, oType, startDate, endDate, symbol))
}
func (this *AssetController) FindWalletBySymbol(ctx *gin.Context, member *transform.AuthMember, symbol string) (result *util.MessageResult) {
	var mr = MessageResult.Success("success")
	mr.SetData(this.WalletService.FindByCoinUnitAndMemberId(symbol, member.GetId()))
	return mr
}
func (this *AssetController) ResetWalletAddress(ctx *gin.Context, member *transform.AuthMember, unit string) (result *util.MessageResult) {
	exception := func() (err error) {
		var json = new(fastjson.JSONObject)
		json.Put("uid", member.GetId())
		this.KafkaTemplate.Send("reset-member-address", unit, json.ToJSONString())
		return MessageResult.Success("提交成功")
		return
	}()
	if exception != nil {
		return MessageResult.Error("未知异常")
	}
}
func (this *AssetController) LockPositionRecordList(ctx *gin.Context, member *transform.AuthMember, status *CommonStatus.CommonStatus, coinUnit string, pageModel *PageModel.PageModel) (result *util.MessageResult) {
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if status != nil {
		booleanExpressions.Add(QLockPositionRecord.LockPositionRecord.Status.Eq(status))
	}
	if coinUnit != nil {
		var coin = this.CoinService.FindByUnit(coinUnit)
		if coin == nil {
			return MessageResult.Error(this.SourceService.GetMessage("COIN_ILLEGAL"))
		}
		booleanExpressions.Add(QLockPositionRecord.LockPositionRecord.Coin.Eq(coin))
	}
	booleanExpressions.Add(QLockPositionRecord.LockPositionRecord.MemberId.Eq(member.GetId()))
	var predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var LockPositionRecordList = this.LockPositionRecordService.FindAll(predicate, pageModel)
	var result = MessageResult.Success()
	result.SetData(this.LockPositionRecordList)
	return result
}
func NewAssetController(walletService *service.MemberWalletService, transactionService *service.MemberTransactionService, coinExchangeFactory *system.CoinExchangeFactory, kafkaTemplate *core.KafkaTemplate, lockPositionRecordService *service.LockPositionRecordService, coinService *service.CoinService, sourceService *service.LocaleMessageSourceService) (ret *AssetController) {
	ret = new(AssetController)
	ret.WalletService = walletService
	ret.TransactionService = transactionService
	ret.CoinExchangeFactory = coinExchangeFactory
	ret.KafkaTemplate = kafkaTemplate
	ret.LockPositionRecordService = lockPositionRecordService
	ret.CoinService = coinService
	ret.SourceService = sourceService
	return
}

type AssetController struct {
	WalletService             *service.MemberWalletService
	TransactionService        *service.MemberTransactionService
	CoinExchangeFactory       *system.CoinExchangeFactory
	KafkaTemplate             *core.KafkaTemplate
	LockPositionRecordService *service.LockPositionRecordService
	CoinService               *service.CoinService
	SourceService             *service.LocaleMessageSourceService
	controller.BaseController
}
