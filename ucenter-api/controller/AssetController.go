package controller

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/TransactionType"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/util/lists/arraylist"
	"strings"
)

func (this *AssetController) FindWallet(member *transform.AuthMember) (result *util.MessageResult) {
	var wallets *arraylist.List[entity.MemberWallet] = this.WalletService.FindAllByMemberId(member.GetId())
	wallets.ForEach(func(wallet interface {
	}) {
		var rate *CoinExchangeFactory.ExchangeRate = this.CoinExchangeFactory.Get(wallet.GetCoin().GetUnit())
		if rate != nil {
			wallet.GetCoin().SetUsdRate(rate.GetUsdRate())
			wallet.GetCoin().SetCnyRate(rate.GetCnyRate())
			wallet.GetCoin().SetSgdRate(rate.GetSgdRate())
		} else {
			log.Infof("unit = %v , rate = null ", wallet.GetCoin().GetUnit())
		}
	})
	var mr *util.MessageResult = MessageResult.Success("success")
	mr.SetData(wallets)
	return mr
}
func (this *AssetController) FindTransaction(member *transform.AuthMember, pageNo int, pageSize int, oType *TransactionType.TransactionType, symbol string, startTime string, endTime string) (result *util.MessageResult, err error) {
	var page *domain.Page = this.TransactionService.QueryByMember(member.GetId(), pageNo, pageSize, oType, startTime, endTime, symbol)
	return success(page)
}
func (this *AssetController) FindTransaction(member *transform.AuthMember, request *http.HttpServletRequest, pageNo int, pageSize int, symbol string) (result *util.MessageResult, err error) {
	var oType *TransactionType.TransactionType
	if StringUtils.IsNotEmpty(request.GetParameter("type")) {
		oType = TransactionType.ValueOfOrdinal(Convert.StrToInt(request.GetParameter("type"), 0))
	}
	var startDate = ""
	var endDate = ""
	if StringUtils.IsNotEmpty(request.GetParameter("dateRange")) {
		var parts []string = strings.Split(request.GetParameter("dateRange"), "~")
		startDate = parts[0].Trim()
		endDate = parts[1].Trim()
	}
	return success(this.TransactionService.QueryByMember(member.GetId(), pageNo, pageSize, oType, startDate, endDate, symbol))
}
func (this *AssetController) FindWalletBySymbol(member *transform.AuthMember, symbol string) (result *util.MessageResult) {
	var mr *util.MessageResult = MessageResult.Success("success")
	mr.SetData(this.WalletService.FindByCoinUnitAndMemberId(symbol, member.GetId()))
	return mr
}
func (this *AssetController) ResetWalletAddress(member *transform.AuthMember, unit string) (result *util.MessageResult) {
	exception := func() (err error) {
		var json *fastjson.JSONObject = new(fastjson.JSONObject)
		json.Put("uid", member.GetId())
		this.KafkaTemplate.Send("reset-member-address", unit, json.ToJSONString())
		return MessageResult.Success("提交成功")
		return
	}()
	if exception != nil {
		return MessageResult.Error("未知异常")
	}
}
func (this *AssetController) LockPositionRecordList(member *transform.AuthMember, status *CommonStatus.CommonStatus, coinUnit string, pageModel *PageModel.PageModel) (result *util.MessageResult) {
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if status != nil {
		booleanExpressions.Add(QLockPositionRecord.LockPositionRecord.Status.Eq(status))
	}
	if coinUnit != nil {
		var coin *entity.Coin = this.CoinService.FindByUnit(coinUnit)
		if coin == nil {
			return MessageResult.Error(this.SourceService.GetMessage("COIN_ILLEGAL"))
		}
		booleanExpressions.Add(QLockPositionRecord.LockPositionRecord.Coin.Eq(coin))
	}
	booleanExpressions.Add(QLockPositionRecord.LockPositionRecord.MemberId.Eq(member.GetId()))
	var predicate *types.Predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var LockPositionRecordList *domain.Page = this.LockPositionRecordService.FindAll(predicate, pageModel)
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(this.LockPositionRecordList)
	return result
}

type AssetController struct {
	WalletService             *service.MemberWalletService
	TransactionService        *service.MemberTransactionService
	CoinExchangeFactory       *system.CoinExchangeFactory
	KafkaTemplate             *core.KafkaTemplate
	LockPositionRecordService *service.LockPositionRecordService
	CoinService               *service.CoinService
	SourceService             *service.LocaleMessageSourceService
	BaseController
}
