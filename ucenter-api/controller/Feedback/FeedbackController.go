package Feedback

import (
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/gin-gonic/gin"
)

func (this *FeedbackController) Feedback(ctx *gin.Context, user *transform.AuthMember, remark string) (result *util.MessageResult) {
	var Feedback = new(entity.Feedback)
	var member = this.MemberService.FindOne(user.GetId())
	this.Feedback.SetMember(member)
	this.Feedback.SetRemark(remark)
	var feedback1 = this.FeedbackService.Save(this.Feedback)
	if feedback1 != nil {
		return MessageResult.Success()
	} else {
		return MessageResult.Error(this.MsService.GetMessage("SYSTEM_ERROR"))
	}
}
func NewFeedbackController(feedbackService *service.FeedbackService, memberService *service.MemberService, msService *service.LocaleMessageSourceService) (ret *FeedbackController) {
	ret = new(FeedbackController)
	ret.FeedbackService = feedbackService
	ret.MemberService = memberService
	ret.MsService = msService
	return
}

type FeedbackController struct {
	FeedbackService *service.FeedbackService
	MemberService   *service.MemberService
	MsService       *service.LocaleMessageSourceService
}
