package service

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/SysHelpClassification"
	"bitrade/core/dao"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *SysHelpService) Save(sysHelp *entity.SysHelp) (result *entity.SysHelp, err error) {
	return this.SysHelpDao.Save(sysHelp)
}
func (this *SysHelpService) FindAll(sort *domain.Sort) (result arraylist.List[entity.SysHelp], err error) {
	return this.SysHelpDao.FindAll(sort)
}
func (this *SysHelpService) FindOne(id int64) (result *entity.SysHelp, err error) {
	return this.SysHelpDao.FindById(id).Get()
}
func (this *SysHelpService) DeleteBatch(ids []int64) (err error) {
	for _, id := range ids {
		this.SysHelpDao.DeleteById(id)
	}
}
func (this *SysHelpService) GetMaxSort() (result int, err error) {
	return this.SysHelpDao.FindMaxSort()
}
func (this *SysHelpService) FindBySysHelpClassification(sysHelpClassification *SysHelpClassification.SysHelpClassification) (result arraylist.List[entity.SysHelp], err error) {
	return this.SysHelpDao.FindAllBySysHelpClassificationAndStatusNot(sysHelpClassification, CommonStatus.ILLEGAL)
}
func (this *SysHelpService) QueryWhereOrPage(booleanExpressionList arraylist.List[dsl.BooleanExpression], pageNo int, pageSize int) (result pagination.PageResult[entity.SysHelp], err error) {
	var jpaQuery = queryFactory.SelectFrom(QSysHelp.SysHelp)
	if booleanExpressionList != nil {
		jpaQuery.Where(booleanExpressionList.ToArray(make([]BooleanExpression, 0)))
	}
	jpaQuery.OrderBy(QSysHelp.SysHelp.CreateTime.Desc())
	var list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	var count = jpaQuery.FetchCount()
	var page = PageResult(list, pageNo, pageSize, count)
	return page
}
func (this *SysHelpService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result pagination.PageResult[entity.SysHelp], err error) {
	return this.SysHelpDao.FindAll(predicate, pageable)
}
func (this *SysHelpService) FindByCondition(pageNo int, pageSize int, cate SysHelpClassification.SysHelpClassification) (result pagination.PageResult[entity.SysHelp], err error) {
	var sort = Sort.By(Sort.Order(Sort.Direction.DESC, "sort"))
	var pageable = PageRequest.Of(pageNo-1, pageSize, sort)
	var specification = new(domain.Specification)
	return this.SysHelpDao.FindAll(specification, pageable)
}
func (this *SysHelpService) GetgetCateTops(cate string) (result arraylist.List[entity.SysHelp], err error) {
	return this.SysHelpDao.GetCateTop(cate)
}
func (this *SysHelpService) FindAllByStatusNotAndSort() (result arraylist.List[entity.SysHelp], err error) {
	return this.SysHelpDao.FindAllByStatusNotAndSort()
}
func NewSysHelpService(sysHelpDao dao.SysHelpDao) (ret *SysHelpService) {
	ret = new(SysHelpService)
	ret.SysHelpDao = sysHelpDao
	return
}

type SysHelpService struct {
	SysHelpDao dao.SysHelpDao
	Base.BaseService[entity.SysHelp, dao.SysHelpDao]
}
