package service

func (this *StatisticsService) GetStatistics(startTime string, endTime string, sql string) (result arraylist.List[interface {
}], err error) {
	var query = this.Em.CreateNativeQuery(sql)
	query.SetParameter("startTime", startTime)
	query.SetParameter("endTime", endTime)
	var resultList = query.GetResultList()
	return resultList
}
func (this *StatisticsService) GetLatelyRegMember(day int) (result int, err error) {
	var startTime = DateUtil.StrToDate(DateUtil.GetPastDate(day) + " 00:00:00")
	var endTime = DateUtil.StrToDate(DateUtil.GetDate() + " 23:59:59")
	return this.MemberDao.CountByRegistrationTimeBetween(startTime, endTime)
}
func (this *StatisticsService) GetLatelyOrder(startTime string, endTime string, status int) (result int, err error) {
	var startTimeDate = DateUtil.StrToDate(startTime + " 00:00:00")
	var endTimeDate = DateUtil.StrToDate(endTime + " 23:59:59")
	if status < 0 {
		return this.OrderDao.CountByCreateTimeBetween(startTimeDate, endTimeDate)
	}
	var orderStatus = EnumHelperUtil.IndexOf(OrderStatus.OrderStatus, status)
	return this.OrderDao.CountByStatusAndCreateTimeBetween(orderStatus, startTimeDate, endTimeDate)
}
func (this *StatisticsService) GetLatelyOrder(status *OrderStatus.OrderStatus) (result int, err error) {
	return this.OrderDao.CountByStatus(status)
}
func (this *StatisticsService) GetLatelyAdvertise(i int) (result int, err error) {
	return 0
}
func NewStatisticsService(em *persistence.EntityManager, memberDao *dao.MemberDao, orderDao *dao.OrderDao) (ret *StatisticsService) {
	ret = new(StatisticsService)
	ret.Em = em
	ret.MemberDao = memberDao
	ret.OrderDao = orderDao
	return
}

type StatisticsService struct {
	Em        *persistence.EntityManager
	MemberDao dao.MemberDao
	OrderDao  dao.OrderDao
	Base.BaseService
}
