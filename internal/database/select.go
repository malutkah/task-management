package database

var ExpSelectAll = []Expression{{"*"}}

func (t *Table) Select(cols []Expression, condition ...Condition) [][]any {
	// TODO: validation for cols and condition
	
	var res [][]any
	
	for i := 0; i < len(t.rows); i++ {
		if t.rowMatch(condition, i) {
			res = append(res, t.filterColumns(cols, i))
		}
	}
	
	return res
}
