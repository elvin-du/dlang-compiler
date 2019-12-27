package parser

import "dlang-compiler/lexer"

/*
CFG（Context Free Grammar）
E -> 1Ep | 1
Ep -> + E
pop the first token
 */
func ShiftTokens(tokens *[]*lexer.Token) *lexer.Token {
	token := (*tokens)[0]

	*tokens = (*tokens)[1:len(*tokens)]

	return token
}

func Parse(ast *ASTTreeNode, tokens []*lexer.Token) {
	token := ShiftTokens(&tokens)
	if "1" == token.Value {
		ast.Left = "1"
		rAst := ASTTreeNode{}
		parse(&rAst, tokens)
	} else {
		panic("syntax err: should be  1")
	}
}

func parse(ast *ASTTreeNode, tokens []*lexer.Token) {
	token := ShiftTokens(&tokens)
	if "+" == token.Value {

		Parse(ast, tokens)
	} else {
		panic("syntax err:should be +")
	}
}
