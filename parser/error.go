package parser

import "fmt"

type SyntaxError struct {
	Msg string
}

func NewSyntaxError(msg string) *SyntaxError {
	return &SyntaxError{msg}
}

func (se *SyntaxError) Error() string {
	return fmt.Sprintf(`{"msg":%s}`, se.Msg)
}
