package qs

import "strings"

// INSERT INTO table_name(column1, column2, …)
// VALUES (value1, value2, …);

type insert string

type val struct {
	table    string
	col, val []string
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
