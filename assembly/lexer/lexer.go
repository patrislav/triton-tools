package lexer

type TokenKind int

const (
	TokenEOF TokenKind = iota
	TokenIllegal
	TokenIdent           // ident
	TokenLabel           // ident:
	TokenDirective       // !ident
	TokenAddressAbsolute // &ident
	TokenAddressRelative // *ident
	TokenTernary         // %01T01T
	TokenNonary          // $20B
	TokenDecimal         // #123
	TokenLParen          // (
	TokenRParen          // )
	TokenComment         // ; comment
)

type Token struct {
	Kind TokenKind
	Val  string
}

type Lexer struct {
	input   string
	pos     int // current position in input (current char)
	readPos int // current reading position in input (after current char)
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken lexes the character at current position and advances the pointer
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	if len(l.input) <= l.pos {
		return Token{Kind: TokenEOF}
	}

	switch l.ch {
	case '!':
		l.readChar()
		tok.Val = l.readIdentifier()
		tok.Kind = TokenDirective
	case '&':
		l.readChar()
		tok.Val = l.readIdentifier()
		tok.Kind = TokenAddressAbsolute
	case '*':
		l.readChar()
		tok.Val = l.readIdentifier()
		tok.Kind = TokenAddressRelative
	case '%':
		l.readChar()
		tok.Val = l.readTernary()
		tok.Kind = TokenTernary
	case '$':
		l.readChar()
		tok.Val = l.readNonary()
		tok.Kind = TokenNonary
	case '#':
		l.readChar()
		tok.Val = l.readNumber()
		tok.Kind = TokenDecimal
	case '(':
		tok.Kind = TokenLParen
	case ')':
		tok.Kind = TokenRParen
	case ';':
		l.readChar()
		tok.Val = l.readComment()
		tok.Kind = TokenComment
	default:
		if isLetter(l.ch) {
			tok.Val = l.readIdentifier()
			if l.ch == ':' {
				tok.Kind = TokenLabel
				l.readChar()
			} else {
				tok.Kind = TokenIdent
			}
			return tok
		} else {
			// tok.Kind = TokenIllegal
			tok.Kind = TokenEOF
			return tok
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isAlphanum(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNonary() string {
	pos := l.pos
	for isNonary(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readTernary() string {
	pos := l.pos
	for isTernary(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readString() string {
	pos := l.pos + 1
	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readComment() string {
	pos := l.pos
	for {
		l.readChar()
		if l.ch == '\n' {
			break
		}
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '.'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isTernary(ch byte) bool {
	return ch == '0' || ch == '1' || ch == 'T' || ch == 't'
}

func isNonary(ch byte) bool {
	return 'a' <= ch && ch <= 'd' || 'A' <= ch && ch <= 'D' || '0' <= ch && ch <= '4'
}

func isAlphanum(ch byte) bool {
	return isLetter(ch) || isDigit(ch)
}
