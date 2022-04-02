package controller

import (
	"bitrade/core/dto"
	"bitrade/core/service"
	"bitrade/core/util"
)

func (this *TestController) TestSms() (result *util.MessageResult) {
	var smsDTO *dto.SmsDTO = this.SmsService.GetByStatus()
	System.Out.Println(smsDTO)
	System.Out.Println(success("succss", smsDTO))
	return success(smsDTO)
}

type TestController struct {
	SmsService          *service.SmsService
	MemberWalletService *service.MemberWalletService
	BaseController
}
