package service

func (this *GiftConfigService) Save(giftRecord *entity.GiftConfig) (result *entity.GiftConfig, err error) {
	return this.GiftConfigDao.Save(giftRecord)
}
func (this *GiftConfigService) FindById(id int64) (result *entity.GiftConfig, err error) {
	return this.GiftConfigDao.FindById(id).Get()
}
func (this *GiftConfigService) GetByPage(giftConfigVO *vo.GiftConfigVO) (result domain.Page[entity.GiftConfig], err error) {
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if StringUtils.IsNotEmpty(giftConfigVO.GetStartTime()) {
		booleanExpressions.Add(QGiftConfig.GiftConfig.CreateTime.Goe(DateUtil.StringToDate(giftConfigVO.GetStartTime())))
	}
	if StringUtils.IsNotEmpty(giftConfigVO.GetEndTime()) {
		booleanExpressions.Add(QGiftConfig.GiftConfig.CreateTime.Loe(DateUtil.StringToDate(giftConfigVO.GetEndTime())))
	}
	if StringUtils.IsNotEmpty(giftConfigVO.GetGiftName()) {
		booleanExpressions.Add(QGiftConfig.GiftConfig.GiftName.Loe("%" + giftConfigVO.GetGiftName() + "%"))
	}
	if giftConfigVO.GetId() != nil {
		booleanExpressions.Add(QGiftConfig.GiftConfig.Id.Eq(giftConfigVO.GetId()))
	}
	var predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var sort = Sort.By(Sort.Direction.DESC, "id")
	var pageable = PageRequest.Of(giftConfigVO.GetPageNum()-1, giftConfigVO.GetPageSize(), sort)
	return this.GiftConfigDao.FindAll(predicate, pageable)
}
func NewGiftConfigService(giftConfigDao *dao.GiftConfigDao) (ret *GiftConfigService) {
	ret = new(GiftConfigService)
	ret.GiftConfigDao = giftConfigDao
	return
}

type GiftConfigService struct {
	GiftConfigDao dao.GiftConfigDao
	Base.BaseService
}
