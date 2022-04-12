package pagination

import (
	"bitrade/core/dao/types"
	"github.com/qauzy/chocolate/lists/arraylist"
)

func (this *QueryDslContext) SetExpressions(Expressions *arraylist.List[types.Expression]) (result *QueryDslContext) {
	this.Expressions = Expressions
	return this
}
func (this *QueryDslContext) GetExpressions() (Expressions *arraylist.List[types.Expression]) {
	return this.Expressions
}
func (this *QueryDslContext) SetEntityPaths(EntityPaths *arraylist.List[types.EntityPath]) (result *QueryDslContext) {
	this.EntityPaths = EntityPaths
	return this
}
func (this *QueryDslContext) GetEntityPaths() (EntityPaths *arraylist.List[types.EntityPath]) {
	return this.EntityPaths
}
func (this *QueryDslContext) SetPredicates(Predicates *arraylist.List[types.Predicate]) (result *QueryDslContext) {
	this.Predicates = Predicates
	return this
}
func (this *QueryDslContext) GetPredicates() (Predicates *arraylist.List[types.Predicate]) {
	return this.Predicates
}
func (this *QueryDslContext) SetOrderSpecifiers(OrderSpecifiers *arraylist.List[types.OrderSpecifier]) (result *QueryDslContext) {
	this.OrderSpecifiers = OrderSpecifiers
	return this
}
func (this *QueryDslContext) GetOrderSpecifiers() (OrderSpecifiers *arraylist.List[types.OrderSpecifier]) {
	return this.OrderSpecifiers
}
func NewQueryDslContext() (this *QueryDslContext) {
	this = new(QueryDslContext)
	this.Expressions = arraylist.New[types.Expression]()
	this.EntityPaths = arraylist.New[types.EntityPath]()
	this.Predicates = arraylist.New[types.Predicate]()
	this.OrderSpecifiers = arraylist.New[types.OrderSpecifier]()
	return
}
func (this *QueryDslContext) GetExpressions() (result *arraylist.List[types.Expression]) {
	return this.Expressions
}
func (this *QueryDslContext) Add(expression *types.Expression) {
	this.Expressions.Add(expression)
}
func (this *QueryDslContext) Add(entityPath *types.EntityPath) {
	this.EntityPaths.Add(entityPath)
}
func (this *QueryDslContext) Add(predicate *types.Predicate) {
	this.Predicates.Add(predicate)
}
func (this *QueryDslContext) Add(orderSpecifier *types.OrderSpecifier) {
	this.OrderSpecifiers.Add(orderSpecifier)
}
func (this *QueryDslContext) ExpressionToArray() (result []types.Expression) {
	return this.Expressions.ToArray(make([]Expression, 0))
}
func (this *QueryDslContext) EntityPathToArray() (result []types.EntityPath) {
	return this.EntityPaths.ToArray(make([]EntityPath, 0))
}
func (this *QueryDslContext) PredicatesToArray() (result []types.Predicate) {
	return this.Predicates.ToArray(make([]Predicate, 0))
}
func (this *QueryDslContext) OrderSpecifiersToArray() (result []types.OrderSpecifier) {
	return this.OrderSpecifiers.ToArray(make([]OrderSpecifier, 0))
}

type QueryDslContext struct {
	Expressions     *arraylist.List[types.Expression]     `gorm:"column:expressions" json:"expressions"`
	EntityPaths     *arraylist.List[types.EntityPath]     `gorm:"column:entity_paths" json:"entityPaths"`
	Predicates      *arraylist.List[types.Predicate]      `gorm:"column:predicates" json:"predicates"`
	OrderSpecifiers *arraylist.List[types.OrderSpecifier] `gorm:"column:order_specifiers" json:"orderSpecifiers"`
}
