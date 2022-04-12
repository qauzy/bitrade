package Gift

import (
	"bitrade/core/controller"
	"bitrade/core/entity/transform"
	"bitrade/core/log"
	"bitrade/core/service"
	"bitrade/core/util/MessageResult"
	"bitrade/core/vo"
	"github.com/gin-gonic/gin"
)

func (this *GiftController) GetGiftRecord(ctx *gin.Context, user *transform.AuthMember, giftRecordVO *vo.GiftRecordVO) (result *MessageResult.MessageResult, err error) {
	log.Info("-------查询个人糖果记录----" + JSONObject.ToJSONString(giftRecordVO))
	giftRecordVO.SetUserId(user.GetId())
	var result = this.GiftRecordService.GetByPage(giftRecordVO)
	return successDataAndTotal(result.GetContent(), result.GetTotalElements())
}
func NewGiftController(giftRecordService *service.GiftRecordService) (ret *GiftController) {
	ret = new(GiftController)
	ret.GiftRecordService = giftRecordService
	return
}

type GiftController struct {
	GiftRecordService *service.GiftRecordService
	controller.BaseController
}
