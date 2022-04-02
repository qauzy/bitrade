package controller

import (
	"bitrade/core/entity"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/fastjson"
	"github.com/qauzy/math"
	"github.com/shopspring/decimal"
)

func (this *TripartiteRechargeController) Recharge(body string) (result string, err error) {
	var json *fastjson.JSONObject = new(fastjson.JSONObject)
	var jsonObject *fastjson.JSONObject = new(fastjson.JSONObject)
	exception := func() (err error) {
		//解密
		var resJson *fastjson.JSONObject = JSONObject.ParseObject(DESUtil.Decrypt(body.Trim(), this.Key))
		if resJson != nil {
			var data *fastjson.JSONObject = resJson.GetJSONObject("data")
			//验证签名
			if Md5.Md5Digest(data.ToJSONString() + this.Smac).Equals(resJson.GetString("sign")) {
				var phone string = data.GetString("user")
				var amount *math.BigDecimal = data.GetBigDecimal("amount")
				var unit string = data.GetString("unit")
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
					var memberWallet *entity.MemberWallet = this.MemberWalletService.FindByCoinUnitAndMemberId(unit, this.MemberService.FindByPhone(phone).GetId())
					var messageResult *util.MessageResult = this.MemberWalletService.Recharge(memberWallet, amount)
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
	var ciphertext string = DESUtil.ENCRYPTMethod(jsonObject.ToJSONString(), this.Key).ToUpperCase()
	return ciphertext
}

type TripartiteRechargeController struct {
	SourceService       *service.LocaleMessageSourceService
	Url                 string
	Key                 string
	Smac                string
	MemberService       *service.MemberService
	MemberWalletService *service.MemberWalletService
}
