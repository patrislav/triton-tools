package lexer

import (
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := `!defmacro over
	swap dup rot swap
!endmacro

main:
	#123 $ABC %T10T10 ; this is a comment
	*hello &__runtime__.print jsr
	(inc over)`

	tests := []struct{
		expKind TokenKind
		expVal  string
	}{
		{TokenDirective, "defmacro"},
		{TokenIdent, "over"},
		{TokenIdent, "swap"},
		{TokenIdent, "dup"},
		{TokenIdent, "rot"},
		{TokenIdent, "swap"},
		{TokenDirective, "endmacro"},
		{TokenLabel, "main"},
		{TokenDecimal, "123"},
		{TokenNonary, "ABC"},
		{TokenTernary, "T10T10"},
		{TokenComment, " this is a comment"},
		{TokenAddressRelative, "hello"},
		{TokenAddressAbsolute, "__runtime__.print"},
		{TokenIdent, "jsr"},
		{TokenLParen, ""},
		{TokenIdent, "inc"},
		{TokenIdent, "over"},
		{TokenRParen, ""},
		{TokenEOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Kind != tt.expKind {
			t.Fatalf("tests[%d] tok.Kind wrong - expected=%q, got=%q", i, tt.expKind, tok.Kind)
		}
		if tok.Val != tt.expVal {
			t.Fatalf("tests[%d] tok.Val wrong - expected=%q, got=%q", i, tt.expVal, tok.Val)
		}
	}
}
