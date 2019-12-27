package main

import (
	"log"
	"dlang-compiler/lexer"
)
const suorceCode = `123 789;`

func main() {
	log.Printf("%+v",lexer.Lex(suorceCode))

}
