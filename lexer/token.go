package lexer

import "fmt"

type Token struct {
	Type  string
	Value string
}

func MakeToken(tp, val string) *Token {
	return &Token{tp, val}
}

func (t *Token) String() string {
	return fmt.Sprintf(`{"type":"%s","value":"%s"}`, t.Type, t.Value)
}
