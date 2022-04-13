package types

import "github.com/jinzhu/gorm"

// Predicate is a string that acts as a condition in the where clause
type Predicate string

const (
	InPredicate                 = Predicate("IN")
	EqualPredicate              = Predicate("=")
	NotEqualPredicate           = Predicate("<>")
	GreaterThanPredicate        = Predicate(">")
	GreaterThanOrEqualPredicate = Predicate(">=")
	SmallerThanPredicate        = Predicate("<")
	SmallerThanOrEqualPredicate = Predicate("<=")
	LikePredicate               = Predicate("LIKE")
)

func NewQueryParam() *QueryParam {
	return &QueryParam{
		limit:  10,
		offset: 0,
	}
}

type QueryParam struct {
	order []string
	where []struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}
	limit  int
	offset int
}

func (qb *QueryParam) Offset(offset int) *QueryParam {
	qb.offset = offset
	return qb
}

func (qb *QueryParam) Limit(limit int) *QueryParam {
	qb.limit = limit
	return qb
}
func (qb *QueryParam) Neq(field string, value interface{}) *QueryParam {
	qb.where = append(qb.where, struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}{
		prefix:    field,
		predicate: NotEqualPredicate,
		value:     value,
	})
	return qb
}
func (qb *QueryParam) Eq(field string, value interface{}) *QueryParam {
	qb.where = append(qb.where, struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}{
		prefix:    field,
		predicate: EqualPredicate,
		value:     value,
	})
	return qb
}
func (qb *QueryParam) Gt(field string, value interface{}) *QueryParam {
	qb.where = append(qb.where, struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}{
		prefix:    field,
		predicate: GreaterThanPredicate,
		value:     value,
	})
	return qb

}
func (qb *QueryParam) Gte(field string, value interface{}) *QueryParam {
	qb.where = append(qb.where, struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}{
		prefix:    field,
		predicate: GreaterThanOrEqualPredicate,
		value:     value,
	})
	return qb
}
func (qb *QueryParam) Lt(field string, value interface{}) *QueryParam {
	qb.where = append(qb.where, struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}{
		prefix:    field,
		predicate: SmallerThanPredicate,
		value:     value,
	})
	return qb
}
func (qb *QueryParam) Lte(field string, value interface{}) *QueryParam {
	qb.where = append(qb.where, struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}{
		prefix:    field,
		predicate: SmallerThanOrEqualPredicate,
		value:     value,
	})
	return qb
}
func (qb *QueryParam) Like(field string, value interface{}) *QueryParam {
	qb.where = append(qb.where, struct {
		prefix    string
		predicate Predicate
		value     interface{}
	}{
		prefix:    field,
		predicate: LikePredicate,
		value:     value,
	})
	return qb
}

func (qb *QueryParam) BuildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}
