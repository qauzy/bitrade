package service

func (this *IntegrationRecordService) FindRecord4Page(queryVo *vo.IntegrationRecordVO) (result domain.Page[entity.IntegrationRecord], err error) {
	var qIntegrationRecord = QIntegrationRecord.IntegrationRecord
	var booleanExpressions = arraylist.New[dsl.BooleanExpression]()
	if queryVo.GetUserId() != nil {
		booleanExpressions.Add(qIntegrationRecord.MemberId.Eq(queryVo.GetUserId()))
	}
	if queryVo.GetType() != nil {
		booleanExpressions.Add(qIntegrationRecord.Type.Eq(queryVo.GetType()))
	}
	if StringUtils.IsNotEmpty(queryVo.GetCreateStartTime()) {
		booleanExpressions.Add(qIntegrationRecord.CreateTime.Gt(DateUtil.StrToDate(queryVo.GetCreateStartTime())))
	}
	if StringUtils.IsNotEmpty(queryVo.GetCreateEndTime()) {
		booleanExpressions.Add(qIntegrationRecord.CreateTime.Lt(DateUtil.StrToDate(queryVo.GetCreateEndTime())))
	}
	var predicate = PredicateUtils.GetPredicate(booleanExpressions)
	var sort = Sort.By(Sort.Direction.DESC, "id")
	var pageable = PageRequest.Of(queryVo.GetPageNum()-1, queryVo.GetPageSize(), sort)
	if predicate == nil {
		predicate = new(core.BooleanBuilder)
	}
	return this.Dao.FindAll(predicate, pageable)
}
func (this *IntegrationRecordService) Save(record *entity.IntegrationRecord) (result *entity.IntegrationRecord, err error) {
	return this.Dao.Save(record)
}
func NewIntegrationRecordService(dao *dao.IntegrationRecordDao) (ret *IntegrationRecordService) {
	ret = new(IntegrationRecordService)
	ret.Dao = dao
	return
}

type IntegrationRecordService struct {
	Dao dao.IntegrationRecordDao
	Base.BaseService[entity.IntegrationRecord]
}
