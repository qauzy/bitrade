package service

import (
	"bitrade/core/constant/CommonStatus"
	"bitrade/core/constant/SysAdvertiseLocation"
	"bitrade/core/dao"
	"bitrade/core/dao/types"
	"bitrade/core/entity"
	"bitrade/core/pagination"
	"bitrade/core/service/Base"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *SysAdvertiseService) Query(predicateList arraylist.List[types.Predicate], pageNo int, pageSize int) (result pagination.PageResult[entity.SysAdvertise], err error) {
	var list arraylist.List[entity.SysAdvertise]
	var jpaQuery = queryFactory.SelectFrom(sysAdvertise)
	if predicateList != nil {
		jpaQuery.Where(predicateList.ToArray(make([]Predicate, 0)))
	}
	if pageNo != nil && pageSize != nil {
		list = jpaQuery.OrderBy(OrderSpecifier(Order.DESC, sysAdvertise.CreateTime)).Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.OrderBy(OrderSpecifier(Order.DESC, sysAdvertise.CreateTime)).Fetch()
	}
	return PageResult(list, jpaQuery.FetchCount())
}
func (this *SysAdvertiseService) FindOne(serialNumber int64) (result *entity.SysAdvertise, err error) {
	return this.SysAdvertiseDao.FindById(serialNumber)
}
func (this *SysAdvertiseService) GetMaxSort() (result int, err error) {
	return this.SysAdvertiseDao.FindMaxSort()
}
func (this *SysAdvertiseService) Save(sysAdvertise *entity.SysAdvertise) (result *entity.SysAdvertise, err error) {
	return this.SysAdvertiseDao.Save(sysAdvertise)
}
func (this *SysAdvertiseService) FindAll() (result arraylist.List[entity.SysAdvertise], err error) {
	return this.SysAdvertiseDao.FindAll()
}
func (this *SysAdvertiseService) DeleteOne(serialNumber int64) (err error) {
	this.SysAdvertiseDao.DeleteById(serialNumber)
}
func (this *SysAdvertiseService) DeleteBatch(array []int64) (err error) {
	for _, serialNumber := range array {
		this.SysAdvertiseDao.DeleteById(serialNumber)
	}
}
func (this *SysAdvertiseService) FindAllNormal(sysAdvertiseLocation *SysAdvertiseLocation.SysAdvertiseLocation) (result arraylist.List[entity.SysAdvertise], err error) {
	return this.SysAdvertiseDao.FindAllByStatusAndSysAdvertiseLocationOrderBySortDesc(CommonStatus.NORMAL, sysAdvertiseLocation)
}
func (this *SysAdvertiseService) FindAll(predicateList arraylist.List[types.Predicate]) (result arraylist.List[entity.SysAdvertise], err error) {
	var list arraylist.List[entity.SysAdvertise]
	var jpaQuery = queryFactory.SelectFrom(sysAdvertise)
	if predicateList != nil {
		jpaQuery.Where(predicateList.ToArray(make([]Predicate, 0)))
	}
	list = jpaQuery.OrderBy(OrderSpecifier(Order.DESC, sysAdvertise.CreateTime)).Fetch()
	return list
}
func (this *SysAdvertiseService) FindAll(predicate *types.Predicate, pageable *domain.Pageable) (result domain.Page[entity.SysAdvertise], err error) {
	return this.SysAdvertiseDao.FindAll(predicate, pageable)
}
func NewSysAdvertiseService(sysAdvertiseDao *dao.SysAdvertiseDao) (ret *SysAdvertiseService) {
	ret = new(SysAdvertiseService)
	ret.SysAdvertiseDao = sysAdvertiseDao
	return
}

type SysAdvertiseService struct {
	SysAdvertiseDao dao.SysAdvertiseDao
	Base.BaseService
}
