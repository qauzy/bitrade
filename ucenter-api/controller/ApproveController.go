package controller

import (
	"bitrade/core/constant/AdvertiseControlStatus"
	"bitrade/core/constant/AuditStatus"
	"bitrade/core/constant/CertifiedBusinessStatus"
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/MemberLevelEnum"
	"bitrade/core/constant/RealNameStatus"
	SysConstant "bitrade/core/constant/SysConstants"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/pagination"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/util/lists/arraylist"
	"time"
)

var Logger slf4j.Logger = LoggerFactory.GetLogger(ApproveController)

func (this *ApproveController) Update(user *transform.AuthMember, url string) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	member.SetAvatar(url)
	return MessageResult.Success()
}
func (this *ApproveController) SecuritySetting(user *transform.AuthMember) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var idNumber string = member.GetIdNumber()
	var memberSecurity *entity.MemberSecurity = new(MemberSecurity).SetUsername(member.GetUsername()).SetCreateTime(member.GetRegistrationTime()).SetId(member.GetId()).SetEmailVerified(func() {
		if StringUtils.IsEmpty(member.GetEmail()) {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetEmail(member.GetEmail()).SetMobilePhone(member.GetMobilePhone()).SetFundsVerified(func() {
		if StringUtils.IsEmpty(member.GetJyPassword()) {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetLoginVerified(IS_TRUE).SetPhoneVerified(func() {
		if StringUtils.IsEmpty(member.GetMobilePhone()) {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetRealName(member.GetRealName()).SetIdCard(func() {
		if StringUtils.IsEmpty(idNumber) {
			return nil
		} else {
			return idNumber.Substring(0, 2) + "**********" + idNumber.Substring(len(idNumber)-2)
		}
	}()).SetRealVerified(func() {
		if StringUtils.IsEmpty(member.GetRealName()) {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetRealAuditing(func() {
		if member.GetRealNameStatus().Equals(RealNameStatus.AUDITING) {
			return IS_TRUE
		} else {
			return IS_FALSE
		}
	}()).SetAvatar(member.GetAvatar()).SetAccountVerified(func() {
		if member.GetBankInfo() == nil && member.GetAlipay() == nil && member.GetWechatPay() == nil {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetGoogleStatus(member.GetGoogleState()).SetTransactions(member.GetTransactions()).SetTransactionTime(member.GetTransactionTime()).SetLevel(member.GetMemberLevel().GetOrdinal()).SetIntegration(member.GetIntegration()).SetKycStatus(member.GetKycStatus()).SetMemberGradeId(member.GetMemberGradeId()).SetGoogleState(member.GetGoogleState()).SetMemberLevel(member.GetMemberLevel())
	if memberSecurity.GetRealAuditing().Equals(IS_FALSE) && memberSecurity.GetRealVerified().Equals(IS_FALSE) {
		var memberApplication *arraylist.List[entity.MemberApplication] = this.MemberApplicationService.FindLatelyReject(member)
		memberSecurity.SetRealNameRejectReason(func() {
			if memberApplication == nil || memberApplication.Size() == 0 {
				return nil
			} else {
				return memberApplication.Get(0).GetRejectReason()
			}
		}())
	} else if member.GetKycStatus() == 3 {
		var kycStatus *arraylist.List[int] = Arrays.AsList(3)
		var memberApplication *entity.MemberApplication = this.MemberApplicationService.FindMemberApplicationByKycStatusIn(kycStatus, member)
		if memberApplication != nil {
			memberSecurity.SetRealNameRejectReason(memberApplication.GetRejectReason())
		}
	}
	var result *util.MessageResult = MessageResult.Success("success")
	result.SetData(memberSecurity)
	return result
}
func (this *ApproveController) ApproveTransaction(jyPassword string, user *transform.AuthMember) (result *util.MessageResult, err error) {
	//校验密码
	isTrue(ValidateUtils.ValidatePassword(jyPassword), this.MsService.GetMessage("JY_PASSWORD_LENGTH_ILLEGAL"))
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetJyPassword() != nil {
		this.MsService.GetMessage("REPEAT_SETTING")
	}
	//生成密码
	var jyPass string = Md5.Md5Digest(jyPassword + member.GetSalt()).ToLowerCase()
	member.SetJyPassword(jyPass)
	return MessageResult.Success(this.MsService.GetMessage("SETTING_JY_PASSWORD"))
}
func (this *ApproveController) UpdateTransaction(oldPassword string, newPassword string, msgCode string, googleCode string, user *transform.AuthMember) (result *util.MessageResult, err error) {
	isTrue(ValidateUtils.ValidatePassword(newPassword), this.MsService.GetMessage("JY_PASSWORD_LENGTH_ILLEGAL"))
	isTrue(org.Apache.Commons.Lang3.StringUtils.IsNotEmpty(msgCode), "请输入验证码")
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetGoogleState() == 1 {
		//谷歌验证
		if org.Apache.Commons.Lang.StringUtils.IsNotEmpty(googleCode) {
			var googleCodes int64 = Long.ParseLong(googleCode)
			var t int64 = System.CurrentTimeMillis()
			var ga *util.GoogleAuthenticatorUtil = new(util.GoogleAuthenticatorUtil)
			var r bool = ga.Check_code(member.GetGoogleKey(), googleCodes, t)
			if !r {
				return MessageResult.Error("谷歌验证失败")
			}
		} else {
			return MessageResult.Error("请输入谷歌验证码")
		}
	}
	isTrue(Md5.Md5Digest(oldPassword+member.GetSalt()).ToLowerCase().Equals(member.GetJyPassword()), this.MsService.GetMessage("ERROR_JYPASSWORD"))
	member.SetJyPassword(Md5.Md5Digest(newPassword + member.GetSalt()).ToLowerCase())
	return MessageResult.Success(this.MsService.GetMessage("SETTING_JY_PASSWORD"))
}
func (this *ApproveController) ResetTransaction(newPassword string, code string, codeMold int, user *transform.AuthMember) (result *util.MessageResult, err error) {
	hasText(newPassword, this.MsService.GetMessage("MISSING_NEW_JY_PASSWORD"))
	hasText(code, this.MsService.GetMessage("MISSING_VERIFICATION_CODE"))
	isTrue(ValidateUtils.ValidatePassword(newPassword), this.MsService.GetMessage("JY_PASSWORD_LENGTH_ILLEGAL"))
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var key = ""
	if codeMold == 0 {
		key = SysConstant.PHONE_UPDATE_PASSWORD_PREFIX
	} else if codeMold == 1 {
		key = SysConstant.RESET_PASSWORD_CODE_PREFIX
	}
	var res *util.MessageResult = this.CheckCode(key, member, code, codeMold)
	if res.GetCode() != 0 {
		return res
	}
	member.SetJyPassword(Md5.Md5Digest(newPassword + member.GetSalt()).ToLowerCase())
	return MessageResult.Success(this.MsService.GetMessage("SETTING_JY_PASSWORD"))
}
func (this *ApproveController) BindPhone(request *http.HttpServletRequest, password string, phone string, code string, user *transform.AuthMember) (result *util.MessageResult, err error) {
	hasText(password, this.MsService.GetMessage("MISSING_LOGIN_PASSWORD"))
	hasText(phone, this.MsService.GetMessage("MISSING_PHONE"))
	hasText(code, this.MsService.GetMessage("MISSING_VERIFICATION_CODE"))
	if user.GetLocation().GetCountry().Equals("中国") {
		if !ValidateUtil.IsMobilePhone(phone.Trim()) {
			return MessageResult.Error(this.MsService.GetMessage("PHONE_FORMAT_ERROR"))
		}
	}
	var cache interface {
	} = this.RedisUtil.Get(SysConstant.PHONE_BIND_CODE_PREFIX + phone)
	notNull(cache, this.MsService.GetMessage("NO_GET_VERIFICATION_CODE"))
	var member1 *entity.Member = this.MemberService.FindByPhone(phone)
	isTrue(member1 == nil, this.MsService.GetMessage("PHONE_ALREADY_BOUND"))
	if !code.Equals(cache.ToString()) {
		return MessageResult.Error(this.MsService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	} else {
		core.DelKey(SysConstant.PHONE_BIND_CODE_PREFIX + phone)
	}
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetMobilePhone() == nil, this.MsService.GetMessage("REPEAT_PHONE_REQUEST"))
	if member.GetPassword().Equals(Md5.Md5Digest(password + member.GetSalt()).ToLowerCase()) {
		member.SetMobilePhone(phone)
		return MessageResult.Success(this.MsService.GetMessage("SETTING_SUCCESS"))
	} else {
		request.RemoveAttribute(SysConstant.SESSION_MEMBER)
		return MessageResult.Error(this.MsService.GetMessage("PASSWORD_ERROR"))
	}
}
func (this *ApproveController) UpdateLoginPassword(request *http.HttpServletRequest, oldPassword string, newPassword string, code string, googleCode string, codeMold int, user *transform.AuthMember) (result *util.MessageResult, err error) {
	log.Info("code=" + code)
	hasText(oldPassword, this.MsService.GetMessage("MISSING_OLD_PASSWORD"))
	hasText(newPassword, this.MsService.GetMessage("MISSING_NEW_PASSWORD"))
	isTrue(ValidateUtils.ValidatePassword(newPassword), this.MsService.GetMessage("PASSWORD_LENGTH_ILLEGAL"))
	hasText(code, this.MsService.GetMessage("MISSING_VERIFICATION_CODE"))
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if member.GetGoogleState() == 1 {
		//谷歌验证
		if org.Apache.Commons.Lang.StringUtils.IsNotEmpty(googleCode) {
			var googleCodes int64 = Long.ParseLong(googleCode)
			var t int64 = System.CurrentTimeMillis()
			var ga *util.GoogleAuthenticatorUtil = new(util.GoogleAuthenticatorUtil)
			var r bool = ga.Check_code(member.GetGoogleKey(), googleCodes, t)
			if !r {
				return MessageResult.Error("谷歌验证失败")
			}
		} else {
			return MessageResult.Error("请输入谷歌验证码")
		}
	}
	isTrue(Md5.Md5Digest(oldPassword+member.GetSalt()).ToLowerCase().Equals(member.GetPassword()), this.MsService.GetMessage("PASSWORD_ERROR"))
	var key = ""
	if codeMold == 0 {
		key = SysConstant.PHONE_UPDATE_PASSWORD_PREFIX
	} else if codeMold == 1 {
		key = SysConstant.RESET_PASSWORD_CODE_PREFIX
	}
	var res *util.MessageResult = this.CheckCode(key, member, code, codeMold)
	if res.GetCode() != 0 {
		return res
	}
	request.RemoveAttribute(SysConstant.SESSION_MEMBER)
	member.SetPassword(Md5.Md5Digest(newPassword + member.GetSalt()).ToLowerCase())
	return MessageResult.Success(this.MsService.GetMessage("SETTING_SUCCESS"))
}
func (this *ApproveController) BindEmail(request *http.HttpServletRequest, password string, code string, email string, user *transform.AuthMember) (result *util.MessageResult, err error) {
	hasText(password, this.MsService.GetMessage("MISSING_LOGIN_PASSWORD"))
	hasText(code, this.MsService.GetMessage("MISSING_VERIFICATION_CODE"))
	hasText(email, this.MsService.GetMessage("MISSING_EMAIL"))
	isTrue(ValidateUtil.IsEmail(email), this.MsService.GetMessage("EMAIL_FORMAT_ERROR"))
	var cache interface {
	} = this.RedisUtil.Get(SysConstant.EMAIL_BIND_CODE_PREFIX + email)
	notNull(cache, this.MsService.GetMessage("NO_GET_VERIFICATION_CODE"))
	isTrue(code.Equals(cache.ToString()), this.MsService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetEmail() == nil, this.MsService.GetMessage("REPEAT_EMAIL_REQUEST"))
	if !Md5.Md5Digest(password + member.GetSalt()).ToLowerCase().Equals(member.GetPassword()) {
		request.RemoveAttribute(SysConstant.SESSION_MEMBER)
		return MessageResult.Error(this.MsService.GetMessage("PASSWORD_ERROR"))
	} else {
		member.SetEmail(email)
		return MessageResult.Success(this.MsService.GetMessage("SETTING_SUCCESS"))
	}
}
func (this *ApproveController) UpdateEmail(user *transform.AuthMember, oldEmailCode string, newEmailCode string, newEmail string) (result *util.MessageResult) {
	hasText(oldEmailCode, this.MsService.GetMessage("MISSING_VERIFICATION_CODE"))
	hasText(newEmailCode, this.MsService.GetMessage("MISSING_VERIFICATION_CODE"))
	hasText(newEmail, this.MsService.GetMessage("MISSING_EMAIL"))
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetEmail() != nil, this.MsService.GetMessage("NOT_BIND_EMAIL"))
	var oldEmailCache interface {
	} = this.RedisUtil.Get(SysConstant.EMAIL_UNTIE_CODE_PREFIX + member.GetEmail())
	var newEmailCache interface {
	} = this.RedisUtil.Get(SysConstant.EMAIL_UPDATE_CODE_PREFIX + newEmail)
	notNull(oldEmailCache, this.MsService.GetMessage("NO_GET_VERIFICATION_CODE"))
	notNull(newEmailCache, this.MsService.GetMessage("NO_GET_VERIFICATION_CODE"))
	isTrue(oldEmailCode.Equals(oldEmailCache.ToString()), this.MsService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	isTrue(newEmailCode.Equals(newEmailCache.ToString()), this.MsService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	core.DelKey(SysConstant.EMAIL_UNTIE_CODE_PREFIX + member.GetEmail())
	core.DelKey(SysConstant.EMAIL_UPDATE_CODE_PREFIX + newEmail)
	member.SetEmail(newEmail)
	return MessageResult.Success()
}
func (this *ApproveController) RealApprove(user *transform.AuthMember, realName string, idCard string, idCardFront string, idCardBack string, handHeldIdCard string, authType int) (result *util.MessageResult) {
	hasText(realName, this.MsService.GetMessage("MISSING_REAL_NAME"))
	hasText(idCard, this.MsService.GetMessage("MISSING_ID_CARD"))
	hasText(idCardFront, this.MsService.GetMessage("MISSING_ID_CARD_FRONT"))
	hasText(idCardBack, this.MsService.GetMessage("MISSING_ID_CARD_BACK"))
	hasText(handHeldIdCard, this.MsService.GetMessage("MISSING_ID_CARD_HAND"))
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	if "China".Equals(member.GetCountry().GetEnName()) && authType == CredentialsType.CARDED.GetOrdinal() {
		isTrue(ValidateUtil.IsChineseName(realName), this.MsService.GetMessage("REAL_NAME_ILLEGAL"))
		isTrue(IdcardValidator.IsValidate18Idcard(idCard), this.MsService.GetMessage("ID_CARD_ILLEGAL"))
	} else {
		isTrue(len(idCard) < 20 && len(idCard) >= 6, this.MsService.GetMessage("ID_CARD_ILLEGAL"))
	}
	isTrue(member.GetRealNameStatus() == RealNameStatus.NOT_CERTIFIED, this.MsService.GetMessage("REPEAT_REAL_NAME_REQUEST"))
	if this.RealTimes != 0 {
		isTrue(this.MemberApplicationService.FindSuccessRealAuthByIdCard(idCard) < this.RealTimes, this.MsService.GetMessage("LIMIT_REAL_NAME_TIMES"))
	}
	var credentialsType *enums.CredentialsType = CredentialsType.GetByValue(authType)
	if credentialsType == nil {
		return MessageResult.Error(this.MsService.GetMessage("ILLEGAL_AUTHENTICATION_TYPE"))
	}
	var memberApplication *entity.MemberApplication = new(entity.MemberApplication)
	//认证类型
	memberApplication.SetType(credentialsType)
	memberApplication.SetAuditStatus(AuditStatus.AUDIT_ING)
	memberApplication.SetRealName(realName)
	memberApplication.SetIdCard(idCard)
	memberApplication.SetMember(member)
	memberApplication.SetIdentityCardImgFront(idCardFront)
	memberApplication.SetIdentityCardImgInHand(handHeldIdCard)
	memberApplication.SetIdentityCardImgReverse(idCardBack)
	memberApplication.SetCreateTime(time.Now())
	memberApplication.SetKycStatus(5)
	this.MemberApplicationService.Save(memberApplication)
	member.SetRealNameStatus(RealNameStatus.AUDITING)
	return MessageResult.Success(this.MsService.GetMessage("REAL_APPLY_SUCCESS"))
}
func (this *ApproveController) RealApproveVideo(user *transform.AuthMember, videoStr string, random string) (result *util.MessageResult) {
	hasText(videoStr, "URL为空")
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetKycStatus() == 1 || member.GetKycStatus() == 3, "请先完成实名认证")
	//查询待二级审核或二级审核失败的
	var status *arraylist.List[int] = Arrays.AsList(1, 3)
	var memberApplication *entity.MemberApplication = this.MemberApplicationService.FindMemberApplicationByKycStatusIn(status, member)
	if memberApplication == nil {
		return MessageResult.Error("请先完成实名认证")
	}
	//认证类型
	memberApplication.SetKycStatus(6)
	memberApplication.SetVideoUrl(videoStr)
	memberApplication.SetVideoRandom(random)
	this.MemberApplicationService.Save(memberApplication)
	member.SetKycStatus(6)
	this.MemberService.Save(member)
	return MessageResult.Success("提交成功，等待审核")
}
func (this *ApproveController) RealNameApproveDetail(user *transform.AuthMember) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var predicateList = arraylist.New[types.Predicate]()
	predicateList.Add(QMemberApplication.MemberApplication.Member.Eq(member))
	var memberApplicationPageResult *pagination.PageResult = this.MemberApplicationService.Query(predicateList, nil, nil)
	var memberApplication *entity.MemberApplication = new(entity.MemberApplication)
	if memberApplicationPageResult != nil && memberApplicationPageResult.GetContent() != nil && memberApplicationPageResult.GetContent().Size() > 0 {
		memberApplication = memberApplicationPageResult.GetContent().Get(0)
	}
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(memberApplication)
	return result
}
func (this *ApproveController) AccountSetting(user *transform.AuthMember) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	hasText(member.GetIdNumber(), this.MsService.GetMessage("NO_REAL_NAME"))
	hasText(member.GetJyPassword(), this.MsService.GetMessage("NO_JY_PASSWORD"))
	var memberAccount *entity.MemberAccount = new(MemberAccount).SetAlipay(member.GetAlipay()).SetAliVerified(func() {
		if member.GetAlipay() == nil {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetBankInfo(member.GetBankInfo()).SetBankVerified(func() {
		if member.GetBankInfo() == nil {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetWechatPay(member.GetWechatPay()).SetWechatVerified(func() {
		if member.GetWechatPay() == nil {
			return IS_FALSE
		} else {
			return IS_TRUE
		}
	}()).SetRealName(member.GetRealName()).SetMemberLevel(member.GetMemberLevel().GetOrdinal())
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(memberAccount)
	return result
}
func (this *ApproveController) BindBank(BindBank *entity.BindBank, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetBankInfo() == nil, this.MsService.GetMessage("REPEAT_SETTING"))
	return this.DoBank(this.BindBank, bindingResult, user)
}
func (this *ApproveController) DoBank(BindBank *entity.BindBank, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var result *util.MessageResult = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(Md5.Md5Digest(this.BindBank.GetJyPassword()+member.GetSalt()).ToLowerCase().Equals(member.GetJyPassword()), this.MsService.GetMessage("ERROR_JYPASSWORD"))
	var bankInfo *entity.BankInfo = new(entity.BankInfo)
	bankInfo.SetBank(this.BindBank.GetBank())
	bankInfo.SetBranch(this.BindBank.GetBranch())
	bankInfo.SetCardNo(this.BindBank.GetCardNo())
	member.SetBankInfo(bankInfo)
	return MessageResult.Success(this.MsService.GetMessage("SETTING_SUCCESS"))
}
func (this *ApproveController) UpdateBank(BindBank *entity.BindBank, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	return this.DoBank(this.BindBank, bindingResult, user)
}
func (this *ApproveController) BindAli(BindAli *entity.BindAli, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetAlipay() == nil, this.MsService.GetMessage("REPEAT_SETTING"))
	return this.DoAli(this.BindAli, bindingResult, user)
}
func (this *ApproveController) DoAli(BindAli *entity.BindAli, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var result *util.MessageResult = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(Md5.Md5Digest(this.BindAli.GetJyPassword()+member.GetSalt()).ToLowerCase().Equals(member.GetJyPassword()), this.MsService.GetMessage("ERROR_JYPASSWORD"))
	var alipay *entity.Alipay = new(entity.Alipay)
	alipay.SetAliNo(this.BindAli.GetAli())
	alipay.SetQrCodeUrl(this.BindAli.GetQrCodeUrl())
	member.SetAlipay(alipay)
	return MessageResult.Success(this.MsService.GetMessage("SETTING_SUCCESS"))
}
func (this *ApproveController) UpdateAli(BindAli *entity.BindAli, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	return this.DoAli(this.BindAli, bindingResult, user)
}
func (this *ApproveController) BindWechat(BindWechat *entity.BindWechat, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(member.GetWechatPay() == nil, this.MsService.GetMessage("REPEAT_SETTING"))
	return this.DoWechat(this.BindWechat, bindingResult, user)
}
func (this *ApproveController) DoWechat(BindWechat *entity.BindWechat, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var result *util.MessageResult = BindingResultUtil.Validate(bindingResult)
	if result != nil {
		return result
	}
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	isTrue(Md5.Md5Digest(this.BindWechat.GetJyPassword()+member.GetSalt()).ToLowerCase().Equals(member.GetJyPassword()), this.MsService.GetMessage("ERROR_JYPASSWORD"))
	var wechatPay *entity.WechatPay = new(entity.WechatPay)
	wechatPay.SetWechat(this.BindWechat.GetWechat())
	wechatPay.SetQrWeCodeUrl(this.BindWechat.GetQrCodeUrl())
	member.SetWechatPay(wechatPay)
	return MessageResult.Success(this.MsService.GetMessage("SETTING_SUCCESS"))
}
func (this *ApproveController) UpdateWechat(BindWechat *entity.BindWechat, bindingResult *validation.BindingResult, user *transform.AuthMember) (result *util.MessageResult, err error) {
	return this.DoWechat(this.BindWechat, bindingResult, user)
}
func (this *ApproveController) CertifiedBusinessStatus(user *transform.AuthMember) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var certifiedBusinessInfo *entity.CertifiedBusinessInfo = new(entity.CertifiedBusinessInfo)
	certifiedBusinessInfo.SetCertifiedBusinessStatus(member.GetCertifiedBusinessStatus())
	certifiedBusinessInfo.SetEmail(member.GetEmail())
	certifiedBusinessInfo.SetMemberLevel(member.GetMemberLevel())
	Logger.Info("会员状态信息:{}", certifiedBusinessInfo)
	if member.GetCertifiedBusinessStatus().Equals(CertifiedBusinessStatus.FAILED) {
		var businessAuthApplyList *arraylist.List[entity.BusinessAuthApply] = this.BusinessAuthApplyService.FindByMemberAndCertifiedBusinessStatus(member, member.GetCertifiedBusinessStatus())
		Logger.Info("会员申请商家认证信息:{}", businessAuthApplyList)
		if businessAuthApplyList != nil && businessAuthApplyList.Size() > 0 {
			certifiedBusinessInfo.SetCertifiedBusinessStatus(businessAuthApplyList.Get(0).GetCertifiedBusinessStatus())
			Logger.Info("会员申请商家认证最新信息:{}", businessAuthApplyList.Get(0))
			certifiedBusinessInfo.SetDetail(businessAuthApplyList.Get(0).GetDetail())
		}
	}
	var businessCancelApplies *arraylist.List[entity.BusinessCancelApply] = this.BusinessCancelApplyService.FindByMember(member)
	if businessCancelApplies != nil && businessCancelApplies.Size() > 0 {
		if businessCancelApplies.Get(0).GetStatus() == RETURN_SUCCESS {
			if member.GetCertifiedBusinessStatus() != VERIFIED {
				certifiedBusinessInfo.SetCertifiedBusinessStatus(RETURN_SUCCESS)
			}
		} else if businessCancelApplies.Get(0).GetStatus() == RETURN_FAILED {
			certifiedBusinessInfo.SetCertifiedBusinessStatus(RETURN_FAILED)
			certifiedBusinessInfo.SetReason(businessCancelApplies.Get(0).GetDetail())
		} else {
			certifiedBusinessInfo.SetCertifiedBusinessStatus(CANCEL_AUTH)
		}
	}
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(certifiedBusinessInfo)
	return result
}
func (this *ApproveController) CertifiedBusiness(user *transform.AuthMember, json string, businessAuthDepositId int64) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	//只有未认证和认证失败的用户，可以发起认证申请
	isTrue(member.GetCertifiedBusinessStatus().Equals(CertifiedBusinessStatus.NOT_CERTIFIED) || member.GetCertifiedBusinessStatus().Equals(CertifiedBusinessStatus.FAILED), this.MsService.GetMessage("REPEAT_APPLICATION"))
	isTrue(member.GetMemberLevel().Equals(MemberLevelEnum.REALNAME), this.MsService.GetMessage("NO_REAL_NAME"))
	//hasText(member.getEmail(), msService.getMessage("NOT_BIND_EMAIL"));
	//检查有没有资产证明和交易证明图片
	var assetImg string = JSONObject.ParseObject(json).GetString("assetData").Trim()
	var tradeDataImg string = JSONObject.ParseObject(json).GetString("tradeData").Trim()
	if assetImg == nil || assetImg.EqualsIgnoreCase("null") || assetImg.Equals("") || tradeDataImg == nil || tradeDataImg.EqualsIgnoreCase("null") || tradeDataImg.Equals("") {
		return MessageResult.Error(this.MsService.GetMessage("IMG_NOT_NULL"))
	}
	var depositList *arraylist.List[entity.BusinessAuthDeposit] = this.BusinessAuthDepositService.FindAllByStatus(CommonStatus.NORMAL)
	//如果当前有启用的保证金类型，必须选择一种保证金才可以申请商家认证
	var businessAuthDeposit *entity.BusinessAuthDeposit
	if depositList != nil && depositList.Size() > 0 {
		if businessAuthDepositId == nil {
			return MessageResult.Error(this.MsService.GetMessage("MUST_SELECT_BUSINESS_DEPOSIT"))
		}
		var flag = false
		for _, deposit := range depositList {
			if deposit.GetId().Equals(businessAuthDepositId) {
				businessAuthDeposit = deposit
				flag = true
			}
		}
		if !flag {
			return MessageResult.Error(this.MsService.GetMessage("BUSINESS_DEPOSIT_NOT_FOUND"))
		}
		var memberWallet *entity.OtcWallet = this.OtcWalletService.FindByCoinAndMember(member.GetId(), businessAuthDeposit.GetCoin())
		if memberWallet == nil {
			return error("请先划转法币账户")
		}
		if memberWallet.GetBalance().CompareTo(businessAuthDeposit.GetAmount()) < 0 {
			return MessageResult.Error(this.MsService.GetMessage("BALANCE_RUNNING_LOW"))
		}
		//冻结保证金需要的金额
		memberWallet.SetBalance(memberWallet.GetBalance().Subtract(businessAuthDeposit.GetAmount()))
		memberWallet.SetFrozenBalance(memberWallet.GetFrozenBalance().Add(businessAuthDeposit.GetAmount()))
	}
	//申请记录
	var businessAuthApply *entity.BusinessAuthApply = new(entity.BusinessAuthApply)
	businessAuthApply.SetCreateTime(time.Now())
	businessAuthApply.SetAuthInfo(json)
	businessAuthApply.SetCertifiedBusinessStatus(CertifiedBusinessStatus.AUDITING)
	businessAuthApply.SetMember(member)
	//不一定会有保证金策略
	if businessAuthDeposit != nil {
		businessAuthApply.SetBusinessAuthDeposit(businessAuthDeposit)
		businessAuthApply.SetAmount(businessAuthDeposit.GetAmount())
	}
	this.BusinessAuthApplyService.Create(businessAuthApply)
	member.SetCertifiedBusinessApplyTime(time.Now())
	member.SetCertifiedBusinessStatus(CertifiedBusinessStatus.AUDITING)
	var certifiedBusinessInfo *entity.CertifiedBusinessInfo = new(entity.CertifiedBusinessInfo)
	certifiedBusinessInfo.SetCertifiedBusinessStatus(member.GetCertifiedBusinessStatus())
	certifiedBusinessInfo.SetEmail(member.GetEmail())
	certifiedBusinessInfo.SetMemberLevel(member.GetMemberLevel())
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(certifiedBusinessInfo)
	return result
}
func (this *ApproveController) ListBusinessAuthDepositList() (result *util.MessageResult) {
	var depositList *arraylist.List[entity.BusinessAuthDeposit] = this.BusinessAuthDepositService.FindAllByStatus(CommonStatus.NORMAL)
	depositList.ForEach(func(deposit interface {
	}) {
		deposit.SetAdmin(nil)
	})
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(depositList)
	return result
}
func (this *ApproveController) ChangePhone(request *http.HttpServletRequest, password string, phone string, code string, user *transform.AuthMember) (result *util.MessageResult, err error) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	hasText(password, this.MsService.GetMessage("MISSING_LOGIN_PASSWORD"))
	hasText(phone, this.MsService.GetMessage("MISSING_PHONE"))
	hasText(code, this.MsService.GetMessage("MISSING_VERIFICATION_CODE"))
	var member1 *entity.Member = this.MemberService.FindByPhone(phone)
	isTrue(member1 == nil, this.MsService.GetMessage("PHONE_ALREADY_BOUND"))
	var cache interface {
	} = this.RedisUtil.Get(SysConstant.PHONE_CHANGE_CODE_PREFIX + member.GetMobilePhone())
	notNull(cache, this.MsService.GetMessage("NO_GET_VERIFICATION_CODE"))
	if member.GetCountry().GetAreaCode().Equals("86") {
		if !ValidateUtil.IsMobilePhone(phone.Trim()) {
			return MessageResult.Error(this.MsService.GetMessage("PHONE_FORMAT_ERROR"))
		}
	}
	if member.GetPassword().Equals(Md5.Md5Digest(password + member.GetSalt()).ToLowerCase()) {
		if !code.Equals(cache.ToString()) {
			return MessageResult.Error(this.MsService.GetMessage("VERIFICATION_CODE_INCORRECT"))
		} else {
			core.DelKey(SysConstant.PHONE_CHANGE_CODE_PREFIX + member.GetMobilePhone())
		}
		member.SetMobilePhone(phone)
		return MessageResult.Success(this.MsService.GetMessage("SETTING_SUCCESS"))
	} else {
		request.RemoveAttribute(SysConstant.SESSION_MEMBER)
		return MessageResult.Error(this.MsService.GetMessage("PASSWORD_ERROR"))
	}
}
func (this *ApproveController) CancelBusiness(user *transform.AuthMember, detail string) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	Logger.Info("申请退保，原因={}", detail)
	var advertiseNum int = this.AdvertiseService.CountByMemberAndStatus(member, AdvertiseControlStatus.PUT_ON_SHELVES)
	log.Errorf("advertiseNum=%v", advertiseNum)
	if advertiseNum > 0 {
		return MessageResult.Error("请先下架或者删除您所挂出的广告")
	}
	var orderNum int64 = this.OrderService.CountByMemberProcessing(member.GetId())
	log.Errorf("orderNum=%v", orderNum)
	if orderNum > 0 {
		return MessageResult.Error("请先处理未完成订单")
	}
	if member.GetCertifiedBusinessStatus() == CANCEL_AUTH {
		return MessageResult.Error("退保审核中，请勿重复提交......")
	}
	if !member.GetCertifiedBusinessStatus().Equals(CertifiedBusinessStatus.VERIFIED) {
		return MessageResult.Error("you are not verified business")
	}
	var businessAuthApplyList *arraylist.List[entity.BusinessAuthApply] = this.BusinessAuthApplyService.FindByMemberAndCertifiedBusinessStatus(member, CertifiedBusinessStatus.VERIFIED)
	if businessAuthApplyList == nil || businessAuthApplyList.Size() < 1 {
		return MessageResult.Error("you are not verified business,business auth apply not exist......")
	}
	if businessAuthApplyList.Get(0).GetCertifiedBusinessStatus() != CertifiedBusinessStatus.VERIFIED {
		return MessageResult.Error("data exception, state inconsistency(CertifiedBusinessStatus in BusinessAuthApply and Member)")
	}
	member.SetCertifiedBusinessStatus(CANCEL_AUTH)
	log.Infof("会员状态:%v", member.GetCertifiedBusinessStatus())
	this.MemberService.Save(member)
	log.Infof("会员状态:%v", member.GetCertifiedBusinessStatus())
	var cancelApply *entity.BusinessCancelApply = new(entity.BusinessCancelApply)
	cancelApply.SetDepositRecordId(businessAuthApplyList.Get(0).GetDepositRecordId())
	cancelApply.SetMember(businessAuthApplyList.Get(0).GetMember())
	cancelApply.SetStatus(CANCEL_AUTH)
	cancelApply.SetReason(detail)
	log.Infof("退保申请状态:%v", cancelApply.GetStatus())
	this.BusinessCancelApplyService.Save(cancelApply)
	log.Infof("退保申请状态:%v", cancelApply.GetStatus())
	return MessageResult.Success()
}
func (this *ApproveController) CheckCode(key string, member *entity.Member, code string, codeMold int) (result *util.MessageResult) {
	if this.CodeType == 0 {
		var fullKey = ""
		if codeMold == 1 {
			fullKey = key + member.GetEmail()
		} else {
			fullKey = key + member.GetMobilePhone()
		}
		var cache interface {
		} = this.RedisUtil.Get(fullKey)
		notNull(cache, this.MsService.GetMessage("NO_GET_VERIFICATION_CODE"))
		if !code.Equals(cache.ToString()) {
			return MessageResult.Error(this.MsService.GetMessage("VERIFICATION_CODE_INCORRECT"))
		} else {
			core.DelKey(fullKey)
		}
	} else if this.CodeType == 1 {
		if member.GetGoogleState() == 1 {
			var r bool = GoogleAuthenticatorUtil.CheckCodes(code, member.GetGoogleKey())
			if !r {
				return MessageResult.Error(this.MsService.GetMessage("GOOGLE_AUTH_FAILD"))
			}
		} else {
			return MessageResult.Error(this.MsService.GetMessage("BIND_GOOGLE_FIRST"))
		}
	}
	return MessageResult.Success()
}
func (this *ApproveController) VideoRandom(user *transform.AuthMember) (result *util.MessageResult) {
	var jsonResult *fastjson.JSONObject = new(fastjson.JSONObject)
	var memberId = user.GetId() + ""
	var random int = GeneratorUtil.GetRandomNumber(100000, 999999)
	jsonResult.Put("memberId", memberId)
	jsonResult.Put("random", random)
	Logger.Info("=====获取视频随机码====" + jsonResult.ToJSONString())
	var result *util.MessageResult = MessageResult.Success()
	result.SetData(jsonResult)
	return result
}

type ApproveController struct {
	MemberService              *service.MemberService
	RedisUtil                  *util.RedisUtil
	MsService                  *service.LocaleMessageSourceService
	MemberApplicationService   *service.MemberApplicationService
	BusinessAuthDepositService *service.BusinessAuthDepositService
	BusinessCancelApplyService *service.BusinessCancelApplyService
	BusinessAuthApplyService   *service.BusinessAuthApplyService
	DepositRecordService       *service.DepositRecordService
	OrderService               *service.OrderService
	AdvertiseService           *service.AdvertiseService
	OtcWalletService           *service.OtcWalletService
	RealTimes                  int
	CodeType                   int
	BaseController
}
