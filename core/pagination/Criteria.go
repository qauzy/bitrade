
package pagination

import (
	"github.com/qauzy/chocolate/lists/arraylist"
	"strings"
)

func (this *Criteria[T]) SetCriteria(Criteria *arraylist.List[Criterion]) (result *Criteria[T]) {
	this.Criteria = Criteria
	return this
}
func (this *Criteria[T]) GetCriteria() (Criteria *arraylist.List[Criterion]) {
	return this.Criteria
}
func (this *Criteria[T]) ToPredicate(root Root[T], query CriteriaQuery[?], builder CriteriaBuilder) (result Predicate) {
	if !this.Criteria.IsEmpty() {
		var predicates = arraylist.New[Predicate]()
		for _, c := range this.Criteria {
			predicates.Add(c.ToPredicate(root, query, builder))
		}
		// 将所有条件用 and 联合起来
		if predicates.Size() > 0 {
			return builder.And(predicates.ToArray(make([]Predicate, 0)))
		}
	}
	return builder.Conjunction()
}
func (this *Criteria[T]) Add(criterion Criterion) {
	if criterion != nil {
		this.Criteria.Add(criterion)
	}
}
func (this *Criteria[T]) Sort(. String) (result *domain.Sort) {
	var orders = arraylist.New[>]()
	for _, f := range fields {
		orders.Add(GenerateOrderStatic(f))
	}
	return Sort.By(orders)
}
func SortStatic(. String) (result *domain.Sort) {
	var orders = arraylist.New[>]()
	for _, f := range fields {
		orders.Add(GenerateOrderStatic(f))
	}
	return Sort.By(orders)
}
func GenerateOrderStatic(f string) (result generateOrderStatic) {
	var order *Sort.Order
	var ff []string = strings.Split(f, "\\.")
	if len(ff) > 1 {
		if ff[1].Equals("desc") {
			order = Sort.Order(Sort.Direction.DESC, ff[0])
		} else {
			order = Sort.Order(Sort.Direction.ASC, ff[0])
		}
		return order
	}
	order = Sort.Order(Sort.Direction.DESC, f)
	return order
}

type Criteria[T any] struct {
	Criteria *arraylist.List[Criterion] `gorm:"column:criteria" json:"criteria"`
}

