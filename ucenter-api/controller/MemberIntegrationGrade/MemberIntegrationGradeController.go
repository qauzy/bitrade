package MemberIntegrationGrade

import (
	"bitrade/core/controller"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/vo"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/fastjson"
)

func (this *MemberIntegrationGradeController) QueryGradeInfo(ctx *gin.Context) (result *util.MessageResult) {
	return success(this.GradeService.FindAll())
}
func (this *MemberIntegrationGradeController) QueryDayWithdrawLimit(ctx *gin.Context, user *transform.AuthMember) (result *util.MessageResult) {
	var count = this.RedisUtil.Get(SysConstant.CUSTOMER_DAY_WITHDRAW_TOTAL_COUNT + user.GetId())
	if count == nil {
		count = 0
	} else {
		count = count
	}
	var amount = this.RedisUtil.Get(SysConstant.CUSTOMER_DAY_WITHDRAW_COVER_USD_AMOUNT + user.GetId())
	if amount == nil {
		amount = 0
	} else {
		amount = amount
	}
	var result = new(fastjson.JSONObject)
	result.Put("count", count)
	result.Put("amount", amount)
	return this.SuccessWithData(result)
}
func (this *MemberIntegrationGradeController) QueryIntegration4PageQuery(ctx *gin.Context, queryVo *vo.IntegrationRecordVO, user *transform.AuthMember) (result *util.MessageResult) {
	var mr = MessageResult(0, "success")
	queryVo.SetUserId(user.GetId())
	var page = this.RecordService.FindRecord4Page(queryVo)
	mr.SetTotal(page.GetTotalElements())
	mr.SetData(page.GetContent())
	return mr
}
func NewMemberIntegrationGradeController(gradeService *service.MemberGradeService, redisUtil *util.RedisUtil, recordService *service.IntegrationRecordService) (ret *MemberIntegrationGradeController) {
	ret = new(MemberIntegrationGradeController)
	ret.GradeService = gradeService
	ret.RedisUtil = redisUtil
	ret.RecordService = recordService
	return
}

type MemberIntegrationGradeController struct {
	GradeService  *service.MemberGradeService
	RedisUtil     *util.RedisUtil
	RecordService *service.IntegrationRecordService
	controller.BaseController
}
