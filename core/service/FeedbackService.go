package service

func (this *FeedbackService) Save(feedback *entity.Feedback) (result *entity.Feedback, err error) {
	return this.FeedbackDao.Save(feedback)
}
func NewFeedbackService(feedbackDao *dao.FeedbackDao) (ret *FeedbackService) {
	ret = new(FeedbackService)
	ret.FeedbackDao = feedbackDao
	return
}

type FeedbackService struct {
	FeedbackDao dao.FeedbackDao
	Base.BaseService
}
