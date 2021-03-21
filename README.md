# go-lexer

Simple lexer in golang. It uses:

- regular expressions to tokenize the strings
- greedy approach to find the longest possible token

```go
var lexer func(text string) ([]Token, error) = NewLexer().
    Add(Token{Name: "NUMBER", Pattern: "^[0-9]+$"}).
    Add(Token{Name: "PLUS", Pattern: "^\\+$"}).
    Add(Token{Name: "SPACE", Pattern: "^\\s+$", Ignore: true}).
    Build()
    
str := "1 + 2"
tokens := lexer(str)  // should return 3 tokens
```
