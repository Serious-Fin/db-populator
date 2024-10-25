package parser

import (
	"fmt"
	"testing"
)

// Correctly parses create table query without columns defined
func TestEmptyCreateQuery(t *testing.T) {
	// Arrange
	gotText := "CREATE TABLE Persons ();"
	want := Database{
		Tables: []Table{
			{
				Name: "MyTestTable",
			},
		},
	}

	// Act
	got, err := Parse(gotText)

	// Assert
	if err != nil {
		t.Fatalf(`Method "Parse" returned an error: %v`, err)
	}

	assertDatabaseEquality(&want, &got, t)
}

func assertDatabaseEquality(want, got *Database, t *testing.T) {
	errors := make(map[string][]string)

	for _, table := range want.Tables {
		found, actualTable := getTableFromDatabase(table.Name, got)

		if !found {
			errors[table.Name] = append(errors[table.Name], fmt.Sprintf(`Table with name %q was not found`, table.Name))
			continue
		}

		areEqual, columnErrors := areTableColumnsEqual(table, *actualTable)
		if !areEqual {
			errors[table.Name] = columnErrors
		}
	}

	if len(errors) > 0 {
		prettyLogErrors(&errors, t)
	}
}

func prettyLogErrors(errors *map[string][]string, t *testing.T) {
	for key, val := range *errors {
		t.Logf("Table %q", key)

		for _, msg := range val {
			t.Logf("  %v", msg)
		}
	}
	t.Fail()
}

func getTableFromDatabase(name string, db *Database) (bool, *Table) {
	for _, table := range db.Tables {
		if name == table.Name {
			return true, &table
		}
	}
	return false, &Table{}
}

func areTableColumnsEqual(want, got Table) (bool, []string) {
	var errors []string

	// TODO: verify columns

	if len(errors) > 0 {
		return false, errors
	}
	return true, nil
}

// Correctly parses query when unknown query is above

// Correctly parses query when unknown query is bellow

// Checks if program can correctly parse CREATE table command with one number datatype

// Checks if program can correctly parse CREATE table command with one text datatype
// with size specified

// Checks if program can correctly parse CREATE table command with multiple datatypes

// Checks if program can correctly parse CREATE table command with NOT NULL rule set

// Checks if program can correctly parse CREATE table command with CONSTRAINT set

// Checks if program can correctly parse CREATE table command with NOT NULL and
// PRIMARY KEY rule set

// Checks if program can correctly parse CREATE table command with PRIMARY KEY rule
// set in a different row (separately)
