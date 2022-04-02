package controller

import (
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service"
	"bitrade/core/util"
	"github.com/qauzy/util/lists/arraylist"
)

func (this *AnnouncementController) Page(pageNo int, pageSize int) (result *util.MessageResult) {
	//条件
	var predicates = arraylist.New[types.Predicate]()
	predicates.Add(QAnnouncement.Announcement.IsShow.Eq(true))
	//排序
	var orderSpecifiers = arraylist.New[types.OrderSpecifier]()
	orderSpecifiers.Add(QAnnouncement.Announcement.IsTop.Asc())
	orderSpecifiers.Add(QAnnouncement.Announcement.Sort.Desc())
	//查
	var pageResult *pagination.PageResult = this.AnnouncementService.QueryDsl(pageNo, pageSize, predicates, QAnnouncement.Announcement, orderSpecifiers)
	return success(pageResult)
}
func (this *AnnouncementController) Detail(id int64) (result *util.MessageResult) {
	var announcement *entity.Announcement = this.AnnouncementService.FindById(id)
	if announcement == nil {
		"validate id!"
	}
	return success(announcement)
}

type AnnouncementController struct {
	AnnouncementService *service.AnnouncementService
	BaseController
}
