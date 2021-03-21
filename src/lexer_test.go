package src

import "testing"

var lexer func(text string) []Token

func init() {
	/* load test data */
	lexer = NewLexer().
		add(Token{Name: "NUMBER", Pattern: "^[0-9]+$"}).
		add(Token{Name: "PLUS", Pattern: "^\\+$"}).
		add(Token{Name: "MINUS", Pattern: "^\\-$"}).
		add(Token{Name: "MULTIPLY", Pattern: "^\\*$"}).
		add(Token{Name: "DIVIDE", Pattern: "^\\/$"}).
		add(Token{Name: "EXPONENT", Pattern: "^\\^$"}).
		build()

	return
}

func TestPlus(t *testing.T) {
	str := "1+2"
	tokens := lexer(str)

	if len(tokens) != 3 || tokens[0].Name != "NUMBER" || tokens[1].Name != "PLUS" || tokens[2].Name != "NUMBER" {
		t.Errorf("Test failed for %s", str)
	}
}
