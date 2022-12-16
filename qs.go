package qs

import (
	"strconv"
	"strings"
)

type selct string
type insert string
type delete string
type update string

// conditions
type where string
type not string
type condition string

// select
type from string
type orderBy string
type nextOrderBy string
type limit string
type offset string

// update
type updateColumn string
type updateWhere string
type updateNot string
type updateCondition string

type returning string

type val struct {
	table    string
	col, val []string
}

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

func InsertInto(table string) insert {
	return insert(table)
}

func (v insert) Value(col, value string) val {
	return val{table: string(v), col: []string{col}, val: []string{value}}
}
func (v val) Value(col, value string) val {
	return val{table: v.table, col: append(v.col, col), val: append(v.val, value)}
}
func (v val) String() string {
	return "INSERT INTO " + v.table + "(" + strings.Join(v.col, ", ") + ") VALUES (" + strings.Join(v.val, ", ") + ")"
}
func (v val) Returning(col ...string) returning {
	q := "INSERT INTO " + v.table + "(" + strings.Join(v.col, ", ") + ") VALUES (" + strings.Join(v.val, ", ") + ")"
	return returning(q + " RETURNING " + strings.Join(col, ", "))
}
func (r returning) String() string {
	return string(r)
}

func Delete(table string) delete {
	return delete("DELETE FROM " + table)
}

func (d delete) String() string {
	return string(d)
}
func (d delete) Returning(cols ...string) returning {
	return returning(d) + returning(" RETURNING "+strings.Join(cols, ", "))
}

func (d delete) Where(col string) updateWhere {
	return updateWhere(d) + " WHERE " + updateWhere(col)
}

func Update(table string) update {
	return update("UPDATE " + table + " SET ")
}

func (u update) Col(col string, val string) updateColumn {
	return updateColumn(u) + " = " + updateColumn(val)
}
func (u updateColumn) Col(col string, val string) updateColumn {
	return u + updateColumn(", ") + updateColumn(col) + " = " + updateColumn(val)
}

func (u updateColumn) Returning(cols ...string) returning {
	return returning(u) + returning(" RETURNING "+strings.Join(cols, ", "))
}

func (u updateColumn) String() string {
	return string(u)
}

func (f updateColumn) Where(col string) updateWhere {
	return updateWhere(f) + " WHERE " + updateWhere(col)
}

func (w updateWhere) Equals(val string) updateCondition {
	return updateCondition(w) + " = " + updateCondition(val)
}

func (w updateWhere) GreaterThan(val string) updateCondition {
	return updateCondition(w) + " > " + updateCondition(val)
}

func (w updateWhere) LessThan(val string) updateCondition {
	return updateCondition(w) + " < " + updateCondition(val)
}

func (w updateWhere) GreaterOrEqualThan(val string) updateCondition {
	return updateCondition(w) + " >= " + updateCondition(val)
}

func (w updateWhere) LessOrEqualThan(val string) updateCondition {
	return updateCondition(w) + " <= " + updateCondition(val)
}

func (w updateWhere) NotEqual(val string) updateCondition {
	return updateCondition(w) + " != " + updateCondition(val)
}

func (w updateWhere) In(val ...string) updateCondition {
	v := "(" + strings.Join(val, ",") + ")"
	return updateCondition(w) + " IN " + updateCondition(v)
}

func (w updateWhere) Like(val string) updateCondition {
	return updateCondition(w) + " LIKE " + updateCondition(val)
}

func (w updateWhere) ILike(val string) updateCondition {
	return updateCondition(w) + " ILIKE " + updateCondition(val)
}

func (w updateWhere) IsNull() updateCondition {
	return updateCondition(w) + " IS NULL"
}

func (w updateWhere) IsNotNull() updateCondition {
	return updateCondition(w) + " IS NOT NULL"
}

func (w updateWhere) Not() updateNot {
	return updateNot(w) + updateNot(" NOT")
}

func (w updateWhere) Between(val1, val2 string) updateCondition {
	return updateCondition(w) + " BETWEEN " + updateCondition(val1+" AND "+val2)
}

func (w updateNot) Equals(val string) updateCondition {
	return updateCondition(w) + " = " + updateCondition(val)
}

func (w updateNot) GreaterThan(val string) updateCondition {
	return updateCondition(w) + " > " + updateCondition(val)
}

func (w updateNot) LessThan(val string) updateCondition {
	return updateCondition(w) + " < " + updateCondition(val)
}

func (w updateNot) GreaterOrEqualThan(val string) updateCondition {
	return updateCondition(w) + " >= " + updateCondition(val)
}

func (w updateNot) LessOrEqualThan(val string) updateCondition {
	return updateCondition(w) + " <= " + updateCondition(val)
}

func (w updateNot) In(val ...string) updateCondition {
	v := "(" + strings.Join(val, ",") + ")"
	return updateCondition(w) + " IN " + updateCondition(v)
}

func (w updateNot) Like(val string) updateCondition {
	return updateCondition(w) + " LIKE " + updateCondition(val)
}

func (w updateNot) ILike(val string) updateCondition {
	return updateCondition(w) + " ILIKE " + updateCondition(val)
}

func (w updateNot) Between(val1, val2 string) updateCondition {
	return updateCondition(w) + " BETWEEN " + updateCondition(val1+" AND "+val2)
}

func (b updateNot) And(val string) updateCondition {
	return updateCondition(b) + " AND " + updateCondition(val)
}
func (u updateCondition) Returning(cols ...string) returning {
	return returning(u) + returning(" RETURNING "+strings.Join(cols, ", "))
}
func (c updateCondition) String() string {
	return string(c)
}

func (c updateCondition) And(col string) updateWhere {
	return updateWhere(c) + " AND " + updateWhere(col)
}

func (c updateCondition) Or(col string) updateWhere {
	return updateWhere(c) + " OR " + updateWhere(col)
}
