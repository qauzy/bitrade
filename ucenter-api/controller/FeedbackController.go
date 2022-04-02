package controller

import (
	"bitrade/core/entity"
	"bitrade/core/entity/transform"
	"bitrade/core/service"
	"bitrade/core/util"
)

func (this *FeedbackController) Feedback(user *transform.AuthMember, remark string) (result *util.MessageResult) {
	var Feedback *entity.Feedback = new(entity.Feedback)
	var member *entity.Member = this.MemberService.FindOne(user.GetId())
	this.Feedback.SetMember(member)
	this.Feedback.SetRemark(remark)
	var feedback1 *entity.Feedback = this.FeedbackService.Save(this.Feedback)
	if feedback1 != nil {
		return MessageResult.Success()
	} else {
		return MessageResult.Error(this.MsService.GetMessage("SYSTEM_ERROR"))
	}
}

type FeedbackController struct {
	FeedbackService *service.FeedbackService
	MemberService   *service.MemberService
	MsService       *service.LocaleMessageSourceService
}
