package controller

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util"
)

func (this *MemberController) SignIn(user *transform.AuthMember) (result *util.MessageResult) {
	//校验 签到活动 币种 会员 会员钱包
	if user == nil {
		"The login timeout!"
	}
	var sign *entity.Sign = this.SignService.FetchUnderway()
	if sign == nil {
		"The check-in activity is over!"
	}
	var coin *entity.Coin = sign.GetCoin()
	if coin.GetStatus() == CommonStatus.NORMAL == false {
		"coin disabled!"
	}
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member == nil {
		"validate member id!"
	}
	if member.GetSignInAbility() == true == false {
		"Have already signed in!"
	}
	var memberWallet *entity.MemberWallet = this.WalletService.FindByCoinAndMember(coin, member)
	if memberWallet == nil {
		"Member wallet does not exist!"
	}
	if memberWallet.GetIsLock() == BooleanEnum.IS_FALSE == false {
		"Wallet locked!"
	}
	//签到事件
	this.MemberService.SignInIncident(member, memberWallet, sign)
	return success()
}

type MemberController struct {
	MemberService *service.MemberService
	SignService   *service.SignService
	WalletService *service.MemberWalletService
	BaseController
}
