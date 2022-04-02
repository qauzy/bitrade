package controller

import (
	SysConstant "bitrade/core/constant/SysConstants"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/fastjson"
	"time"
)

func (this *GoogleAuthenticationController) Yzgoogle(user *transform.AuthMember, codes string) (result *util.MessageResult) {
	// enter the code shown on device. Edit this and run it fast before the
	// code expires!
	var code int64 = Long.ParseLong(codes)
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var t int64 = System.CurrentTimeMillis()
	var ga *util.GoogleAuthenticatorUtil = new(util.GoogleAuthenticatorUtil)
	//  ga.setWindowSize(0); // should give 5 * 30 seconds of grace...
	var r bool = ga.Check_code(member.GetGoogleKey(), code, t)
	log.Info("rrrr=" + r)
	if !r {
		return MessageResult.Error("验证失败")
	} else {
		return MessageResult.Success("验证通过")
	}
}
func (this *GoogleAuthenticationController) Sendgoogle(user *transform.AuthMember) (result *util.MessageResult) {
	/*for(int i = 0;i<50;i++){
	    log.info("######################       开始循环次数={}    ######################",i+1);
	    GoogleAuthenticatorUtil.generateSecretKey();
	    log.info("######################       结束循环次数={}    ######################",i+1);
	}*/
	log.Infof("开始进入用户id=%v", user.GetId())
	var current int64 = System.CurrentTimeMillis()
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	log.Infof("查询完毕 耗时=%v", System.CurrentTimeMillis()-current)
	if member == nil {
		return MessageResult.Error("未登录")
	}
	var secret string = GoogleAuthenticatorUtil.GenerateSecretKey()
	log.Infof("secret完毕 耗时=%v", System.CurrentTimeMillis()-current)
	var qrBarcodeURL string = GoogleAuthenticatorUtil.GetQRBarcodeURL(member.GetMobilePhone(), this.GoogleAuthUrl, secret)
	log.Infof("qrBarcodeURL完毕 耗时=%v", System.CurrentTimeMillis()-current)
	var jsonObject *fastjson.JSONObject = new(fastjson.JSONObject)
	jsonObject.Put("link", qrBarcodeURL)
	jsonObject.Put("secret", secret)
	log.Infof("jsonObject完毕 耗时=%v", System.CurrentTimeMillis()-current)
	var messageResult *util.MessageResult = new(util.MessageResult)
	messageResult.SetData(jsonObject)
	messageResult.SetMessage("获取成功")
	log.Infof("执行完毕 耗时=%v", System.CurrentTimeMillis()-current)
	return messageResult
}
func (this *GoogleAuthenticationController) Jcgoogle(codes string, user *transform.AuthMember, smsCode string) (result *util.MessageResult) {
	// enter the code shown on device. Edit this and run it fast before the
	// code expires!
	//String GoogleKey = (String) request.getSession().getAttribute("googleKey");
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var GoogleKey string = member.GetGoogleKey()
	//        if(StringUtils.isEmpty(password)){
	//            return MessageResult.error("密码不能为空");
	//        }
	//        try {
	//            if(!(Md5.md5Digest(password + member.getSalt()).toLowerCase().equals(member.getPassword().toLowerCase()))){
	//                return MessageResult.error("密码错误");
	//            }
	//        } catch (Exception e) {
	//            e.printStackTrace();
	//        }
	if !smsCode.Equals(this.RedisUtil.Get(SysConstant.RESET_GOOGLE_CODE_PREFIX + member.GetMobilePhone())) {
		return error(this.SourceService.GetMessage("VERIFICATION_CODE_INCORRECT"))
	} else {
		core.DelKey(SysConstant.RESET_GOOGLE_CODE_PREFIX + member.GetMobilePhone())
	}
	var code int64 = Long.ParseLong(codes)
	var t int64 = System.CurrentTimeMillis()
	var ga *util.GoogleAuthenticatorUtil = new(util.GoogleAuthenticatorUtil)
	// ga.setWindowSize(0); // should give 5 * 30 seconds of grace...
	var r bool = ga.Check_code(GoogleKey, code, t)
	if !r {
		return MessageResult.Error("验证失败")
	} else {
		member.SetGoogleDate(time.Now())
		member.SetGoogleState(0)
		var result *entity.Member = this.MemberService.Save(member)
		if result != nil {
			return MessageResult.Success("解绑成功")
		} else {
			return MessageResult.Error("解绑失败")
		}
	}
}
func (this *GoogleAuthenticationController) GoogleAuth(codes string, user *transform.AuthMember, secret string) (result *util.MessageResult) {
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	var code int64 = Long.ParseLong(codes)
	var t int64 = System.CurrentTimeMillis()
	var ga *util.GoogleAuthenticatorUtil = new(util.GoogleAuthenticatorUtil)
	var r bool = ga.Check_code(secret, code, t)
	if !r {
		return MessageResult.Error("验证失败")
	} else {
		member.SetGoogleState(1)
		member.SetGoogleKey(secret)
		member.SetGoogleDate(time.Now())
		var result *entity.Member = this.MemberService.Save(member)
		if result != nil {
			return MessageResult.Success("绑定成功")
		} else {
			return MessageResult.Error("绑定失败")
		}
	}
}

type GoogleAuthenticationController struct {
	MemberService *service.MemberService
	GoogleAuthUrl string
	RedisUtil     *util.RedisUtil
	SourceService *service.LocaleMessageSourceService
	BaseController
}
