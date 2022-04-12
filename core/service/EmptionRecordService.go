package service

func (this *EmptionRecordService) Save(emptionRecord *entity.EmptionRecord) (result *entity.EmptionRecord, err error) {
	return this.EmptionRecordDao.Save(emptionRecord)
}
func (this *EmptionRecordService) GetByPage(emptionRecrodVO *vo.EmptionRecrodVO) (result domain.Page[entity.EmptionRecord], err error) {
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if StringUtils.IsNotEmpty(emptionRecrodVO.GetStartTime()) {
		booleanExpressions.Add(QEmptionRecord.EmptionRecord.CreateTime.Goe(DateUtil.StringToDate(emptionRecrodVO.GetStartTime())))
	}
	if StringUtils.IsNotEmpty(emptionRecrodVO.GetEndTime()) {
		booleanExpressions.Add(QEmptionRecord.EmptionRecord.CreateTime.Loe(DateUtil.StringToDate(emptionRecrodVO.GetEndTime())))
	}
	if StringUtils.IsNotEmpty(emptionRecrodVO.GetIeoName()) {
		booleanExpressions.Add(QEmptionRecord.EmptionRecord.IeoName.Like("%" + emptionRecrodVO.GetIeoName() + "%"))
	}
	if StringUtils.IsNotEmpty(emptionRecrodVO.GetUserName()) {
		booleanExpressions.Add(QEmptionRecord.EmptionRecord.UserName.Like("%" + emptionRecrodVO.GetUserName() + "%"))
	}
	if StringUtils.IsNotEmpty(emptionRecrodVO.GetUserMobile()) {
		booleanExpressions.Add(QEmptionRecord.EmptionRecord.UserMobile.Eq(emptionRecrodVO.GetUserMobile()))
	}
	if StringUtils.IsNotEmpty(emptionRecrodVO.GetStatus()) {
		booleanExpressions.Add(QEmptionRecord.EmptionRecord.Status.Eq(emptionRecrodVO.GetStatus()))
	}
	if emptionRecrodVO.GetUserId() != nil {
		booleanExpressions.Add(QEmptionRecord.EmptionRecord.UserId.Eq(emptionRecrodVO.GetUserId()))
	}
	var predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var sort = Sort.By(Sort.Direction.DESC, "id")
	var pageable = PageRequest.Of(emptionRecrodVO.GetPageNum()-1, emptionRecrodVO.GetPageSize(), sort)
	return this.EmptionRecordDao.FindAll(predicate, pageable)
}
func (this *EmptionRecordService) FindById(id int64) (result *entity.EmptionRecord, err error) {
	return this.EmptionRecordDao.FindById(id).Get()
}
func NewEmptionRecordService(emptionRecordDao *dao.EmptionRecordDao) (ret *EmptionRecordService) {
	ret = new(EmptionRecordService)
	ret.EmptionRecordDao = emptionRecordDao
	return
}

type EmptionRecordService struct {
	EmptionRecordDao dao.EmptionRecordDao
	Base.BaseService
}
