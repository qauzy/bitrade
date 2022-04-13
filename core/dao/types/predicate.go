package types

import (
	"bitrade/core/dao/types/Direction"
	"github.com/jinzhu/gorm"
)

// Predicate is a string that acts as a condition in the where clause
type Predicate string

func (p Predicate) String() string {
	return string(p)
}

const (
	InPredicate                 = Predicate(" IN ?")
	EqualPredicate              = Predicate(" = ?")
	NotEqualPredicate           = Predicate(" <> ?")
	GreaterThanPredicate        = Predicate(" > ?")
	GreaterThanOrEqualPredicate = Predicate(" >= ?")
	SmallerThanPredicate        = Predicate(" < ?")
	SmallerThanOrEqualPredicate = Predicate(" <= ?")
	LikePredicate               = Predicate(" LIKE ?")
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
	in []struct {
		prefix string
		values []interface{}
	}
	limit  int
	offset int
}

func (q *QueryParam) In(prefix string, values []interface{}) *QueryParam {
	q.in = append(q.in, struct {
		prefix string
		values []interface{}
	}{prefix, values})
	return q
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
func (qb *QueryParam) Order(field string, direction Direction.Direction) *QueryParam {
	qb.order = append(qb.order, field+" "+string(direction))
	return qb
}
func (qb *QueryParam) BuildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix+where.predicate.String(), where.value)
	}
	for _, in := range qb.in {
		ret = ret.Where(in.prefix+InPredicate.String(), in.values)
	}

	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}
