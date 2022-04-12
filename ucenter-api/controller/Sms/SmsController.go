package Sms

import (
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/BigDecimalUtils"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/xtime"
	"time"
)

func (this *SmsController) SendCheckCode(ctx *gin.Context, phone string, country string) (result *util.MessageResult, err error) {
	if this.MemberService.PhoneIsExist(phone) {
		this.LocaleMessageSourceService.GetMessage("PHONE_ALREADY_EXISTS")
	}
	if country == nil {
		this.LocaleMessageSourceService.GetMessage("REQUEST_ILLEGAL")
	}
	var country1 = this.CountryService.FindOne(country)
	if country1 == nil {
		this.LocaleMessageSourceService.GetMessage("REQUEST_ILLEGAL")
	}
	var key = SysConstant.PHONE_REG_CODE_PREFIX + phone
	var code = this.RedisUtil.Get(key)
	if code != nil {
		//判断如果请求间隔小于一分钟则请求失败
		if !BigDecimalUtils.Compare(DateUtil.DiffMinute(this.RedisUtil.Get(key+"Time").(xtime.Xtime)), BigDecimal.ONE) {
			return error(this.LocaleMessageSourceService.GetMessage("FREQUENTLY_REQUEST"))
		}
	}
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	log.Info("fakeParameters====" + randomCode)
	var result *util.MessageResult
	//253国际短信，可以发国内号码，都要加上区域号
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(country1.GetAreaCode()+phone, randomCode)
	} else {
		if country1.GetAreaCode().Equals("86") {
			if ValidateUtil.IsMobilePhone(phone.Trim()) == false {
				this.LocaleMessageSourceService.GetMessage("PHONE_EMPTY_OR_INCORRECT")
			}
			result = this.SmsProvider.SendVerifyMessage(phone, randomCode)
		} else {
			result = this.SmsProvider.SendInternationalMessage(randomCode, country1.GetAreaCode()+phone)
		}
	}
	if result.GetCode() == 0 {
		this.RedisUtil.Delete(key)
		this.RedisUtil.Delete(key + "Time")
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		this.RedisUtil.Set(key+"Time", time.Now(), 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) SendSMSCode(ctx *gin.Context, member *entity.Member, prefix string) (result *util.MessageResult) {
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	var result *util.MessageResult
	exception := func() (err error) {
		if this.DriverName.EqualsIgnoreCase("two_five_three") {
			result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
		} else if member.GetCountry().GetAreaCode().Equals("86") {
			result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
		} else {
			result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
		}
		if result.GetCode() == 0 {
			var key = prefix + member.GetMobilePhone()
			this.RedisUtil.Delete(key)
			// 缓存验证码
			this.RedisUtil.Set(key, randomCode, 10, time.Minute)
			return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
		} else {
			return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
		}
		return
	}()
	if exception != nil {
		return MessageResult.Error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) SendResetTransactionCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	var result *util.MessageResult
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if member.GetCountry().GetAreaCode().Equals("86") {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_UPDATE_PASSWORD_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) SetBindPhoneCode(ctx *gin.Context, country string, phone string, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() != nil {
		this.LocaleMessageSourceService.GetMessage("REPEAT_PHONE_REQUEST")
	}
	var result *util.MessageResult
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	// 修改所在国家
	if StringUtils.IsNotBlank(country) {
		var one = this.CountryService.FindOne(country)
		if one != nil {
			member.SetCountry(one)
			this.MemberService.SaveAndFlush(member)
		}
	}
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+phone, randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		if !ValidateUtil.IsMobilePhone(phone.Trim()) {
			return error(this.LocaleMessageSourceService.GetMessage("PHONE_EMPTY_OR_INCORRECT"))
		}
		result = this.SmsProvider.SendVerifyMessage(phone, randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+phone)
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_BIND_CODE_PREFIX + phone
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) UpdatePasswordCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var key = SysConstant.PHONE_UPDATE_PASSWORD_PREFIX + member.GetMobilePhone()
	var code = this.RedisUtil.Get(key)
	if code != nil {
		//判断如果请求间隔小于一分钟则请求失败
		if !BigDecimalUtils.Compare(DateUtil.DiffMinute(this.RedisUtil.Get(key+"Time").(xtime.Xtime)), BigDecimal.ONE) {
			return error(this.LocaleMessageSourceService.GetMessage("FREQUENTLY_REQUEST"))
		}
	}
	var result *util.MessageResult
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		this.RedisUtil.Delete(key)
		this.RedisUtil.Delete(key + "Time")
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		this.RedisUtil.Set(key+"Time", time.Now(), 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) AddAddressCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_ADD_ADDRESS_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) ResetPasswordCode(ctx *gin.Context, account string) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindByPhone(account)
	if member == nil {
		this.LocaleMessageSourceService.GetMessage("MEMBER_NOT_EXISTS")
	}
	var result *util.MessageResult
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.RESET_PASSWORD_CODE_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) ResetPhoneCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_CHANGE_CODE_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) ResetGoogleCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member == nil {
		this.LocaleMessageSourceService.GetMessage("MEMBER_NOT_EXISTS")
	}
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.RESET_GOOGLE_CODE_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) WithdrawCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	log.Info("===提币验证码发送===mobile：" + member.GetMobilePhone())
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_WITHDRAW_MONEY_CODE_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) TradeCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	log.Info("===交易密码验证码发送===mobile：" + member.GetMobilePhone())
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_trade_CODE_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) BindApiSendCode(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	log.Info("===交易密码验证码发送===mobile：" + member.GetMobilePhone())
	var randomCode = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.API_BIND_CODE_PREFIX + member.GetMobilePhone()
		this.RedisUtil.Delete(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func NewSmsController(smsProvider *provider.SMSProvider, redisUtil *util.RedisUtil, memberService *service.MemberService, localeMessageSourceService *service.LocaleMessageSourceService, countryService *service.CountryService, driverName string) (ret *SmsController) {
	ret = new(SmsController)
	ret.SmsProvider = smsProvider
	ret.RedisUtil = redisUtil
	ret.MemberService = memberService
	ret.LocaleMessageSourceService = localeMessageSourceService
	ret.CountryService = countryService
	ret.DriverName = driverName
	return
}

type SmsController struct {
	SmsProvider                provider.SMSProvider
	RedisUtil                  *util.RedisUtil
	MemberService              *service.MemberService
	LocaleMessageSourceService *service.LocaleMessageSourceService
	CountryService             *service.CountryService
	DriverName                 string
}
