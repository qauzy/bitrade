package LegalWalletWithdraw

import (
	"bitrade/core/constant/PageModel"
	"bitrade/core/constant/WithdrawStatus"
	"bitrade/core/controller"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util/BigDecimalUtils"
	"bitrade/core/util/MessageResult"
	"github.com/gin-gonic/gin"
)

func (this *LegalWalletWithdrawController) Page(ctx *gin.Context, pageModel *PageModel.PageModel, status *WithdrawStatus.WithdrawStatus, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var eq = QLegalWalletWithdraw.LegalWalletWithdraw.Member.Id.Eq(user.GetId())
	if status != nil {
		eq.And(QLegalWalletWithdraw.LegalWalletWithdraw.Status.Eq(status))
	}
	var Page = this.LegalWalletWithdrawService.FindAll(eq, pageModel)
	return success(this.Page)
}
func (this *LegalWalletWithdrawController) Post(ctx *gin.Context, model *entity.LegalWalletWithdrawModel, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var result = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	// 合法币种
	var coin = this.CoinService.FindByUnit(model.GetUnit())
	if coin == nil {
		"validate coin name!"
	}
	if coin.GetHasLegal() == false {
		"validate coin name!"
	}
	//用户提现的币种钱包
	var wallet = this.WalletService.FindOneByCoinNameAndMemberId(coin.GetName(), user.GetId())
	if wallet == nil {
		"wallet null!"
	}
	if BigDecimalUtils.Compare(wallet.GetBalance(), model.GetAmount()) == false {
		"insufficient closeBalance!"
	}
	//提现人
	var member = this.MemberService.FindOne(user.GetId())
	if member == nil {
		"validate login user!"
	}
	//创建 提现
	var legalWalletWithdraw = new(entity.LegalWalletWithdraw)
	legalWalletWithdraw.SetMember(member)
	legalWalletWithdraw.SetAccount(model.GetAccount())
	legalWalletWithdraw.SetCoin(coin)
	legalWalletWithdraw.SetAmount(model.GetAmount())
	legalWalletWithdraw.SetPayMode(model.GetPayMode())
	legalWalletWithdraw.SetStatus(WithdrawStatus.PROCESSING)
	legalWalletWithdraw.SetRemark(model.GetRemark())
	//提现操作
	this.LegalWalletWithdrawService.Withdraw(wallet, legalWalletWithdraw)
	return success()
}
func (this *LegalWalletWithdrawController) Detail(ctx *gin.Context, id int64, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var one = this.LegalWalletWithdrawService.FindDetailWeb(id, user.GetId())
	if one == nil {
		"validate id!"
	}
	return this.SuccessWithData(one)
}
func NewLegalWalletWithdrawController(legalWalletWithdrawService *service.LegalWalletWithdrawService, memberService *service.MemberService, coinService *service.CoinService, walletService *service.MemberWalletService) (ret *LegalWalletWithdrawController) {
	ret = new(LegalWalletWithdrawController)
	ret.LegalWalletWithdrawService = legalWalletWithdrawService
	ret.MemberService = memberService
	ret.CoinService = coinService
	ret.WalletService = walletService
	return
}

type LegalWalletWithdrawController struct {
	LegalWalletWithdrawService *service.LegalWalletWithdrawService
	MemberService              *service.MemberService
	CoinService                *service.CoinService
	WalletService              *service.MemberWalletService
	controller.BaseController
}
