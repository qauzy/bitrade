package controller

import (
	SysConstant "bitrade/core/constant/SysConstants"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/RequestUtil"
	"github.com/qauzy/util/sets/hashset"
	"time"
)

func (this *LoginController) Login(request *http.HttpServletRequest, username string, password string, code string, ticket string, randStr string) (result *util.MessageResult, err error) {
	if username == "" {
		this.MessageSourceService.GetMessage("MISSING_USERNAME")
	}
	if password == "" {
		this.MessageSourceService.GetMessage("MISSING_PASSWORD")
	}
	var ip string = getRemoteIp(request)
	var member *entity.Member = this.MemberService.FindByPhoneOrEmail(username)
	var host string = RequestUtil.RemoteIp(request)
	log.Info("host:" + host)
	//防水验证
	//        boolean result = TengXunWatherProofUtil.watherProof(ticket, randStr, ip);
	//        if (!result) {
	//            return error("验证失败");
	//        }
	if member == nil {
		return error("用户不存在")
	}
	if member.GetGoogleState() == 1 {
		//谷歌验证
		if StringUtils.IsNotEmpty(code) {
			var googleCode int64 = Long.ParseLong(code)
			var t int64 = System.CurrentTimeMillis()
			var ga *util.GoogleAuthenticatorUtil = new(util.GoogleAuthenticatorUtil)
			//  ga.setWindowSize(0); // should give 5 * 30 seconds of grace...
			var r bool = ga.Check_code(member.GetGoogleKey(), googleCode, t)
			if !r {
				return MessageResult.Error("谷歌验证失败")
			}
		} else {
			return MessageResult.Error("请输入谷歌验证码")
		}
	}
	//从session中获取userid
	var userid = request.GetSession().GetAttribute("userid").(string)
	//自定义参数,可选择添加
	var param = hashset.New[string]()
	param.Put("user_id", userid)
	//网站用户id
	param.Put("client_type", "web")
	//web:电脑上的浏览器；h5:手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生SDK植入APP应用的方式
	param.Put("ip_address", ip)
	//传输用户请求验证时所携带的IP
	exception := func() (err error) {
		var loginInfo *entity.LoginInfo = this.GetLoginInfo(username, password, ip, request)
		return success(loginInfo)
		return
	}()
	if exception != nil {
		return error(e.GetMessage())
	}
}
func (this *LoginController) GetLoginInfo(username string, password string, ip string, request *http.HttpServletRequest) (result *entity.LoginInfo, err error) {
	var member *entity.Member = this.MemberService.Login(username, password)
	if member == nil {
		var key = SysConstant.LOGIN_LOCK + username
		var code interface {
		} = this.RedisUtil.Get(key)
		if code == nil {
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
	var token string = request.GetHeader("access-auth-token")
	if !StringUtils.IsBlank(token) && token.Equals(request.GetSession().GetId()) {
		member.SetToken(token)
		var calendar *util.Calendar = Calendar.GetInstance()
		calendar.Add(Calendar.HOUR_OF_DAY, 24*7)
		member.SetTokenExpireTime(calendar.GetTime())
	} else {
		member.SetToken(request.GetSession().GetId())
		var calendar *util.Calendar = Calendar.GetInstance()
		calendar.Add(Calendar.HOUR_OF_DAY, 24*7)
		member.SetTokenExpireTime(calendar.GetTime())
	}
	// 签到活动是否进行
	var sign *entity.Sign = this.SignService.FetchUnderway()
	var loginInfo *entity.LoginInfo
	if sign == nil {
		loginInfo = LoginInfo.GetLoginInfo(member, request.GetSession().GetId(), false, this.PromotePrefix)
	} else {
		loginInfo = LoginInfo.GetLoginInfo(member, request.GetSession().GetId(), true, this.PromotePrefix)
	}
	if this.LoginSms == 1 && member.GetMobilePhone() != nil {
		var phone string = member.GetMobilePhone()
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
func (this *LoginController) LoginOut2(request *http.HttpServletRequest, user *transform.AuthMember) (result *util.MessageResult) {
	var messageResult *util.MessageResult = new(util.MessageResult)
	log.Info(">>>>>退出登陆接口开始>>>>>")
	exception := func() (err error) {
		request.GetSession().RemoveAttribute(SysConstant.SESSION_MEMBER)
		var member *entity.Member = this.MemberService.FindOne(user.GetId())
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
func (this *LoginController) CheckLogin(request *http.HttpServletRequest) (result *util.MessageResult) {
	var authMember = request.GetSession().GetAttribute(SESSION_MEMBER).(transform.AuthMember)
	var result *util.MessageResult = MessageResult.Success()
	if authMember != nil {
		result.SetData(true)
	} else {
		result.SetData(false)
	}
	return result
}
func (this *LoginController) CheckLogin(mobile string) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindByPhone(mobile)
	if member != nil {
		return success(member.GetGoogleState())
	}
	var emailMember *entity.Member = this.MemberService.FindByEmail(mobile)
	if emailMember != nil {
		return success(emailMember.GetGoogleState())
	}
	return error("用户名错误")
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
	BaseController
}
