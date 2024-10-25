package parser

import (
	"strings"
)

func Parse(text string) Database {
	db := Database{}

	parseAll(&text, &db)

	return db
}

func parseAll(text *string, db *Database) {
	splitQueries := strings.Split(*text, ";")

	for _, query := range splitQueries {
		parseQuery(&query, db)
	}
}

func parseQuery(query *string, db *Database) {
	tokens := TokenizeQuery(*query)

	if len(tokens) < 2 {
		return
	}

	if tokens[1].Type != TABLE {
		return
	}

	switch tokens[0].Type {
	case CREATE:
		handleCreate(&tokens, db)
	}
}

func handleCreate(tokens *[]Token, db *Database) {
	db.Tables = append(db.Tables, Table{Name: (*tokens)[2].Value})
}

type Database struct {
	Tables []Table
}

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name            string
	Type            DataType
	Size            int
	DecimalPoints   int
	IsPrimaryKey    bool
	IsAutoIncrement bool
	IsForeignKey    bool
	IsNotNull       bool
	IsDefault       bool
	IsCheck         bool
	IsUnique        bool
}

const (
	// String data types
	VARCHAR DataType = "varchar"
	CHAR    DataType = "char"
	TEXT    DataType = "text"

	// Numeric data types
	INT     DataType = "int"
	FLOAT   DataType = "float"
	DOUBLE  DataType = "double"
	DECIMAL DataType = "decimal"
	BOOL    DataType = "bool"

	// Date and time data types
	DATE      DataType = "date"
	DATETIME  DataType = "datetime"
	TIMESTAMP DataType = "timestamp"
	TIME      DataType = "time"
	YEAR      DataType = "year"
)

type DataType string
