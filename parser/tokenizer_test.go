package parser

import "testing"

// should tokenize simple create query without column definitions
func TestEmptyCreateQueryTokenization(t *testing.T) {
	gotText := "CREATE TABLE Persons ();"
	want := []Token{
		{Type: CREATE, Value: "CREATE"},
		{Type: TABLE, Value: "TABLE"},
		{Type: NAME, Value: "Persons"},
		{Type: LPAREN, Value: "("},
		{Type: RPAREN, Value: ")"},
		{Type: SEMICOLON, Value: ";"},
		{Type: END, Value: "0"},
	}

	gotTokens := TokenizeQuery(gotText)

	assertTokens(want, gotTokens, t)
}

func assertTokens(want, got []Token, t *testing.T) {
	if len(want) != len(got) {
		t.Log("Wrong number of tokens generated")
		t.Logf("Expected %v tokens: %+v", len(want), want)
		t.Logf("Got %v tokens: %+v", len(got), got)
		t.Fail()
		return
	}

	for index, token := range want {
		if token.Type != got[index].Type {
			t.Fatalf("TokenType mismatch. Expected %q but got %q", token.Type, got[index].Type)
		}

		if token.Value != got[index].Value {
			t.Fatalf("Token value mismatch. Expected %q but got %q", token.Value, got[index].Value)
		}
	}
}
