package controller

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/TransactionType"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/BigDecimalUtils"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/math"
	"github.com/qauzy/util/lists/arraylist"
	"github.com/qauzy/util/sets/hashset"
)

func (this *TransferController) TransferAddress(unit string, user *transform.AuthMember) (result *util.MessageResult) {
	var coin *entity.Coin = this.CoinService.FindByUnit(unit)
	var memberWallet *entity.MemberWallet = this.MemberWalletService.FindByCoinAndMemberId(coin, user.GetId())
	var list *arraylist.List[entity.TransferAddress] = this.TransferAddressService.FindByCoin(coin)
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(new(TransferAddressInfo).SetBalance(memberWallet.GetBalance()).SetInfo(list.Stream().Map(func(x interface {
	}) {
		var oMap = hashset.New[interface {
		}]()
		oMap.Put("address", x.GetAddress())
		oMap.Put("rate", x.GetTransferFee())
		oMap.Put("minAmount", x.GetMinAmount())
		return oMap
	}).Collect(Collectors.ToList())))
	return result
}
func (this *TransferController) Transfer(user *transform.AuthMember, unit string, address string, amount *math.BigDecimal, fee *math.BigDecimal, jyPassword string, remark string) (result *util.MessageResult, err error) {
	hasText(jyPassword, this.SourceService.GetMessage("MISSING_JYPASSWORD"))
	hasText(unit, this.SourceService.GetMessage("MISSING_COIN_TYPE"))
	var coin *entity.Coin = this.CoinService.FindByUnit(unit)
	notNull(coin, this.SourceService.GetMessage("COIN_ILLEGAL"))
	isTrue(coin.GetStatus().Equals(CommonStatus.NORMAL) && coin.GetCanTransfer().Equals(BooleanEnum.IS_TRUE), this.SourceService.GetMessage("COIN_NOT_SUPPORT"))
	var TransferAddress *entity.TransferAddress = this.TransferAddressService.FindOnlyOne(coin, address)
	isTrue(this.TransferAddress != nil, this.SourceService.GetMessage("WRONG_ADDRESS"))
	isTrue(fee.CompareTo(BigDecimalUtils.MulRound(amount, BigDecimalUtils.GetRate(this.TransferAddress.GetTransferFee()))) == 0, this.SourceService.GetMessage("FEE_ERROR"))
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var mbPassword string = member.GetJyPassword()
	if mbPassword == "" {
		this.SourceService.GetMessage("NO_SET_JYPASSWORD")
	}
	if Md5.Md5Digest(jyPassword+member.GetSalt()).ToLowerCase().Equals(mbPassword) == false {
		this.SourceService.GetMessage("ERROR_JYPASSWORD")
	}
	var memberWallet *entity.MemberWallet = this.MemberWalletService.FindByCoinAndMemberId(coin, user.GetId())
	isTrue(compare(memberWallet.GetBalance(), BigDecimalUtils.Add(amount, fee)), this.SourceService.GetMessage("INSUFFICIENT_BALANCE"))
	var result int = this.MemberWalletService.DeductBalance(memberWallet, BigDecimalUtils.Add(amount, fee))
	if result <= 0 {
		return InformationExpiredException("Information Expired")
	}
	var transferRecord *entity.TransferRecord = new(entity.TransferRecord)
	transferRecord.SetAmount(amount)
	transferRecord.SetCoin(coin)
	transferRecord.SetMemberId(user.GetId())
	transferRecord.SetFee(fee)
	transferRecord.SetAddress(address)
	transferRecord.SetRemark(remark)
	var transferRecord1 *entity.TransferRecord = this.TransferRecordService.Save(transferRecord)
	var memberTransaction *entity.MemberTransaction = new(entity.MemberTransaction)
	memberTransaction.SetAddress(address)
	memberTransaction.SetAmount(BigDecimalUtils.Add(fee, amount))
	memberTransaction.SetMemberId(user.GetId())
	memberTransaction.SetSymbol(coin.GetUnit())
	memberTransaction.SetCreateTime(transferRecord1.GetCreateTime())
	memberTransaction.SetType(TransactionType.TRANSFER_ACCOUNTS)
	memberTransaction.SetFee(transferRecord.GetFee())
	this.MemberTransactionService.Save(memberTransaction)
	if transferRecord1 == nil {
		return InformationExpiredException("Information Expired")
	} else {
		var json *fastjson.JSONObject = new(fastjson.JSONObject)
		//会员id
		json.Put("uid", user.GetId())
		//转账数目
		json.Put("amount", amount)
		//转账手续费
		json.Put("fee", fee)
		//币种单位
		json.Put("coin", coin.GetUnit())
		//转账地址
		json.Put("address", address)
		//转账记录ID
		json.Put("recordId", transferRecord1.GetId())
		var jsonObject *fastjson.JSONObject = new(fastjson.JSONObject)
		jsonObject.Put("data", json)
		jsonObject.Put("sign", Md5.Md5Digest(json.ToJSONString()+this.Smac))
		var ciphertext string = DESUtil.ENCRYPTMethod(jsonObject.ToJSONString(), this.Key).ToUpperCase()
		var response string = this.RestTemplate.PostForEntity(this.Url, ciphertext, String).GetBody()
		var resJson *fastjson.JSONObject = JSONObject.ParseObject(DESUtil.Decrypt(response.Trim(), this.Key))
		if resJson != nil {
			//验证签名
			if Md5.Md5Digest(resJson.GetJSONObject("data").ToJSONString() + this.Smac).Equals(resJson.GetString("sign")) {
				if resJson.GetJSONObject("data").GetIntValue("returnCode") == 1 {
					transferRecord1.SetOrderSn(resJson.GetJSONObject("data").GetString("returnMsg"))
					log.Info("============》转账成功")
				}
			}
		}
		return MessageResult.Success()
	}
}
func (this *TransferController) PageWithdraw(user *transform.AuthMember, pageNo int, pageSize int) (result *util.MessageResult) {
	var mr *util.MessageResult = MessageResult(0, "success")
	var records *domain.Page = this.TransferRecordService.FindAllByMemberId(user.GetId(), pageNo, pageSize)
	mr.SetData(records)
	return mr
}
func (this *TransferController) GetSupportTransferCoin(coinUnit string) (result *util.MessageResult) {
	var otcCoins = arraylist.New[entity.OtcCoin]()
	var leverCoins *arraylist.List[entity.LeverCoin]
	if !StringUtils.IsEmpty(coinUnit) {
		var otcCoin *entity.OtcCoin = this.OtcCoinService.FindOtcCoinByUnitAndStatus(coinUnit, CommonStatus.NORMAL)
		otcCoins.Add(otcCoin)
		leverCoins = this.LeverCoinService.FindLeverCoinByCoinUnitAndEnable(coinUnit, BooleanEnum.IS_TRUE)
	} else {
		otcCoins = this.OtcCoinService.GetNormalCoin()
		leverCoins = this.LeverCoinService.FindByEnable(BooleanEnum.IS_TRUE)
	}
	var resultJson *fastjson.JSONObject = new(fastjson.JSONObject)
	resultJson.Put("supportOtcCoins", otcCoins)
	resultJson.Put("supportLeverCoins", leverCoins)
	return success(resultJson)
}

type TransferController struct {
	SourceService            *service.LocaleMessageSourceService
	TransferAddressService   *service.TransferAddressService
	CoinService              *service.CoinService
	MemberWalletService      *service.MemberWalletService
	TransferRecordService    *service.TransferRecordService
	RestTemplate             *client.RestTemplate
	Url                      string
	Key                      string
	Smac                     string
	MemberTransactionService *service.MemberTransactionService
	MemberService            *service.MemberService
	OtcCoinService           *service.OtcCoinService
	LeverCoinService         *service.LeverCoinService
	BaseController
}
