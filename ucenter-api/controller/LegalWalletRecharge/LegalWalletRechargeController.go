package LegalWalletRecharge

import (
	"bitrade/core/constant/LegalWalletState"
	"bitrade/core/constant/PageModel"
	"bitrade/core/controller"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util/MessageResult"
	"github.com/gin-gonic/gin"
)

func (this *LegalWalletRechargeController) Recharge(ctx *gin.Context, model *entity.LegalWalletRechargeModel, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var result = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	var coin = this.CoinService.FindByUnit(model.GetUnit())
	if coin == nil {
		"validate coin name!"
	}
	if coin.GetHasLegal() == false {
		"validate coin legal!"
	}
	//登录用户
	var member = this.MemberService.FindOne(user.GetId())
	//新建充值
	var legalWalletRecharge = new(entity.LegalWalletRecharge)
	legalWalletRecharge.SetMember(member)
	//所属会员
	legalWalletRecharge.SetCoin(coin)
	//充值币种
	legalWalletRecharge.SetAmount(model.GetAmount())
	//充值金额
	legalWalletRecharge.SetPaymentInstrument(model.GetPaymentInstrument())
	//支付凭证
	legalWalletRecharge.SetState(LegalWalletState.APPLYING)
	//状态
	legalWalletRecharge.SetPayMode(model.GetPayMode())
	//支付方式
	legalWalletRecharge.SetRemark(model.GetRemark())
	//备注
	this.LegalWalletRechargeService.Save(legalWalletRecharge)
	return success()
}
func (this *LegalWalletRechargeController) Page(ctx *gin.Context, pageModel *PageModel.PageModel, screen *screen.LegalWalletScreen, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var coin *entity.Coin
	if StringUtils.IsNotBlank(screen.GetCoinName()) {
		coin = this.CoinService.FindOne(screen.GetCoinName())
		if coin == nil {
			"validate coin name!"
		}
		if coin.GetHasLegal() == false {
			"validate coin legal!"
		}
	}
	var eq = QLegalWalletRecharge.LegalWalletRecharge.Member.Id.Eq(user.GetId())
	if coin != nil {
		eq.And(QLegalWalletRecharge.LegalWalletRecharge.Coin.Name.Eq(screen.GetCoinName()))
	}
	if screen.GetState() != nil {
		eq.And(QLegalWalletRecharge.LegalWalletRecharge.State.Eq(screen.GetState()))
	}
	var Page = this.LegalWalletRechargeService.FindAll(eq, pageModel)
	return success(this.Page)
}
func (this *LegalWalletRechargeController) Get(ctx *gin.Context, user *transform.AuthMember, id int64) (result *MessageResult.MessageResult) {
	var one = this.LegalWalletRechargeService.FindOneByIdAndMemberId(id, user.GetId())
	if one == nil {
		"validate id!"
	}
	return this.SuccessWithData(one)
}
func NewLegalWalletRechargeController(legalWalletRechargeService *service.LegalWalletRechargeService, memberService *service.MemberService, coinService *service.CoinService) (ret *LegalWalletRechargeController) {
	ret = new(LegalWalletRechargeController)
	ret.LegalWalletRechargeService = legalWalletRechargeService
	ret.MemberService = memberService
	ret.CoinService = coinService
	return
}

type LegalWalletRechargeController struct {
	LegalWalletRechargeService *service.LegalWalletRechargeService
	MemberService              *service.MemberService
	CoinService                *service.CoinService
	controller.BaseController
}
