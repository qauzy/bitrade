package OtcWallet

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/controller"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/vo"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func (this *OtcWalletController) GetUserOtcWallet(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult) {
	log.Info("---------查询用户法币账户:" + user.GetId())
	var result = this.OtcWalletService.FindByMemberId(user.GetId())
	return this.SuccessWithData(result)
}
func (this *OtcWalletController) TransferOtcWallet(ctx *gin.Context, user *transform.AuthMember, otcWalletVO *vo.OtcWalletVO) (result *util.MessageResult, err error) {
	log.Info("---------币币账户到法币账户互转:userId=" + user.GetId() + "," + JSONObject.ToJSONString(otcWalletVO))
	//        String jyPassword = otcWalletVO.getJyPassword();
	//        hasText(jyPassword, sourceService.getMessage("MISSING_JYPASSWORD"));
	var coin = this.OtcCoinService.FindByUnit(otcWalletVO.GetCoinName())
	if coin == nil {
		return error("不支持的法币币种")
	}
	var amount = otcWalletVO.GetAmount().SetScale(coin.GetCoinScale(), BigDecimal.ROUND_DOWN)
	var member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetMemberLevel() != MemberLevelEnum.GENERAL, "请先进行实名认证!")
	//        Assert.isTrue(Md5.md5Digest(jyPassword + member.getSalt()).toLowerCase().equals(member.getJyPassword()),
	//                sourceService.getMessage("ERROR_JYPASSWORD"));
	isTrue(compare(amount, decimal.Zero), this.SourceService.GetMessage("参数异常"))
	var memberCoin = this.CoinService.FindByUnit(coin.GetUnit())
	//查询用户币币账户
	var memberWallet = this.MemberWalletService.FindByCoinAndMemberId(memberCoin, user.GetId())
	isTrue(memberWallet.GetIsLock() == BooleanEnum.IS_FALSE, "钱包已锁定")
	//查询用户法币账户
	var otcWallet = this.OtcWalletService.FindByOtcCoinAndMemberId(member.GetId(), coin)
	if otcWallet == nil {
		//如果法币账户不存在新建
		var otcWalletNew = new(entity.OtcWallet)
		otcWalletNew.SetCoin(memberCoin)
		otcWalletNew.SetIsLock(0)
		otcWalletNew.SetMemberId(member.GetId())
		otcWalletNew.SetBalance(decimal.Zero)
		otcWalletNew.SetFrozenBalance(decimal.Zero)
		otcWalletNew.SetReleaseBalance(decimal.Zero)
		otcWalletNew.SetVersion(0)
		otcWallet = this.OtcWalletService.Save(otcWalletNew)
		if otcWallet == nil {
			return error("法币账户创建失败，请联系客服")
		}
	}
	isTrue(otcWallet.GetIsLock() == 0, "钱包已锁定")
	if "0".Equals(otcWalletVO.GetDirection()) {
		//币币转法币
		isTrue(compare(memberWallet.GetBalance(), amount), this.SourceService.GetMessage("INSUFFICIENT_BALANCE"))
		var subResult = this.OtcWalletService.Coin2Otc(memberWallet, otcWallet, amount)
		if subResult == 1 {
			return success("划转成功")
		}
		return error("划转失败")
	} else if "1".Equals(otcWalletVO.GetDirection()) {
		//法币转币币
		isTrue(compare(otcWallet.GetBalance(), amount), this.SourceService.GetMessage("INSUFFICIENT_BALANCE"))
		var addResult = this.OtcWalletService.Otc2Coin(memberWallet, otcWallet, amount)
		if addResult == 1 {
			return success("划转成功")
		}
		return error("划转失败")
	} else {
		return error("参数异常")
	}
}
func NewOtcWalletController(otcWalletService *service.OtcWalletService, memberService *service.MemberService, memberWalletService *service.MemberWalletService, sourceService *service.LocaleMessageSourceService, coinService *service.CoinService, otcCoinService *service.OtcCoinService) (ret *OtcWalletController) {
	ret = new(OtcWalletController)
	ret.OtcWalletService = otcWalletService
	ret.MemberService = memberService
	ret.MemberWalletService = memberWalletService
	ret.SourceService = sourceService
	ret.CoinService = coinService
	ret.OtcCoinService = otcCoinService
	return
}

type OtcWalletController struct {
	OtcWalletService    *service.OtcWalletService
	MemberService       *service.MemberService
	MemberWalletService *service.MemberWalletService
	SourceService       *service.LocaleMessageSourceService
	CoinService         *service.CoinService
	OtcCoinService      *service.OtcCoinService
	controller.BaseController
}
