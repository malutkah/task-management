package database

func (t *Table) Update(set []Set, condition []Condition) {
	for i := 0; i < len(t.rows); i++ {
		if t.rowMatch(condition, i) {
			for _, s := range set {
				t.rows[i][t.getColumnIndex(s.Column)] = s.Value
			}
		}
	}
}
