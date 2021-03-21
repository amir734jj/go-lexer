package go_lexer

import "testing"

var lexer func(text string) ([]Token, error)

func init() {
	/* load test data */
	lexer = NewLexer().
		Add(Token{Name: "NUMBER", Pattern: "^[0-9]+$"}).
		Add(Token{Name: "PLUS", Pattern: "^\\+$"}).
		Add(Token{Name: "MINUS", Pattern: "^\\-$"}).
		Add(Token{Name: "MULTIPLY", Pattern: "^\\*$"}).
		Add(Token{Name: "DIVIDE", Pattern: "^\\/$"}).
		Add(Token{Name: "EXPONENT", Pattern: "^\\^$"}).
		Add(Token{Name: "SPACE", Pattern: "^\\s+$", Ignore: true}).
		Build()

	return
}

func TestPlus(t *testing.T) {
	str := "1 + 2"
	tokens, err := lexer(str)

	if err != nil {
		t.Error(err.Error())
	}

	if len(tokens) != 3 || tokens[0].Name != "NUMBER" || tokens[1].Name != "PLUS" || tokens[2].Name != "NUMBER" {
		t.Errorf("Test failed for %s", str)
	}
}
