package TripartiteRecharge

import (
	"bitrade/core/service"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/fastjson"
	"github.com/shopspring/decimal"
)

func (this *TripartiteRechargeController) Recharge(ctx *gin.Context, body string) (result string, err error) {
	var json = new(fastjson.JSONObject)
	var jsonObject = new(fastjson.JSONObject)
	exception := func() (err error) {
		//解密
		var resJson = JSONObject.ParseObject(DESUtil.Decrypt(body.Trim(), this.Key))
		if resJson != nil {
			var data = resJson.GetJSONObject("data")
			//验证签名
			if Md5.Md5Digest(data.ToJSONString() + this.Smac).Equals(resJson.GetString("sign")) {
				var phone = data.GetString("user")
				var amount = data.GetBigDecimal("amount")
				var unit = data.GetString("unit")
				if !(StringUtils.HasText(phone) && this.MemberService.FindByPhone(phone) != nil) {
					json.Put("code", 1003)
					json.Put("message", "无此用户")
				} else if amount.CompareTo(decimal.Zero) <= 0 {
					json.Put("code", 1004)
					json.Put("message", "充值金额不合法")
				} else if !(StringUtils.HasText(unit) && this.MemberWalletService.FindByCoinUnitAndMemberId(unit, this.MemberService.FindByPhone(phone).GetId()) != nil) {
					json.Put("code", 1005)
					json.Put("message", "不支持该币种")
				} else {
					var memberWallet = this.MemberWalletService.FindByCoinUnitAndMemberId(unit, this.MemberService.FindByPhone(phone).GetId())
					var messageResult = this.MemberWalletService.Recharge(memberWallet, amount)
					if messageResult.GetCode() == 0 {
						json.Put("code", 0)
						json.Put("message", "成功")
						json.Put("user", phone)
						json.Put("unit", unit)
						json.Put("amount", amount)
					} else {
						json.Put("code", 1006)
						json.Put("message", "失败")
					}
				}
			} else {
				json.Put("code", 1002)
				json.Put("message", "签名错误，数据可能被篡改")
			}
		}
		return
	}()
	if exception != nil {
		json.Put("code", 1001)
		json.Put("message", "加密错误")
		e.PrintStackTrace()
	}
	jsonObject.Put("data", json)
	//制作签名
	jsonObject.Put("sign", Md5.Md5Digest(json.ToJSONString()+this.Smac))
	//加密
	var ciphertext = DESUtil.ENCRYPTMethod(jsonObject.ToJSONString(), this.Key).ToUpperCase()
	return ciphertext
}
func NewTripartiteRechargeController(sourceService *service.LocaleMessageSourceService, url string, key string, smac string, memberService *service.MemberService, memberWalletService *service.MemberWalletService) (ret *TripartiteRechargeController) {
	ret = new(TripartiteRechargeController)
	ret.SourceService = sourceService
	ret.Url = url
	ret.Key = key
	ret.Smac = smac
	ret.MemberService = memberService
	ret.MemberWalletService = memberWalletService
	return
}

type TripartiteRechargeController struct {
	SourceService       *service.LocaleMessageSourceService
	Url                 string
	Key                 string
	Smac                string
	MemberService       *service.MemberService
	MemberWalletService *service.MemberWalletService
}
