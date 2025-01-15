package database

func (t *Table) Delete(condition []Condition) {
	for i := 0; i < len(t.rows); i++ {
		if t.rowMatch(condition, i) {
			t.rows[i] = nil
		}
	}
}
