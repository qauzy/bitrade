package controller

import (
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util"
	"bitrade/core/vo"
)

func (this *GiftController) GetGiftRecord(user *transform.AuthMember, giftRecordVO *vo.GiftRecordVO) (result *util.MessageResult, err error) {
	log.Info("-------查询个人糖果记录----" + JSONObject.ToJSONString(giftRecordVO))
	giftRecordVO.SetUserId(user.GetId())
	var result *domain.Page = this.GiftRecordService.GetByPage(giftRecordVO)
	return successDataAndTotal(result.GetContent(), result.GetTotalElements())
}

type GiftController struct {
	GiftRecordService *service.GiftRecordService
	BaseController
}
