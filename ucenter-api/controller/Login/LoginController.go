package Login

import (
	"bitrade/core/constant/SysConstant"
	"bitrade/core/controller"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/event"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/MessageResult"
	"bitrade/core/util/RequestUtil"
	"bitrade/libs/common-lang/StringUtils"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/maps/hashmap"
	"time"
)

func (this *LoginController) Login(ctx *gin.Context, username string, password string, code string, ticket string, randStr string) (result *MessageResult.MessageResult) {
	if username == "" {
		this.MessageSourceService.GetMessage("MISSING_USERNAME")
	}
	if password == "" {
		this.MessageSourceService.GetMessage("MISSING_PASSWORD")
	}
	var ip = this.GetRemoteIp(ctx.Request)
	var member, err = this.MemberService.FindByPhoneOrEmail(username)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	var host = RequestUtil.RemoteIp(ctx.Request)
	log.Info("host:" + host)
	//防水验证
	//        boolean result = TengXunWatherProofUtil.watherProof(ticket, randStr, ip);
	//        if (!result) {
	//            return error("验证失败");
	//        }
	if member == nil {
		return MessageResult.Error("用户不存在")
	}
	if member.GetGoogleState() == 1 {
		//谷歌验证
		if StringUtils.IsNotEmpty(code) {
			var googleCode = Long.ParseLong(code)
			var t = System.CurrentTimeMillis()
			var ga = new(util.GoogleAuthenticatorUtil)
			//  ga.setWindowSize(0); // should give 5 * 30 seconds of grace...
			var r = ga.Check_code(member.GetGoogleKey(), googleCode, t)
			if !r {
				return MessageResult.Error("谷歌验证失败")
			}
		} else {
			return MessageResult.Error("请输入谷歌验证码")
		}
	}
	//从session中获取userid
	var userid = ctx.Request.GetSession().GetAttribute("userid").(string)
	//自定义参数,可选择添加
	var param = hashmap.New[string, string]()
	param.Put("user_id", userid)
	//网站用户id
	param.Put("client_type", "web")
	//web:电脑上的浏览器；h5:手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生SDK植入APP应用的方式
	param.Put("ip_address", ip)
	//传输用户请求验证时所携带的IP
	exception := func() (err error) {
		var loginInfo = this.GetLoginInfo(ctx, username, password, ip)
		return this.SuccessWithData(loginInfo)
		return
	}()
	if exception != nil {
		return error(e.GetMessage())
	}
}
func (this *LoginController) GetLoginInfo(ctx *gin.Context, username string, password string, ip string) (result *entity.LoginInfo) {
	var member = this.MemberService.Login(username, password)
	if member == nil {
		var key = SysConstant.LOGIN_LOCK + username
		var code = this.RedisUtil.Get(key)
		if code == "" {
			code = 0
		}
		var codeNum = code.(int)
		codeNum += 1
		if codeNum < 10 {
			this.RedisUtil.Set(key, codeNum, 3, time.Minute)
		} else {
			this.MemberService.Lock(username)
		}
		return AuthenticationException("账号或密码错误")
	}
	this.MemberEvent.OnLoginSuccess(member, ip)
	request.GetSession().SetAttribute(SysConstant.SESSION_MEMBER, AuthMember.ToAuthMember(member))
	var token = request.GetHeader("access-auth-token")
	if !StringUtils.IsBlank(token) && token.Equals(request.GetSession().GetId()) {
		member.SetToken(token)
		var calendar = Calendar.GetInstance()
		calendar.Add(Calendar.HOUR_OF_DAY, 24*7)
		member.SetTokenExpireTime(calendar.GetTime())
	} else {
		member.SetToken(request.GetSession().GetId())
		var calendar = Calendar.GetInstance()
		calendar.Add(Calendar.HOUR_OF_DAY, 24*7)
		member.SetTokenExpireTime(calendar.GetTime())
	}
	// 签到活动是否进行
	var sign = this.SignService.FetchUnderway()
	var loginInfo *entity.LoginInfo
	if sign == nil {
		loginInfo = LoginInfo.GetLoginInfo(member, request.GetSession().GetId(), false, this.PromotePrefix)
	} else {
		loginInfo = LoginInfo.GetLoginInfo(member, request.GetSession().GetId(), true, this.PromotePrefix)
	}
	if this.LoginSms == 1 && member.GetMobilePhone() != nil {
		var phone = member.GetMobilePhone()
		//253国际短信，可以发国内号码，都要加上区域号
		if this.DriverName.EqualsIgnoreCase("two_five_three") {
			this.SmsProvider.SendLoginMessage(ip, member.GetCountry().GetAreaCode()+phone)
		} else {
			if member.GetCountry().GetAreaCode().Equals("86") {
				this.SmsProvider.SendLoginMessage(ip, phone)
			} else {
				this.SmsProvider.SendLoginMessage(ip, member.GetCountry().GetAreaCode()+phone)
			}
		}
	}
	return loginInfo
}
func (this *LoginController) LoginOut2(ctx *gin.Context, request *http.HttpServletRequest, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var messageResult = new(MessageResult.MessageResult)
	log.Info(">>>>>退出登陆接口开始>>>>>")
	exception := func() (err error) {
		request.GetSession().RemoveAttribute(SysConstant.SESSION_MEMBER)
		var member = this.MemberService.FindOne(user.GetId())
		member.SetToken(nil)
		if request.GetSession().GetAttribute(SysConstant.SESSION_MEMBER) != nil {
			messageResult = error(this.MessageSourceService.GetMessage("LOGOUT_FAILED"))
		} else {
			messageResult = success(this.MessageSourceService.GetMessage("LOGOUT_SUCCESS"))
		}
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
		log.Info(">>>>>登出失败>>>>>" + e)
	}
	log.Info(">>>>>退出登陆接口结束>>>>>")
	return messageResult
}
func (this *LoginController) CheckLogin(ctx *gin.Context, request *http.HttpServletRequest) (result *MessageResult.MessageResult) {
	var authMember = request.GetSession().GetAttribute(SESSION_MEMBER).(transform.AuthMember)
	var result = MessageResult.Success()
	if authMember != nil {
		result.SetData(true)
	} else {
		result.SetData(false)
	}
	return result
}
func (this *LoginController) CheckLogin(ctx *gin.Context, mobile string) (result *MessageResult.MessageResult) {
	var member = this.MemberService.FindByPhone(mobile)
	if member != nil {
		return success(member.GetGoogleState())
	}
	var emailMember = this.MemberService.FindByEmail(mobile)
	if emailMember != nil {
		return success(emailMember.GetGoogleState())
	}
	return this.Error("用户名错误")
}
func NewLoginController(memberService *service.MemberService, memberEvent *event.MemberEvent, messageSourceService *service.LocaleMessageSourceService, msService *service.LocaleMessageSourceService, gtSdk *system.GeetestLib, signService *service.SignService, redisUtil *util.RedisUtil, smsProvider *provider.SMSProvider, promotePrefix string, ip138ApiUrl string, ip138Key string, ip138Value string, driverName string, loginSms int) (ret *LoginController) {
	ret = new(LoginController)
	ret.MemberService = memberService
	ret.MemberEvent = memberEvent
	ret.MessageSourceService = messageSourceService
	ret.MsService = msService
	ret.GtSdk = gtSdk
	ret.SignService = signService
	ret.RedisUtil = redisUtil
	ret.SmsProvider = smsProvider
	ret.PromotePrefix = promotePrefix
	ret.Ip138ApiUrl = ip138ApiUrl
	ret.Ip138Key = ip138Key
	ret.Ip138Value = ip138Value
	ret.DriverName = driverName
	ret.LoginSms = loginSms
	return
}

type LoginController struct {
	MemberService        *service.MemberService
	MemberEvent          *event.MemberEvent
	MessageSourceService *service.LocaleMessageSourceService
	MsService            *service.LocaleMessageSourceService
	GtSdk                *system.GeetestLib
	SignService          *service.SignService
	RedisUtil            *util.RedisUtil
	SmsProvider          provider.SMSProvider
	PromotePrefix        string
	Ip138ApiUrl          string
	Ip138Key             string
	Ip138Value           string
	DriverName           string
	LoginSms             int
	controller.BaseController
}
