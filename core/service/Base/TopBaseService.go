package Base

import (
	"bitrade/core/constant/PageModel"
	"bitrade/core/dao/base"
	"bitrade/core/dao/types"
	"bitrade/core/dto"
	"github.com/qauzy/chocolate/lists/arraylist"
	"github.com/qauzy/math"
)

func (this *TopBaseService[E, D]) FindById(id int64) (result E, err error) {
	return this.Dao.FindById(id)
}
func (this *TopBaseService[E, D]) FindAll() (result arraylist.List[E], err error) {
	return this.Dao.FindAll()
}
func (this *TopBaseService[E, D]) Delete(id int64) (err error) {
	this.Dao.Delete(id)
}
func (this *TopBaseService[E, D]) Deletes(ids []int64) (err error) {
	for _, id := range ids {
		this.Delete(id)
	}
}
func (this *TopBaseService[E, D]) Save(e E) (result E, err error) {
	return this.Dao.Save(e)
}

}
func (this *TopBaseService[E, D]) Update(updateAbility ability.UpdateAbility[E], e E) (result E, err error) {
	return this.Dao.Save(updateAbility.Transformation(e)).(E)
}
func (this *TopBaseService[E, D]) FindAll(predicate *types.Predicate, pageModel *PageModel.PageModel) (result Page[E], err error) {
	return this.Dao.FindAll(predicate, pageModel.GetPageable())
}
func (this *TopBaseService[E, D]) FindAllScreen(screenAbility *ability.ScreenAbility, pageModel *PageModel.PageModel) (result Page[E], err error) {
	return this.Dao.FindAll(screenAbility.GetPredicate(), pageModel.GetPageable())
}
func (this *TopBaseService[E, D]) PageQuery(pagenation *dto.Pagenation, predicate *types.Predicate) (result dto.Pagenation[E], err error) {
	// FIXME
	var sort = Sort.By(pagenation.GetPageParam().GetDirection(), pagenation.GetPageParam().GetOrders().ToString())
	var pageable = PageRequest.Of(pagenation.GetPageParam().GetPageNo()-1, pagenation.GetPageParam().GetPageSize(), sort)
	var page = this.Dao.FindAll(predicate, pageable)
	return pagenation.SetData(page.GetContent(), page.GetTotalElements(), page.GetTotalPages())
}
//func (this *TopBaseService[E, D]) CreateNativePageQuery(countSql StringBuilder, sql StringBuilder, pageModel *PageModel.PageModel, result *transform.ResultTransformer) (result Page, err error) {
//	var query1 = this.EntityManager.CreateNativeQuery(countSql.ToString())
//	var count = query1.GetSingleResult().(math.BigInteger).LongValue()
//	if pageModel.GetProperty() != nil && pageModel.GetProperty().Size() > 0 && pageModel.GetDirection().Size() == pageModel.GetProperty().Size() {
//		sql.Append(" order by")
//		for i := 0; i < pageModel.GetProperty().Size(); i += 1 {
//			sql.Append(" " + pageModel.GetProperty().Get(i) + " " + pageModel.GetDirection().Get(i) + " ")
//			if i < pageModel.GetProperty().Size()-1 {
//				sql.Append(",")
//			}
//		}
//	}
//	sql.Append(" limit " + pageModel.GetPageSize()*(pageModel.GetPageNo()-1) + " , " + pageModel.GetPageSize())
//	var query2 = this.EntityManager.CreateNativeQuery(sql.ToString())
//	query2.Unwrap(hibernate.SQLQuery).SetResultTransformer(result)
//	var list = query2.GetResultList()
//	return PageImpl(list, pageModel.GetPageable(), count)
//}
func NewTopBaseService(entityManager *persistence.EntityManager, dao D) (ret *TopBaseService) {
	ret = new(TopBaseService)
	ret.EntityManager = entityManager
	ret.Dao = dao
	return
}

type TopBaseService[E any, D base.BaseDao[E]] struct {
	EntityManager *persistence.EntityManager
	Dao D
}
