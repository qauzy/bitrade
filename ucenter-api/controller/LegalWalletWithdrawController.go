package controller

import (
	"bitrade/core/constant/WithdrawStatus"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/BigDecimalUtils"
)

func (this *LegalWalletWithdrawController) Page(pageModel *PageModel.PageModel, status *WithdrawStatus.WithdrawStatus, user *transform.AuthMember) (result *util.MessageResult) {
	var eq *dsl.BooleanExpression = QLegalWalletWithdraw.LegalWalletWithdraw.Member.Id.Eq(user.GetId())
	if status != nil {
		eq.And(QLegalWalletWithdraw.LegalWalletWithdraw.Status.Eq(status))
	}
	var Page *domain.Page = this.LegalWalletWithdrawService.FindAll(eq, pageModel)
	return success(this.Page)
}
func (this *LegalWalletWithdrawController) Post(model *entity.LegalWalletWithdrawModel, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult) {
	var result *util.MessageResult = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	// 合法币种
	var coin *entity.Coin = this.CoinService.FindByUnit(model.GetUnit())
	if coin == nil {
		"validate coin name!"
	}
	if coin.GetHasLegal() == false {
		"validate coin name!"
	}
	//用户提现的币种钱包
	var wallet *entity.MemberWallet = this.WalletService.FindOneByCoinNameAndMemberId(coin.GetName(), user.GetId())
	if wallet == nil {
		"wallet null!"
	}
	if BigDecimalUtils.Compare(wallet.GetBalance(), model.GetAmount()) == false {
		"insufficient closeBalance!"
	}
	//提现人
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member == nil {
		"validate login user!"
	}
	//创建 提现
	var legalWalletWithdraw *entity.LegalWalletWithdraw = new(entity.LegalWalletWithdraw)
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
func (this *LegalWalletWithdrawController) Detail(id int64, user *transform.AuthMember) (result *util.MessageResult) {
	var one *entity.LegalWalletWithdraw = this.LegalWalletWithdrawService.FindDetailWeb(id, user.GetId())
	if one == nil {
		"validate id!"
	}
	return success(one)
}

type LegalWalletWithdrawController struct {
	LegalWalletWithdrawService *service.LegalWalletWithdrawService
	MemberService              *service.MemberService
	CoinService                *service.CoinService
	WalletService              *service.MemberWalletService
	BaseController
}
