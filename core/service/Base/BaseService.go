package Base

import (
	"bitrade/core/dao/base"
	"bitrade/core/dao/types"
	"bitrade/core/pagination"
	"github.com/qauzy/chocolate/lists/arraylist"
	"strings"
)

func (this *BaseService[T]) QueryDsl(pageNo int, pageSize int, predicateList arraylist.List[types.Predicate], entityPathBase dsl.EntityPathBase[T], orderSpecifierList arraylist.List[types.OrderSpecifier]) (result pagination.PageResult[T], err error) {
	var list arraylist.List[T]
	//查询表
	var jpaQuery = this.QueryFactory.SelectFrom(entityPathBase)
	//查询条件
	if predicateList != nil && predicateList.Size() > 0 {
		jpaQuery.Where(predicateList.ToArray(make([]Predicate, 0)))
	}
	//排序方式
	if orderSpecifierList != nil && orderSpecifierList.Size() > 0 {
		jpaQuery.OrderBy(orderSpecifierList.ToArray(make([]OrderSpecifier, 0)))
	}
	//分页查询
	if pageNo != 0 && pageSize != 0 {
		list = jpaQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	} else {
		list = jpaQuery.Fetch()
	}
	return pagination.NewPageResultEx(list, pageNo, pageSize, jpaQuery.FetchCount())
}
func (this *BaseService[T]) QueryOneDsl(predicate *types.Predicate, entityPathBase dsl.EntityPathBase[T]) (result T, err error) {
	return this.QueryFactory.SelectFrom(entityPathBase).Where(predicate).FetchFirst()
}
func (this *BaseService[T]) QueryDslForPageListResult(expressions arraylist.List[types.Expression], entityPaths arraylist.List[types.EntityPath], predicates arraylist.List[types.Predicate], orderSpecifierList arraylist.List[types.OrderSpecifier], pageNo int, pageSize int) (result *pagination.PageListMapResult, err error) {
	var jpaQuery = this.QueryFactory.Select(expressions.ToArray(make([]Expression, 0))).From(entityPaths.ToArray(make([]EntityPath, 0))).Where(predicates.ToArray(make([]Predicate, 0)))
	var tuples = jpaQuery.OrderBy(orderSpecifierList.ToArray(make([]OrderSpecifier, 0))).Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	var list = new(util.LinkedList)
	//返回结果
	//封装结果
	for i := 0; i < tuples.Size(); i += 1 {
		//遍历tuples
		var oMap = new(util.LinkedHashMap)
		//一条信息
		for _, expression := range expressions {
			oMap.Put(strings.Split(expression.ToString(), " as ")[1], tuples.Get(i).Get(expression))
			//获取结果
		}
		list.Add(oMap)
	}
	var pageListMapResult = PageListMapResult(list, pageNo, pageSize, jpaQuery.FetchCount())
	//分页封装
	return pageListMapResult
}
func (this *BaseService[T]) QueryDslForPageListResult(qdc *pagination.QueryDslContext, pageNo int, pageSize int) (result *pagination.PageListMapResult, err error) {
	var jpaQuery = this.QueryFactory.Select(qdc.ExpressionToArray()).From(qdc.EntityPathToArray()).Where(qdc.PredicatesToArray())
	var tuples = jpaQuery.OrderBy(qdc.OrderSpecifiersToArray()).Offset((pageNo - 1) * pageSize).Limit(pageSize).Fetch()
	var list = new(util.LinkedList)
	//返回结果
	//封装结果
	for i := 0; i < tuples.Size(); i += 1 {
		//遍历tuples
		var oMap = new(util.LinkedHashMap)
		//一条信息
		for _, expression := range qdc.GetExpressions() {
			oMap.Put(strings.Split(expression.ToString(), " as ")[1], tuples.Get(i).Get(expression))
			//获取结果
		}
		list.Add(oMap)
	}
	var pageListMapResult = PageListMapResult(list, pageNo, pageSize, jpaQuery.FetchCount())
	//分页封装
	return pageListMapResult
}
func NewBaseService(queryFactory *impl.JPAQueryFactory) (ret *BaseService) {
	ret = new(BaseService)
	ret.QueryFactory = queryFactory
	return
}

type BaseService[T any, D base.BaseDao] struct {
	QueryFactory *impl.JPAQueryFactory
	TopBaseService[T, D]
}
