package Test

import (
	"bitrade/core/controller"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/gin-gonic/gin"
)

func (this *TestController) TestSms(ctx *gin.Context) (result *util.MessageResult) {
	var smsDTO = this.SmsService.GetByStatus()
	System.Out.Println(smsDTO)
	System.Out.Println(success("succss", smsDTO))
	return this.SuccessWithData(smsDTO)
}
func NewTestController(smsService *service.SmsService, memberWalletService *service.MemberWalletService) (ret *TestController) {
	ret = new(TestController)
	ret.SmsService = smsService
	ret.MemberWalletService = memberWalletService
	return
}

type TestController struct {
	SmsService          *service.SmsService
	MemberWalletService *service.MemberWalletService
	controller.BaseController
}
