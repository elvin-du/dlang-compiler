package lexer

import "fmt"

type LexicalError struct {
	Msg string
}

func NewLexicalError(msg string) *LexicalError {
	return &LexicalError{msg}
}

func (le *LexicalError) Error() string {
	return fmt.Sprintf(`{"msg":%s}`, le.Msg)
}
