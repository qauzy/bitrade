package service

func (this *GiftRecordService) Save(giftRecord *entity.GiftRecord) (result *entity.GiftRecord, err error) {
	return this.GiftRecordDao.Save(giftRecord)
}
func (this *GiftRecordService) FindById(id int64) (result *entity.GiftRecord, err error) {
	return this.GiftRecordDao.FindById(id).Get()
}
func (this *GiftRecordService) GetByPage(giftRecordVO *vo.GiftRecordVO) (result domain.Page[entity.GiftRecord], err error) {
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if StringUtils.IsNotEmpty(giftRecordVO.GetStartTime()) {
		booleanExpressions.Add(QGiftRecord.GiftRecord.CreateTime.Goe(DateUtil.StringToDate(giftRecordVO.GetStartTime())))
	}
	if StringUtils.IsNotEmpty(giftRecordVO.GetEndTime()) {
		booleanExpressions.Add(QGiftRecord.GiftRecord.CreateTime.Loe(DateUtil.StringToDate(giftRecordVO.GetEndTime())))
	}
	if StringUtils.IsNotEmpty(giftRecordVO.GetUserName()) {
		booleanExpressions.Add(QGiftRecord.GiftRecord.UserName.Loe("%" + giftRecordVO.GetUserName() + "%"))
	}
	if StringUtils.IsNotEmpty(giftRecordVO.GetMobile()) {
		booleanExpressions.Add(QGiftRecord.GiftRecord.UserMobile.Loe(giftRecordVO.GetMobile()))
	}
	if giftRecordVO.GetUserId() != nil {
		booleanExpressions.Add(QGiftRecord.GiftRecord.UserId.Eq(giftRecordVO.GetUserId()))
	}
	var predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var sort = Sort.By(Sort.Direction.DESC, "id")
	var pageable = PageRequest.Of(giftRecordVO.GetPageNum()-1, giftRecordVO.GetPageSize(), sort)
	return this.GiftRecordDao.FindAll(predicate, pageable)
}
func NewGiftRecordService(giftRecordDao *dao.GiftRecordDao) (ret *GiftRecordService) {
	ret = new(GiftRecordService)
	ret.GiftRecordDao = giftRecordDao
	return
}

type GiftRecordService struct {
	GiftRecordDao dao.GiftRecordDao
	Base.BaseService
}
