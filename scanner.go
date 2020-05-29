package main

const (
	EMPTY     = iota // no token
	LITERAL          // letter(s)
	LBRACKET         // (
	RBRACKET         // )
	LBRACE           // {
	RBRACE           // }
	LSBRACKET        // [
	RSBRACKET        // ]
	BAD              // bad input
	EOF              // end of buffer
)

type TokenType int

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

func read(str *string, pos *int, token *TokenType, value *string) {
	if *pos >= len(*str) {
		if *token == EMPTY {
			*token = EOF
		}
		return
	}
	curSymbol := rune((*str)[*pos])
	if isLetter(curSymbol) {
		if *token == EMPTY {
			*token = LITERAL
			*value = string(curSymbol)
			*pos = *pos + 1
			read(str, pos, token, value)
		} else if *token == LITERAL {
			*value += string(curSymbol)
			*pos = *pos + 1
			read(str, pos, token, value)
		} else {
			// unread symbol
			*pos = *pos - 1
		}
	} else {
		if *token != LITERAL {
			switch curSymbol {
			case '[':
				{
					*token = LSBRACKET
					*value = "["
				}
			case ']':
				{
					*token = RSBRACKET
					*value = "]"
				}
			case '{':
				{
					*token = LBRACE
					*value = "{"
				}
			case '}':
				{
					*token = RBRACE
					*value = "}"
				}
			case '(':
				{
					*token = LBRACKET
					*value = "("
				}
			case ')':
				{
					*token = RBRACKET
					*value = ")"
				}
			default:
				{
					*token = BAD
					*value = ""
				}
			}
			*pos = *pos + 1
		}
	}
}

func (s *Scanner) Read() (TokenType, string) {
	var token TokenType = EMPTY
	value := ""
	read(s.buf, &s.pos, &token, &value)

	return token, value
}
