package controller

import (
	SysConstant "bitrade/core/constant/SysConstants"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/util/BigDecimalUtils"
	"time"
)

func (this *SmsController) SendCheckCode(phone string, country string) (result *util.MessageResult, err error) {
	if this.MemberService.PhoneIsExist(phone) {
		this.LocaleMessageSourceService.GetMessage("PHONE_ALREADY_EXISTS")
	}
	if country == nil {
		this.LocaleMessageSourceService.GetMessage("REQUEST_ILLEGAL")
	}
	var country1 *entity.Country = this.CountryService.FindOne(country)
	if country1 == nil {
		this.LocaleMessageSourceService.GetMessage("REQUEST_ILLEGAL")
	}
	var key = SysConstant.PHONE_REG_CODE_PREFIX + phone
	var code interface {
	} = this.RedisUtil.Get(key)
	if code != nil {
		//判断如果请求间隔小于一分钟则请求失败
		if !BigDecimalUtils.Compare(DateUtil.DiffMinute(this.RedisUtil.Get(key+"Time").(time.Time)), BigDecimal.ONE) {
			return error(this.LocaleMessageSourceService.GetMessage("FREQUENTLY_REQUEST"))
		}
	}
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
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
		core.DelKey(key)
		core.DelKey(key + "Time")
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		this.RedisUtil.Set(key+"Time", time.Now(), 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) SendSMSCode(member *entity.Member, prefix string) (result *util.MessageResult) {
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
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
			core.DelKey(key)
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
func (this *SmsController) SendResetTransactionCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
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
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) SetBindPhoneCode(country string, phone string, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() != nil {
		this.LocaleMessageSourceService.GetMessage("REPEAT_PHONE_REQUEST")
	}
	var result *util.MessageResult
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	// 修改所在国家
	if StringUtils.IsNotBlank(country) {
		var one *entity.Country = this.CountryService.FindOne(country)
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
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) UpdatePasswordCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var key = SysConstant.PHONE_UPDATE_PASSWORD_PREFIX + member.GetMobilePhone()
	var code interface {
	} = this.RedisUtil.Get(key)
	if code != nil {
		//判断如果请求间隔小于一分钟则请求失败
		if !BigDecimalUtils.Compare(DateUtil.DiffMinute(this.RedisUtil.Get(key+"Time").(time.Time)), BigDecimal.ONE) {
			return error(this.LocaleMessageSourceService.GetMessage("FREQUENTLY_REQUEST"))
		}
	}
	var result *util.MessageResult
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		core.DelKey(key)
		core.DelKey(key + "Time")
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		this.RedisUtil.Set(key+"Time", time.Now(), 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) AddAddressCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_ADD_ADDRESS_PREFIX + member.GetMobilePhone()
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) ResetPasswordCode(account string) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindByPhone(account)
	if member == nil {
		this.LocaleMessageSourceService.GetMessage("MEMBER_NOT_EXISTS")
	}
	var result *util.MessageResult
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.RESET_PASSWORD_CODE_PREFIX + member.GetMobilePhone()
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) ResetPhoneCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_CHANGE_CODE_PREFIX + member.GetMobilePhone()
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) ResetGoogleCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member == nil {
		this.LocaleMessageSourceService.GetMessage("MEMBER_NOT_EXISTS")
	}
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if this.DriverName.EqualsIgnoreCase("two_five_three") {
		result = this.SmsProvider.SendVerifyMessage(member.GetCountry().GetAreaCode()+member.GetMobilePhone(), randomCode)
	} else if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.RESET_GOOGLE_CODE_PREFIX + member.GetMobilePhone()
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) WithdrawCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	log.Info("===提币验证码发送===mobile：" + member.GetMobilePhone())
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_WITHDRAW_MONEY_CODE_PREFIX + member.GetMobilePhone()
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) TradeCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	log.Info("===交易密码验证码发送===mobile：" + member.GetMobilePhone())
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.PHONE_trade_CODE_PREFIX + member.GetMobilePhone()
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}
func (this *SmsController) BindApiSendCode(user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetMobilePhone() == "" {
		this.LocaleMessageSourceService.GetMessage("NOT_BIND_PHONE")
	}
	var result *util.MessageResult
	log.Info("===交易密码验证码发送===mobile：" + member.GetMobilePhone())
	var randomCode string = String.ValueOf(GeneratorUtil.GetRandomNumber(100000, 999999))
	if "86".Equals(member.GetCountry().GetAreaCode()) {
		result = this.SmsProvider.SendVerifyMessage(member.GetMobilePhone(), randomCode)
	} else {
		result = this.SmsProvider.SendInternationalMessage(randomCode, member.GetCountry().GetAreaCode()+member.GetMobilePhone())
	}
	if result.GetCode() == 0 {
		var key = SysConstant.API_BIND_CODE_PREFIX + member.GetMobilePhone()
		core.DelKey(key)
		// 缓存验证码
		this.RedisUtil.Set(key, randomCode, 10, time.Minute)
		return success(this.LocaleMessageSourceService.GetMessage("SEND_SMS_SUCCESS"))
	} else {
		return error(this.LocaleMessageSourceService.GetMessage("SEND_SMS_FAILED"))
	}
}

type SmsController struct {
	SmsProvider                provider.SMSProvider
	RedisUtil                  *util.RedisUtil
	MemberService              *service.MemberService
	LocaleMessageSourceService *service.LocaleMessageSourceService
	CountryService             *service.CountryService
	DriverName                 string
}
