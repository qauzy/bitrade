
package pagination

import "github.com/qauzy/chocolate/lists/arraylist"

func (this *LogicalExpression) SetCriterion(Criterion []Criterion) (result *LogicalExpression) {
	this.Criterion = Criterion
	return this
}
func (this *LogicalExpression) GetCriterion() (Criterion []Criterion) {
	return this.Criterion
}
func (this *LogicalExpression) SetOperator(Operator *Operator) (result *LogicalExpression) {
	this.Operator = Operator
	return this
}
func (this *LogicalExpression) GetOperator() (Operator *Operator) {
	return this.Operator
}
func NewLogicalExpression(pageCriteria []Criterion, Operator Operator) (this *LogicalExpression) {
	this = new(LogicalExpression)
	this.Criterion = pageCriteria
	this.Operator = this.Operator
	return
}
func (this *LogicalExpression) ToPredicate(root Root[?], query CriteriaQuery[?], builder CriteriaBuilder) (result Predicate) {
	var predicates = arraylist.New[Predicate]()
	for i := 0; i < len(this.Criterion); i += 1 {
		predicates.Add(this.Criterion[i].ToPredicate(root, query, builder))
	}
	switch this.Operator {
	case OR:
		return builder.Or(predicates.ToArray(make([]Predicate, 0)))
		default:
			return nil
	}
}

type LogicalExpression struct {
	Criterion []Criterion `gorm:"column:criterion" json:"criterion"`
	Operator  *Operator   `gorm:"column:operator" json:"operator"`
}

