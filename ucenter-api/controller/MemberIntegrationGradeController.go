package controller

import (
	SysConstant "bitrade/core/constant/SysConstants"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/vo"
	"github.com/qauzy/fastjson"
)

func (this *MemberIntegrationGradeController) QueryGradeInfo() (result *util.MessageResult) {
	return success(this.GradeService.FindAll())
}
func (this *MemberIntegrationGradeController) QueryDayWithdrawLimit(user *transform.AuthMember) (result *util.MessageResult) {
	var count interface {
	} = this.RedisUtil.Get(SysConstant.CUSTOMER_DAY_WITHDRAW_TOTAL_COUNT + user.GetId())
	if count == nil {
		count = 0
	} else {
		count = count
	}
	var amount interface {
	} = this.RedisUtil.Get(SysConstant.CUSTOMER_DAY_WITHDRAW_COVER_USD_AMOUNT + user.GetId())
	if amount == nil {
		amount = 0
	} else {
		amount = amount
	}
	var result *fastjson.JSONObject = new(fastjson.JSONObject)
	result.Put("count", count)
	result.Put("amount", amount)
	return success(result)
}
func (this *MemberIntegrationGradeController) QueryIntegration4PageQuery(queryVo *vo.IntegrationRecordVO, user *transform.AuthMember) (result *util.MessageResult) {
	var mr *util.MessageResult = MessageResult(0, "success")
	queryVo.SetUserId(user.GetId())
	var page *domain.Page = this.RecordService.FindRecord4Page(queryVo)
	mr.SetTotal(page.GetTotalElements())
	mr.SetData(page.GetContent())
	return mr
}

type MemberIntegrationGradeController struct {
	GradeService  *service.MemberGradeService
	RedisUtil     *util.RedisUtil
	RecordService *service.IntegrationRecordService
	BaseController
}
