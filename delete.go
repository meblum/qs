package qs

import "strings"

type delete string

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
