package parser

import (
	"errors"
	"strings"
)

func TokenizeQuery(query string) []Token {
	tokenizer, err := createTokenizer(query)
	if err != nil {
		return nil
	}

	for {
		if tokenizer.CurrValue == 0 {
			tokenizer.Tokens = append(tokenizer.Tokens, Token{Type: END, Value: "0"})
			break
		}

		if isEmptySpace(tokenizer.CurrValue) {
			tokenizer.readNextSymbol()
			continue
		}

		tokenizer.readToken()

		tokenizer.readNextSymbol()
	}
	return tokenizer.Tokens
}

func (t *tokenizer) readToken() {
	switch string(t.CurrValue) {
	case LPAREN:
		t.Tokens = append(t.Tokens, Token{Type: LPAREN, Value: string(t.CurrValue)})

	case RPAREN:
		t.Tokens = append(t.Tokens, Token{Type: RPAREN, Value: string(t.CurrValue)})

	case SEMICOLON:
		t.Tokens = append(t.Tokens, Token{Type: SEMICOLON, Value: string(t.CurrValue)})

	default:
		t.readWord()
	}
}

func (t *tokenizer) readWord() {
	var value = ""
	for {
		value += string(t.CurrValue)
		if isAlphanumeric(t.NextValue) {
			t.readNextSymbol()
		} else {
			lower := strings.ToLower(value)
			switch lower {
			case CREATE:
				t.Tokens = append(t.Tokens, Token{Type: CREATE, Value: value})
			case TABLE:
				t.Tokens = append(t.Tokens, Token{Type: TABLE, Value: value})
			default:
				t.Tokens = append(t.Tokens, Token{Type: NAME, Value: value})
			}
			break
		}
	}
}

func createTokenizer(query string) (*tokenizer, error) {
	if len(query) < 2 {
		return nil, errors.New("Invalid query: too short")
	}

	return &tokenizer{
		Text:            query,
		CurrentLocation: 0,
		Tokens:          make([]Token, 0),
		CurrValue:       query[0],
		NextValue:       query[1],
	}, nil
}

func (t *tokenizer) readNextSymbol() {
	t.CurrentLocation++
	t.CurrValue = t.NextValue

	if t.CurrentLocation+1 >= len(t.Text) {
		t.NextValue = 0
	} else {
		t.NextValue = t.Text[t.CurrentLocation+1]
	}
}

func isEmptySpace(c byte) bool {
	return c == '\n' || c == '\t' || c == ' '
}

func isAlphanumeric(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '_'
}

type tokenizer struct {
	Tokens          []Token
	Text            string
	CurrentLocation int
	CurrValue       byte
	NextValue       byte
}

type Token struct {
	Type  string
	Value string
}

const (
	// Words
	CREATE string = "create"
	TABLE  string = "table"
	NAME   string = "name"

	// Symbols
	LPAREN    string = "("
	RPAREN    string = ")"
	SEMICOLON string = ";"

	// Special
	END string = "end"
)
