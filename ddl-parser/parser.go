package parser

func Parse(text string) (Database, error) {
	// create database object (we will be passing its pointer)
	// create text object (we will be passing its pointer)

	// call parseAll method to do all the parsing and to return a database object

	// handle any errors
	// return
	return Database{}, nil
}

func parseAll(text *string, db *Database) error {
	// split all text into queries by semicolons
	// call parseQuery on each query
	return nil
}

func parseQuery(query *string, db *Database) error {
	// tokenize query
	// have a switch that calls appropriate method on different queries (CREATE, APPEND, DELETE)
	return nil
}

func handleCreate(db *Database) {
	// all logic to correctly parse CREATE query to the database
}

func TokenizeQuery(query *string) ([]Token, error) {
	return []Token{}, nil
}
