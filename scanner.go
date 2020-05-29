package main

type TokenType NodeType

var tokenMap = map[TokenType]string{
	EMPTY:     "EMPTY",     // no token
	LITERAL:   "LITERAL",   // letter(s)
	LBRACKET:  "LBRACKET",  // (
	RBRACKET:  "RBRACKET",  // )
	LBRACE:    "LBRACE",    // {
	RBRACE:    "RBRACE",    // }
	LSBRACKET: "LSBRACKET", // [
	RSBRACKET: "RSBRACKET", // ]
	BAD:       "BAD",       // bad input
	EOF:       "EOF",       // end of buffer
}

type Token struct {
	Type  TokenType
	Value string
}

type Scanner struct {
	buf *string
	pos int
}

func NewScanner(str *string) *Scanner {
	return &Scanner{buf: str, pos: 0}
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func read(str *string, pos *int, token *Token) {
	if *pos >= len(*str) {
		if token.Type == EMPTY {
			token.Type = EOF
		}
		return
	}
	curSymbol := rune((*str)[*pos])
	if isLetter(curSymbol) {
		if token.Type == EMPTY {
			token.Type = LITERAL
			token.Value = string(curSymbol)
			*pos = *pos + 1
			read(str, pos, token)
		} else if token.Type == LITERAL {
			token.Value += string(curSymbol)
			*pos = *pos + 1
			read(str, pos, token)
		} else {
			// unread symbol
			*pos = *pos - 1
		}
	} else {
		if token.Type != LITERAL {
			switch curSymbol {
			case '[':
				{
					token.Type = LSBRACKET
					token.Value = "["
				}
			case ']':
				{
					token.Type = RSBRACKET
					token.Value = "]"
				}
			case '{':
				{
					token.Type = LBRACE
					token.Value = "{"
				}
			case '}':
				{
					token.Type = RBRACE
					token.Value = "}"
				}
			case '(':
				{
					token.Type = LBRACKET
					token.Value = "("
				}
			case ')':
				{
					token.Type = RBRACKET
					token.Value = ")"
				}
			default:
				{
					token.Type = BAD
					token.Value = string(curSymbol)
				}
			}
			*pos = *pos + 1
		}
	}
}

func (s *Scanner) Read() Token {
	var token = Token{Type: EMPTY, Value: ""}
	read(s.buf, &s.pos, &token)

	return token
}
