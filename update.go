package qs

import (
	"strings"
)

// q := fmt.Sprint("SELECT id, street, city, state, county, zip, property_type, ",
// 		"construction_type, lockbox, first_name, last_name, cell_phone, email, status, ",
// 		"created_by, updated_at FROM projects ", whereQ, "ORDER BY updated_at ", orderBy,
// 		" LIMIT $1 OFFSET $2", ";")

type update string
type column string
type updateWhere string
type updateNot string
type updateCondition string

func Update(table string) update {
	return update("UPDATE " + table + " SET ")
}

func (u update) Col(col string, val string) column {
	return column(u) + " = " + column(val)
}
func (u column) Col(col string, val string) column {
	return column(", ") + u + " = " + column(val)
}

func (u column) String() string {
	return string(u)
}

func (f column) Where(col string) updateWhere {
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

func (c updateCondition) String() string {
	return string(c)
}

func (c updateCondition) And(col string) updateWhere {
	return updateWhere(c) + " AND " + updateWhere(col)
}

func (c updateCondition) Or(col string) updateWhere {
	return updateWhere(c) + " OR " + updateWhere(col)
}
