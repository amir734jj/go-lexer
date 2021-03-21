# go-lexer

Simple lexer in golang. It uses:

- regular expressions to tokenize the strings
- greedy approach to find the longest possible token

```go
var lexer func(text string) ([]Token, error) = NewLexer().
    add(Token{Name: "NUMBER", Pattern: "^[0-9]+$"}).
    add(Token{Name: "PLUS", Pattern: "^\\+$"}).
    add(Token{Name: "SPACE", Pattern: "^\\s+$", Ignore: true}).
    build()
    
    str := "1 + 2"
    tokens := lexer(str)  // should return 3 tokens
```
