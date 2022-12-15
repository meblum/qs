package qs

type delete string

func Delete(table string) delete {
	return delete("DELETE FROM " + table)
}

func (d delete) String() string {
	return string(d)
}

func (d delete) Where(col string) updateWhere {
	return updateWhere(d) + " WHERE " + updateWhere(col)
}
