package sss

type TokenType NodeType // use tree node types as token types. It simplifies parser code

// map to convert token constant to string
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
	Type  TokenType // token constant
	Value string    // token value
}

type Scanner struct {
	buf *string // placeholder for input string
	pos int     // current reader position
}

// Get new scanner
func NewScanner(str *string) *Scanner {
	return &Scanner{buf: str, pos: 0}
}

// Check is rune from letters range
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// recursively tokenize input string
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
				token.Type = LSBRACKET
			case ']':
				token.Type = RSBRACKET
			case '{':
				token.Type = LBRACE
			case '}':
				token.Type = RBRACE
			case '(':
				token.Type = LBRACKET
			case ')':
				token.Type = RBRACKET
			default:
				token.Type = BAD
			}
			token.Value = string(curSymbol)
			*pos = *pos + 1
		}
	}
}

// Read current token from input string and move to next
func (s *Scanner) Read() Token {
	var token = Token{Type: EMPTY, Value: ""}
	read(s.buf, &s.pos, &token)

	return token
}
