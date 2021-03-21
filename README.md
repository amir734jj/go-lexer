# go-lexer

Simple lexer in golang

```go
var lexer func(text string) []Token = NewLexer().
    add(Token{Name: "NUMBER", Pattern: "^[0-9]+$"}).
    add(Token{Name: "PLUS", Pattern: "^\\+$"}).
    add(Token{Name: "MINUS", Pattern: "^\\-$"}).
    add(Token{Name: "MULTIPLY", Pattern: "^\\*$"}).
    add(Token{Name: "DIVIDE", Pattern: "^\\/$"}).
    add(Token{Name: "EXPONENT", Pattern: "^\\^$"}).
    add(Token{Name: "SPACE", Pattern: "^\\s+$", Ignore: true}).
    build()
```