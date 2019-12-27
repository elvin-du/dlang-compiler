package lexer

func Lex(sourceCode string) []*Token {
	tokens := []*Token{}

	for i := 0; i < len(sourceCode); {
		c := sourceCode[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			token, err := number(sourceCode, i)
			if nil != err {
				panic(err)
			}
			tokens = append(tokens, token)

			i += len(token.Value)
		case ' ':
			i++
		case ';':
			return tokens
		default:
			panic("lex error")
		}
	}

	return tokens
}

//func literal(sourceCode string,i int64) Token {
//
//}
//
func number(sourceCode string, i int) (*Token, error) {
	state := 0
	start := i

	for i < len(sourceCode) {
		c := sourceCode[i]

		if '0' <= c && '9' >= c {
			state = 1
			i++
		} else {
			if 0 == state {
				return nil, NewLexicalError("not a number")
			} else {
				//num,err := strconv.Atoi(sourceCode[start:i])
				//if nil != err{
				//	panic(err)
				//}
				return MakeToken("number", sourceCode[start:i]), nil
			}
		}
	}

	return nil, NewLexicalError("index out of source, not a number")
}

//
//func relop(sourceCode string,i int64) Token {
//
//}
//
//func arithmeticop(sourceCode string, i int64) (*Token, error) {
//
//}

