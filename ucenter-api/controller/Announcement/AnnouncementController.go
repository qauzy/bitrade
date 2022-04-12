package Announcement

import (
	"bitrade/core/controller"
	"bitrade/core/dao/types"
	"bitrade/core/service"
	"bitrade/core/util/MessageResult"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *AnnouncementController) Page(ctx *gin.Context, pageNo int, pageSize int) (result *MessageResult.MessageResult) {
	//条件
	var predicates = arraylist.New[types.Predicate]()
	predicates.Add(QAnnouncement.Announcement.IsShow.Eq(true))
	//排序
	var orderSpecifiers = arraylist.New[types.OrderSpecifier]()
	orderSpecifiers.Add(QAnnouncement.Announcement.IsTop.Asc())
	orderSpecifiers.Add(QAnnouncement.Announcement.Sort.Desc())
	//查
	var pageResult = this.AnnouncementService.QueryDsl(pageNo, pageSize, predicates, QAnnouncement.Announcement, orderSpecifiers)
	return this.SuccessWithData(pageResult)
}
func (this *AnnouncementController) Detail(ctx *gin.Context, id int64) (result *MessageResult.MessageResult) {
	var announcement, err = this.AnnouncementService.FindById(id)
	if err != nil {
		return MessageResult.Error(err.Error())
	}
	if announcement == nil {
		return MessageResult.Error("validate id!")
	}
	return this.SuccessWithData(announcement)
}
func NewAnnouncementController(announcementService *service.AnnouncementService) (ret *AnnouncementController) {
	ret = new(AnnouncementController)
	ret.AnnouncementService = announcementService
	return
}

type AnnouncementController struct {
	AnnouncementService *service.AnnouncementService
	controller.BaseController
}
