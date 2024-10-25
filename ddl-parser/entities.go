package parser

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

type Token struct {
	Type  TokenType
	Value string
}

const (
	// Words
	CREATE TokenType = "create"
	TABLE  TokenType = "table"
	NAME   TokenType = "name"

	// Symbols
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	SEMICOLON TokenType = ";"
)

type TokenType string
