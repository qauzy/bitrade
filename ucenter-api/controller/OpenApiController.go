package controller

import (
	SysConstant "bitrade/core/constant/SysConstants"
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/util/lists/arraylist"
	"time"
)

func (this *OpenApiController) QueryApiKey(member *transform.AuthMember) (result *util.MessageResult) {
	var result *arraylist.List[entity.MemberApiKey] = this.MemberApiKeyService.FindAllByMemberId(member.GetId())
	return success(result)
}
func (this *OpenApiController) SaveApiKey(member *transform.AuthMember, memberApiKey *entity.MemberApiKey) (result *util.MessageResult) {
	log.Info("-------新增API-key:" + JSONObject.ToJSONString(memberApiKey))
	var code string = memberApiKey.GetCode()
	if StringUtils.IsNotEmpty(code) == false {
		"请输入验证码"
	}
	var cacheCode interface {
	} = this.RedisUtil.Get(SysConstant.API_BIND_CODE_PREFIX + member.GetMobilePhone())
	if cacheCode == nil {
		return MessageResult.Error("验证码已过期")
	}
	if !code.EqualsIgnoreCase(cacheCode.ToString()) {
		return MessageResult.Error("验证码不正确")
	}
	var all *arraylist.List[entity.MemberApiKey] = this.MemberApiKeyService.FindAllByMemberId(member.GetId())
	if all.IsEmpty() || all.Size() < 5 {
		memberApiKey.SetId(nil)
		if StringUtils.IsBlank(memberApiKey.GetBindIp()) {
			//不绑定IP时默认90天过期
			var calendar *util.Calendar = Calendar.GetInstance()
			calendar.Add(Calendar.DAY_OF_MONTH, 90)
			memberApiKey.SetExpireTime(calendar.GetTime())
		}
		memberApiKey.SetApiName(member.GetId() + "")
		memberApiKey.SetApiKey(GeneratorUtil.GetUUID())
		var secret string = GeneratorUtil.GetUUID()
		memberApiKey.SetSecretKey(secret)
		memberApiKey.SetMemberId(member.GetId())
		memberApiKey.SetCreateTime(time.Now())
		this.MemberApiKeyService.Save(memberApiKey)
		return success("新增成功", secret)
	} else {
		return error("数量超过最大限制")
	}
}
func (this *OpenApiController) UpdateApiKey(member *transform.AuthMember, memberApiKey *entity.MemberApiKey) (result *util.MessageResult) {
	log.Info("-------修改API-key:" + JSONObject.ToJSONString(memberApiKey))
	if memberApiKey.GetId() != nil {
		var findMemberApiKey *entity.MemberApiKey = this.MemberApiKeyService.FindByMemberIdAndId(member.GetId(), memberApiKey.GetId())
		if findMemberApiKey != nil {
			if !memberApiKey.GetRemark().Equals(findMemberApiKey.GetRemark()) {
				findMemberApiKey.SetRemark(memberApiKey.GetRemark())
			}
			if StringUtils.IsNotEmpty(memberApiKey.GetBindIp()) {
				findMemberApiKey.SetBindIp(memberApiKey.GetBindIp())
			} else {
				findMemberApiKey.SetBindIp(nil)
			}
			this.MemberApiKeyService.Save(findMemberApiKey)
			return success("修改成功")
		} else {
			return error("记录不存在")
		}
	} else {
		return error("记录不存在")
	}
}
func (this *OpenApiController) UpdateApiKey(member *transform.AuthMember, id int64) (result *util.MessageResult) {
	log.Infof("------删除api-key：memberId=%v,id=%v", member.GetId(), id)
	this.MemberApiKeyService.Del(id)
	return success("删除成功")
}

type OpenApiController struct {
	MemberApiKeyService *service.MemberApiKeyService
	RedisUtil           *util.RedisUtil
	BaseController
}
