package service

import (
	"bitrade/core/constant/PageModel"
	"bitrade/core/dao"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *AdminAccessLogService) SaveLog(adminAccessLog *entity.AdminAccessLog) (result *entity.AdminAccessLog, err error) {
	return this.AdminAccessLogDao.Save(adminAccessLog)
}
func (this *AdminAccessLogService) QueryAll() (result arraylist.List[*entity.AdminAccessLog], err error) {
	return this.AdminAccessLogDao.FindAll()
}
func (this *AdminAccessLogService) QueryById(id int64) (result *entity.AdminAccessLog, err error) {
	return this.AdminAccessLogDao.FindById(id)
}
func (this *AdminAccessLogService) Query(predicateList arraylist.List[types.Predicate], pageNo int, pageSize int) (result pagination.PageResult[entity.AdminAccessLog], err error) {
	var list arraylist.List[entity.AdminAccessLog]
	var jpaQuery = queryFactory.SelectFrom(QAdminAccessLog.AdminAccessLog)
	if predicateList != nil && predicateList.Size() > 0 {
		jpaQuery.Where(predicateList.ToArray(make([]Predicate, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
	//添加总条数
}
func (this *AdminAccessLogService) Page(predicates arraylist.List[dsl.BooleanExpression], pageModel *PageModel.PageModel) (result domain.Page[entity.AdminAccessLog], err error) {
	var Query = queryFactory.Select(Projections.Fields(entity.AdminAccessLog, QAdminAccessLog.AdminAccessLog.Id.As("id"), QAdminAccessLog.AdminAccessLog.AccessIp.As("accessIp"), QAdminAccessLog.AdminAccessLog.AccessMethod.As("accessMethod"), QAdminAccessLog.AdminAccessLog.AccessTime.As("accessTime"), QAdminAccessLog.AdminAccessLog.AdminId.As("adminId"), QAdminAccessLog.AdminAccessLog.Uri.As("uri"), QAdminAccessLog.AdminAccessLog.Module.As("module"), QAdminAccessLog.AdminAccessLog.Operation.As("operation"), QAdmin.Admin.Username.As("adminName"))).From(QAdminAccessLog.AdminAccessLog, QAdmin.Admin).Where(predicates.ToArray(make([]BooleanExpression, 0)))
	var orderSpecifiers = pageModel.GetOrderSpecifiers()
	this.Query.OrderBy(orderSpecifiers.ToArray(make([]OrderSpecifier, 0)))
	var total = this.Query.FetchCount()
	this.Query.Offset(pageModel.GetPageSize() * (pageModel.GetPageNo() - 1)).Limit(pageModel.GetPageSize())
	var list = this.Query.Fetch()
	return PageImpl(list, pageModel.GetPageable(), total)
}
func NewAdminAccessLogService(adminAccessLogDao dao.AdminAccessLogDao) (ret *AdminAccessLogService) {
	ret = new(AdminAccessLogService)
	ret.AdminAccessLogDao = adminAccessLogDao
	return
}

type AdminAccessLogService struct {
	AdminAccessLogDao dao.AdminAccessLogDao
	Base.BaseService
}
