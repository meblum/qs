# qs
**Work in progress**


qs is a simple Go package to buld sql statements using a fluent api.

For example:

```go
qs.Select("col_1", "col_2", "col_3").From("table").Where("col_3").Between("2").And("7").Or("col_1").LessThan("32").OrderBy().Acs("col_2").Limit(2).Offset(30).String()
