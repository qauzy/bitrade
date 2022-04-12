package Member

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/controller"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util/MessageResult"
	"github.com/gin-gonic/gin"
)

func (this *MemberController) SignIn(ctx *gin.Context, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	//校验 签到活动 币种 会员 会员钱包
	if user == nil {
		"The login timeout!"
	}
	var sign = this.SignService.FetchUnderway()
	if sign == nil {
		"The check-in activity is over!"
	}
	var coin = sign.GetCoin()
	if coin.GetStatus() == CommonStatus.NORMAL == false {
		"coin disabled!"
	}
	var member = this.MemberService.FindOne(user.GetId())
	if member == nil {
		"validate member id!"
	}
	if member.GetSignInAbility() == true == false {
		"Have already signed in!"
	}
	var memberWallet = this.WalletService.FindByCoinAndMember(coin, member)
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
func NewMemberController(memberService *service.MemberService, signService *service.SignService, walletService *service.MemberWalletService) (ret *MemberController) {
	ret = new(MemberController)
	ret.MemberService = memberService
	ret.SignService = signService
	ret.WalletService = walletService
	return
}

type MemberController struct {
	MemberService *service.MemberService
	SignService   *service.SignService
	WalletService *service.MemberWalletService
	controller.BaseController
}
