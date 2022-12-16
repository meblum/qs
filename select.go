package qs

import (
	"strconv"
	"strings"
)

type selct string
type from string
type where string
type not string
type condition string
type orderBy string
type nextOrderBy string
type limit string
type offset string

func Select(cols ...string) selct {
	q := "SELECT " + strings.Join(cols, ", ") + " FROM "
	return selct(q)
}

func (s selct) From(table string) from {
	return from(s) + from(table)
}

func (f from) String() string {
	return string(f)
}

func (f from) Offset(val int) offset {
	if val < 0 {
		val = 0
	}
	return offset(f) + offset(" OFFSET "+strconv.Itoa(val))
}

func (f from) Where(col string) where {
	return where(f) + " WHERE " + where(col)
}

func (f from) OrderBy(col string) orderBy {
	return orderBy(f) + orderBy(" ORDER BY ")
}

func (f from) Limit(val int) limit {
	if val < 0 {
		val = 0
	}
	return limit(f) + limit(" LIMIT "+strconv.Itoa(val))
}

func (w where) Equals(val string) condition {
	return condition(w) + " = " + condition(val)
}

func (w where) GreaterThan(val string) condition {
	return condition(w) + " > " + condition(val)
}

func (w where) LessThan(val string) condition {
	return condition(w) + " < " + condition(val)
}

func (w where) GreaterOrEqualThan(val string) condition {
	return condition(w) + " >= " + condition(val)
}

func (w where) LessOrEqualThan(val string) condition {
	return condition(w) + " <= " + condition(val)
}

func (w where) NotEqual(val string) condition {
	return condition(w) + " != " + condition(val)
}

func (w where) In(val ...string) condition {
	v := "(" + strings.Join(val, ",") + ")"
	return condition(w) + " IN " + condition(v)
}

func (w where) Like(val string) condition {
	return condition(w) + " LIKE " + condition(val)
}

func (w where) ILike(val string) condition {
	return condition(w) + " ILIKE " + condition(val)
}

func (w where) IsNull() condition {
	return condition(w) + " IS NULL"
}

func (w where) IsNotNull() condition {
	return condition(w) + " IS NOT NULL"
}

func (w where) Not() not {
	return not(w) + not(" NOT")
}

func (w where) Between(val1, val2 string) condition {
	return condition(w) + " BETWEEN " + condition(val1+" AND "+val2)
}

func (w not) Equals(val string) condition {
	return condition(w) + " = " + condition(val)
}

func (w not) GreaterThan(val string) condition {
	return condition(w) + " > " + condition(val)
}

func (w not) LessThan(val string) condition {
	return condition(w) + " < " + condition(val)
}

func (w not) GreaterOrEqualThan(val string) condition {
	return condition(w) + " >= " + condition(val)
}

func (w not) LessOrEqualThan(val string) condition {
	return condition(w) + " <= " + condition(val)
}

func (w not) In(val ...string) condition {
	v := "(" + strings.Join(val, ",") + ")"
	return condition(w) + " IN " + condition(v)
}

func (w not) Like(val string) condition {
	return condition(w) + " LIKE " + condition(val)
}

func (w not) ILike(val string) condition {
	return condition(w) + " ILIKE " + condition(val)
}

func (w not) Between(val1, val2 string) condition {
	return condition(w) + " BETWEEN " + condition(val1+" AND "+val2)
}

func (c condition) String() string {
	return string(c)
}

func (c condition) And(col string) where {
	return where(c) + " AND " + where(col)
}

func (c condition) Or(col string) where {
	return where(c) + " OR " + where(col)
}

func (c condition) OrderBy() orderBy {
	return orderBy(c) + orderBy(" ORDER BY ")
}

func (c condition) Limit(val int) limit {
	if val < 0 {
		val = 0
	}
	return limit(c) + limit(" LIMIT "+strconv.Itoa(val))
}

func (c condition) Offset(val int) offset {
	if val < 0 {
		val = 0
	}
	return offset(c) + offset(" OFFSET "+strconv.Itoa(val))
}

func (o orderBy) Asc(col string) nextOrderBy {
	return nextOrderBy(o) + nextOrderBy(col+" ASC")
}
func (o orderBy) Desc(col string) nextOrderBy {
	return nextOrderBy(o) + nextOrderBy(col+" DESC")
}

func (o nextOrderBy) Asc(col string) nextOrderBy {
	return nextOrderBy(o) + nextOrderBy(", "+col+" DESC")
}

func (o nextOrderBy) Desc(col string) nextOrderBy {
	return nextOrderBy(o) + nextOrderBy(", "+col+" DESC")
}

func (o nextOrderBy) Limit(val int) limit {
	if val < 0 {
		val = 0
	}
	return limit(o) + limit(" LIMIT "+strconv.Itoa(val))
}

func (o nextOrderBy) Offset(val int) offset {
	if val < 0 {
		val = 0
	}
	return offset(o) + offset(" OFFSET "+strconv.Itoa(val))
}

func (o nextOrderBy) String() string {
	return string(o)
}

func (l limit) String() string {
	return string(l)
}

func (l limit) Offset(val int) offset {
	if val < 0 {
		val = 0
	}
	return offset(l) + offset(" OFFSET "+strconv.Itoa(val))
}

func (o offset) String() string {
	return string(o)
}
