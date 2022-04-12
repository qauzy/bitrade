package Register

import (
	"bitrade/core/constant/BooleanEnum"
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/controller"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/MessageResult"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/sets/hashset"
	"reflect"
	"time"
)

func (this *RegisterController) AllCountry(ctx *gin.Context) (result *MessageResult.MessageResult) {
	var result = new(MessageResult.success)
	var list = this.CountryService.GetAllCountry()
	result.SetData(list)
	return result
}
func (this *RegisterController) CheckUsername(ctx *gin.Context, username string) (result *MessageResult.MessageResult) {
	var result = new(MessageResult.success)
	if this.MemberService.UsernameIsExist(username) {
		result.SetCode(500)
		result.SetMessage(this.LocaleMessageSourceService.GetMessage("ACTIVATION_FAILS_USERNAME"))
	}
	return result
}
func (this *RegisterController) Activate(ctx *gin.Context, key string, request *http.HttpServletRequest) (result string, err error) {
	if StringUtils.IsEmpty(key) {
		request.SetAttribute("result", this.LocaleMessageSourceService.GetMessage("INVALID_LINK"))
		return "registeredResult"
	}
	var info = this.RedisUtil.Get(key)
	var loginByEmail *entity.LoginByEmail
	if reflect.TypeOf(info) == LoginByEmail {
		loginByEmail = info.(entity.LoginByEmail)
	}
	if loginByEmail == nil {
		request.SetAttribute("result", this.LocaleMessageSourceService.GetMessage("INVALID_LINK"))
		return "registeredResult"
	}
	if this.MemberService.EmailIsExist(loginByEmail.GetEmail()) {
		request.SetAttribute("result", this.LocaleMessageSourceService.GetMessage("ACTIVATION_FAILS_EMAIL"))
		return "registeredResult"
	} else if this.MemberService.UsernameIsExist(loginByEmail.GetUsername()) {
		request.SetAttribute("result", this.LocaleMessageSourceService.GetMessage("ACTIVATION_FAILS_USERNAME"))
		return "registeredResult"
	}
	//删除redis里存的键值
	this.RedisUtil.Delete(key)
	this.RedisUtil.Delete(loginByEmail.GetEmail())
	//不可重复随机数
	var loginNo = String.ValueOf(this.IdWorkByTwitter.NextId())
	//盐
	var credentialsSalt = ByteSource.Util.Bytes(loginNo).ToHex()
	//生成密码
	var password = Md5.Md5Digest(loginByEmail.GetPassword() + credentialsSalt).ToLowerCase()
	var member = new(entity.Member)
	member.SetMemberLevel(MemberLevelEnum.GENERAL)
	var location = new(entity.Location)
	location.SetCountry(loginByEmail.GetCountry())
	member.SetLocation(location)
	var country = new(entity.Country)
	country.SetZhName(loginByEmail.GetCountry())
	member.SetCountry(country)
	member.SetUsername(loginByEmail.GetUsername())
	member.SetPassword(password)
	member.SetEmail(loginByEmail.GetEmail())
	member.SetSalt(credentialsSalt)
	member.SetKycStatus(0)
	var member1 = this.MemberService.Save(member)
	if member1 != nil {
		member1.SetPromotionCode(String.Format(this.UserNameFormat, member1.GetId()) + GeneratorUtil.GetNonceString(2))
		this.MemberEvent.OnRegisterSuccess(member1, loginByEmail.GetPromotion())
		request.SetAttribute("result", this.LocaleMessageSourceService.GetMessage("ACTIVATION_SUCCESSFUL"))
	} else {
		request.SetAttribute("result", this.LocaleMessageSourceService.GetMessage("ACTIVATION_FAILS"))
	}
	return "registeredResult"
}
func (this *RegisterController) RegisterByEmail(ctx *gin.Context, loginByEmail *entity.LoginByEmail, request *http.HttpServletRequest, bindingResult *validation.BindingResult) (result *MessageResult.MessageResult, err error) {
	var result = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	var ip = getRemoteIp(request)
	//防水验证
	//        boolean resultProof = TengXunWatherProofUtil.watherProof(loginByEmail.getTicket(), loginByEmail.getRandStr(), ip);
	//        if (!resultProof) {
	//            return error("验证码错误");
	//        }
	var email = loginByEmail.GetEmail()
	isTrue(!this.MemberService.EmailIsExist(email), this.LocaleMessageSourceService.GetMessage("EMAIL_ALREADY_BOUND"))
	isTrue(!this.MemberService.UsernameIsExist(loginByEmail.GetUsername()), this.LocaleMessageSourceService.GetMessage("USERNAME_ALREADY_EXISTS"))
	if this.RedisUtil.Get(email) != nil {
		return error(this.LocaleMessageSourceService.GetMessage("LOGIN_EMAIL_ALREADY_SEND"))
	}
	exception := func() (err error) {
		log.Info("send==================================")
		this.SentEmail(loginByEmail, email)
		log.Info("success===============================")
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
		return error(this.LocaleMessageSourceService.GetMessage("REGISTRATION_FAILED"))
	}
	return success(this.LocaleMessageSourceService.GetMessage("SEND_LOGIN_EMAIL_SUCCESS"))
}
func (this *RegisterController) SentEmail(ctx *gin.Context, loginByEmail *entity.LoginByEmail, email string) (err error) {
	//缓存邮箱和注册信息
	var token = UUID.RandomUUID().ToString().Replace("-", "")
	var mimeMessage = this.JavaMailSender.CreateMimeMessage()
	var helper *javamail.MimeMessageHelper
	helper = MimeMessageHelper(mimeMessage, true)
	helper.SetFrom(this.From)
	helper.SetTo(email)
	helper.SetSubject(this.Company)
	var model = hashset.New[interface {
	}]()
	model.Put("username", loginByEmail.GetUsername())
	model.Put("token", token)
	model.Put("host", this.Host)
	model.Put("name", this.Company)
	var cfg = Configuration(Configuration.VERSION_2_3_26)
	cfg.SetClassForTemplateLoading(this.GetClass(), "/templates")
	var template = cfg.GetTemplate("activateEmail.ftl")
	var html = FreeMarkerTemplateUtils.ProcessTemplateIntoString(template, model)
	helper.SetText(html, true)
	//发送邮件
	this.JavaMailSender.Send(mimeMessage)
	this.RedisUtil.Set(token, loginByEmail, 12, time.Hour)
	this.RedisUtil.Set(email, "", 12, time.Hour)
}
func (this *RegisterController) LoginByPhone(ctx *gin.Context, loginByPhone *entity.LoginByPhone, bindingResult *validation.BindingResult, request *http.HttpServletRequest) (result *MessageResult.MessageResult, err error) {
	var result = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	var ip = getRemoteIp(request)
	//防水验证
	//        boolean resultProof = TengXunWatherProofUtil.watherProof(loginByPhone.getTicket(), loginByPhone.getRandStr(), ip);
	//        if (!resultProof) {
	//            return error("验证码验证失败,请重新获取验证码");
	//        }
	if this.LoginByPhone.GetCountry().Equals("中国") {
		if ValidateUtil.IsMobilePhone(this.LoginByPhone.GetPhone().Trim()) == false {
			this.LocaleMessageSourceService.GetMessage("PHONE_EMPTY_OR_INCORRECT")
		}
	}
	var phone = this.LoginByPhone.GetPhone()
	var code = this.RedisUtil.Get(SysConstant.PHONE_REG_CODE_PREFIX + phone)
	isTrue(!this.MemberService.PhoneIsExist(phone), this.LocaleMessageSourceService.GetMessage("PHONE_ALREADY_EXISTS"))
	isTrue(!this.MemberService.UsernameIsExist(this.LoginByPhone.GetUsername()), this.LocaleMessageSourceService.GetMessage("USERNAME_ALREADY_EXISTS"))
	notNull(code, this.LocaleMessageSourceService.GetMessage("VERIFICATION_CODE_NOT_EXISTS"))
	if !code.ToString().Equals(this.LoginByPhone.GetCode()) {
		return error(this.LocaleMessageSourceService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	} else {
		this.RedisUtil.Delete(SysConstant.PHONE_REG_CODE_PREFIX + phone)
	}
	//不可重复随机数
	var loginNo = String.ValueOf(this.IdWorkByTwitter.NextId())
	//盐
	var credentialsSalt = ByteSource.Util.Bytes(loginNo).ToHex()
	//生成密码
	var password = Md5.Md5Digest(this.LoginByPhone.GetPassword() + credentialsSalt).ToLowerCase()
	var member = new(entity.Member)
	member.SetMemberLevel(MemberLevelEnum.GENERAL)
	var location = new(entity.Location)
	location.SetCountry(this.LoginByPhone.GetCountry())
	var country = new(entity.Country)
	country.SetZhName(this.LoginByPhone.GetCountry())
	member.SetCountry(country)
	member.SetLocation(location)
	member.SetUsername(this.LoginByPhone.GetUsername())
	member.SetPassword(password)
	member.SetMobilePhone(phone)
	member.SetSalt(credentialsSalt)
	member.SetKycStatus(0)
	var member1 = this.MemberService.Save(member)
	if member1 != nil {
		member1.SetPromotionCode(String.Format(this.UserNameFormat, member1.GetId()) + GeneratorUtil.GetNonceString(2))
		this.MemberEvent.OnRegisterSuccess(member1, this.LoginByPhone.GetPromotion())
		return success(this.LocaleMessageSourceService.GetMessage("REGISTRATION_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("REGISTRATION_FAILED"))
	}
}
func (this *RegisterController) SendBindEmail(ctx *gin.Context, email string, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	if ValidateUtil.IsEmail(email) == false {
		this.LocaleMessageSourceService.GetMessage("WRONG_EMAIL")
	}
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetEmail() != nil {
		this.LocaleMessageSourceService.GetMessage("BIND_EMAIL_REPEAT")
	}
	if this.MemberService.EmailIsExist(email) {
		this.LocaleMessageSourceService.GetMessage("EMAIL_ALREADY_BOUND")
	}
	var code = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.RedisUtil.Get(EMAIL_BIND_CODE_PREFIX+email) != nil {
		return error(this.LocaleMessageSourceService.GetMessage("EMAIL_ALREADY_SEND"))
	}
	exception := func() (err error) {
		this.SentEmailCode(email, code)
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
		return error(this.LocaleMessageSourceService.GetMessage("SEND_FAILED"))
	}
	return success(this.LocaleMessageSourceService.GetMessage("SENT_SUCCESS_TEN"))
}
func (this *RegisterController) SentEmailCode(ctx *gin.Context, email string, code string) (err error) {
	var mimeMessage = this.JavaMailSender.CreateMimeMessage()
	var helper *javamail.MimeMessageHelper
	helper = MimeMessageHelper(mimeMessage, true)
	helper.SetFrom(this.From)
	helper.SetTo(email)
	helper.SetSubject(this.Company)
	var model = hashset.New[interface {
	}]()
	model.Put("code", code)
	var cfg = Configuration(Configuration.VERSION_2_3_26)
	cfg.SetClassForTemplateLoading(this.GetClass(), "/templates")
	var template = cfg.GetTemplate("bindCodeEmail.ftl")
	var html = FreeMarkerTemplateUtils.ProcessTemplateIntoString(template, model)
	helper.SetText(html, true)
	//发送邮件
	this.JavaMailSender.Send(mimeMessage)
	log.Infof("send email for %v,content:%v", email, html)
	this.RedisUtil.Set(EMAIL_BIND_CODE_PREFIX+email, code, 10, time.Minute)
}
func (this *RegisterController) SendAddAddress(ctx *gin.Context, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var code = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	var member = this.MemberService.FindOne(user.GetId())
	var email = member.GetEmail()
	if email == nil {
		return error(this.LocaleMessageSourceService.GetMessage("NOT_BIND_EMAIL"))
	}
	if this.RedisUtil.Get(ADD_ADDRESS_CODE_PREFIX+email) != nil {
		return error(this.LocaleMessageSourceService.GetMessage("EMAIL_ALREADY_SEND"))
	}
	exception := func() (err error) {
		this.SentEmailAddCode(email, code)
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
		return error(this.LocaleMessageSourceService.GetMessage("SEND_FAILED"))
	}
	return success(this.LocaleMessageSourceService.GetMessage("SENT_SUCCESS_TEN"))
}
func (this *RegisterController) SentEmailAddCode(ctx *gin.Context, email string, code string) (err error) {
	var mimeMessage = this.JavaMailSender.CreateMimeMessage()
	var helper *javamail.MimeMessageHelper
	helper = MimeMessageHelper(mimeMessage, true)
	helper.SetFrom(this.From)
	helper.SetTo(email)
	helper.SetSubject(this.Company)
	var model = hashset.New[interface {
	}]()
	model.Put("code", code)
	var cfg = Configuration(Configuration.VERSION_2_3_26)
	cfg.SetClassForTemplateLoading(this.GetClass(), "/templates")
	var template = cfg.GetTemplate("addAddressCodeEmail.ftl")
	var html = FreeMarkerTemplateUtils.ProcessTemplateIntoString(template, model)
	helper.SetText(html, true)
	//发送邮件
	this.JavaMailSender.Send(mimeMessage)
	this.RedisUtil.Set(ADD_ADDRESS_CODE_PREFIX+email, code, 10, time.Minute)
}
func (this *RegisterController) SendResetPasswordCode(ctx *gin.Context, account string) (result *MessageResult.MessageResult) {
	var member = this.MemberService.FindByEmail(account)
	if member == nil {
		this.LocaleMessageSourceService.GetMessage("MEMBER_NOT_EXISTS")
	}
	if this.RedisUtil.Get(RESET_PASSWORD_CODE_PREFIX+account) != nil {
		return error(this.LocaleMessageSourceService.GetMessage("EMAIL_ALREADY_SEND"))
	}
	exception := func() (err error) {
		var code = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
		this.SentResetPassword(account, code, RESET_PASSWORD_CODE_PREFIX)
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
		return error(this.LocaleMessageSourceService.GetMessage("SEND_FAILED"))
	}
	return success(this.LocaleMessageSourceService.GetMessage("SENT_SUCCESS_TEN"))
}
func (this *RegisterController) SentResetPassword(ctx *gin.Context, email string, code string, prefix string) (err error) {
	var mimeMessage = this.JavaMailSender.CreateMimeMessage()
	var helper *javamail.MimeMessageHelper
	helper = MimeMessageHelper(mimeMessage, true)
	helper.SetFrom(this.From)
	helper.SetTo(email)
	helper.SetSubject(this.Company)
	var model = hashset.New[interface {
	}]()
	model.Put("code", code)
	var cfg = Configuration(Configuration.VERSION_2_3_26)
	cfg.SetClassForTemplateLoading(this.GetClass(), "/templates")
	var template = cfg.GetTemplate("resetPasswordCodeEmail.ftl")
	var html = FreeMarkerTemplateUtils.ProcessTemplateIntoString(template, model)
	helper.SetText(html, true)
	//发送邮件
	this.JavaMailSender.Send(mimeMessage)
	this.RedisUtil.Set(prefix+email, code, 10, time.Minute)
}
func (this *RegisterController) ForgetPassword(ctx *gin.Context, mode int, account string, code string, password string) (result *MessageResult.MessageResult, err error) {
	isTrue(ValidateUtils.ValidatePassword(password), this.MsService.GetMessage("PASSWORD_LENGTH_ILLEGAL"))
	var member *entity.Member
	var redisCode = this.RedisUtil.Get(SysConstant.RESET_PASSWORD_CODE_PREFIX + account)
	notNull(redisCode, this.LocaleMessageSourceService.GetMessage("VERIFICATION_CODE_NOT_EXISTS"))
	if mode == 0 {
		member = this.MemberService.FindByPhone(account)
	} else if mode == 1 {
		member = this.MemberService.FindByEmail(account)
	}
	isTrue(len(password) >= 6 && len(password) <= 20, this.LocaleMessageSourceService.GetMessage("PASSWORD_LENGTH_ILLEGAL"))
	notNull(member, this.LocaleMessageSourceService.GetMessage("MEMBER_NOT_EXISTS"))
	if !code.Equals(redisCode.ToString()) {
		return error(this.LocaleMessageSourceService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	} else {
		this.RedisUtil.Delete(SysConstant.RESET_PASSWORD_CODE_PREFIX + account)
	}
	//生成密码
	var newPassword = Md5.Md5Digest(password + member.GetSalt()).ToLowerCase()
	member.SetPassword(newPassword)
	member.SetLoginLock(BooleanEnum.IS_FALSE)
	return new(MessageResult.success)
}
func (this *RegisterController) UntieEmailCode(ctx *gin.Context, user *transform.AuthMember) (result *MessageResult.MessageResult) {
	var member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetEmail() != nil, this.MsService.GetMessage("NOT_BIND_EMAIL"))
	var cache = this.RedisUtil.Get(SysConstant.EMAIL_UNTIE_CODE_PREFIX + member.GetEmail())
	if cache != nil {
		return error(this.LocaleMessageSourceService.GetMessage("EMAIL_ALREADY_SEND"))
	}
	var code = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	exception := func() (err error) {
		this.SentResetPassword(member.GetEmail(), code, EMAIL_UNTIE_CODE_PREFIX)
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
	}
	return new(MessageResult.success)
}
func (this *RegisterController) UpdateEmailCode(ctx *gin.Context, user *transform.AuthMember, email string) (result *MessageResult.MessageResult) {
	if this.MemberService.EmailIsExist(email) {
		return MessageResult.Error(this.MsService.GetMessage("REPEAT_EMAIL_REQUEST"))
	}
	var member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetEmail() != nil, this.MsService.GetMessage("NOT_BIND_EMAIL"))
	var cache = this.RedisUtil.Get(SysConstant.EMAIL_UPDATE_CODE_PREFIX + email)
	if cache != nil {
		return error(this.LocaleMessageSourceService.GetMessage("EMAIL_ALREADY_SEND"))
	}
	var code = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	exception := func() (err error) {
		this.SentResetPassword(email, code, EMAIL_UPDATE_CODE_PREFIX)
		return
	}()
	if exception != nil {
		e.PrintStackTrace()
	}
	return new(MessageResult.success)
}
func NewRegisterController(javaMailSender *javamail.JavaMailSender, from string, host string, company string, redisUtil *util.RedisUtil, memberService *service.MemberService, idWorkByTwitter *util.IdWorkByTwitter, memberEvent *event.MemberEvent, countryService *service.CountryService, msService *service.LocaleMessageSourceService, localeMessageSourceService *service.LocaleMessageSourceService, userNameFormat string) (ret *RegisterController) {
	ret = new(RegisterController)
	ret.JavaMailSender = javaMailSender
	ret.From = from
	ret.Host = host
	ret.Company = company
	ret.RedisUtil = redisUtil
	ret.MemberService = memberService
	ret.IdWorkByTwitter = idWorkByTwitter
	ret.MemberEvent = memberEvent
	ret.CountryService = countryService
	ret.MsService = msService
	ret.LocaleMessageSourceService = localeMessageSourceService
	ret.UserNameFormat = userNameFormat
	return
}

type RegisterController struct {
	JavaMailSender             *javamail.JavaMailSender
	From                       string
	Host                       string
	Company                    string
	RedisUtil                  *util.RedisUtil
	MemberService              *service.MemberService
	IdWorkByTwitter            *util.IdWorkByTwitter
	MemberEvent                *event.MemberEvent
	CountryService             *service.CountryService
	MsService                  *service.LocaleMessageSourceService
	LocaleMessageSourceService *service.LocaleMessageSourceService
	UserNameFormat             string
	controller.BaseController
}
