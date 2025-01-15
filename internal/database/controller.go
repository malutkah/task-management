package database

func NewDatabase() *Database {
	return &Database{
		tables: make(map[string]Table),
	}
}

// NewTable creates a new Table with the given name and columns.
// It initializes the column index map for quick lookups by column name.
func NewTable(name string, cols []Column) Table {
	t := Table{
		Name:    name,
		columns: cols,
	}
	t.colIdx = make(map[string]int, len(cols))
	for i, col := range cols {
		t.colIdx[col.Name] = i
	}
	return t
}
