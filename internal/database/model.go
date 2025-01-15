package database

import (
	"fmt"
	"github.com/google/uuid"
)

const (
	ColumnInt = 1 << iota
	ColumnVarchar
	ColumnBool
	ColumnDate
	ColumnDateTime
)

type Database struct {
	tables map[string]Table
}

// CreateTable creates a new table with the given name and columns.
// It returns an error if a table with the same name already exists.
func (db *Database) CreateTable(name string, cols []Column) (*Table, error) {
	if db.ifTableExists(name) {
		return &Table{}, fmt.Errorf("a table with the name '%s' already exists", name)
	}

	newTable := NewTable(name, cols)
	db.tables[newTable.Name] = newTable
	return &newTable, nil
}

func (db *Database) ifTableExists(name string) bool {
	_, ok := db.tables[name]
	return ok
}

type Column struct {
	Name string
	Type int
}

type Table struct {
	Name    string
	columns []Column
	rows    [][]any
	colIdx  map[string]int
}

func (t *Table) Insert(data ...any) {
	t.rows = append(t.rows, data)
}

func (t *Table) getColumnIndex(name string) int {
	return t.colIdx[name]
}

type Condition struct {
	Column string
	Equals string
	Value  any
}

type Expression struct {
	Column string
}

type Set struct {
	Column string
	Value  any
}

func GetRandomID() string {
	return uuid.New().String()
}
