package database

import (
	"fmt"
	"log"
	"testing"
	
	"github.com/google/go-cmp/cmp"
)

func TestDatabase(t *testing.T) {
	db := NewDatabase()
	
	// CREATE TABLE cities (id int, name varchar, population int);
	tab, err := db.CreateTable("cities", []Column{
		{"id", ColumnInt},
		{"name", ColumnVarchar},
		{"population", ColumnInt},
		{"good_weather", ColumnBool},
	})
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	_, err = db.CreateTable("cities", []Column{
		{"id", ColumnInt},
		{"name", ColumnVarchar},
		{"population", ColumnInt},
		{"good_weather", ColumnBool},
	})
	
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(tab)
}

func TestTable_Insert(t *testing.T) {
	db := NewDatabase()
	
	// CREATE TABLE cities (id int, name varchar, population int);
	tab, err := db.CreateTable("cities", []Column{
		{"id", ColumnInt},
		{"name", ColumnVarchar},
		{"population", ColumnInt},
		{"good_weather", ColumnBool},
	})
	
	if diff := cmp.Diff(len(db.tables), 1); diff != "" {
		t.Errorf("Select: (-want +got)\n%s", diff)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	
	tab.Insert(1, "Tokyo", 37, true)
	tab.Insert(2, "London", 20, false)
	tab.Insert(3, "Berlin", 4, false)
	
	fmt.Println(tab)
}

func TestTable_Select(t *testing.T) {
	db := NewDatabase()
	
	tab, err := db.CreateTable("cities", []Column{
		{"id", ColumnInt},
		{"name", ColumnVarchar},
		{"population", ColumnInt},
		{"good_weather", ColumnBool},
	})
	
	if diff := cmp.Diff(len(db.tables), 1); diff != "" {
		t.Errorf("Select: (-want +got)\n%s", diff)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	
	tab.Insert(1, "Tokyo", 37, true)
	tab.Insert(2, "London", 20, false)
	tab.Insert(3, "Berlin", 4, false)
	
	var got, want [][]any
	
	// select * from cities
	got = tab.Select([]Expression{{"*"}})
	want = [][]any{{1, "Tokyo", 37, true}, {2, "London", 20, false}, {3, "Berlin", 4, false}}
	
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Select (-want +got)\n%s", diff)
	}
	
	fmt.Println(got)
	
	// select * population from cities where name = "Tokyo"
	got = tab.Select(
		[]Expression{{"population"}},
		[]Condition{{"name", ConditionEqual, "Tokyo"}}...,
	)
	want = [][]any{{37}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Select (-want +got)\n%s", diff)
	}
	fmt.Println(got)
	
	// select * from cities where != "London" and population = 4
	got = tab.Select([]Expression{{"*"}}, []Condition{{"name", ConditionUnequal, "London"}, {"population", ConditionEqual, 4}}...)
	
	fmt.Println(got)
}

func TestTable_Update(t *testing.T) {
	db := NewDatabase()
	
	tab, err := db.CreateTable("cities", []Column{
		{"id", ColumnInt},
		{"name", ColumnVarchar},
		{"population", ColumnInt},
		{"good_weather", ColumnBool},
	})
	
	if diff := cmp.Diff(len(db.tables), 1); diff != "" {
		t.Errorf("Select: (-want +got)\n%s", diff)
	}
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	tab.Insert(1, "Tokyo", 37, true)
	tab.Insert(2, "London", 20, false)
	tab.Insert(3, "Berlin", 4, false)
	tab.Insert(4, "Cologne", 1, true)
	
	// update cities set population = 15 where name = "London"
	tab.Update([]Set{{"population", 15}}, []Condition{{"name", ConditionEqual, "London"}})
	got := tab.Select([]Expression{{"*"}})
	fmt.Println(got)
}

func TestTable_Delete(t *testing.T) {
	db := NewDatabase()
	
	tab, err := db.CreateTable("cities", []Column{
		{"id", ColumnInt},
		{"name", ColumnVarchar},
		{"population", ColumnInt},
		{"good_weather", ColumnBool},
	})
	
	if diff := cmp.Diff(len(db.tables), 1); diff != "" {
		t.Errorf("Select: (-want +got)\n%s", diff)
	}
	
	if err != nil {
		log.Fatal(err.Error())
	}
	
	tab.Insert(1, "Tokyo", 37, true)
	tab.Insert(2, "London", 20, false)
	tab.Insert(3, "Berlin", 4, false)
	tab.Insert(4, "Cologne", 1, true)
	
	tab.Delete([]Condition{{"population", ConditionGreater, 4}})
	
	got := tab.Select(ExpSelectAll)
	fmt.Println(got)
}
